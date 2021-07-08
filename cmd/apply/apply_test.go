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
package apply

import (
	"fmt"
	"testing"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"github.com/golang/mock/gomock"
	"github.com/proofzero/kmdr/api"
)

const validSchema string = `
manifests: {
	foo: {
		kind:       "Space"
		apiVersion: "kubelt.com/v1alpha1"
	}
}
`

const invalidSchema string = `
manifests: {
	foo: {
		kind:       "Space"
		apiVersion: "kublt.com/v1"
		invalidfield: "hi"
	}
}
`

const validDefinition string = `
#manifests: [_=string]: {
    apiVersion: string
    kind: string
	metadata: {
		name: string
		[_]: _
	}
	data?: {
		[_]: _
	}
}
`

func Test_runValidation(t *testing.T) {
	type args struct {
		applyStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "validate applied manifest",
			args: args{
				applyStr: validSchema,
			},
			wantErr: false,
		},
		{
			name: "invalidate applied manifest",
			args: args{
				applyStr: invalidSchema,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			context := cuecontext.New()
			d := context.CompileString(validDefinition)
			v := context.CompileString(tt.args.applyStr)
			m := api.NewMockCueAPI(ctrl)
			m.EXPECT().CompileSchemaFromString(tt.args.applyStr).Return(v, nil).AnyTimes()
			if tt.wantErr {
				m.EXPECT().FetchSchema(gomock.Any()).Return(cue.Value{}, fmt.Errorf("bad")).AnyTimes()
				m.EXPECT().ValidateResource(gomock.Any(), gomock.Any()).Return(fmt.Errorf("invalid")).AnyTimes()
			} else {
				m.EXPECT().FetchSchema("kubelt://kmdr").Return(d, nil).AnyTimes()
				m.EXPECT().ValidateResource(gomock.Any(), gomock.Any()).AnyTimes()
			}

			got, err := runValidation(tt.args.applyStr, m)
			if (err != nil) != tt.wantErr {
				t.Errorf("runValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Err() != nil && tt.wantErr == false {
				t.Errorf("runValidation() = %v", got)
			}
		})
	}
}
