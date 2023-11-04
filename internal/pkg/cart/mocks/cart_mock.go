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

// MockCartUsecase is a mock of CartUsecase interface.
type MockCartUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCartUsecaseMockRecorder
}

// MockCartUsecaseMockRecorder is the mock recorder for MockCartUsecase.
type MockCartUsecaseMockRecorder struct {
	mock *MockCartUsecase
}

// NewMockCartUsecase creates a new mock instance.
func NewMockCartUsecase(ctrl *gomock.Controller) *MockCartUsecase {
	mock := &MockCartUsecase{ctrl: ctrl}
	mock.recorder = &MockCartUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartUsecase) EXPECT() *MockCartUsecaseMockRecorder {
	return m.recorder
}

// AddProduct mocks base method.
func (m *MockCartUsecase) AddProduct(arg0 context.Context, arg1 models.Cart, arg2 models.CartProductUpdate) (models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProduct", arg0, arg1, arg2)
	ret0, _ := ret[0].(models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProduct indicates an expected call of AddProduct.
func (mr *MockCartUsecaseMockRecorder) AddProduct(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockCartUsecase)(nil).AddProduct), arg0, arg1, arg2)
}

// DeleteProduct mocks base method.
func (m *MockCartUsecase) DeleteProduct(arg0 context.Context, arg1 models.Cart, arg2 models.CartProductDelete) (models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1, arg2)
	ret0, _ := ret[0].(models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockCartUsecaseMockRecorder) DeleteProduct(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockCartUsecase)(nil).DeleteProduct), arg0, arg1, arg2)
}

// GetCart mocks base method.
func (m *MockCartUsecase) GetCart(arg0 context.Context, arg1 uuid.UUID) (models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCart", arg0, arg1)
	ret0, _ := ret[0].(models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCart indicates an expected call of GetCart.
func (mr *MockCartUsecaseMockRecorder) GetCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockCartUsecase)(nil).GetCart), arg0, arg1)
}

// UpdateCart mocks base method.
func (m *MockCartUsecase) UpdateCart(arg0 context.Context, arg1 models.Cart) (models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", arg0, arg1)
	ret0, _ := ret[0].(models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCart indicates an expected call of UpdateCart.
func (mr *MockCartUsecaseMockRecorder) UpdateCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockCartUsecase)(nil).UpdateCart), arg0, arg1)
}

// MockCartRepo is a mock of CartRepo interface.
type MockCartRepo struct {
	ctrl     *gomock.Controller
	recorder *MockCartRepoMockRecorder
}

// MockCartRepoMockRecorder is the mock recorder for MockCartRepo.
type MockCartRepoMockRecorder struct {
	mock *MockCartRepo
}

// NewMockCartRepo creates a new mock instance.
func NewMockCartRepo(ctrl *gomock.Controller) *MockCartRepo {
	mock := &MockCartRepo{ctrl: ctrl}
	mock.recorder = &MockCartRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartRepo) EXPECT() *MockCartRepoMockRecorder {
	return m.recorder
}

// AddProduct mocks base method.
func (m *MockCartRepo) AddProduct(arg0 context.Context, arg1 models.Cart, arg2 models.CartProductUpdate) (models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProduct", arg0, arg1, arg2)
	ret0, _ := ret[0].(models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProduct indicates an expected call of AddProduct.
func (mr *MockCartRepoMockRecorder) AddProduct(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockCartRepo)(nil).AddProduct), arg0, arg1, arg2)
}

// CheckCart mocks base method.
func (m *MockCartRepo) CheckCart(arg0 context.Context, arg1 uuid.UUID) (models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCart", arg0, arg1)
	ret0, _ := ret[0].(models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckCart indicates an expected call of CheckCart.
func (mr *MockCartRepoMockRecorder) CheckCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCart", reflect.TypeOf((*MockCartRepo)(nil).CheckCart), arg0, arg1)
}

// CreateCart mocks base method.
func (m *MockCartRepo) CreateCart(arg0 context.Context, arg1 uuid.UUID) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCart", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCart indicates an expected call of CreateCart.
func (mr *MockCartRepoMockRecorder) CreateCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCart", reflect.TypeOf((*MockCartRepo)(nil).CreateCart), arg0, arg1)
}

// DeleteCart mocks base method.
func (m *MockCartRepo) DeleteCart(ctx context.Context, cartID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCart", ctx, cartID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCart indicates an expected call of DeleteCart.
func (mr *MockCartRepoMockRecorder) DeleteCart(ctx, cartID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCart", reflect.TypeOf((*MockCartRepo)(nil).DeleteCart), ctx, cartID)
}

// DeleteProduct mocks base method.
func (m *MockCartRepo) DeleteProduct(arg0 context.Context, arg1 models.Cart, arg2 models.CartProductDelete) (models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1, arg2)
	ret0, _ := ret[0].(models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockCartRepoMockRecorder) DeleteProduct(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockCartRepo)(nil).DeleteProduct), arg0, arg1, arg2)
}

// ReadCart mocks base method.
func (m *MockCartRepo) ReadCart(arg0 context.Context, arg1 uuid.UUID) (models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCart", arg0, arg1)
	ret0, _ := ret[0].(models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCart indicates an expected call of ReadCart.
func (mr *MockCartRepoMockRecorder) ReadCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCart", reflect.TypeOf((*MockCartRepo)(nil).ReadCart), arg0, arg1)
}

// ReadCartProducts mocks base method.
func (m *MockCartRepo) ReadCartProducts(arg0 context.Context, arg1 models.Cart) (models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCartProducts", arg0, arg1)
	ret0, _ := ret[0].(models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCartProducts indicates an expected call of ReadCartProducts.
func (mr *MockCartRepoMockRecorder) ReadCartProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCartProducts", reflect.TypeOf((*MockCartRepo)(nil).ReadCartProducts), arg0, arg1)
}

// UpdateCart mocks base method.
func (m *MockCartRepo) UpdateCart(arg0 context.Context, arg1 models.Cart) (models.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCart", arg0, arg1)
	ret0, _ := ret[0].(models.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCart indicates an expected call of UpdateCart.
func (mr *MockCartRepoMockRecorder) UpdateCart(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCart", reflect.TypeOf((*MockCartRepo)(nil).UpdateCart), arg0, arg1)
}
