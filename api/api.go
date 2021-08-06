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
	"errors"
	"fmt"

	"cuelang.org/go/cue"
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

	authapi, _ := newAuthApi()
	configapi, _ := newConfigAPI()
	cueapi, _ := newCueAPI()
	ktrlapi, _ := newKtrlAPI()
	planapi, _ := newPlanAPI()

	kmdrapi = &kmdrAPI{
		auth:   authapi,
		config: configapi,
		cue:    cueapi,
		ktrl:   ktrlapi,
		plan:   planapi,
	}
	return *kmdrapi, nil
}

// KmdrApi
type KmdrAPI interface {
	SetupUser(string) error
	Apply(string) error
}

// kmdrAPI
type kmdrAPI struct {
	auth   AuthAPI
	config ConfigAPI
	cue    CueAPI
	ktrl   KtrlAPI
	plan   PlanAPI
}

func (api kmdrAPI) SetupUser(username string) error {
	if err := api.auth.AddKeys(username); err != nil {
		return err
	}

	_ = api.config.InitConfig()
	_ = api.config.AddContext("default", true)
	_ = api.config.AddUser(username, true)
	if err := api.config.Commit(); err != nil {
		return err
	}

	user := make(map[string]interface{})
	user["data.name"] = username
	user["data.publicEncyrptionKey"] = api.auth.EncryptionKey()
	user["data.publicSignatureKey"] = api.auth.AuthKey()
	if userVal, err := api.cue.GenerateCueSpec("#user", user); err != nil {
		return err
	} else {
		_, err = api.ktrl.Apply([]interface{}{userVal})
		return err
	}
}

func (api kmdrAPI) Apply(applyStr string) error {
	// convert the input to a cue value
	applySchemas, err := api.cue.CompileSchemaFromString(applyStr)
	if err != nil {
		return err
	}

	// validate that the input is correct
	manifests := applySchemas.Value().LookupPath(cue.ParsePath("manifests"))
	if val, err := api.cue.ValidateResource("#manifests", manifests); err != nil {
		return errors.New(fmt.Sprintf(err.Error(), val, manifests.Eval()))
	}

	// query for the worls as it relates to the manifests input
	res, err := api.ktrl.Query(manifests)
	if err != nil {
		return err
	}

	// Validate and plan and return the commands
	plan, err := api.plan.Plan(res)
	if err != nil {
		return err
	}

	// Sign the changes
	for _, cmd := range plan {
		api.auth.SignNode(cmd.([]byte))
	}

	// Apply the changes
	resp, err := api.ktrl.Apply(plan)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}
	for _, v := range resp.Resources.Cue {
		fmt.Println(v)
	}

	return nil
}
