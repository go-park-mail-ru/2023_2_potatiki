package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart"
	uuid "github.com/satori/go.uuid"
)

func TestCartUsecase_GetCart(t *testing.T) {
	type fields struct {
		repo cart.CartRepo
	}
	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &CartUsecase{
				repo: tt.fields.repo,
			}
			got, err := uc.GetCart(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartUsecase.GetCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartUsecase.GetCart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartUsecase_AddProduct(t *testing.T) {
	type fields struct {
		repo cart.CartRepo
	}
	type args struct {
		ctx     context.Context
		cart    models.Cart
		product models.CartProductUpdate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &CartUsecase{
				repo: tt.fields.repo,
			}
			got, err := uc.AddProduct(tt.args.ctx, tt.args.cart, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartUsecase.AddProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartUsecase.AddProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartUsecase_DeleteProduct(t *testing.T) {
	type fields struct {
		repo cart.CartRepo
	}
	type args struct {
		ctx     context.Context
		cart    models.Cart
		product models.CartProductDelete
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &CartUsecase{
				repo: tt.fields.repo,
			}
			got, err := uc.DeleteProduct(tt.args.ctx, tt.args.cart, tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartUsecase.DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartUsecase.DeleteProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartUsecase_UpdateCart(t *testing.T) {
	type fields struct {
		repo cart.CartRepo
	}
	type args struct {
		ctx  context.Context
		cart models.Cart
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Cart
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &CartUsecase{
				repo: tt.fields.repo,
			}
			got, err := uc.UpdateCart(tt.args.ctx, tt.args.cart)
			if (err != nil) != tt.wantErr {
				t.Errorf("CartUsecase.UpdateCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartUsecase.UpdateCart() = %v, want %v", got, tt.want)
			}
		})
	}
}
