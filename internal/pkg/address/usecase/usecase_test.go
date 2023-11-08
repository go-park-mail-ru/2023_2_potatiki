package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address/mocks"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

func TestAddressUsecase_AddAddress(t *testing.T) {

	type args struct {
		ctx         context.Context //nolint:containedctx
		userID      uuid.UUID
		addressInfo models.AddressPayload
	}
	type fields struct {
		repo *mock.MockAddressRepo
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *AddressUsecase
		args    args
		want    models.Address
		wantErr bool
	}{
		{
			name: "TestAddressUsecase_AddAddress good",
			prepare: func(f *fields) {
				f.repo.EXPECT().CreateAddress(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), models.AddressPayload{}).Return(models.Address{}, nil)

			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), models.AddressPayload{}},
			want:    models.Address{},
			wantErr: false,
		},
		{
			name: "TestAddressUsecase_AddAddress bad",
			prepare: func(f *fields) {
				f.repo.EXPECT().CreateAddress(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), models.AddressPayload{}).Return(models.Address{}, errors.New("dummyError"))

			},
			args:    args{context.Background(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), models.AddressPayload{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockAddressRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewAddressUsecase(f.repo)

			got, err := tt.uc.AddAddress(tt.args.ctx, tt.args.userID, tt.args.addressInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressUsecase.AddAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
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
	type fields struct {
		repo *mock.MockAddressRepo
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *AddressUsecase
		args    args
		want    models.Address
		wantErr bool
	}{
		/*{
			name: "TestAddressUsecase_UpdateAddress good",
			prepare: func(f *fields) {
				f.repo.EXPECT().CreateAddress(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), models.AddressPayload{}).Return(models.Address{}, nil)

			},
			args:    args{context.Background(), models.Address{}},
			want:    models.Address{},
			wantErr: false,
		},
		{
			name: "TestAddressUsecase_UpdateAddress bad",
			prepare: func(f *fields) {
				f.repo.EXPECT().CreateAddress(gomock.Any(), uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a"), models.AddressPayload{}).Return(models.Address{}, errors.New("dummyError"))

			},
			args:    args{context.Background(), models.Address{}},
			wantErr: true,
		},*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockAddressRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewAddressUsecase(f.repo)

			got, err := tt.uc.UpdateAddress(tt.args.ctx, tt.args.addressInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressUsecase.UpdateAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
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
	type fields struct {
		repo *mock.MockAddressRepo
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *AddressUsecase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockAddressRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewAddressUsecase(f.repo)

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
	type fields struct {
		repo *mock.MockAddressRepo
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *AddressUsecase
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockAddressRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewAddressUsecase(f.repo)

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
	type fields struct {
		repo *mock.MockAddressRepo
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *AddressUsecase
		args    args
		want    models.Address
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockAddressRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewAddressUsecase(f.repo)

			got, err := tt.uc.GetCurrentAddress(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressUsecase.GetCurrentAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
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
	type fields struct {
		repo *mock.MockAddressRepo
	}
	tests := []struct {
		name    string
		fields  fields
		prepare func(f *fields)
		uc      *AddressUsecase
		args    args
		want    []models.Address
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				repo: mock.NewMockAddressRepo(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.uc = NewAddressUsecase(f.repo)

			got, err := tt.uc.GetAllAddresses(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddressUsecase.GetAllAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddressUsecase.GetAllAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
