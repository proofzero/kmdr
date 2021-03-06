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
	"path"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
)

// ConfigAPI
type ConfigAPI interface {
	InitConfig() error
	GetCurrentUser() (*UserConfig, error)
	AddUser(id string, username string, email string, isDefault ...bool) error
	RemoveUser(user string) error
	SetDefaultUser(user string) error
	CommitConfig() error
}

// NOTE: We should consider doing an upstream to cobra/viper to support the use of CUElang
// To avoid using toml
// NOTE:PR has been open to Viper for the above note

// configAPI for managing ktrl configs
type configAPI struct {
	configDir   string
	CurrentUser string                `toml:"current_user"`
	Ktrl        ktrlConfig            `toml:"ktrl"`
	Users       map[string]UserConfig `toml:"users"`
}

type UserConfig struct {
	ID       string `toml:"id"`
	Username string `toml:"username"`
	Email    string `toml:"email"`
}

// NewConfigAPI returns a new ConfigAPI
func newConfigAPI() ConfigAPI {
	configDir := path.Join(xdg.ConfigHome, "kubelt")
	c := &configAPI{
		configDir:   configDir,
		CurrentUser: "",
		Users:       make(map[string]UserConfig),
		Ktrl: ktrlConfig{
			Server: serverConfig{
				Protocol: "tcp",
				Port:     ":50051",
			},
		},
	}
	return c
}

// initConfig bootstraps the ktrl config
func (c *configAPI) InitConfig() error {
	if configPath, err := xdg.ConfigFile("kubelt/config.toml"); err != nil {
		return err
	} else {
		// check if the config already exists
		if _, err := os.Stat(configPath); !errors.Is(err, fs.ErrNotExist) {
			// if it does, bootstrap the config struct
			f, _ := ioutil.ReadFile(configPath)
			if _, err := toml.Decode(string(f), &c); err != nil {
				return err
			}
		}
		return nil
	}
}

func (c *configAPI) GetCurrentUser() (*UserConfig, error) {
	if c.CurrentUser == "" {
		return nil, fmt.Errorf("No user set.")
	}
	user := c.Users[c.CurrentUser]
	return &user, nil
}

// AddUser adds a context to the config
func (c *configAPI) AddUser(id string, username string, email string, isDefault ...bool) error {
	// TODO: check if user already exists
	// If the user already exist return an error

	c.Users[username] = UserConfig{
		ID:       id,
		Username: username,
		Email:    email,
	}
	if len(isDefault) > 0 && isDefault[0] {
		c.CurrentUser = username
	}
	return nil
}

// RemoveUser adds a context to the config
func (c *configAPI) RemoveUser(user string) error {
	if _, ok := c.Users[user]; !ok {
		return fmt.Errorf("No user with name %s exist", user)
	}
	delete(c.Users, user)
	return nil
}

// SetDefaultUser adds a context to the config
func (c *configAPI) SetDefaultUser(user string) error {
	c.CurrentUser = user
	return nil
}

// Commit write the config to disk
func (c *configAPI) CommitConfig() error {
	configPath := fmt.Sprintf("%s/config.toml", c.configDir)
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
