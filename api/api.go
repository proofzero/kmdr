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
	"embed"
)

// this api is used multuple times during an execution
// so we create a singleton
var (
	kmdrapi  *kmdrAPI
	StaticFS embed.FS
)

// NewAPI creates a instance of KmdrAPI
func NewAPI() (KmdrAPI, error) {
	if kmdrapi != nil {
		return *kmdrapi, nil
	}

	configapi, err := newConfigAPI()
	if err != nil {
		return nil, err
	}
	cueapi, err := newCueAPI()
	if err != nil {
		return nil, err
	}
	ktrlapi, err := newKtrlAPI()
	if err != nil {
		return nil, err
	}

	kmdrapi = &kmdrAPI{
		config: configapi,
		cue:    cueapi,
		ktrl:   ktrlapi,
	}
	return *kmdrapi, nil
}

// KmdrApi
type KmdrAPI interface {
	Config() ConfigAPI
	Cue() CueAPI
	Ktrl() KtrlAPI
}

// kmdrAPI
type kmdrAPI struct {
	config ConfigAPI
	cue    CueAPI
	ktrl   KtrlAPI
}

// Config returns an instance of the CueAPI
func (api kmdrAPI) Config() ConfigAPI {
	return api.config
}

// Cue returns an instance of the CueAPI
func (api kmdrAPI) Cue() CueAPI {
	return api.cue
}

// Ktrl returns an instance of the KtrlAPI
func (api kmdrAPI) Ktrl() KtrlAPI {
	return api.ktrl
}
