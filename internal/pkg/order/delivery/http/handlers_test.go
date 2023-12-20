package http

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/delivery/grpc/gen"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/mocks"
	orderRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/repo"
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

func TestOrderHandler_CreateOrder(t *testing.T) {
	id := uuid.NewV4()
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockOrderClient)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
		jsonString     string
	}{
		{
			name: "UnsuccessfulCreateOrder/GrpcError",
			mockUsecaseFn: func(client *mock.MockOrderClient) {
				client.EXPECT().CreateOrder(gomock.Any(), &gen.CreateOrderRequest{
					Id:           id.String(),
					DeliveryDate: "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f",
				}).Return(&gen.CreateOrderResponse{}, errors.New("error grpc"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"deliveryDate": "08ddc2cf-fbf9-491a-b0b3-a8d5f698d81f"
			}`,
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(client *mock.MockOrderClient) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
		},
		{
			name: "UnsuccessfulAddAddress/NoBody",
			mockUsecaseFn: func(client *mock.MockOrderClient) {
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockOrderUsecase(ctrl)
			client := mock.NewMockOrderClient(ctrl)

			tc.mockUsecaseFn(client)

			req := httptest.NewRequest(http.MethodPost, "http://zuzu-market.ru/api/order/create",
				strings.NewReader(tc.jsonString))
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewOrderHandler(client, logger.Set("local", os.Stdout), mockUsecase)
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

			client := mock.NewMockOrderClient(ctrl)

			req = req.WithContext(ctx)
			addressHandler := NewOrderHandler(client, logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.GetCurrentOrder(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestOrderHandler_GetOrders(t *testing.T) {
	id := uuid.NewV4()
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockOrderClient)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
	}{
		{
			name: "SuccessfulGetOrders",
			mockUsecaseFn: func(client *mock.MockOrderClient) {
				client.EXPECT().GetOrders(gomock.Any(), &gen.OrdersRequest{
					Id: id.String(),
				}).Return(&gen.OrdersResponse{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(client *mock.MockOrderClient) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
		},
		{
			name: "UnsuccessfulGetCurrentOrder",
			mockUsecaseFn: func(client *mock.MockOrderClient) {
				client.EXPECT().GetOrders(gomock.Any(), &gen.OrdersRequest{
					Id: id.String(),
				}).Return(&gen.OrdersResponse{}, errors.New("error grpc"))
			},

			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockOrderUsecase(ctrl)
			client := mock.NewMockOrderClient(ctrl)

			tc.mockUsecaseFn(client)

			req := httptest.NewRequest(http.MethodGet, "http://zuzu-market.ru/api/order/get_all", nil)
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewOrderHandler(client, logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.GetOrders(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}
