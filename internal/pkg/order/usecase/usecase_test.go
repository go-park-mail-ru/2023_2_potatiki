package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address"
	mockAddress "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart"
	mockCart "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order"
	mockOrder "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/mocks"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

func TestOrderUsecase_CreateOrder(t *testing.T) {
	type fields struct {
		repoOrder   *mockOrder.MockOrderRepo
		repoCart    *mockCart.MockCartRepo
		repoAddress *mockAddress.MockAddressRepo
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
				f.repoOrder.EXPECT().CreateOrder(gomock.Any(), models.Cart{}, uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), 1).Return(models.Order{}, nil)
				f.repoCart.EXPECT().DeleteCart(gomock.Any(), gomock.Any()).Return(nil)
			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")},
			want:    models.Order{Address: models.Address{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}},
			wantErr: false,
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
			tt.uc = NewOrderUsecase(f.repoOrder, f.repoCart, f.repoAddress)

			got, err := tt.uc.CreateOrder(tt.args.ctx, tt.args.userID)
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
			tt.uc = NewOrderUsecase(f.repoOrder, f.repoCart, f.repoAddress)

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
		repoOrder   order.OrderRepo
		repoCart    cart.CartRepo
		repoAddress address.AddressRepo
	}
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &OrderUsecase{
				repoOrder:   tt.fields.repoOrder,
				repoCart:    tt.fields.repoCart,
				repoAddress: tt.fields.repoAddress,
			}
			got, err := uc.GetOrders(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderUsecase.GetOrders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderUsecase.GetOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
