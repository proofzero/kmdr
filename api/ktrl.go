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
	"context"
	"errors"
	"fmt"
	"os/exec"
	"path"

	"cuelang.org/go/cue"
	"github.com/adrg/xdg"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	kb "github.com/proofzero/proto/pkg/v1alpha1"
)

// KtrlAPI
type KtrlAPI interface {
	InitConfig() error
	IsAvailable() error
	Apply(cueValue cue.Value) (*kb.ApplyDefault, error)
}

// ktrlAPI for managing the ktrl grpc service
type ktrlAPI struct {
	Connection *grpc.ClientConn
	Client     kb.KubeltClient
	Config     ktrlConfig
}

// ktrlConfig
type ktrlConfig struct {
	Server         serverConfig           `toml:"server"`
	CurrentContext string                 `toml:"current_context"`
	Contexts       map[string]*kb.Context `toml:"contexts"`
}

// serverConfig
type serverConfig struct {
	Port     string `toml:"port"`
	Protocol string `toml:"protocol"`
}

// NewKtrlAPI returns a new Client
func newKtrlAPI(options ...grpc.DialOption) (KtrlAPI, error) {
	ktrl := ktrlAPI{
		Config: ktrlConfig{
			Server: serverConfig{},
		},
	}
	err := ktrl.InitConfig()
	if err != nil {
		return ktrl, err
	}

	if len(options) == 0 {
		options = []grpc.DialOption{
			grpc.WithInsecure(),
		}
	}

	conn, err := grpc.Dial(ktrl.Config.Server.Port, options...)
	if err != nil {
		return ktrl, fmt.Errorf("fail to dial: %v", err)
	}
	ktrl.Connection = conn

	ktrl.Client = kb.NewKubeltClient(conn)

	return ktrl, err
}

// init reads in configurations for the kubelt config directory to setup a ktrlClient
func (ktrl ktrlAPI) InitConfig() error {
	parentName := "kubelt"
	fileName := "config"
	configType := "toml"

	configDir := path.Join(xdg.ConfigHome, parentName)

	// Will be uppercased automatically. Environment variables must
	// have this prefix to be treated as configuration sources.
	viper.SetEnvPrefix(fileName)

	viper.SetDefault("ktrl.server.protocol", "tcp")

	viper.SetDefault("ktrl.server.port", ":50051")

	// name of config file (without extension)
	viper.SetConfigName(fileName)
	// REQUIRED if the config file does not have the extension in the name.
	viper.SetConfigType(configType)
	// Path to look for the config file in. Call multiple times to
	// add many search paths.
	viper.AddConfigPath(configDir)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired.
			fmt.Printf("missing config file: %s", err)
		} else {
			// Config file was found but another error was produced.
			return fmt.Errorf("error loading configuration: %s", err)
		}
	}

	err := viper.Unmarshal(&ktrl.Config)
	if err != nil {
		return fmt.Errorf("could not unmarshal config: %s", err)
	}

	return nil
}

// IsAvailable checks if the ktrl daemon is installed and running
func (ktrl ktrlAPI) IsAvailable() error {
	if _, err := exec.LookPath("ktrl"); err != nil {
		return errors.New("ktrl is not installed and running")
	}
	// check if ktrl is running
	if !checkKtrlProcess() {
		return errors.New("ktrl is not running")
	}
	return nil
}

// Apply calls out to ktrl to mutate the kubelt graph using values supplied by the user
func (ktrl ktrlAPI) Apply(cueValue cue.Value) (*kb.ApplyDefault, error) {
	ctx := ktrl.Config.Contexts[ktrl.Config.CurrentContext]
	cueString := fmt.Sprint(cueValue)
	resource := &kb.Cue{
		Cue: cueString,
	}
	request := &kb.ApplyRequest{Resources: resource, Context: ctx}
	r, err := ktrl.Client.Apply(context.Background(), request)
	return r, err
}
