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

// MockOrderUsecase is a mock of OrderUsecase interface.
type MockOrderUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockOrderUsecaseMockRecorder
}

// MockOrderUsecaseMockRecorder is the mock recorder for MockOrderUsecase.
type MockOrderUsecaseMockRecorder struct {
	mock *MockOrderUsecase
}

// NewMockOrderUsecase creates a new mock instance.
func NewMockOrderUsecase(ctrl *gomock.Controller) *MockOrderUsecase {
	mock := &MockOrderUsecase{ctrl: ctrl}
	mock.recorder = &MockOrderUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderUsecase) EXPECT() *MockOrderUsecaseMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderUsecase) CreateOrder(arg0 context.Context, arg1 uuid.UUID, arg2, arg3, arg4 string) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderUsecaseMockRecorder) CreateOrder(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderUsecase)(nil).CreateOrder), arg0, arg1, arg2, arg3, arg4)
}

// GetCurrentOrder mocks base method.
func (m *MockOrderUsecase) GetCurrentOrder(arg0 context.Context, arg1 uuid.UUID) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentOrder", arg0, arg1)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentOrder indicates an expected call of GetCurrentOrder.
func (mr *MockOrderUsecaseMockRecorder) GetCurrentOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentOrder", reflect.TypeOf((*MockOrderUsecase)(nil).GetCurrentOrder), arg0, arg1)
}

// GetOrders mocks base method.
func (m *MockOrderUsecase) GetOrders(arg0 context.Context, arg1 uuid.UUID) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", arg0, arg1)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockOrderUsecaseMockRecorder) GetOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrderUsecase)(nil).GetOrders), arg0, arg1)
}

// MockOrderRepo is a mock of OrderRepo interface.
type MockOrderRepo struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepoMockRecorder
}

// MockOrderRepoMockRecorder is the mock recorder for MockOrderRepo.
type MockOrderRepoMockRecorder struct {
	mock *MockOrderRepo
}

// NewMockOrderRepo creates a new mock instance.
func NewMockOrderRepo(ctrl *gomock.Controller) *MockOrderRepo {
	mock := &MockOrderRepo{ctrl: ctrl}
	mock.recorder = &MockOrderRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepo) EXPECT() *MockOrderRepoMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderRepo) CreateOrder(arg0 context.Context, arg1 models.Cart, arg2, arg3 uuid.UUID, arg4 int64, arg5, arg6 string) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderRepoMockRecorder) CreateOrder(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderRepo)(nil).CreateOrder), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// GetUpdates mocks base method.
func (m *MockOrderRepo) GetUpdates(arg0 context.Context, arg1 uuid.UUID, arg2 time.Time) ([]models.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdates", arg0, arg1, arg2)
	ret0, _ := ret[0].([]models.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUpdates indicates an expected call of GetUpdates.
func (mr *MockOrderRepoMockRecorder) GetUpdates(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdates", reflect.TypeOf((*MockOrderRepo)(nil).GetUpdates), arg0, arg1, arg2)
}

// ReadOrder mocks base method.
func (m *MockOrderRepo) ReadOrder(arg0 context.Context, arg1 uuid.UUID) (models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadOrder", arg0, arg1)
	ret0, _ := ret[0].(models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadOrder indicates an expected call of ReadOrder.
func (mr *MockOrderRepoMockRecorder) ReadOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadOrder", reflect.TypeOf((*MockOrderRepo)(nil).ReadOrder), arg0, arg1)
}

// ReadOrderID mocks base method.
func (m *MockOrderRepo) ReadOrderID(arg0 context.Context, arg1 uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadOrderID", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadOrderID indicates an expected call of ReadOrderID.
func (mr *MockOrderRepoMockRecorder) ReadOrderID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadOrderID", reflect.TypeOf((*MockOrderRepo)(nil).ReadOrderID), arg0, arg1)
}

// ReadOrdersID mocks base method.
func (m *MockOrderRepo) ReadOrdersID(arg0 context.Context, arg1 uuid.UUID) ([]uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadOrdersID", arg0, arg1)
	ret0, _ := ret[0].([]uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadOrdersID indicates an expected call of ReadOrdersID.
func (mr *MockOrderRepoMockRecorder) ReadOrdersID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadOrdersID", reflect.TypeOf((*MockOrderRepo)(nil).ReadOrdersID), arg0, arg1)
}

// SetPromoOrder mocks base method.
func (m *MockOrderRepo) SetPromoOrder(arg0 context.Context, arg1 int, arg2 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPromoOrder", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPromoOrder indicates an expected call of SetPromoOrder.
func (mr *MockOrderRepoMockRecorder) SetPromoOrder(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPromoOrder", reflect.TypeOf((*MockOrderRepo)(nil).SetPromoOrder), arg0, arg1, arg2)
}
