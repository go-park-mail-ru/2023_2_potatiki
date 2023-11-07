package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

func TestAddressUsecase_AddAddress(t *testing.T) {
	type args struct {
		ctx         context.Context //nolint:containedctx
		userID      uuid.UUID
		addressInfo models.AddressInfo
	}
	tests := []struct {
		name    string
		uc      *AddressUsecase
		args    args
		want    models.Address
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.AddAddress(tt.args.ctx, tt.args.userID, tt.args.addressInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressUsecase.AddAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddressUsecase.AddAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddressUsecase_UpdateAddress(t *testing.T) {
	type args struct {
		ctx         context.Context //nolint:containedctx
		addressInfo models.Address
	}
	tests := []struct {
		name    string
		uc      *AddressUsecase
		args    args
		want    models.Address
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.UpdateAddress(tt.args.ctx, tt.args.addressInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressUsecase.UpdateAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddressUsecase.UpdateAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddressUsecase_DeleteAddress(t *testing.T) {
	type args struct {
		ctx         context.Context //nolint:containedctx
		addressInfo models.AddressDelete
	}
	tests := []struct {
		name    string
		uc      *AddressUsecase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uc.DeleteAddress(tt.args.ctx, tt.args.addressInfo); (err != nil) != tt.wantErr {
				t.Errorf("AddressUsecase.DeleteAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddressUsecase_MakeCurrentAddress(t *testing.T) {
	type args struct {
		ctx         context.Context
		addressInfo models.AddressMakeCurrent
	}
	tests := []struct {
		name    string
		uc      *AddressUsecase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uc.MakeCurrentAddress(tt.args.ctx, tt.args.addressInfo); (err != nil) != tt.wantErr {
				t.Errorf("AddressUsecase.MakeCurrentAddress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddressUsecase_GetCurrentAddress(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		uc      *AddressUsecase
		args    args
		want    models.Address
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetCurrentAddress(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressUsecase.GetCurrentAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddressUsecase.GetCurrentAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddressUsecase_GetAllAddresses(t *testing.T) {
	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}
	tests := []struct {
		name    string
		uc      *AddressUsecase
		args    args
		want    []models.Address
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.uc.GetAllAddresses(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressUsecase.GetAllAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddressUsecase.GetAllAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
