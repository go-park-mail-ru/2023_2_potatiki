package grpc

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/delivery/grpc/gen"
)

func TestGrpcOrderHandler_CreateOrder(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *gen.CreateOrderRequest
	}
	tests := []struct {
		name    string
		h       GrpcOrderHandler
		args    args
		want    *gen.CreateOrderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
	type args struct {
		ctx context.Context
		in  *gen.OrdersRequest
	}
	tests := []struct {
		name    string
		h       GrpcOrderHandler
		args    args
		want    *gen.OrdersResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
