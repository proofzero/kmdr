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
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/load"
)

var (
	noSchemasErr  string = `No kubelt resources found`
	validationErr string = `
%s

%s
`
)

var (
	KindPath       cue.Path = cue.ParsePath("kind")
	ApiVersionPath cue.Path = cue.ParsePath("apiVersion")
)

// Cue API
//-----------------------------------------------------------------------------------------------

// CueAPI
type CueAPI interface {
	CompileSchemaFromString(apply string) (cue.Value, error)
	GenerateCueSpec(schema string, properties map[string]interface{}) (cue.Value, error)
	ValidateResource(def string, val cue.Value) (cue.Value, error)
}

// cueAPI
type cueAPI struct {
	context         *cue.Context
	definitions     cue.Value
	schemasVersions *map[string]cue.Value // singleton for faster processing during "apply"
	schemaFetcher   func(apiVersion string) ([]byte, error)
}

func (api cueAPI) unifyKmdrModule(input cue.Value) cue.Value {
	return api.definitions.Unify(input)
}

// NewCueAPI returns a new CueAPI
func newCueAPI() (CueAPI, error) {
	// Load definitions from the embedded FS
	overlay := make(map[string]load.Source)
	err := fs.WalkDir(StaticFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d == nil || d.Type().IsDir() || !d.Type().IsRegular() {
			return nil
		}

		if strings.Contains(path, ".cue") {
			f, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			overlay[filepath.Join("/", path)] = load.FromBytes(f)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	config := &load.Config{
		Dir:     "/",
		Overlay: overlay,
	}
	instances := load.Instances([]string{"/static/cue/kmdr"}, config)

	cueContext := cuecontext.New()
	defs := cueContext.BuildInstance(instances[0])

	mustFind := func(v cue.Value) (cue.Value, error) {
		if err = v.Err(); err != nil {
			return v, err
		}
		return v, nil
	}

	c := cueAPI{
		context:     cueContext,
		definitions: defs,
	}

	// Validate the embedded FS load worked
	_, err = mustFind(defs.LookupPath(cue.ParsePath("#manifests")))
	if err != nil {
		return c, err
	}
	return c, nil
}

// CompileShemaFromString accepts a string containg cue then builds and returns a  cue.Value
func (api cueAPI) CompileSchemaFromString(apply string) (cue.Value, error) {
	applySchemas := api.context.CompileString(apply)
	if applySchemas.Err() != nil {
		return applySchemas, errors.New(noSchemasErr)
	}
	return applySchemas, nil
}

// ValidateResource accepts a cue value and validates it against a cue definition
// for the "kind" of the cue value
func (api cueAPI) ValidateResource(def string, val cue.Value) (cue.Value, error) {
	pre := val.Eval()
	definition := api.definitions.LookupPath(cue.ParsePath(def))
	val = val.Unify(definition)
	if err := val.Validate(); err != nil {
		return pre, errors.New(validationErr)
	}
	return val, nil
}

// GenerateCueSpec injects concrete values and into cue value and validates the results before returning the concrete cue value.
func (api cueAPI) GenerateCueSpec(schema string, properties map[string]interface{}) (cue.Value, error) {
	// lookup the schema type from all the available schemas in the cue.Instance
	specSchema := api.definitions.LookupPath(cue.ParsePath(schema))
	if !specSchema.Exists() {
		return cue.Value{}, fmt.Errorf("%s is not a valid schema", schema)
	}

	// set the concrete values in the schema using the paths and values
	for path, value := range properties {
		cuePath := cue.ParsePath(path)
		specSchema = specSchema.FillPath(cuePath, value)
	}

	// return and validate the schema
	return specSchema, specSchema.Validate()
}
