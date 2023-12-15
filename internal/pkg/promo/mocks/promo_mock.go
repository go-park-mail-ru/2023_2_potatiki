// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockPromoUsecase is a mock of PromoUsecase interface.
type MockPromoUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockPromoUsecaseMockRecorder
}

// MockPromoUsecaseMockRecorder is the mock recorder for MockPromoUsecase.
type MockPromoUsecaseMockRecorder struct {
	mock *MockPromoUsecase
}

// NewMockPromoUsecase creates a new mock instance.
func NewMockPromoUsecase(ctrl *gomock.Controller) *MockPromoUsecase {
	mock := &MockPromoUsecase{ctrl: ctrl}
	mock.recorder = &MockPromoUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPromoUsecase) EXPECT() *MockPromoUsecaseMockRecorder {
	return m.recorder
}

// CheckPromocode mocks base method.
func (m *MockPromoUsecase) CheckPromocode(arg0 context.Context, arg1 string) (*models.Promocode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPromocode", arg0, arg1)
	ret0, _ := ret[0].(*models.Promocode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPromocode indicates an expected call of CheckPromocode.
func (mr *MockPromoUsecaseMockRecorder) CheckPromocode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPromocode", reflect.TypeOf((*MockPromoUsecase)(nil).CheckPromocode), arg0, arg1)
}

// MockPromoRepo is a mock of PromoRepo interface.
type MockPromoRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPromoRepoMockRecorder
}

// MockPromoRepoMockRecorder is the mock recorder for MockPromoRepo.
type MockPromoRepoMockRecorder struct {
	mock *MockPromoRepo
}

// NewMockPromoRepo creates a new mock instance.
func NewMockPromoRepo(ctrl *gomock.Controller) *MockPromoRepo {
	mock := &MockPromoRepo{ctrl: ctrl}
	mock.recorder = &MockPromoRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPromoRepo) EXPECT() *MockPromoRepoMockRecorder {
	return m.recorder
}

// ReadPromocode mocks base method.
func (m *MockPromoRepo) ReadPromocode(arg0 context.Context, arg1 string) (*models.Promocode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadPromocode", arg0, arg1)
	ret0, _ := ret[0].(*models.Promocode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadPromocode indicates an expected call of ReadPromocode.
func (mr *MockPromoRepoMockRecorder) ReadPromocode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadPromocode", reflect.TypeOf((*MockPromoRepo)(nil).ReadPromocode), arg0, arg1)
}