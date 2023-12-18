package grpc

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/delivery/grpc/gen"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/go-park-mail-ru/2023_2_potatiki/proto/gmodels"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

func TestGrpcOrderHandler_CreateOrder(t *testing.T) {
	t.Skip()
	id := uuid.NewV4()
	type fields struct {
		uc *mock.MockOrderUsecase
	}
	type args struct {
		ctx context.Context
		in  *gen.CreateOrderRequest
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		h       *GrpcOrderHandler
		args    args
		want    *gen.CreateOrderResponse
		wantErr bool
	}{
		{
			name: "TestGrpcOrderHandler_CreateOrder good",
			prepare: func(f *fields) {
				f.uc.EXPECT().CreateOrder(gomock.Any(), id, "", "", "").Return(
					models.Order{
						Id:      id,
						Address: models.Address{},
					}, nil)
			},
			args: args{context.Background(), &gen.CreateOrderRequest{
				Id: id.String(),
			}},
			want: &gen.CreateOrderResponse{
				Order: &gmodels.Order{
					Id:         id.String(),
					Address:    &gmodels.Address{Id: uuid.Nil.String(), ProfileId: uuid.Nil.String()},
					CreationAt: time.Time{}.Format(time.RFC3339),
				}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				uc: mock.NewMockOrderUsecase(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.h = NewGrpcOrderHandler(f.uc, logger.Set("local", os.Stdout))

			got, err := tt.h.CreateOrder(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GrpcOrderHandler.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GrpcOrderHandler.CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrpcOrderHandler_GetOrders(t *testing.T) {
	t.Skip()
	id := uuid.NewV4()
	type fields struct {
		uc *mock.MockOrderUsecase
	}
	type args struct {
		ctx context.Context
		in  *gen.OrdersRequest
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		h       *GrpcOrderHandler
		args    args
		want    *gen.OrdersResponse
		wantErr bool
	}{
		{
			name: "TestGrpcOrderHandler_GetOrders good",
			prepare: func(f *fields) {
				f.uc.EXPECT().GetOrders(gomock.Any(), id).Return(
					[]models.Order{{
						Id:      id,
						Address: models.Address{},
					}}, nil)
			},
			args: args{context.Background(), &gen.OrdersRequest{
				Id: id.String(),
			}},
			want: &gen.OrdersResponse{
				Orders: []*gmodels.Order{{
					Id:         id.String(),
					Address:    &gmodels.Address{Id: uuid.Nil.String(), ProfileId: uuid.Nil.String()},
					CreationAt: time.Time{}.Format(time.RFC3339),
				}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				uc: mock.NewMockOrderUsecase(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.h = NewGrpcOrderHandler(f.uc, logger.Set("local", os.Stdout))

			got, err := tt.h.GetOrders(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GrpcOrderHandler.GetOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GrpcOrderHandler.GetOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
