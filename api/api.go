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
	"io/fs"
	"io/ioutil"
	"strings"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

// this api is used multuple times during an execution
// so we create a singleton
var (
	kmdrapi  *kmdrAPI
	StaticFS embed.FS
)

// NewAPI creates a instance of KmdrAPI
func NewAPI() KmdrAPI {
	if kmdrapi != nil {
		return *kmdrapi
	}

	config := load.Config{
		Overlay: make(map[string]load.Source),
	}
	fs.WalkDir(StaticFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.Contains(path, ".cue") {
			f, err := ioutil.ReadFile(path)
			config.Overlay[path] = load.FromString(string(f))
			return err
		}
		if d != nil {
			return nil
		}
		return err
	})
	instances := load.Instances([]string{"kmdr"}, &config)

	cueContext := cuecontext.New()

	// kmdrInstance := instances[0]
	// instance := cueContext.BuildInstance(kmdrInstance)

	cue := cueAPI{
		context:   cueContext,
		instances: instances,
	}
	cue.schemaFetcher = cue.fetchSchema

	ktrl, _ := newKtrlAPI()

	kmdrapi = &kmdrAPI{
		cue:  cue,
		ktrl: ktrl,
	}
	return *kmdrapi
}

// KmdrApi
type KmdrAPI interface {
	Cue() CueAPI
	Ktrl() KtrlAPI
}

// kmdrAPI
type kmdrAPI struct {
	cue  CueAPI
	ktrl KtrlAPI
}

// Cue returns an instance of the CueAPI
func (api kmdrAPI) Cue() CueAPI {
	return api.cue
}

// Ktrl returns an instance of the KtrlAPI
func (api kmdrAPI) Ktrl() KtrlAPI {
	return api.ktrl
}
