package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/repo"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

func TestCartUsecase_GetCart(t *testing.T) {
	type fields struct {
		repo *mock.MockCartRepo
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *CartUsecase
		args    args
		want    models.Cart
		wantErr bool
	}{
		{
			name: "TestCartUsecase_GetCart good",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			want:    models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: false,
		},
		{
			name: "TestCartUsecase_GetCart bad ReadCart",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, errors.New("dummyError"))
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
		{
			name: "TestCartUsecase_GetCart bad CreateCart",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, repo.ErrCartNotFound)
				f.repo.EXPECT().CreateCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), errors.New("dummyError"))

			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
		{
			name: "TestCartUsecase_GetCart bad good",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, repo.ErrCartNotFound)
				f.repo.EXPECT().CreateCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), nil)

			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			want:    models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), IsCurrent: true, Products: make([]models.CartProduct, 0)},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockCartRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewCartUsecase(f.repo)

			got, err := tt.uc.GetCart(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartUsecase.GetCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartUsecase.GetCart() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestCartUsecase_AddProduct(t *testing.T) {
	type fields struct {
		repo *mock.MockCartRepo
	}
	type args struct {
		ctx     context.Context
		cart    models.Cart
		product models.CartProductUpdate
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *CartUsecase
		args    args
		want    models.Cart
		wantErr bool
	}{
		{
			name: "TestCartUsecase_AddProduct good",
			prepare: func(f *fields) {
				f.repo.EXPECT().CheckCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
				f.repo.EXPECT().AddProduct(gomock.Any(), models.Cart{ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, models.CartProductUpdate{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}).Return(models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
			},
			args:    args{context.Background(), models.Cart{ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, models.CartProductUpdate{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}},
			want:    models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: false,
		},
		{
			name: "TestCartUsecase_AddProduct bad",
			prepare: func(f *fields) {
				f.repo.EXPECT().CheckCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, errors.New("dummyError"))
			},
			args:    args{context.Background(), models.Cart{ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, models.CartProductUpdate{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}},
			wantErr: true,
		},
		{
			name: "TestCartUsecase_AddProduct bad CheckCart",
			prepare: func(f *fields) {
				f.repo.EXPECT().CheckCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, repo.ErrCartNotFound)
				f.repo.EXPECT().CreateCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), nil)
				f.repo.EXPECT().AddProduct(gomock.Any(),
					models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), IsCurrent: true},
					models.CartProductUpdate{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}).
					Return(models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
			},
			args:    args{context.Background(), models.Cart{ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, models.CartProductUpdate{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}},
			want:    models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: false,
		},
		{
			name: "TestCartUsecase_AddProduct bad-baad",
			prepare: func(f *fields) {
				f.repo.EXPECT().CheckCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, repo.ErrCartNotFound)
				f.repo.EXPECT().CreateCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(uuid.UUID{}, errors.New("dummyError"))
			},
			args:    args{context.Background(), models.Cart{ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, models.CartProductUpdate{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}},
			want:    models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockCartRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewCartUsecase(f.repo)

			got, err := tt.uc.AddProduct(tt.args.ctx, tt.args.cart, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartUsecase.AddProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartUsecase.AddProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartUsecase_DeleteProduct(t *testing.T) {
	type fields struct {
		repo *mock.MockCartRepo
	}
	type args struct {
		ctx     context.Context
		cart    models.Cart
		product models.CartProductDelete
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *CartUsecase
		args    args
		want    models.Cart
		wantErr bool
	}{
		{
			name: "TestCartUsecase_DeleteProduct good",
			prepare: func(f *fields) {
				f.repo.EXPECT().CheckCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
				f.repo.EXPECT().DeleteProduct(gomock.Any(), models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, models.CartProductDelete{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}).Return(models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)

			},
			args:    args{context.Background(), models.Cart{ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, models.CartProductDelete{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}},
			want:    models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: false,
		},
		{
			name: "TestCartUsecase_DeleteProduct bad",
			prepare: func(f *fields) {
				f.repo.EXPECT().CheckCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, errors.New("dummyError"))

			},
			args:    args{context.Background(), models.Cart{ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, models.CartProductDelete{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockCartRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewCartUsecase(f.repo)

			got, err := tt.uc.DeleteProduct(tt.args.ctx, tt.args.cart, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartUsecase.DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartUsecase.DeleteProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartUsecase_UpdateCart(t *testing.T) {
	type fields struct {
		repo *mock.MockCartRepo
	}
	type args struct {
		ctx  context.Context
		cart models.Cart
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *CartUsecase
		args    args
		want    models.Cart
		wantErr bool
	}{
		{
			name: "TestCartUsecase_UpdateCart good",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
				f.repo.EXPECT().UpdateCart(gomock.Any(), models.Cart{ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}).
					Return(models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
			},
			args:    args{context.Background(), models.Cart{ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}},
			want:    models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: false,
		},
		{
			name: "TestCartUsecase_UpdateCart bad",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, errors.New("dummyError"))
			},
			args:    args{context.Background(), models.Cart{ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}},
			wantErr: true,
		},
		{
			name: "TestCartUsecase_UpdateCart bad-bad",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, repo.ErrCartNotFound)
				f.repo.EXPECT().CreateCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(uuid.UUID{}, errors.New("dummyError"))
			},
			args:    args{context.Background(), models.Cart{ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}},
			wantErr: true,
		},
		{
			name: "TestCartUsecase_UpdateCart bad-good",
			prepare: func(f *fields) {
				f.repo.EXPECT().ReadCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, repo.ErrCartNotFound)
				f.repo.EXPECT().CreateCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), nil)
				f.repo.EXPECT().UpdateCart(gomock.Any(),
					models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"),
						ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), IsCurrent: true}).
					Return(models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)

			},
			args:    args{context.Background(), models.Cart{ProfileId: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}},
			want:    models.Cart{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockCartRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewCartUsecase(f.repo)

			got, err := tt.uc.UpdateCart(tt.args.ctx, tt.args.cart)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartUsecase.UpdateCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartUsecase.UpdateCart() = %v, want %v", got, tt.want)
			}
		})
	}
}
