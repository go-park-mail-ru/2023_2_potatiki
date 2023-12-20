package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mockAddress "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address/mocks"
	addressRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address/repo"
	mockCart "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/mocks"
	cartRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/repo"
	mockOrder "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/mocks"

	mockPromo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/promo/mocks"

	orderRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/repo"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

func TestOrderUsecase_CreateOrder(t *testing.T) {
	t.Skip()
	type fields struct {
		repoOrder   *mockOrder.MockOrderRepo
		repoCart    *mockCart.MockCartRepo
		repoAddress *mockAddress.MockAddressRepo
		repoPromo   *mockPromo.MockPromoRepo
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *OrderUsecase
		args    args
		want    models.Order
		wantErr bool
	}{
		{
			name: "TestOrderUsecase_CreateOrder good",
			prepare: func(f *fields) {
				f.repoAddress.EXPECT().ReadCurrentAddress(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Address{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
				f.repoCart.EXPECT().ReadCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, nil)
				f.repoOrder.EXPECT().CreateOrder(gomock.Any(), models.Cart{}, uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), int64(1), "", "").Return(models.Order{}, nil)
				f.repoCart.EXPECT().DeleteCart(gomock.Any(), gomock.Any()).Return(nil)
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			want:    models.Order{Address: models.Address{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}},
			wantErr: false,
		},
		{
			name: "TestOrderUsecase_CreateOrder bad ReadCurrentAddress",
			prepare: func(f *fields) {
				f.repoAddress.EXPECT().ReadCurrentAddress(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Address{}, errors.New("dummyErrors"))
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
		{
			name: "TestOrderUsecase_CreateOrder bad ReadCart",
			prepare: func(f *fields) {
				f.repoAddress.EXPECT().ReadCurrentAddress(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Address{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
				f.repoCart.EXPECT().ReadCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, errors.New("dummyError"))
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
		{
			name: "TestOrderUsecase_CreateOrder bad CreateOrder",
			prepare: func(f *fields) {
				f.repoAddress.EXPECT().ReadCurrentAddress(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Address{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
				f.repoCart.EXPECT().ReadCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, nil)
				f.repoOrder.EXPECT().CreateOrder(gomock.Any(), models.Cart{}, uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), int64(1), "", "").Return(models.Order{}, errors.New("dummyError"))
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
		{
			name: "TestOrderUsecase_CreateOrder bad DeleteCart",
			prepare: func(f *fields) {
				f.repoAddress.EXPECT().ReadCurrentAddress(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Address{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
				f.repoCart.EXPECT().ReadCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, nil)
				f.repoOrder.EXPECT().CreateOrder(gomock.Any(), models.Cart{}, uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), int64(1), "", "").Return(models.Order{}, nil)
				f.repoCart.EXPECT().DeleteCart(gomock.Any(), gomock.Any()).Return(errors.New("dummyError"))
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
		{
			name: "TestOrderUsecase_CreateOrder bad ReadCurrentAddress",
			prepare: func(f *fields) {
				f.repoAddress.EXPECT().ReadCurrentAddress(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Address{}, addressRepo.ErrAddressNotFound)
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
		{
			name: "TestOrderUsecase_CreateOrder bad ReadCart",
			prepare: func(f *fields) {
				f.repoAddress.EXPECT().ReadCurrentAddress(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Address{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
				f.repoCart.EXPECT().ReadCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, cartRepo.ErrCartNotFound)
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
		{
			name: "TestOrderUsecase_CreateOrder bad CreateOrder",
			prepare: func(f *fields) {
				f.repoAddress.EXPECT().ReadCurrentAddress(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Address{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
				f.repoCart.EXPECT().ReadCart(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Cart{}, nil)
				f.repoOrder.EXPECT().CreateOrder(gomock.Any(), models.Cart{}, uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), int64(1), "", "").Return(models.Order{}, orderRepo.ErrPoductNotFound)
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
				repoOrder:   mockOrder.NewMockOrderRepo(ctrl),
				repoCart:    mockCart.NewMockCartRepo(ctrl),
				repoAddress: mockAddress.NewMockAddressRepo(ctrl),
				repoPromo:   mockPromo.NewMockPromoRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewOrderUsecase(f.repoOrder, f.repoCart, f.repoAddress, f.repoPromo)

			got, err := tt.uc.CreateOrder(tt.args.ctx, tt.args.userID, "", "", "")
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderUsecase.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderUsecase.CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderUsecase_GetCurrentOrder(t *testing.T) {
	type fields struct {
		repoOrder   *mockOrder.MockOrderRepo
		repoCart    *mockCart.MockCartRepo
		repoAddress *mockAddress.MockAddressRepo
		repoPromo   *mockPromo.MockPromoRepo
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *OrderUsecase
		args    args
		want    models.Order
		wantErr bool
	}{
		{
			name: "TestOrderUsecase_GetCurrentOrder good",
			prepare: func(f *fields) {
				f.repoOrder.EXPECT().ReadOrderID(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), nil)
				f.repoOrder.EXPECT().ReadOrder(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Order{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			want:    models.Order{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: false,
		},
		{
			name: "TestOrderUsecase_GetCurrentOrder bad ReadOrderID",
			prepare: func(f *fields) {
				f.repoOrder.EXPECT().ReadOrderID(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), errors.New("dummyError"))
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
		{
			name: "TestOrderUsecase_GetCurrentOrder bad ReadOrder",
			prepare: func(f *fields) {
				f.repoOrder.EXPECT().ReadOrderID(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), nil)
				f.repoOrder.EXPECT().ReadOrder(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Order{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, errors.New("dummyError"))
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
		{
			name: "TestOrderUsecase_GetCurrentOrder bad ReadOrderID",
			prepare: func(f *fields) {
				f.repoOrder.EXPECT().ReadOrderID(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), orderRepo.ErrOrderNotFound)
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
		{
			name: "TestOrderUsecase_GetCurrentOrder bad ReadOrder",
			prepare: func(f *fields) {
				f.repoOrder.EXPECT().ReadOrderID(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), nil)
				f.repoOrder.EXPECT().ReadOrder(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Order{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, orderRepo.ErrPoductsInOrderNotFound)
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
				repoOrder:   mockOrder.NewMockOrderRepo(ctrl),
				repoCart:    mockCart.NewMockCartRepo(ctrl),
				repoAddress: mockAddress.NewMockAddressRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewOrderUsecase(f.repoOrder, f.repoCart, f.repoAddress, f.repoPromo)

			got, err := tt.uc.GetCurrentOrder(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderUsecase.GetCurrentOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderUsecase.GetCurrentOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderUsecase_GetOrders(t *testing.T) {
	type fields struct {
		repoOrder   *mockOrder.MockOrderRepo
		repoCart    *mockCart.MockCartRepo
		repoAddress *mockAddress.MockAddressRepo
		repoPromo   *mockPromo.MockPromoRepo
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *OrderUsecase
		args    args
		want    []models.Order
		wantErr bool
	}{
		{
			name: "TestOrderUsecase_GetOrders good",
			prepare: func(f *fields) {
				f.repoOrder.EXPECT().ReadOrdersID(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return([]uuid.UUID{uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
				f.repoOrder.EXPECT().ReadOrder(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Order{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			want:    []models.Order{{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}},
			wantErr: false,
		},
		{
			name: "TestOrderUsecase_GetOrders bad ReadOrdersID",
			prepare: func(f *fields) {
				f.repoOrder.EXPECT().ReadOrdersID(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return([]uuid.UUID{}, errors.New("dummyError"))
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
		{
			name: "TestOrderUsecase_GetOrders bad ReadOrder",
			prepare: func(f *fields) {
				f.repoOrder.EXPECT().ReadOrdersID(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return([]uuid.UUID{uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
				f.repoOrder.EXPECT().ReadOrder(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Order{}, errors.New("dummyError"))
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
		{
			name: "TestOrderUsecase_GetOrders bad ReadOrdersID",
			prepare: func(f *fields) {
				f.repoOrder.EXPECT().ReadOrdersID(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return([]uuid.UUID{}, orderRepo.ErrOrdersNotFound)
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			wantErr: true,
		},
		{
			name: "TestOrderUsecase_GetOrders bad ReadOrder",
			prepare: func(f *fields) {
				f.repoOrder.EXPECT().ReadOrdersID(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return([]uuid.UUID{uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, nil)
				f.repoOrder.EXPECT().ReadOrder(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")).Return(models.Order{}, orderRepo.ErrPoductsInOrderNotFound)
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
				repoOrder:   mockOrder.NewMockOrderRepo(ctrl),
				repoCart:    mockCart.NewMockCartRepo(ctrl),
				repoAddress: mockAddress.NewMockAddressRepo(ctrl),
				repoPromo:   mockPromo.NewMockPromoRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewOrderUsecase(f.repoOrder, f.repoCart, f.repoAddress, f.repoPromo)

			got, err := tt.uc.GetOrders(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderUsecase.GetOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderUsecase.GetOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
