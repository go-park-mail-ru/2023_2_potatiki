package http

import (
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address"
	"log/slog"
	"net/http"
	"reflect"
	"testing"
)

func TestAddressHandler_AddAddress(t *testing.T) {
	type fields struct {
		log *slog.Logger
		uc  address.AddressUsecase
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
			h := &AddressHandler{
				log: tt.fields.log,
				uc:  tt.fields.uc,
			}
			h.AddAddress(tt.args.w, tt.args.r)
		})
	}
}

func TestAddressHandler_DeleteAddress(t *testing.T) {
	type fields struct {
		log *slog.Logger
		uc  address.AddressUsecase
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
			h := &AddressHandler{
				log: tt.fields.log,
				uc:  tt.fields.uc,
			}
			h.DeleteAddress(tt.args.w, tt.args.r)
		})
	}
}

func TestAddressHandler_GetAllAddresses(t *testing.T) {
	type fields struct {
		log *slog.Logger
		uc  address.AddressUsecase
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
			h := &AddressHandler{
				log: tt.fields.log,
				uc:  tt.fields.uc,
			}
			h.GetAllAddresses(tt.args.w, tt.args.r)
		})
	}
}

func TestAddressHandler_GetCurrentAddress(t *testing.T) {
	type fields struct {
		log *slog.Logger
		uc  address.AddressUsecase
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
			h := &AddressHandler{
				log: tt.fields.log,
				uc:  tt.fields.uc,
			}
			h.GetCurrentAddress(tt.args.w, tt.args.r)
		})
	}
}

func TestAddressHandler_MakeCurrentAddress(t *testing.T) {
	type fields struct {
		log *slog.Logger
		uc  address.AddressUsecase
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
			h := &AddressHandler{
				log: tt.fields.log,
				uc:  tt.fields.uc,
			}
			h.MakeCurrentAddress(tt.args.w, tt.args.r)
		})
	}
}

func TestAddressHandler_UpdateAddress(t *testing.T) {
	type fields struct {
		log *slog.Logger
		uc  address.AddressUsecase
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
			h := &AddressHandler{
				log: tt.fields.log,
				uc:  tt.fields.uc,
			}
			h.UpdateAddress(tt.args.w, tt.args.r)
		})
	}
}

func TestNewAddressHandler(t *testing.T) {
	type args struct {
		log *slog.Logger
		uc  address.AddressUsecase
	}
	tests := []struct {
		name string
		args args
		want AddressHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAddressHandler(tt.args.log, tt.args.uc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAddressHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
