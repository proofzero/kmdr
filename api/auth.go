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

	"github.com/adrg/xdg"
	"github.com/proofzero/kmdr/util"
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
// Update the key structs to support the multiformat

// TODO: armor the keys with our multihash library
// Store the keys to disk after encoding them to a text format

// Implementation
// -----------------------------------------------------------------------------

// Auth

type auth struct {
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
	return &auth{
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

	// Save signing keys
	if secSigFile, err := xdg.ConfigFile(fmt.Sprintf("kubelt/keys/signing_%s", username)); err != nil {
		return err
	} else {
		err = ioutil.WriteFile(secSigFile, privateSig[:], 0600)
		if err != nil {
			return err
		}
	}
	if pubSigFile, err := xdg.ConfigFile(fmt.Sprintf("kubelt/keys/signing_%s.pub", username)); err != nil {
		return err
	} else {
		err = ioutil.WriteFile(pubSigFile, publicSig[:], 0600)
		if err != nil {
			return err
		}
	}

	// Save encryption keys
	if secKeyFile, err := xdg.ConfigFile(fmt.Sprintf("kubelt/keys/%s", username)); err != nil {
		return err
	} else {
		err = ioutil.WriteFile(secKeyFile, key[:], 0600)
		if err != nil {
			return err
		}
	}
	if pubKeyFile, err := xdg.ConfigFile(fmt.Sprintf("kubelt/keys/%s.pub", username)); err != nil {
		return err
	} else {
		err = ioutil.WriteFile(pubKeyFile, pk[:], 0600)
		if err != nil {
			return err
		}
	}

	a.SigningKeys.privateSigningKey = privateSig
	a.SigningKeys.PublicSigningKey = publicSig

	a.EncrptionKeys.privateEncryptionKey = key
	a.EncrptionKeys.PublicEncryptionKey = pk

	return nil
}

func (a *auth) LoadKeys(username string) error {
	// Read signing keys
	var privateSig64 [64]byte
	var publicSig32 [32]byte

	if pubSigFile, err := xdg.ConfigFile(fmt.Sprintf("kubelt/keys/signing_%s.pub", username)); err != nil {
		return err
	} else {
		if publicSig, err := ioutil.ReadFile(pubSigFile); err != nil {
			return err
		} else {
			copy(publicSig32[:], publicSig)
		}
	}

	if secSigFile, err := xdg.ConfigFile(fmt.Sprintf("kubelt/keys/signing_%s", username)); err != nil {
		return err
	} else {
		if secretSig, err := ioutil.ReadFile(secSigFile); err != nil {
			return err
		} else {
			copy(privateSig64[:], secretSig)
		}
	}

	// Read encryption keys
	var key32 [32]byte
	var pk32 [32]byte

	if pubKeyFile, err := xdg.ConfigFile(fmt.Sprintf("kubelt/keys/%s.pub", username)); err != nil {
		return err
	} else {
		if publicKey, err := ioutil.ReadFile(pubKeyFile); err != nil {
			return err
		} else {
			copy(key32[:], publicKey)
		}
	}

	if secKeyFile, err := xdg.ConfigFile(fmt.Sprintf("kubelt/keys/%s", username)); err != nil {
		return err
	} else {
		if secretKey, err := ioutil.ReadFile(secKeyFile); err != nil {
			return err
		} else {
			copy(key32[:], secretKey)
		}
	}

	a.SigningKeys.privateSigningKey = &privateSig64
	a.SigningKeys.PublicSigningKey = &publicSig32

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
	return util.Sign(h, a.SigningKeys.privateSigningKey), nil
}

func (a *auth) ValidateNode(signature []byte) bool {
	return util.ValidateSignature(signature, a.SigningKeys.PublicSigningKey)
}
