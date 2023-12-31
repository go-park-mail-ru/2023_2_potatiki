// Code generated by MockGen. DO NOT EDIT.
// Source: ./gen/products_grpc.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gen "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/delivery/grpc/gen"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockProductsClient is a mock of ProductsClient interface.
type MockProductsClient struct {
	ctrl     *gomock.Controller
	recorder *MockProductsClientMockRecorder
}

// MockProductsClientMockRecorder is the mock recorder for MockProductsClient.
type MockProductsClientMockRecorder struct {
	mock *MockProductsClient
}

// NewMockProductsClient creates a new mock instance.
func NewMockProductsClient(ctrl *gomock.Controller) *MockProductsClient {
	mock := &MockProductsClient{ctrl: ctrl}
	mock.recorder = &MockProductsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductsClient) EXPECT() *MockProductsClientMockRecorder {
	return m.recorder
}

// GetCategory mocks base method.
func (m *MockProductsClient) GetCategory(ctx context.Context, in *gen.CategoryRequest, opts ...grpc.CallOption) (*gen.CategoryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCategory", varargs...)
	ret0, _ := ret[0].(*gen.CategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockProductsClientMockRecorder) GetCategory(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockProductsClient)(nil).GetCategory), varargs...)
}

// GetProduct mocks base method.
func (m *MockProductsClient) GetProduct(ctx context.Context, in *gen.ProductRequest, opts ...grpc.CallOption) (*gen.ProductResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProduct", varargs...)
	ret0, _ := ret[0].(*gen.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockProductsClientMockRecorder) GetProduct(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockProductsClient)(nil).GetProduct), varargs...)
}

// GetProducts mocks base method.
func (m *MockProductsClient) GetProducts(ctx context.Context, in *gen.ProductsRequest, opts ...grpc.CallOption) (*gen.ProductsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProducts", varargs...)
	ret0, _ := ret[0].(*gen.ProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockProductsClientMockRecorder) GetProducts(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockProductsClient)(nil).GetProducts), varargs...)
}

// MockProductsServer is a mock of ProductsServer interface.
type MockProductsServer struct {
	ctrl     *gomock.Controller
	recorder *MockProductsServerMockRecorder
}

// MockProductsServerMockRecorder is the mock recorder for MockProductsServer.
type MockProductsServerMockRecorder struct {
	mock *MockProductsServer
}

// NewMockProductsServer creates a new mock instance.
func NewMockProductsServer(ctrl *gomock.Controller) *MockProductsServer {
	mock := &MockProductsServer{ctrl: ctrl}
	mock.recorder = &MockProductsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductsServer) EXPECT() *MockProductsServerMockRecorder {
	return m.recorder
}

// GetCategory mocks base method.
func (m *MockProductsServer) GetCategory(arg0 context.Context, arg1 *gen.CategoryRequest) (*gen.CategoryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", arg0, arg1)
	ret0, _ := ret[0].(*gen.CategoryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockProductsServerMockRecorder) GetCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockProductsServer)(nil).GetCategory), arg0, arg1)
}

// GetProduct mocks base method.
func (m *MockProductsServer) GetProduct(arg0 context.Context, arg1 *gen.ProductRequest) (*gen.ProductResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(*gen.ProductResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockProductsServerMockRecorder) GetProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockProductsServer)(nil).GetProduct), arg0, arg1)
}

// GetProducts mocks base method.
func (m *MockProductsServer) GetProducts(arg0 context.Context, arg1 *gen.ProductsRequest) (*gen.ProductsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", arg0, arg1)
	ret0, _ := ret[0].(*gen.ProductsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockProductsServerMockRecorder) GetProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockProductsServer)(nil).GetProducts), arg0, arg1)
}

// mustEmbedUnimplementedProductsServer mocks base method.
func (m *MockProductsServer) mustEmbedUnimplementedProductsServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedProductsServer")
}

// mustEmbedUnimplementedProductsServer indicates an expected call of mustEmbedUnimplementedProductsServer.
func (mr *MockProductsServerMockRecorder) mustEmbedUnimplementedProductsServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedProductsServer", reflect.TypeOf((*MockProductsServer)(nil).mustEmbedUnimplementedProductsServer))
}

// MockUnsafeProductsServer is a mock of UnsafeProductsServer interface.
type MockUnsafeProductsServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeProductsServerMockRecorder
}

// MockUnsafeProductsServerMockRecorder is the mock recorder for MockUnsafeProductsServer.
type MockUnsafeProductsServerMockRecorder struct {
	mock *MockUnsafeProductsServer
}

// NewMockUnsafeProductsServer creates a new mock instance.
func NewMockUnsafeProductsServer(ctrl *gomock.Controller) *MockUnsafeProductsServer {
	mock := &MockUnsafeProductsServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeProductsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeProductsServer) EXPECT() *MockUnsafeProductsServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedProductsServer mocks base method.
func (m *MockUnsafeProductsServer) mustEmbedUnimplementedProductsServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedProductsServer")
}

// mustEmbedUnimplementedProductsServer indicates an expected call of mustEmbedUnimplementedProductsServer.
func (mr *MockUnsafeProductsServerMockRecorder) mustEmbedUnimplementedProductsServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedProductsServer", reflect.TypeOf((*MockUnsafeProductsServer)(nil).mustEmbedUnimplementedProductsServer))
}
