package http

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	addressRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address/repo"
	cartRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/repo"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/mocks"
	orderRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/repo"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestOrderHandler_CreateOrder(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockOrderUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
	}{
		{
			name: "SuccessfulCreateOrder",
			mockUsecaseFn: func(mockUsecase *mock.MockOrderUsecase) {
				mockUsecase.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(models.Order{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockOrderUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
		},
		{
			name: "UnsuccessfulAddAddress",
			mockUsecaseFn: func(mockUsecase *mock.MockOrderUsecase) {
				mockUsecase.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(models.Order{}, errors.New("error in create order"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name: "CartNotFound",
			mockUsecaseFn: func(mockUsecase *mock.MockOrderUsecase) {
				mockUsecase.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(models.Order{}, cartRepo.ErrCartNotFound)
			},
			expectedStatus: http.StatusNotFound,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name: "ProductNotFound",
			mockUsecaseFn: func(mockUsecase *mock.MockOrderUsecase) {
				mockUsecase.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(models.Order{}, orderRepo.ErrPoductNotFound)
			},
			expectedStatus: http.StatusNotFound,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name: "AddressNotFound",
			mockUsecaseFn: func(mockUsecase *mock.MockOrderUsecase) {
				mockUsecase.EXPECT().CreateOrder(gomock.Any(), gomock.Any()).Return(models.Order{}, addressRepo.ErrAddressNotFound)
			},
			expectedStatus: http.StatusNotFound,
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

			mockUsecase := mock.NewMockOrderUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodPost, "http://zuzu-market.ru/api/order/create", nil)
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewOrderHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.CreateOrder(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestOrderHandler_GetCurrentOrder(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockOrderUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
	}{
		{
			name: "SuccessfulGetCurrentOrder",
			mockUsecaseFn: func(mockUsecase *mock.MockOrderUsecase) {
				mockUsecase.EXPECT().GetCurrentOrder(gomock.Any(), gomock.Any()).Return(models.Order{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockOrderUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
		},
		{
			name: "UnsuccessfulGetCurrentOrder",
			mockUsecaseFn: func(mockUsecase *mock.MockOrderUsecase) {
				mockUsecase.EXPECT().GetCurrentOrder(gomock.Any(), gomock.Any()).Return(models.Order{}, errors.New("error in get current order"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name: "OrderNotFound",
			mockUsecaseFn: func(mockUsecase *mock.MockOrderUsecase) {
				mockUsecase.EXPECT().GetCurrentOrder(gomock.Any(), gomock.Any()).Return(models.Order{}, orderRepo.ErrOrderNotFound)
			},
			expectedStatus: http.StatusNotFound,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name: "ProductsInOrderNotFound",
			mockUsecaseFn: func(mockUsecase *mock.MockOrderUsecase) {
				mockUsecase.EXPECT().GetCurrentOrder(gomock.Any(), gomock.Any()).Return(models.Order{}, orderRepo.ErrPoductsInOrderNotFound)
			},
			expectedStatus: http.StatusNotFound,
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

			mockUsecase := mock.NewMockOrderUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodGet, "http://zuzu-market.ru/api/order/get_current", nil)
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewOrderHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.GetCurrentOrder(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestOrderHandler_GetOrders(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockOrderUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
	}{
		{
			name: "SuccessfulGetOrders",
			mockUsecaseFn: func(mockUsecase *mock.MockOrderUsecase) {
				mockUsecase.EXPECT().GetOrders(gomock.Any(), gomock.Any()).Return([]models.Order{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockOrderUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
		},
		{
			name: "UnsuccessfulGetCurrentOrder",
			mockUsecaseFn: func(mockUsecase *mock.MockOrderUsecase) {
				mockUsecase.EXPECT().GetOrders(gomock.Any(), gomock.Any()).Return([]models.Order{}, errors.New("error in get orders"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name: "OrdersNotFound",
			mockUsecaseFn: func(mockUsecase *mock.MockOrderUsecase) {
				mockUsecase.EXPECT().GetOrders(gomock.Any(), gomock.Any()).Return([]models.Order{}, orderRepo.ErrOrdersNotFound)
			},
			expectedStatus: http.StatusNotFound,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name: "ProductsInOrderNotFound",
			mockUsecaseFn: func(mockUsecase *mock.MockOrderUsecase) {
				mockUsecase.EXPECT().GetOrders(gomock.Any(), gomock.Any()).Return([]models.Order{}, orderRepo.ErrPoductsInOrderNotFound)
			},
			expectedStatus: http.StatusNotFound,
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

			mockUsecase := mock.NewMockOrderUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodGet, "http://zuzu-market.ru/api/order/get_all", nil)
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewOrderHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.GetOrders(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}
