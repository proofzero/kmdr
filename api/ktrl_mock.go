// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/ktrl.go

// Package mock_api is a generated GoMock package.
package api

import (
	reflect "reflect"

	cue "cuelang.org/go/cue"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/proofzero/proto/pkg/v1alpha1"
)

// MockKtrlAPI is a mock of KtrlAPI interface.
type MockKtrlAPI struct {
	ctrl     *gomock.Controller
	recorder *MockKtrlAPIMockRecorder
}

// MockKtrlAPIMockRecorder is the mock recorder for MockKtrlAPI.
type MockKtrlAPIMockRecorder struct {
	mock *MockKtrlAPI
}

// NewMockKtrlAPI creates a new mock instance.
func NewMockKtrlAPI(ctrl *gomock.Controller) *MockKtrlAPI {
	mock := &MockKtrlAPI{ctrl: ctrl}
	mock.recorder = &MockKtrlAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKtrlAPI) EXPECT() *MockKtrlAPIMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockKtrlAPI) Apply(cueValue cue.Value) (*v1alpha1.ApplyDefault, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", cueValue)
	ret0, _ := ret[0].(*v1alpha1.ApplyDefault)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockKtrlAPIMockRecorder) Apply(cueValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockKtrlAPI)(nil).Apply), cueValue)
}

// InitConfig mocks base method.
func (m *MockKtrlAPI) InitConfig() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitConfig")
	ret0, _ := ret[0].(error)
	return ret0
}

// InitConfig indicates an expected call of InitConfig.
func (mr *MockKtrlAPIMockRecorder) InitConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitConfig", reflect.TypeOf((*MockKtrlAPI)(nil).InitConfig))
}

// IsAvailable mocks base method.
func (m *MockKtrlAPI) IsAvailable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAvailable")
	ret0, _ := ret[0].(error)
	return ret0
}

// IsAvailable indicates an expected call of IsAvailable.
func (mr *MockKtrlAPIMockRecorder) IsAvailable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAvailable", reflect.TypeOf((*MockKtrlAPI)(nil).IsAvailable))
}