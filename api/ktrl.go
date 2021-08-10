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
	"fmt"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	kb "github.com/proofzero/proto/pkg/v1alpha1"
)

// TODO: implement query
// Once we've defiend the query interfaces we need to implement this stub

// TODO: implement apply query
// Once we've defiend the apply interface we need to update the apply implementation

// KtrlAPI
type KtrlAPI interface {
	isAvailable() (bool, error)
	initConfig() error
	Query(interface{}) (interface{}, error)
	Apply([]interface{}) (*kb.ApplyResponse, error)
}

// ktrlAPI for managing the ktrl grpc service
type ktrlAPI struct {
	Connection *grpc.ClientConn
	Client     kb.KubeltClient
	Config     ktrlConfig
}

// ktrlConfig
type ktrlConfig struct {
	Server serverConfig `toml:"server"`
}

// serverConfig
type serverConfig struct {
	Port     string `toml:"port"`
	Protocol string `toml:"protocol"`
}

// NewKtrlAPI returns a new Client
func newKtrlAPI(options ...grpc.DialOption) (KtrlAPI, error) {
	ktrl := &ktrlAPI{
		Config: ktrlConfig{
			Server: serverConfig{},
		},
	}
	err := ktrl.initConfig()
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
func (ktrl *ktrlAPI) initConfig() error {
	err := viper.Unmarshal(&ktrl.Config)
	if err != nil {
		return fmt.Errorf("could not unmarshal config: %s", err)
	}

	return nil
}

// IsAvailable checks if the ktrl daemon is installed and running
func (ktrl *ktrlAPI) isAvailable() (bool, error) {
	request := &kb.HealthCheckRequest{}
	if res, err := ktrl.Client.HealthCheck(context.Background(), request); err != nil {
		return false, fmt.Errorf("KTRL is not running.")
	} else {
		return res.Error == nil, nil
	}
}

// Run a query against the data grid
func (ktrl *ktrlAPI) Query(q interface{}) (interface{}, error) {
	if ok, err := ktrl.isAvailable(); !ok {
		return nil, err
	}
	return nil, nil
}

// Apply calls out to ktrl to mutate the kubelt graph using values supplied by the user
func (ktrl *ktrlAPI) Apply([]interface{}) (*kb.ApplyResponse, error) {
	if ok, err := ktrl.isAvailable(); !ok {
		return nil, err
	}
	return nil, nil
}
