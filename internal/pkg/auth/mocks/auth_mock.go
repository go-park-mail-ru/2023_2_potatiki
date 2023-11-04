// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	models "github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

// MockAuthUsecase is a mock of AuthUsecase interface.
type MockAuthUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockAuthUsecaseMockRecorder
}

// MockAuthUsecaseMockRecorder is the mock recorder for MockAuthUsecase.
type MockAuthUsecaseMockRecorder struct {
	mock *MockAuthUsecase
}

// NewMockAuthUsecase creates a new mock instance.
func NewMockAuthUsecase(ctrl *gomock.Controller) *MockAuthUsecase {
	mock := &MockAuthUsecase{ctrl: ctrl}
	mock.recorder = &MockAuthUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthUsecase) EXPECT() *MockAuthUsecaseMockRecorder {
	return m.recorder
}

// CheckAuth mocks base method.
func (m *MockAuthUsecase) CheckAuth(arg0 context.Context, arg1 uuid.UUID) (*models.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAuth", arg0, arg1)
	ret0, _ := ret[0].(*models.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAuth indicates an expected call of CheckAuth.
func (mr *MockAuthUsecaseMockRecorder) CheckAuth(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAuth", reflect.TypeOf((*MockAuthUsecase)(nil).CheckAuth), arg0, arg1)
}

// SignIn mocks base method.
func (m *MockAuthUsecase) SignIn(arg0 context.Context, arg1 *models.SignInPayload) (*models.Profile, string, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", arg0, arg1)
	ret0, _ := ret[0].(*models.Profile)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(time.Time)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// SignIn indicates an expected call of SignIn.
func (mr *MockAuthUsecaseMockRecorder) SignIn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockAuthUsecase)(nil).SignIn), arg0, arg1)
}

// SignUp mocks base method.
func (m *MockAuthUsecase) SignUp(arg0 context.Context, arg1 *models.SignUpPayload) (*models.Profile, string, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", arg0, arg1)
	ret0, _ := ret[0].(*models.Profile)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(time.Time)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// SignUp indicates an expected call of SignUp.
func (mr *MockAuthUsecaseMockRecorder) SignUp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockAuthUsecase)(nil).SignUp), arg0, arg1)
}
