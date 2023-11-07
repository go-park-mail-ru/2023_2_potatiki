package http

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestAddressHandler_AddAddress(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockAddressUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
		jsonString     string
	}{
		{
			name: "SuccessfulAddAddress",
			mockUsecaseFn: func(mockUsecase *mock.MockAddressUsecase) {
				mockUsecase.EXPECT().AddAddress(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Address{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"addressId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f",
				"city": "Москва",
				"street": "Туристская",
				"house": "33",
				"flat": "1000",
				"addressIsCurrent": true
			}`,
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockAddressUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
			jsonString: `{
				"addressId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f",
				"city": "Москва",
				"street": "Туристская",
				"house": "33",
				"flat": "1000",
				"addressIsCurrent": true
			}`,
		},
		{
			name:           "EmptyBody",
			mockUsecaseFn:  func(mockUsecase *mock.MockAddressUsecase) {},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: ``,
		},
		{
			name: "UnsuccessfulAddAddress",
			mockUsecaseFn: func(mockUsecase *mock.MockAddressUsecase) {
				mockUsecase.EXPECT().AddAddress(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Address{}, errors.New("error in add address"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"addressId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f",
				"city": "Москва",
				"street": "Туристская",
				"house": "33",
				"flat": "1000",
				"addressIsCurrent": true
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockAddressUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodPost, "http://zuzu-market.ru/api/address/add",
				strings.NewReader(tc.jsonString))
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewAddressHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.AddAddress(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestAddressHandler_DeleteAddress(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockAddressUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
		jsonString     string
	}{
		{
			name: "SuccessfulDeleteAddress",
			mockUsecaseFn: func(mockUsecase *mock.MockAddressUsecase) {
				mockUsecase.EXPECT().DeleteAddress(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"addressId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f"
			}`,
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockAddressUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
			jsonString: `{
				"addressId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f"
			}`,
		},
		{
			name:           "EmptyBody",
			mockUsecaseFn:  func(mockUsecase *mock.MockAddressUsecase) {},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: ``,
		},
		{
			name: "UnsuccessfulDeleteAddress",
			mockUsecaseFn: func(mockUsecase *mock.MockAddressUsecase) {
				mockUsecase.EXPECT().DeleteAddress(gomock.Any(), gomock.Any()).Return(errors.New("error in delete address"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"addressId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f"
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockAddressUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodDelete, "http://zuzu-market.ru/api/address/delete",
				strings.NewReader(tc.jsonString))
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewAddressHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.DeleteAddress(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestAddressHandler_GetAllAddresses(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockAddressUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
	}{
		{
			name: "SuccessfulGetAllAddresses",
			mockUsecaseFn: func(mockUsecase *mock.MockAddressUsecase) {
				mockUsecase.EXPECT().GetAllAddresses(gomock.Any(), gomock.Any()).Return([]models.Address{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockAddressUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
		},
		{
			name: "UnsuccessfulGetAllAddresses",
			mockUsecaseFn: func(mockUsecase *mock.MockAddressUsecase) {
				mockUsecase.EXPECT().GetAllAddresses(gomock.Any(), gomock.Any()).Return([]models.Address{}, errors.New("error in get all addresses"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockAddressUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodGet, "http://zuzu-market.ru/api/address/get_all", nil)
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewAddressHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.GetAllAddresses(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestAddressHandler_GetCurrentAddress(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockAddressUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
	}{
		{
			name: "SuccessfulGetCurrentAddress",
			mockUsecaseFn: func(mockUsecase *mock.MockAddressUsecase) {
				mockUsecase.EXPECT().GetCurrentAddress(gomock.Any(), gomock.Any()).Return(models.Address{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockAddressUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
		},
		{
			name: "UnsuccessfulGetAllAddresses",
			mockUsecaseFn: func(mockUsecase *mock.MockAddressUsecase) {
				mockUsecase.EXPECT().GetCurrentAddress(gomock.Any(), gomock.Any()).Return(models.Address{}, errors.New("error in get current address"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockAddressUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodGet, "http://zuzu-market.ru/api/address/get_current", nil)
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewAddressHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.GetCurrentAddress(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestAddressHandler_MakeCurrentAddress(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockAddressUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
		jsonString     string
	}{
		{
			name: "SuccessfulMakeCurrentAddress",
			mockUsecaseFn: func(mockUsecase *mock.MockAddressUsecase) {
				mockUsecase.EXPECT().MakeCurrentAddress(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"addressId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f"
			}`,
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockAddressUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
			jsonString: `{
				"addressId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f"
			}`,
		},
		{
			name:           "EmptyBody",
			mockUsecaseFn:  func(mockUsecase *mock.MockAddressUsecase) {},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: ``,
		},
		{
			name: "UnsuccessfulMakeCurrentAddress",
			mockUsecaseFn: func(mockUsecase *mock.MockAddressUsecase) {
				mockUsecase.EXPECT().MakeCurrentAddress(gomock.Any(), gomock.Any()).Return(errors.New("error in add address"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"addressId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f"
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockAddressUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodPost, "http://zuzu-market.ru/api/address/add",
				strings.NewReader(tc.jsonString))
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewAddressHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.MakeCurrentAddress(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestAddressHandler_UpdateAddress(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockAddressUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
		jsonString     string
	}{
		{
			name: "SuccessfulUpdateAddress",
			mockUsecaseFn: func(mockUsecase *mock.MockAddressUsecase) {
				mockUsecase.EXPECT().UpdateAddress(gomock.Any(), gomock.Any()).Return(models.Address{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"addressId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f",
				"city": "Москва",
				"street": "Туристская",
				"house": "33",
				"flat": "1000",
				"addressIsCurrent": true
			}`,
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockAddressUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
			jsonString: `{
				"addressId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f",
				"city": "Москва",
				"street": "Туристская",
				"house": "33",
				"flat": "1000",
				"addressIsCurrent": true
			}`,
		},
		{
			name:           "EmptyBody",
			mockUsecaseFn:  func(mockUsecase *mock.MockAddressUsecase) {},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: ``,
		},
		{
			name: "UnsuccessfulAddAddress",
			mockUsecaseFn: func(mockUsecase *mock.MockAddressUsecase) {
				mockUsecase.EXPECT().UpdateAddress(gomock.Any(), gomock.Any()).Return(models.Address{}, errors.New("error in update address"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"addressId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f",
				"city": "Москва",
				"street": "Туристская",
				"house": "33",
				"flat": "1000",
				"addressIsCurrent": true
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockAddressUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodPost, "http://zuzu-market.ru/api/address/update",
				strings.NewReader(tc.jsonString))
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewAddressHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.UpdateAddress(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}
