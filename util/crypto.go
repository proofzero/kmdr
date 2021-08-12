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
package util

import (
	"crypto/rand"

	"golang.org/x/crypto/nacl/box"
	"golang.org/x/crypto/nacl/sign"
)

func GenerateSigningKeys() (*[32]byte, *[64]byte, error) {
	public, private, err := sign.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return public, private, nil
}

func GenerateEncryptionKeys() (*[32]byte, *[32]byte, error) {
	public, private, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return public, private, nil
}

// GenerateKeyEncryptionKey uses the maintainers secret key to generate a shared encryption key.
// The shared encryption key is used to encrypt content. The receipient can then use
// their private key with the maintainers public key to generate the decryption key
// To generate encryption key the author/sender will use their private key and reciever's public key
// To generate a decrtyption key the reciver will use their private key and the author/sender's public key
func GenerateSharedEncryptionKey(privateKey *[32]byte, publicKey *[32]byte) *[32]byte {
	sharedEncryptKey := new([32]byte)
	box.Precompute(sharedEncryptKey, publicKey, privateKey)
	return sharedEncryptKey
}

func Sign(node []byte, key *[64]byte) []byte {
	return sign.Sign(nil, node, key)
}

func ValidateSignature(signature []byte, key *[32]byte) bool {
	_, ok := sign.Open(nil, signature, key)
	return ok
}

func Bytes32ToBytes(key *[32]byte) []byte {
	var b []byte
	copy(b, key[:])
	return b
}
