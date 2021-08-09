// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/ktrl.go

// Package mock_api is a generated GoMock package.
package api

import (
	reflect "reflect"

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
func (m *MockKtrlAPI) Apply(arg0 []interface{}) (*v1alpha1.ApplyDefault, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0)
	ret0, _ := ret[0].(*v1alpha1.ApplyDefault)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockKtrlAPIMockRecorder) Apply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockKtrlAPI)(nil).Apply), arg0)
}

// Query mocks base method.
func (m *MockKtrlAPI) Query(arg0 interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockKtrlAPIMockRecorder) Query(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockKtrlAPI)(nil).Query), arg0)
}

// initConfig mocks base method.
func (m *MockKtrlAPI) initConfig() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "initConfig")
	ret0, _ := ret[0].(error)
	return ret0
}

// initConfig indicates an expected call of initConfig.
func (mr *MockKtrlAPIMockRecorder) initConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "initConfig", reflect.TypeOf((*MockKtrlAPI)(nil).initConfig))
}

// isAvailable mocks base method.
func (m *MockKtrlAPI) isAvailable() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "isAvailable")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// isAvailable indicates an expected call of isAvailable.
func (mr *MockKtrlAPIMockRecorder) isAvailable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isAvailable", reflect.TypeOf((*MockKtrlAPI)(nil).isAvailable))
}
