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
package ktrl

import (
	"log"
	"path"

	"github.com/adrg/xdg"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	kb "github.com/proofzero/proto/pkg/v1alpha1"
)

// Client for managing the ktrl grpc service
type Client struct {
	Connection *grpc.ClientConn
	Client     kb.KubeltClient
	Config     *Config
}

// Config
type Config struct {
	Ktrl           KtrlConfig             `toml:"ktrl"`
	CurrentContext string                 `toml:"current_context"`
	Contexts       map[string]*kb.Context `toml:"contexts"`
}

// KtrlConfig
type KtrlConfig struct {
	Server ServerConfig
}

// ServerConfig
type ServerConfig struct {
	Port     string `toml:"port"`
	Protocol string `toml:"protocol"`
}

var (
	ktrlClient *Client
	ktrlConfig *Config
)

// init reads in configurations for the kubelt config directory to setup a ktrlClient
func init() {
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
			log.Fatal("missing config file:", err)
		} else {
			// Config file was found but another error was produced.
			log.Fatal("error loading configuration:", err)
		}
	}

	err := viper.Unmarshal(&ktrlConfig)
	if err != nil {
		log.Fatal("could not unmarshal config")
	}
}

// NewKtrlClient returns a new Client
func NewKtrlClient(options ...grpc.DialOption) (*Client, error) {
	if ktrlClient != nil {
		return ktrlClient, nil
	}
	if len(options) == 0 {
		options = []grpc.DialOption{
			grpc.WithInsecure(),
		}
	}

	conn, err := grpc.Dial(ktrlConfig.Ktrl.Server.Port, options...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	client := kb.NewKubeltClient(conn)

	ktrlClient = &Client{
		Connection: conn,
		Client:     client,
		Config:     ktrlConfig,
	}
	return ktrlClient, err
}
