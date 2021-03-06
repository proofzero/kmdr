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

	"cuelang.org/go/cue"
	"github.com/proofzero/kmdr/util"
	kb "github.com/proofzero/proto/pkg/v1alpha1"
)

// Interface
// -----------------------------------------------------------------------------

// PlanAPI
type PlanAPI interface {
	Lookup(*cue.Value) (*kb.Manifests, error)
	Plan([]*kb.Node) (*kb.Commands, error)
	User(string, string) *kb.Command
}

// TODO: methods to support multi format
// TODO: define types for the 32 and 64 byte lengths and update signatures

// Implementation
// -----------------------------------------------------------------------------

// Plan

type plan struct {
	context *AuthAPI
}

func newPlanAPI(context *AuthAPI) (PlanAPI, error) {
	return &plan{context: context}, nil
}

func (p *plan) Lookup(*cue.Value) (*kb.Manifests, error) {
	return nil, nil
}

func (p *plan) Plan(nodes []*kb.Node) (*kb.Commands, error) {
	return nil, nil
}

func (p *plan) User(username string, email string) *kb.Command {
	// TODO: encode keys in a better format
	// Saltpack? Multihash base64?
	user := kb.UserData{
		Email:        email,
		PublicKey:    util.Bytes32ToBytes((*p.context).EncryptionKey()),
		SignatureKey: util.Bytes32ToBytes((*p.context).AuthKey()),
	}

	fmt.Println(user.String())
	signature, _ := (*p.context).SignNode(user.String())
	author := kb.Author{
		// ID to be generated by KTRL RegisterUser reciever
		Signature: signature,
	}

	cmd := kb.Command{
		Op: kb.Command_UPSERT,
		Node: &kb.Command_User{
			User: &kb.UserCommand{
				Header: &kb.Header{
					Author: &author,
				},
				Metadata: &kb.Metadata{
					Name: username,
				},
				Settings: &kb.CommandSettings{
					Timeout: 5,
				},
				Data: &user,
			},
		},
	}
	return &cmd
}
