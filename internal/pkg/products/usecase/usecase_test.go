package usecase

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/mocks"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

func TestProductsUsecase_GetProduct(t *testing.T) {
	type fields struct {
		repo *mock.MockProductsRepo
	}
	type args struct {
		ctx context.Context //nolint:containedctx
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		uc      *ProductsUsecase
		args    args
		want    models.Product
		wantErr bool
	}{
		{
			name: "TestProductsUsecase_GetProduct good",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadProduct(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Product{}, nil)
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			want:    models.Product{},
			wantErr: false,
		},
		{
			name: "TestProductsUsecase_GetGetProduct bad",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadProduct(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Product{}, fmt.Errorf("dummyErr"))
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockProductsRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewProductsUsecase(f.repo)

			got, err := tt.uc.GetProduct(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductsUsecase.GetProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductsUsecase.GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductsUsecase_GetProducts(t *testing.T) {
	type fields struct {
		repo *mock.MockProductsRepo
	}
	type args struct {
		ctx    context.Context //nolint:containedctx
		paging int64
		count  int64
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		uc      *ProductsUsecase
		args    args
		want    []models.Product
		wantErr bool
	}{
		{
			name: "TestProductsUsecase_GetProducts good",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadProducts(gomock.Any(), int64(1), int64(1)).Return([]models.Product{}, nil)
			},
			args:    args{context.Background(), 1, 1},
			want:    []models.Product{},
			wantErr: false,
		},
		{
			name: "TestProductsUsecase_GetProducts bad",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadProducts(gomock.Any(), int64(1), int64(1)).Return(nil, fmt.Errorf("dummyErr"))
			},
			args:    args{context.Background(), 1, 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockProductsRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewProductsUsecase(f.repo)

			got, err := tt.uc.GetProducts(tt.args.ctx, tt.args.paging, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductsUsecase.GetProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductsUsecase.GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductsUsecase_GetCategory(t *testing.T) {
	type fields struct {
		repo *mock.MockProductsRepo
	}
	type args struct {
		ctx    context.Context //nolint:containedctx
		id     int
		paging int64
		count  int64
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		uc      *ProductsUsecase
		args    args
		want    []models.Product
		wantErr bool
	}{
		{
			name: "TestProductsUsecase_GetCategory good",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadCategory(gomock.Any(), 1, int64(1), int64(1)).Return([]models.Product{}, nil)
			},
			args:    args{context.Background(), 1, 1, 1},
			want:    []models.Product{},
			wantErr: false,
		},
		{
			name: "TestProductsUsecase_GetCategory bad",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadCategory(gomock.Any(), 1, int64(1), int64(1)).Return(nil, fmt.Errorf("dummyErr"))
			},
			args:    args{context.Background(), 1, 1, 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockProductsRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewProductsUsecase(f.repo)

			got, err := tt.uc.GetCategory(tt.args.ctx, tt.args.id, tt.args.paging, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductsUsecase.GetCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductsUsecase.GetCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}
