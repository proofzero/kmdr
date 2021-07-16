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
	"reflect"
	"testing"

	"cuelang.org/go/cue"
	"github.com/golang/mock/gomock"
	kb "github.com/proofzero/proto/pkg/v1alpha1"
	tkb "github.com/proofzero/proto/pkg/v1alpha1/test"
	"google.golang.org/grpc"
)

const validCue = `
foo: {
	apiVersion: "kubelt.com/v1alpha1"
	kind: "Space"
}
`

func TestKtrlClient_Apply(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		Connection *grpc.ClientConn
		Client     *tkb.MockKubeltClient
		Config     ktrlConfig
	}
	type args struct{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *kb.ApplyDefault
		wantErr bool
	}{
		{
			name: "run apply",
			fields: fields{
				Connection: nil,
				Client:     tkb.NewMockKubeltClient(ctrl),
				Config: ktrlConfig{
					CurrentContext: "",
					Contexts: map[string]*kb.Context{
						"default": {Name: "test"},
					},
				},
			},
			args:    args{},
			want:    &kb.ApplyDefault{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &ktrlAPI{
				Connection: tt.fields.Connection,
				Client:     tt.fields.Client,
				Config:     tt.fields.Config,
			}

			tt.fields.Client.EXPECT().Apply(
				gomock.Any(),
				gomock.Any(),
			).Return(&kb.ApplyDefault{}, nil)

			schema, _ := NewAPI().Cue().CompileSchemaFromString(validCue)
			val := schema.LookupPath(cue.ParsePath("foo"))
			got, err := k.Apply(val)
			if (err != nil) != tt.wantErr {
				t.Errorf("KtrlClient.Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KtrlClient.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}
