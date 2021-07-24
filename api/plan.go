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

// Interface
// -----------------------------------------------------------------------------

// PlanAPI
type PlanAPI interface {
	Plan(interface{}) ([]interface{}, error) // TODO: fix signature
}

// TODO: methods to support multi format
// TODO: define types for the 32 and 64 byte lengths and update signatures

// Implementation
// -----------------------------------------------------------------------------

// Plan

type plan struct{}

func newPlanAPI() (PlanAPI, error) {
	return &plan{}, nil
}

func (p *plan) Plan(graph interface{}) ([]interface{}, error) {
	return nil, nil
}
