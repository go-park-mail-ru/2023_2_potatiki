// Code generated by MockGen. DO NOT EDIT.
// Source: hub.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	websocket "github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// MockHubInterface is a mock of HubInterface interface.
type MockHubInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHubInterfaceMockRecorder
}

// MockHubInterfaceMockRecorder is the mock recorder for MockHubInterface.
type MockHubInterfaceMockRecorder struct {
	mock *MockHubInterface
}

// NewMockHubInterface creates a new mock instance.
func NewMockHubInterface(ctrl *gomock.Controller) *MockHubInterface {
	mock := &MockHubInterface{ctrl: ctrl}
	mock.recorder = &MockHubInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHubInterface) EXPECT() *MockHubInterfaceMockRecorder {
	return m.recorder
}

// AddClient mocks base method.
func (m *MockHubInterface) AddClient(arg0 uuid.UUID, arg1 *websocket.Conn) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddClient", arg0, arg1)
}

// AddClient indicates an expected call of AddClient.
func (mr *MockHubInterfaceMockRecorder) AddClient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddClient", reflect.TypeOf((*MockHubInterface)(nil).AddClient), arg0, arg1)
}

// Run mocks base method.
func (m *MockHubInterface) Run(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run", arg0)
}

// Run indicates an expected call of Run.
func (mr *MockHubInterfaceMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockHubInterface)(nil).Run), arg0)
}
