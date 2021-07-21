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
	"html/template"
	"io/fs"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
	kb "github.com/proofzero/proto/pkg/v1alpha1"
)

// ConfigAPI
type ConfigAPI interface {
	InitConfig() (ConfigAPI, error)
	AddContext(context string, isDefault ...bool) (ConfigAPI, error)
	RemoveContext(context string) (ConfigAPI, error)
	SetDefaultContext(context string) (ConfigAPI, error)
	AddUser(user string, isDefault ...bool) (ConfigAPI, error)
	RemoveUser(user string) (ConfigAPI, error)
	SetDefaultUser(user string) (ConfigAPI, error)
	Commit() error
}

// configAPI for managing ktrl configs
type configAPI struct {
	CurrentUser string                       `toml:"current_user"`
	Ktrl        ktrlConfig                   `toml:"ktrl"`
	Users       map[string]map[string]string `toml:"users"` // TODO: User struct?
}

// NewConfigAPI returns a new ConfigAPI
func newConfigAPI() (ConfigAPI, error) {
	c := &configAPI{
		Ktrl:  ktrlConfig{},
		Users: make(map[string]map[string]string),
	}
	return c, nil
}

// initConfig bootstraps the ktrl config
func (c configAPI) InitConfig() (ConfigAPI, error) {
	home, _ := homedir.Dir()
	configPath := fmt.Sprintf("%s/.config/kubelt/kmdr.toml", home)
	// check if the config already exists
	if _, err := os.Stat(configPath); !errors.Is(err, fs.ErrNotExist) {
		// if it does, bootstrap the config struct
		f, _ := ioutil.ReadFile(configPath)
		if _, err := toml.Decode(string(f), &c); err != nil {
			return c, err
		}
		fmt.Println("before")
		fmt.Println(c)
	}
	return c, nil
}

// AddContext adds a context t o the config
func (c configAPI) AddContext(context string, isDefault ...bool) (ConfigAPI, error) {
	if c.Ktrl.Contexts == nil {
		c.Ktrl.Contexts = make(map[string]*kb.Context)
	}
	// TODO: check if context already exists
	c.Ktrl.Contexts[context] = &kb.Context{Name: context}
	c.Ktrl.Contexts[context].Name = context
	if len(isDefault) > 0 && isDefault[0] {
		c.Ktrl.CurrentContext = context
	}
	return c, nil
}

// RemoveContext adds a context to the config
func (c configAPI) RemoveContext(context string) (ConfigAPI, error) {
	return c, nil
}

// SetDefaultContext adds a context to the config
func (c configAPI) SetDefaultContext(context string) (ConfigAPI, error) {
	return c, nil
}

// AddUser adds a context to the config
func (c configAPI) AddUser(user string, isDefault ...bool) (ConfigAPI, error) {
	// TODO: check if user already exists
	c.Users[user] = make(map[string]string)
	c.Users[user]["Name"] = user
	if len(isDefault) > 0 && isDefault[0] {
		c.CurrentUser = user
	}
	return c, nil
}

// RemoveUser adds a context to the config
func (c configAPI) RemoveUser(user string) (ConfigAPI, error) {
	return c, nil
}

// SetDefaultUser adds a context to the config
func (c configAPI) SetDefaultUser(user string) (ConfigAPI, error) {
	return c, nil
}

// Commit write the config to disk
func (c configAPI) Commit() error {
	fmt.Println("after")
	fmt.Println(c)
	home, _ := homedir.Dir()
	configPath := fmt.Sprintf("%s/.config/kubelt/kmdr.toml", home)
	t := template.Must(template.ParseFS(StaticFS, "static/templates/kmdr.gotmpl"))
	f, err := os.Create(configPath)
	if err != nil {
		return err
	}

	err = t.Execute(f, c)
	if err != nil {
		return err
	}
	return nil
}
