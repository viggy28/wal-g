// Code generated by MockGen. DO NOT EDIT.
// Source: internal/stream_fetch_helper.go

// Package mock_internal is a generated GoMock package.
package testtools

import (
	"github.com/golang/mock/gomock"
	"reflect"
)

// MockLogFetchSettings is a mock of LogFetchSettings interface
type MockLogFetchSettings struct {
	ctrl     *gomock.Controller
	recorder *MockLogFetchSettingsMockRecorder
}

// MockLogFetchSettingsMockRecorder is the mock recorder for MockLogFetchSettings
type MockLogFetchSettingsMockRecorder struct {
	mock *MockLogFetchSettings
}

// NewMockLogFetchSettings creates a new mock instance
func NewMockLogFetchSettings(ctrl *gomock.Controller) *MockLogFetchSettings {
	mock := &MockLogFetchSettings{ctrl: ctrl}
	mock.recorder = &MockLogFetchSettingsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogFetchSettings) EXPECT() *MockLogFetchSettingsMockRecorder {
	return m.recorder
}

// GetEndTsEnv mocks base method
func (m *MockLogFetchSettings) GetEndTsEnv() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEndTsEnv")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEndTsEnv indicates an expected call of GetEndTsEnv
func (mr *MockLogFetchSettingsMockRecorder) GetEndTsEnv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEndTsEnv", reflect.TypeOf((*MockLogFetchSettings)(nil).GetEndTsEnv))
}

// GetDstEnv mocks base method
func (m *MockLogFetchSettings) GetDstEnv() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDstEnv")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDstEnv indicates an expected call of GetDstEnv
func (mr *MockLogFetchSettingsMockRecorder) GetDstEnv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDstEnv", reflect.TypeOf((*MockLogFetchSettings)(nil).GetDstEnv))
}

// GetLogFolderPath mocks base method
func (m *MockLogFetchSettings) GetLogFolderPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogFolderPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetLogFolderPath indicates an expected call of GetLogFolderPath
func (mr *MockLogFetchSettingsMockRecorder) GetLogFolderPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogFolderPath", reflect.TypeOf((*MockLogFetchSettings)(nil).GetLogFolderPath))
}