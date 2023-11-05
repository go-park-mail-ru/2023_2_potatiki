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

// MockAddressUsecase is a mock of AddressUsecase interface.
type MockAddressUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockAddressUsecaseMockRecorder
}

// MockAddressUsecaseMockRecorder is the mock recorder for MockAddressUsecase.
type MockAddressUsecaseMockRecorder struct {
	mock *MockAddressUsecase
}

// NewMockAddressUsecase creates a new mock instance.
func NewMockAddressUsecase(ctrl *gomock.Controller) *MockAddressUsecase {
	mock := &MockAddressUsecase{ctrl: ctrl}
	mock.recorder = &MockAddressUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddressUsecase) EXPECT() *MockAddressUsecaseMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockAddressUsecase) AddAddress(arg0 context.Context, arg1 uuid.UUID, arg2 models.AddressInfo) (models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockAddressUsecaseMockRecorder) AddAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockAddressUsecase)(nil).AddAddress), arg0, arg1, arg2)
}

// DeleteAddress mocks base method.
func (m *MockAddressUsecase) DeleteAddress(arg0 context.Context, arg1 models.AddressDelete) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockAddressUsecaseMockRecorder) DeleteAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockAddressUsecase)(nil).DeleteAddress), arg0, arg1)
}

// MakeCurrentAddress mocks base method.
func (m *MockAddressUsecase) MakeCurrentAddress(ctx context.Context, addressInfo models.AddressMakeCurrent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeCurrentAddress", ctx, addressInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeCurrentAddress indicates an expected call of MakeCurrentAddress.
func (mr *MockAddressUsecaseMockRecorder) MakeCurrentAddress(ctx, addressInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeCurrentAddress", reflect.TypeOf((*MockAddressUsecase)(nil).MakeCurrentAddress), ctx, addressInfo)
}

// UpdateAddress mocks base method.
func (m *MockAddressUsecase) UpdateAddress(arg0 context.Context, arg1 models.Address) (models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", arg0, arg1)
	ret0, _ := ret[0].(models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockAddressUsecaseMockRecorder) UpdateAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockAddressUsecase)(nil).UpdateAddress), arg0, arg1)
}

// MockAddressRepo is a mock of AddressRepo interface.
type MockAddressRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAddressRepoMockRecorder
}

// MockAddressRepoMockRecorder is the mock recorder for MockAddressRepo.
type MockAddressRepoMockRecorder struct {
	mock *MockAddressRepo
}

// NewMockAddressRepo creates a new mock instance.
func NewMockAddressRepo(ctrl *gomock.Controller) *MockAddressRepo {
	mock := &MockAddressRepo{ctrl: ctrl}
	mock.recorder = &MockAddressRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddressRepo) EXPECT() *MockAddressRepoMockRecorder {
	return m.recorder
}

// CreateAddress mocks base method.
func (m *MockAddressRepo) CreateAddress(arg0 context.Context, arg1 uuid.UUID, arg2 models.AddressInfo) (models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAddress indicates an expected call of CreateAddress.
func (mr *MockAddressRepoMockRecorder) CreateAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAddress", reflect.TypeOf((*MockAddressRepo)(nil).CreateAddress), arg0, arg1, arg2)
}

// DeleteAddress mocks base method.
func (m *MockAddressRepo) DeleteAddress(arg0 context.Context, arg1 models.AddressDelete) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockAddressRepoMockRecorder) DeleteAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockAddressRepo)(nil).DeleteAddress), arg0, arg1)
}

// GetAddress mocks base method.
func (m *MockAddressRepo) GetAddress(arg0 context.Context, arg1, arg2 uuid.UUID) (models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddress indicates an expected call of GetAddress.
func (mr *MockAddressRepoMockRecorder) GetAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddress", reflect.TypeOf((*MockAddressRepo)(nil).GetAddress), arg0, arg1, arg2)
}

// MakeCurrentAddress mocks base method.
func (m *MockAddressRepo) MakeCurrentAddress(arg0 context.Context, arg1 models.AddressMakeCurrent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeCurrentAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeCurrentAddress indicates an expected call of MakeCurrentAddress.
func (mr *MockAddressRepoMockRecorder) MakeCurrentAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeCurrentAddress", reflect.TypeOf((*MockAddressRepo)(nil).MakeCurrentAddress), arg0, arg1)
}

// UpdateAddress mocks base method.
func (m *MockAddressRepo) UpdateAddress(arg0 context.Context, arg1 models.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockAddressRepoMockRecorder) UpdateAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockAddressRepo)(nil).UpdateAddress), arg0, arg1)
}
