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
	"context"
	"fmt"

	"cuelang.org/go/cue"
	kb "github.com/proofzero/proto/pkg/v1alpha1"
)

// Apply calls out to ktrl to mutate the kubelt graph using values supplied by the user
func (k *Client) Apply(cueValue cue.Value) (*kb.ApplyDefault, error) {
	ctx := k.Config.Contexts[k.Config.CurrentContext]
	cueString := fmt.Sprint(cueValue)
	fmt.Println(cueString)
	resource := &kb.Cue{
		Cue: cueString,
	}
	request := &kb.ApplyRequest{Resources: resource, Context: ctx}
	r, err := k.Client.Apply(context.Background(), request)
	return r, err
}
