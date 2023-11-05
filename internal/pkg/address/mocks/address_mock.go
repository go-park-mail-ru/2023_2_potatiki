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

// GetAllAddresses mocks base method.
func (m *MockAddressUsecase) GetAllAddresses(arg0 context.Context, arg1 uuid.UUID) ([]models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAddresses", arg0, arg1)
	ret0, _ := ret[0].([]models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAddresses indicates an expected call of GetAllAddresses.
func (mr *MockAddressUsecaseMockRecorder) GetAllAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAddresses", reflect.TypeOf((*MockAddressUsecase)(nil).GetAllAddresses), arg0, arg1)
}

// GetCurrentAddress mocks base method.
func (m *MockAddressUsecase) GetCurrentAddress(arg0 context.Context, arg1 uuid.UUID) (models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentAddress", arg0, arg1)
	ret0, _ := ret[0].(models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentAddress indicates an expected call of GetCurrentAddress.
func (mr *MockAddressUsecaseMockRecorder) GetCurrentAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentAddress", reflect.TypeOf((*MockAddressUsecase)(nil).GetCurrentAddress), arg0, arg1)
}

// MakeCurrentAddress mocks base method.
func (m *MockAddressUsecase) MakeCurrentAddress(arg0 context.Context, arg1 models.AddressMakeCurrent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeCurrentAddress", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// MakeCurrentAddress indicates an expected call of MakeCurrentAddress.
func (mr *MockAddressUsecaseMockRecorder) MakeCurrentAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeCurrentAddress", reflect.TypeOf((*MockAddressUsecase)(nil).MakeCurrentAddress), arg0, arg1)
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

// ReadAddress mocks base method.
func (m *MockAddressRepo) ReadAddress(arg0 context.Context, arg1, arg2 uuid.UUID) (models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAddress", arg0, arg1, arg2)
	ret0, _ := ret[0].(models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAddress indicates an expected call of ReadAddress.
func (mr *MockAddressRepoMockRecorder) ReadAddress(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAddress", reflect.TypeOf((*MockAddressRepo)(nil).ReadAddress), arg0, arg1, arg2)
}

// ReadAllAddresses mocks base method.
func (m *MockAddressRepo) ReadAllAddresses(arg0 context.Context, arg1 uuid.UUID) ([]models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllAddresses", arg0, arg1)
	ret0, _ := ret[0].([]models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllAddresses indicates an expected call of ReadAllAddresses.
func (mr *MockAddressRepoMockRecorder) ReadAllAddresses(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllAddresses", reflect.TypeOf((*MockAddressRepo)(nil).ReadAllAddresses), arg0, arg1)
}

// ReadCurrentAddress mocks base method.
func (m *MockAddressRepo) ReadCurrentAddress(arg0 context.Context, arg1 uuid.UUID) (models.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCurrentAddress", arg0, arg1)
	ret0, _ := ret[0].(models.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCurrentAddress indicates an expected call of ReadCurrentAddress.
func (mr *MockAddressRepoMockRecorder) ReadCurrentAddress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCurrentAddress", reflect.TypeOf((*MockAddressRepo)(nil).ReadCurrentAddress), arg0, arg1)
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
