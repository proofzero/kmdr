/*
Copyright © 2021 kubelt

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package api

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/proofzero/kmdr/util"
	"golang.org/x/crypto/nacl/sign"
	"golang.org/x/crypto/sha3"
)

// Interface
// -----------------------------------------------------------------------------

// AuthAPI
type AuthAPI interface {
	AddKeys(username string) error
	LoadKeys(username string) error
	AuthKey() *[32]byte
	EncryptionKey() *[32]byte
	SignNode(node []byte) ([]byte, error)
	ValidateNode(signature []byte) bool
}

// TODO: methods to support multi format
// TODO: define types for the 32 and 64 byte lengths and update signatures

// Implementation
// -----------------------------------------------------------------------------

// Auth

type auth struct {
	keysDir       string
	SigningKeys   *signingKeys
	EncrptionKeys *encryptionKeys
}

type signingKeys struct {
	PublicSigningKey  *[32]byte
	privateSigningKey *[64]byte
}

type encryptionKeys struct {
	PublicEncryptionKey  *[32]byte
	privateEncryptionKey *[32]byte
}

func newAuthApi() (AuthAPI, error) {
	home, _ := homedir.Dir()
	keysDir := fmt.Sprintf("%s/.config/kubelt/keys", home)
	return &auth{
		keysDir:       keysDir,
		SigningKeys:   &signingKeys{},
		EncrptionKeys: &encryptionKeys{},
	}, nil
}

func (a *auth) AddKeys(username string) error {
	publicSig, privateSig, err := util.GenerateSigningKeys()
	if err != nil {
		return err
	}
	pk, key, err := util.GenerateEncryptionKeys()
	if err != nil {
		return err
	}

	_, err = os.Create(a.keysDir)
	if err != nil {
		return err
	}

	// Save signing keys
	pubSigFile := fmt.Sprintf("%s/signing_%s.pub", a.keysDir, username)
	secSigkFile := fmt.Sprintf("%s/signing_%s", a.keysDir, username)
	err = ioutil.WriteFile(pubSigFile, privateSig[:], 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(secSigkFile, publicSig[:], 0644)
	if err != nil {
		return err
	}

	// Save encryption keys
	pkFile := fmt.Sprintf("%s/%s.pub", a.keysDir, username)
	skFile := fmt.Sprintf("%s/%s", a.keysDir, username)
	err = ioutil.WriteFile(pkFile, pk[:], 0644)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(skFile, key[:], 0644)
	if err != nil {
		return err
	}

	a.SigningKeys.privateSigningKey = privateSig
	a.SigningKeys.PublicSigningKey = publicSig

	a.EncrptionKeys.privateEncryptionKey = key
	a.EncrptionKeys.PublicEncryptionKey = pk

	return nil
}

func (a *auth) LoadKeys(username string) error {
	// Read signing keys
	pubSigFile := fmt.Sprintf("%s/signing_%s.pub", a.keysDir, username)
	secSigkFile := fmt.Sprintf("%s/signing_%s", a.keysDir, username)
	privateSig, err := ioutil.ReadFile(pubSigFile)
	if err != nil {
		return err
	}
	publicSig, err := ioutil.ReadFile(secSigkFile)
	if err != nil {
		return err
	}

	// Read encryption keys
	pkFile := fmt.Sprintf("%s/%s.pub", a.keysDir, username)
	skFile := fmt.Sprintf("%s/%s", a.keysDir, username)
	pk, err := ioutil.ReadFile(pkFile)
	if err != nil {
		return err
	}
	key, err := ioutil.ReadFile(skFile)
	if err != nil {
		return err
	}

	var privateSig64 [64]byte
	copy(privateSig64[:], privateSig)

	var publicSig32 [32]byte
	copy(publicSig32[:], publicSig)

	a.SigningKeys.privateSigningKey = &privateSig64
	a.SigningKeys.PublicSigningKey = &publicSig32

	var key32 [32]byte
	copy(key32[:], key)

	var pk32 [32]byte
	copy(key32[:], pk)

	a.EncrptionKeys.privateEncryptionKey = &key32
	a.EncrptionKeys.PublicEncryptionKey = &pk32

	return nil
}

func (a *auth) AuthKey() *[32]byte {
	return a.SigningKeys.PublicSigningKey
}

func (a *auth) EncryptionKey() *[32]byte {
	return a.EncrptionKeys.PublicEncryptionKey
}

func (a *auth) SignNode(node []byte) ([]byte, error) {
	h := make([]byte, 64)
	sha3.ShakeSum256(h, node)
	return sign.Sign(nil, h, a.SigningKeys.privateSigningKey), nil
}

func (a *auth) ValidateNode(signature []byte) bool {
	_, ok := sign.Open(nil, signature, a.SigningKeys.PublicSigningKey)
	return ok
}
