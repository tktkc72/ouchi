// Code generated by MockGen. DO NOT EDIT.
// Source: collector.go

// Package mock_ouchidashboard is a generated GoMock package.
package mock_ouchidashboard

import (
	gomock "github.com/golang/mock/gomock"
	natureremo "github.com/tenntenn/natureremo"
	reflect "reflect"
)

// MocknatureremoClient is a mock of natureremoClient interface
type MocknatureremoClient struct {
	ctrl     *gomock.Controller
	recorder *MocknatureremoClientMockRecorder
}

// MocknatureremoClientMockRecorder is the mock recorder for MocknatureremoClient
type MocknatureremoClientMockRecorder struct {
	mock *MocknatureremoClient
}

// NewMocknatureremoClient creates a new mock instance
func NewMocknatureremoClient(ctrl *gomock.Controller) *MocknatureremoClient {
	mock := &MocknatureremoClient{ctrl: ctrl}
	mock.recorder = &MocknatureremoClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MocknatureremoClient) EXPECT() *MocknatureremoClientMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MocknatureremoClient) GetAll() ([]*natureremo.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*natureremo.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MocknatureremoClientMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MocknatureremoClient)(nil).GetAll))
}
