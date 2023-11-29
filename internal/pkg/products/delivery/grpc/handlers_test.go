package grpc

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/delivery/grpc/gen"
)

func TestGrpcProductsHandler_GetProduct(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *gen.ProductRequest
	}
	tests := []struct {
		name    string
		h       GrpcProductsHandler
		args    args
		want    *gen.ProductResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.GetProduct(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GrpcProductsHandler.GetProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GrpcProductsHandler.GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrpcProductsHandler_GetProducts(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *gen.ProductsRequest
	}
	tests := []struct {
		name    string
		h       GrpcProductsHandler
		args    args
		want    *gen.ProductsResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.GetProducts(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GrpcProductsHandler.GetProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GrpcProductsHandler.GetProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrpcProductsHandler_GetCategory(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *gen.CategoryRequest
	}
	tests := []struct {
		name    string
		h       GrpcProductsHandler
		args    args
		want    *gen.CategoryResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.GetCategory(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GrpcProductsHandler.GetCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GrpcProductsHandler.GetCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}
