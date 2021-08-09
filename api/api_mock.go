// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/api.go

// Package mock_api is a generated GoMock package.
package api

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockKmdrAPI is a mock of KmdrAPI interface.
type MockKmdrAPI struct {
	ctrl     *gomock.Controller
	recorder *MockKmdrAPIMockRecorder
}

// MockKmdrAPIMockRecorder is the mock recorder for MockKmdrAPI.
type MockKmdrAPIMockRecorder struct {
	mock *MockKmdrAPI
}

// NewMockKmdrAPI creates a new mock instance.
func NewMockKmdrAPI(ctrl *gomock.Controller) *MockKmdrAPI {
	mock := &MockKmdrAPI{ctrl: ctrl}
	mock.recorder = &MockKmdrAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKmdrAPI) EXPECT() *MockKmdrAPIMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockKmdrAPI) Apply(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockKmdrAPIMockRecorder) Apply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockKmdrAPI)(nil).Apply), arg0)
}

// SetupUser mocks base method.
func (m *MockKmdrAPI) SetupUser(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetupUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupUser indicates an expected call of SetupUser.
func (mr *MockKmdrAPIMockRecorder) SetupUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupUser", reflect.TypeOf((*MockKmdrAPI)(nil).SetupUser), arg0)
}