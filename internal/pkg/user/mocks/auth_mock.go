// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUserUsecase is a mock of UserUsecase interface.
type MockUserUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUsecaseMockRecorder
}

// MockUserUsecaseMockRecorder is the mock recorder for MockUserUsecase.
type MockUserUsecaseMockRecorder struct {
	mock *MockUserUsecase
}

// NewMockUserUsecase creates a new mock instance.
func NewMockUserUsecase(ctrl *gomock.Controller) *MockUserUsecase {
	mock := &MockUserUsecase{ctrl: ctrl}
	mock.recorder = &MockUserUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecase) EXPECT() *MockUserUsecaseMockRecorder {
	return m.recorder
}

// GetProfile mocks base method.
func (m *MockUserUsecase) GetProfile(arg0 context.Context, arg1 uuid.UUID) (*models.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", arg0, arg1)
	ret0, _ := ret[0].(*models.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockUserUsecaseMockRecorder) GetProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockUserUsecase)(nil).GetProfile), arg0, arg1)
}

// UpdateInfo mocks base method.
func (m *MockUserUsecase) UpdateInfo(arg0 context.Context, arg1 uuid.UUID, arg2 *models.ProfileInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInfo indicates an expected call of UpdateInfo.
func (mr *MockUserUsecaseMockRecorder) UpdateInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInfo", reflect.TypeOf((*MockUserUsecase)(nil).UpdateInfo), arg0, arg1, arg2)
}

// UpdatePhoto mocks base method.
func (m *MockUserUsecase) UpdatePhoto(ctx context.Context, userID uuid.UUID, filePhotoByte []byte, fileType string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhoto", ctx, userID, filePhotoByte, fileType)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePhoto indicates an expected call of UpdatePhoto.
func (mr *MockUserUsecaseMockRecorder) UpdatePhoto(ctx, userID, filePhotoByte, fileType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhoto", reflect.TypeOf((*MockUserUsecase)(nil).UpdatePhoto), ctx, userID, filePhotoByte, fileType)
}

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// CreateProfile mocks base method.
func (m *MockUserRepo) CreateProfile(arg0 context.Context, arg1 *models.Profile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProfile indicates an expected call of CreateProfile.
func (mr *MockUserRepoMockRecorder) CreateProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfile", reflect.TypeOf((*MockUserRepo)(nil).CreateProfile), arg0, arg1)
}

// GetProfileIdByUser mocks base method.
func (m *MockUserRepo) GetProfileIdByUser(arg0 context.Context, arg1 *models.User) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileIdByUser", arg0, arg1)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileIdByUser indicates an expected call of GetProfileIdByUser.
func (mr *MockUserRepoMockRecorder) GetProfileIdByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileIdByUser", reflect.TypeOf((*MockUserRepo)(nil).GetProfileIdByUser), arg0, arg1)
}

// ReadProfile mocks base method.
func (m *MockUserRepo) ReadProfile(arg0 context.Context, arg1 uuid.UUID) (*models.Profile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadProfile", arg0, arg1)
	ret0, _ := ret[0].(*models.Profile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadProfile indicates an expected call of ReadProfile.
func (mr *MockUserRepoMockRecorder) ReadProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadProfile", reflect.TypeOf((*MockUserRepo)(nil).ReadProfile), arg0, arg1)
}

// UpdatePhoto mocks base method.
func (m *MockUserRepo) UpdatePhoto(arg0 context.Context, arg1 uuid.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePhoto", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePhoto indicates an expected call of UpdatePhoto.
func (mr *MockUserRepoMockRecorder) UpdatePhoto(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePhoto", reflect.TypeOf((*MockUserRepo)(nil).UpdatePhoto), arg0, arg1, arg2)
}

// UpdateProfile mocks base method.
func (m *MockUserRepo) UpdateProfile(arg0 context.Context, arg1 *models.Profile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockUserRepoMockRecorder) UpdateProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUserRepo)(nil).UpdateProfile), arg0, arg1)
}
