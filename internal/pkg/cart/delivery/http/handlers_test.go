package http

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/mocks"
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

func TestCartHandler_AddProduct(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockCartUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
		jsonString     string
	}{
		{
			name: "SuccessfulAddProduct",
			mockUsecaseFn: func(mockUsecase *mock.MockCartUsecase) {
				mockUsecase.EXPECT().AddProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Cart{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"productId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f",
				"quantity": 2
			}`,
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockCartUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
			jsonString: `{
				"productId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f",
				"quantity": 2
			}`,
		},
		{
			name:           "EmptyBody",
			mockUsecaseFn:  func(mockUsecase *mock.MockCartUsecase) {},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: ``,
		},
		{
			name: "UnsuccessfulAddProduct",
			mockUsecaseFn: func(mockUsecase *mock.MockCartUsecase) {
				mockUsecase.EXPECT().AddProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Cart{}, errors.New("error in add product"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"productId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f",
				"quantity": 2
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockCartUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodPost, "http://zuzu-market.ru/api/cart/add_product",
				strings.NewReader(tc.jsonString))
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewCartHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.AddProduct(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestCartHandler_DeleteProduct(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockCartUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
		jsonString     string
	}{
		{
			name: "SuccessfulDeleteProduct",
			mockUsecaseFn: func(mockUsecase *mock.MockCartUsecase) {
				mockUsecase.EXPECT().DeleteProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Cart{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"productId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f"
			}`,
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockCartUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
			jsonString: `{
				"productId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f"
			}`,
		},
		{
			name:           "EmptyBody",
			mockUsecaseFn:  func(mockUsecase *mock.MockCartUsecase) {},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: ``,
		},
		{
			name: "UnsuccessfulDeleteProduct",
			mockUsecaseFn: func(mockUsecase *mock.MockCartUsecase) {
				mockUsecase.EXPECT().DeleteProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Cart{}, errors.New("error in delete product"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"productId": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f"
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockCartUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodDelete, "http://zuzu-market.ru/api/cart/delete_product",
				strings.NewReader(tc.jsonString))
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewCartHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.DeleteProduct(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestCartHandler_GetCart(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockCartUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
	}{
		{
			name: "SuccessfulGetCart",
			mockUsecaseFn: func(mockUsecase *mock.MockCartUsecase) {
				mockUsecase.EXPECT().GetCart(gomock.Any(), gomock.Any()).Return(models.Cart{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockCartUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
		},
		{
			name: "UnsuccessfulGetCart",
			mockUsecaseFn: func(mockUsecase *mock.MockCartUsecase) {
				mockUsecase.EXPECT().GetCart(gomock.Any(), gomock.Any()).Return(models.Cart{}, errors.New("error in get cart"))
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

			mockUsecase := mock.NewMockCartUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodGet, "http://zuzu-market.ru/api/cart/summary", nil)
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewCartHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.GetCart(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestCartHandler_UpdateCart(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockCartUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
		jsonString     string
	}{
		{
			name: "SuccessfulUpdateCart",
			mockUsecaseFn: func(mockUsecase *mock.MockCartUsecase) {
				mockUsecase.EXPECT().UpdateCart(gomock.Any(), gomock.Any()).Return(models.Cart{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"products": [
					{
						"quantity": 1,
						"productId": "007749b5-7e07-4be8-8c91-8db273ace4c3"
					},
					{
						"quantity": 2,
						"productId": "007749b5-7e07-4be8-8c91-8db273ace4c1"
					},
					{
						"quantity": 7,
						"productId": "007749b5-7e07-4be8-8c91-8db273ace4c2"
					}
				]
			}`,
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockCartUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
			jsonString: `{
				"products": [
					{
						"quantity": 1,
						"productId": "007749b5-7e07-4be8-8c91-8db273ace4c3"
					},
					{
						"quantity": 2,
						"productId": "007749b5-7e07-4be8-8c91-8db273ace4c1"
					},
					{
						"quantity": 7,
						"productId": "007749b5-7e07-4be8-8c91-8db273ace4c2"
					}
				]
			}`,
		},
		{
			name:           "EmptyBody",
			mockUsecaseFn:  func(mockUsecase *mock.MockCartUsecase) {},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: ``,
		},
		{
			name: "UnsuccessfulUpdateCart",
			mockUsecaseFn: func(mockUsecase *mock.MockCartUsecase) {
				mockUsecase.EXPECT().UpdateCart(gomock.Any(), gomock.Any()).Return(models.Cart{}, errors.New("error in add product"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"products": [
					{
						"quantity": 1,
						"productId": "007749b5-7e07-4be8-8c91-8db273ace4c3"
					},
					{
						"quantity": 2,
						"productId": "007749b5-7e07-4be8-8c91-8db273ace4c1"
					},
					{
						"quantity": 7,
						"productId": "007749b5-7e07-4be8-8c91-8db273ace4c2"
					}
				]
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockCartUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodPost, "http://zuzu-market.ru/api/cart/update",
				strings.NewReader(tc.jsonString))
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewCartHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.UpdateCart(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}
