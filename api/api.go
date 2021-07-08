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
	"cuelang.org/go/cue/cuecontext"
)

func NewAPI() KmdrAPI {
	// home, _ := homedir.Dir()
	// importPathOption := cue.ImportPath(fmt.Sprintf("%s/.config/kubelt/definitions", home))
	cueContext := cuecontext.New()
	cue := cueAPI{
		context: cueContext,
		// importPath: importPathOption,
	}
	cue.schemaFetcher = cue.fetchSchema

	return kmdrAPI{
		cue: cue,
	}
}

type KmdrAPI interface {
	Cue() CueAPI
}

type kmdrAPI struct {
	cue CueAPI
}

func (api kmdrAPI) Cue() CueAPI {
	return api.cue
}
