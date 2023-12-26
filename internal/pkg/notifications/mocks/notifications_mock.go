// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

// MockNotificationsUsecase is a mock of NotificationsUsecase interface.
type MockNotificationsUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationsUsecaseMockRecorder
}

// MockNotificationsUsecaseMockRecorder is the mock recorder for MockNotificationsUsecase.
type MockNotificationsUsecaseMockRecorder struct {
	mock *MockNotificationsUsecase
}

// NewMockNotificationsUsecase creates a new mock instance.
func NewMockNotificationsUsecase(ctrl *gomock.Controller) *MockNotificationsUsecase {
	mock := &MockNotificationsUsecase{ctrl: ctrl}
	mock.recorder = &MockNotificationsUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationsUsecase) EXPECT() *MockNotificationsUsecaseMockRecorder {
	return m.recorder
}

// GetDayNotifications mocks base method.
func (m *MockNotificationsUsecase) GetDayNotifications(arg0 context.Context, arg1 uuid.UUID) ([]models.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDayNotifications", arg0, arg1)
	ret0, _ := ret[0].([]models.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDayNotifications indicates an expected call of GetDayNotifications.
func (mr *MockNotificationsUsecaseMockRecorder) GetDayNotifications(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDayNotifications", reflect.TypeOf((*MockNotificationsUsecase)(nil).GetDayNotifications), arg0, arg1)
}

// MockNotificationsRepo is a mock of NotificationsRepo interface.
type MockNotificationsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationsRepoMockRecorder
}

// MockNotificationsRepoMockRecorder is the mock recorder for MockNotificationsRepo.
type MockNotificationsRepoMockRecorder struct {
	mock *MockNotificationsRepo
}

// NewMockNotificationsRepo creates a new mock instance.
func NewMockNotificationsRepo(ctrl *gomock.Controller) *MockNotificationsRepo {
	mock := &MockNotificationsRepo{ctrl: ctrl}
	mock.recorder = &MockNotificationsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationsRepo) EXPECT() *MockNotificationsRepoMockRecorder {
	return m.recorder
}

// ReadDayNotifications mocks base method.
func (m *MockNotificationsRepo) ReadDayNotifications(arg0 context.Context, arg1 uuid.UUID) ([]models.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDayNotifications", arg0, arg1)
	ret0, _ := ret[0].([]models.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDayNotifications indicates an expected call of ReadDayNotifications.
func (mr *MockNotificationsRepoMockRecorder) ReadDayNotifications(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDayNotifications", reflect.TypeOf((*MockNotificationsRepo)(nil).ReadDayNotifications), arg0, arg1)
}