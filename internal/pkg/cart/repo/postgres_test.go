package repo

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

func TestCartRepo_CreateCart(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		r       *CartRepo
		args    args
		want    uuid.UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.CreateCart(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartRepo.CreateCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartRepo.CreateCart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartRepo_DeleteCart(t *testing.T) {
	type args struct {
		ctx    context.Context
		cartID uuid.UUID
	}
	tests := []struct {
		name    string
		r       *CartRepo
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.DeleteCart(tt.args.ctx, tt.args.cartID); (err != nil) != tt.wantErr {
				t.Errorf("CartRepo.DeleteCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCartRepo_CheckCart(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		r       *CartRepo
		args    args
		want    models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.CheckCart(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartRepo.CheckCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartRepo.CheckCart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartRepo_ReadCart(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		r       *CartRepo
		args    args
		want    models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadCart(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartRepo.ReadCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartRepo.ReadCart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartRepo_UpdateCart(t *testing.T) {
	type args struct {
		ctx  context.Context
		cart models.Cart
	}
	tests := []struct {
		name    string
		r       *CartRepo
		args    args
		want    models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.UpdateCart(tt.args.ctx, tt.args.cart)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartRepo.UpdateCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartRepo.UpdateCart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartRepo_ReadCartProducts(t *testing.T) {
	type args struct {
		ctx  context.Context
		cart models.Cart
	}
	tests := []struct {
		name    string
		r       *CartRepo
		args    args
		want    models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.ReadCartProducts(tt.args.ctx, tt.args.cart)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartRepo.ReadCartProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartRepo.ReadCartProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartRepo_AddProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		cart    models.Cart
		product models.CartProductUpdate
	}
	tests := []struct {
		name    string
		r       *CartRepo
		args    args
		want    models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.AddProduct(tt.args.ctx, tt.args.cart, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartRepo.AddProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartRepo.AddProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartRepo_DeleteProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		cart    models.Cart
		product models.CartProductDelete
	}
	tests := []struct {
		name    string
		r       *CartRepo
		args    args
		want    models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.DeleteProduct(tt.args.ctx, tt.args.cart, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartRepo.DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartRepo.DeleteProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
