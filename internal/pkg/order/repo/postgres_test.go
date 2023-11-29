package repo

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

func TestOrderRepo_CreateOrder(t *testing.T) {
	type args struct {
		ctx       context.Context
		cart      models.Cart
		addressID uuid.UUID
		userID    uuid.UUID
		statusID  int64
	}
	tests := []struct {
		name    string
		r       *OrderRepo
		args    args
		want    models.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.CreateOrder(tt.args.ctx, tt.args.cart, tt.args.addressID, tt.args.userID, tt.args.statusID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderRepo.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderRepo.CreateOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderRepo_ReadOrderID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		r       *OrderRepo
		args    args
		want    uuid.UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadOrderID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderRepo.ReadOrderID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderRepo.ReadOrderID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderRepo_ReadOrder(t *testing.T) {
	type args struct {
		ctx     context.Context
		orderID uuid.UUID
	}
	tests := []struct {
		name    string
		r       *OrderRepo
		args    args
		want    models.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadOrder(tt.args.ctx, tt.args.orderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderRepo.ReadOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderRepo.ReadOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderRepo_ReadOrdersID(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		r       *OrderRepo
		args    args
		want    []uuid.UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadOrdersID(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderRepo.ReadOrdersID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderRepo.ReadOrdersID() = %v, want %v", got, tt.want)
			}
		})
	}
}
