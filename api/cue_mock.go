// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/cue.go

// Package mock_api is a generated GoMock package.
package api

import (
	reflect "reflect"

	cue "cuelang.org/go/cue"
	gomock "github.com/golang/mock/gomock"
)

// MockCueAPI is a mock of CueAPI interface.
type MockCueAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCueAPIMockRecorder
}

// MockCueAPIMockRecorder is the mock recorder for MockCueAPI.
type MockCueAPIMockRecorder struct {
	mock *MockCueAPI
}

// NewMockCueAPI creates a new mock instance.
func NewMockCueAPI(ctrl *gomock.Controller) *MockCueAPI {
	mock := &MockCueAPI{ctrl: ctrl}
	mock.recorder = &MockCueAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCueAPI) EXPECT() *MockCueAPIMockRecorder {
	return m.recorder
}

// CompileSchemaFromString mocks base method.
func (m *MockCueAPI) CompileSchemaFromString(apply string) (cue.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompileSchemaFromString", apply)
	ret0, _ := ret[0].(cue.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompileSchemaFromString indicates an expected call of CompileSchemaFromString.
func (mr *MockCueAPIMockRecorder) CompileSchemaFromString(apply interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompileSchemaFromString", reflect.TypeOf((*MockCueAPI)(nil).CompileSchemaFromString), apply)
}

// FetchSchema mocks base method.
func (m *MockCueAPI) FetchSchema(apiVersion string) (cue.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchSchema", apiVersion)
	ret0, _ := ret[0].(cue.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSchema indicates an expected call of FetchSchema.
func (mr *MockCueAPIMockRecorder) FetchSchema(apiVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSchema", reflect.TypeOf((*MockCueAPI)(nil).FetchSchema), apiVersion)
}

// GenerateCueSpec mocks base method.
func (m *MockCueAPI) GenerateCueSpec(schema string, properties map[string]string, val cue.Value) (cue.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateCueSpec", schema, properties, val)
	ret0, _ := ret[0].(cue.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateCueSpec indicates an expected call of GenerateCueSpec.
func (mr *MockCueAPIMockRecorder) GenerateCueSpec(schema, properties, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCueSpec", reflect.TypeOf((*MockCueAPI)(nil).GenerateCueSpec), schema, properties, val)
}

// ValidateResource mocks base method.
func (m *MockCueAPI) ValidateResource(val, def cue.Value) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateResource", val, def)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateResource indicates an expected call of ValidateResource.
func (mr *MockCueAPIMockRecorder) ValidateResource(val, def interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateResource", reflect.TypeOf((*MockCueAPI)(nil).ValidateResource), val, def)
}

// fetchSchema mocks base method.
func (m *MockCueAPI) fetchSchema(apiVersion string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "fetchSchema", apiVersion)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// fetchSchema indicates an expected call of fetchSchema.
func (mr *MockCueAPIMockRecorder) fetchSchema(apiVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "fetchSchema", reflect.TypeOf((*MockCueAPI)(nil).fetchSchema), apiVersion)
}
