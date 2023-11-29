package grpc

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/delivery/grpc/gen"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/go-park-mail-ru/2023_2_potatiki/proto/gmodels"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

func TestGrpcProductsHandler_GetProduct(t *testing.T) {
	id := uuid.NewV4()
	type fields struct {
		uc *mock.MockProductsUsecase
	}
	type args struct {
		ctx context.Context
		in  *gen.ProductRequest
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		h       *GrpcProductsHandler
		args    args
		want    *gen.ProductResponse
		wantErr bool
	}{
		{
			name: "TestGrpcProductsHandler_GetProduct good",
			prepare: func(f *fields) {
				f.uc.EXPECT().GetProduct(gomock.Any(), id).Return(
					models.Product{
						Id:          id,
						Name:        "123",
						Description: "123",
						Price:       123,
						Category:    models.Category{},
					}, nil)
			},
			args: args{context.Background(), &gen.ProductRequest{
				Id: id.String(),
			}},
			want: &gen.ProductResponse{
				Product: &gmodels.Product{
					Id:          id.String(),
					Name:        "123",
					Description: "123",
					Price:       123,
					Category:    &gmodels.Category{},
				}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				uc: mock.NewMockProductsUsecase(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.h = NewGrpcProductsHandler(f.uc, logger.Set("local", os.Stdout))

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
	id := uuid.NewV4()
	type fields struct {
		uc *mock.MockProductsUsecase
	}
	type args struct {
		ctx context.Context
		in  *gen.ProductsRequest
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		h       *GrpcProductsHandler
		args    args
		want    *gen.ProductsResponse
		wantErr bool
	}{
		{
			name: "TestGrpcProductsHandler_GetProducts good",
			prepare: func(f *fields) {
				f.uc.EXPECT().GetProducts(gomock.Any(), int64(1), int64(1), "", "").Return([]models.Product{{
					Id:          id,
					Name:        "123",
					Description: "123",
					Price:       123,
					Category:    models.Category{},
				}}, nil)
			},
			args: args{context.Background(), &gen.ProductsRequest{
				Paging:   1,
				Count:    1,
				PriceBy:  "",
				RatingBy: "",
			}},
			want: &gen.ProductsResponse{
				Products: []*gmodels.Product{{
					Id:          id.String(),
					Name:        "123",
					Description: "123",
					Price:       123,
					Category:    &gmodels.Category{},
				}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				uc: mock.NewMockProductsUsecase(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.h = NewGrpcProductsHandler(f.uc, logger.Set("local", os.Stdout))

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
	id := uuid.NewV4()
	type fields struct {
		uc *mock.MockProductsUsecase
	}
	type args struct {
		ctx context.Context
		in  *gen.CategoryRequest
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		h       *GrpcProductsHandler
		args    args
		want    *gen.CategoryResponse
		wantErr bool
	}{
		{
			name: "TestGrpcProductsHandler_GetProduct good",
			prepare: func(f *fields) {
				f.uc.EXPECT().GetCategory(gomock.Any(), 1, int64(1), int64(1), "", "").Return([]models.Product{{
					Id:          id,
					Name:        "123",
					Description: "123",
					Price:       123,
					Category:    models.Category{},
				}}, nil)
			},
			args: args{context.Background(), &gen.CategoryRequest{
				Id:       1,
				Paging:   1,
				Count:    1,
				PriceBy:  "",
				RatingBy: "",
			}},
			want: &gen.CategoryResponse{
				Products: []*gmodels.Product{{
					Id:          id.String(),
					Name:        "123",
					Description: "123",
					Price:       123,
					Category:    &gmodels.Category{},
				}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				uc: mock.NewMockProductsUsecase(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.h = NewGrpcProductsHandler(f.uc, logger.Set("local", os.Stdout))

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
