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

	authapi := newAuthApi()
	configapi := newConfigAPI()
	// TODO: check for errors during api init
	cueapi, _ := newCueAPI()
	ktrlapi, _ := newKtrlAPI()
	planapi, _ := newPlanAPI(&authapi)

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
	SetupUser(string, string) error
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

func (api kmdrAPI) SetupUser(username string, email string) error {
	// TODO: prompt for passphrase
	// How would you do this for NACL keys?

	// request a passphrase when generating keys. Empty == no passphrase
	if err := api.auth.GenerateKeys(username); err != nil {
		return err
	}

	cmd := api.plan.User(username, email)
	user, err := api.ktrl.RegisterUser(cmd)
	if err != nil {
		return err
	}

	// Add the new user to the config
	if err := api.config.InitConfig(); err != nil {
		return err
	}
	if err := api.config.AddUser(user.Header.ID, username, email, true); err != nil {
		return err
	}

	// Commit the configuration changes
	if err := api.auth.CommitKeys(); err != nil {
		return err
	}
	if err := api.config.CommitConfig(); err != nil {
		return err
	}

	return nil
}

func (api kmdrAPI) Apply(applyStr string) error {
	// Load config context
	if err := api.config.InitConfig(); err != nil {
		return err
	}

	// Load keys into context
	if user, err := api.config.GetCurrentUser(); err != nil {
		return err
	} else if err := api.auth.LoadKeys(user.ID); err != nil {
		return err
	}

	// convert the input to a cue value
	applySchemas, err := api.cue.CompileSchemaFromString(applyStr)
	if err != nil {
		return err
	}

	// validate that the input is correct
	manifestsCue := applySchemas.Value().LookupPath(cue.ParsePath("manifests"))
	if val, err := api.cue.ValidateResource("#manifests", manifestsCue); err != nil {
		return fmt.Errorf(fmt.Sprintf(err.Error(), val, manifestsCue.Eval()))
	}

	// query for the worls as it relates to the manifests input
	manifests, err := api.plan.Lookup(&manifestsCue)
	if err != nil {
		return err
	}
	nodes, err := api.ktrl.Query(manifests)
	if err != nil {
		return err
	}

	// Validate and plan and return the commands
	plan, err := api.plan.Plan(nodes)
	if err != nil {
		return err
	}

	// Apply the changes
	_, err = api.ktrl.Apply(plan)
	if err != nil {
		return err
	}

	return nil
}
