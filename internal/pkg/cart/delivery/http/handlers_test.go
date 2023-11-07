package http

import (
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart"
	"log/slog"
	"net/http"
	"testing"
)

func TestCartHandler_AddProduct(t *testing.T) {
	type fields struct {
		log *slog.Logger
		uc  cart.CartUsecase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &CartHandler{
				log: tt.fields.log,
				uc:  tt.fields.uc,
			}
			h.AddProduct(tt.args.w, tt.args.r)
		})
	}
}

func TestCartHandler_DeleteProduct(t *testing.T) {
	type fields struct {
		log *slog.Logger
		uc  cart.CartUsecase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &CartHandler{
				log: tt.fields.log,
				uc:  tt.fields.uc,
			}
			h.DeleteProduct(tt.args.w, tt.args.r)
		})
	}
}

func TestCartHandler_GetCart(t *testing.T) {
	type fields struct {
		log *slog.Logger
		uc  cart.CartUsecase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &CartHandler{
				log: tt.fields.log,
				uc:  tt.fields.uc,
			}
			h.GetCart(tt.args.w, tt.args.r)
		})
	}
}

func TestCartHandler_UpdateCart(t *testing.T) {
	type fields struct {
		log *slog.Logger
		uc  cart.CartUsecase
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &CartHandler{
				log: tt.fields.log,
				uc:  tt.fields.uc,
			}
			h.UpdateCart(tt.args.w, tt.args.r)
		})
	}
}
