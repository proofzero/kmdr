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

// ConfigAPI
type ConfigAPI interface {
	InitConfig() error
	AddContext(context string, isDefault ...bool) error
	RemoveContext(context string) error
	SetDefaultContext(context string) error
	AddUser(user string, isDefault ...bool) error
	RemoveUser(user string) error
	SetDefaultUser(user string) error
}

// configAPI for managing ktrl configs
type configAPI struct{}

// NewConfigAPI returns a new ConfigAPI
func newConfigAPI() (ConfigAPI, error) {
	c := configAPI{}
	return c, nil
}

// initConfig bootstraps the ktrl config
func (c configAPI) InitConfig() error {
	return nil
}

// AddContext adds a context to the config
func (c configAPI) AddContext(context string, isDefault ...bool) error {
	return nil
}

// RemoveContext adds a context to the config
func (c configAPI) RemoveContext(context string) error {
	return nil
}

// SetDefaultContext adds a context to the config
func (c configAPI) SetDefaultContext(context string) error {
	return nil
}

// AddUser adds a context to the config
func (c configAPI) AddUser(user string, isDefault ...bool) error {
	return nil
}

// RemoveUser adds a context to the config
func (c configAPI) RemoveUser(user string) error {
	return nil
}

// SetDefaultUser adds a context to the config
func (c configAPI) SetDefaultUser(user string) error {
	return nil
}
