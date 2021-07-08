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

// NewAPI creates a instance of KmdrAPI
func NewAPI() KmdrAPI {
	cueContext := cuecontext.New()
	cue := cueAPI{
		context: cueContext,
	}
	cue.schemaFetcher = cue.fetchSchema

	return kmdrAPI{
		cue: cue,
	}
}

// KMDRApi
type KmdrAPI interface {
	Cue() CueAPI
}

// kmdrAPI
type kmdrAPI struct {
	cue CueAPI
}

// Cue returns an instance of the CueAPI
func (api kmdrAPI) Cue() CueAPI {
	return api.cue
}
