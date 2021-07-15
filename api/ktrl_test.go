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
	"google.golang.org/grpc"

	kb "github.com/proofzero/proto/pkg/v1alpha1"
	tkb "github.com/proofzero/proto/pkg/v1alpha1/test"
)

func TestNewKtrlClient(t *testing.T) {
	type args struct {
		config  Config
		options []grpc.DialOption
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{
			name: "create ktrl client with no args",
			args: args{config: Config{}, options: nil},
			want: Config{
				Ktrl: KtrlConfig{
					Server: ServerConfig{
						Port:     ":50051",
						Protocol: "tcp",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "create ktrl client with args",
			args: args{config: Config{}, options: []grpc.DialOption{grpc.WithInsecure()}},
			want: Config{
				Ktrl: KtrlConfig{
					Server: ServerConfig{
						Port:     ":50051",
						Protocol: "tcp",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewKtrlClient(tt.args.config, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKtrlClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*got.Config, tt.want) {
				t.Errorf("NewKtrlClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		Config     *Config
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
				Config: &Config{
					Ktrl:           KtrlConfig{},
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
			k := &Client{
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
