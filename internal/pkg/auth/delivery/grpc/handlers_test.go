package grpc

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/grpc/gen"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/go-park-mail-ru/2023_2_potatiki/proto/gmodels"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

func TestGrpcAuthHandler_SignUp(t *testing.T) {
	type fields struct {
		uc *mock.MockAuthUsecase
	}
	type args struct {
		ctx context.Context
		in  *gen.SignUpRequest
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		h       *GrpcAuthHandler
		args    args
		want    *gen.SignUpResponse
		wantErr bool
	}{
		{
			name: "TestGrpcAuthHandler_SignUp good",
			prepare: func(f *fields) {
				f.uc.EXPECT().SignUp(gomock.Any(), &models.SignUpPayload{
					Login:    "User",
					Password: "Dima@gmail.com",
				}).Return(&models.Profile{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, "", time.Time{}, nil)
			},
			args: args{context.Background(), &gen.SignUpRequest{
				Login:    "User",
				Password: "Dima@gmail.com",
			}},
			want: &gen.SignUpResponse{
				Profile: &gmodels.Profile{Id: "fb11fe90-09bb-4e72-98a5-5ffba93aa39a"},
				Expires: time.Time{}.String()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				uc: mock.NewMockAuthUsecase(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.h = NewGrpcAuthHandler(f.uc, logger.Set("local", os.Stdout))

			got, err := tt.h.SignUp(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GrpcAuthHandler.SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GrpcAuthHandler.SignUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrpcAuthHandler_SignIn(t *testing.T) {
	type fields struct {
		uc *mock.MockAuthUsecase
	}
	type args struct {
		ctx context.Context
		in  *gen.SignInRequest
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		h       *GrpcAuthHandler
		args    args
		want    *gen.SignInResponse
		wantErr bool
	}{
		{
			name: "TestGrpcAuthHandler_SignIn good",
			prepare: func(f *fields) {
				f.uc.EXPECT().SignIn(gomock.Any(), &models.SignInPayload{
					Login:    "User",
					Password: "Dima@gmail.com",
				}).Return(&models.Profile{Id: uuid.FromStringOrNil("fb11fe90-09bb-4e72-98a5-5ffba93aa39a")}, "", time.Time{}, nil)
			},
			args: args{context.Background(), &gen.SignInRequest{
				Login:    "User",
				Password: "Dima@gmail.com",
			}},
			want: &gen.SignInResponse{
				Profile: &gmodels.Profile{Id: "fb11fe90-09bb-4e72-98a5-5ffba93aa39a"},
				Expires: time.Time{}.String()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				uc: mock.NewMockAuthUsecase(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.h = NewGrpcAuthHandler(f.uc, logger.Set("local", os.Stdout))

			got, err := tt.h.SignIn(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GrpcAuthHandler.SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GrpcAuthHandler.SignIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrpcAuthHandler_CheckAuth(t *testing.T) {
	id := uuid.NewV4()
	type fields struct {
		uc *mock.MockAuthUsecase
	}
	type args struct {
		ctx context.Context
		in  *gen.CheckAuthRequst
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		h       *GrpcAuthHandler
		args    args
		want    *gen.CheckAuthResponse
		wantErr bool
	}{
		{
			name: "TestGrpcAuthHandler_CheckAuth good",
			prepare: func(f *fields) {
				f.uc.EXPECT().CheckAuth(gomock.Any(), id).Return(&models.Profile{Id: id}, nil)
			},
			args: args{context.Background(), &gen.CheckAuthRequst{
				ID: id.String(),
			}},
			want: &gen.CheckAuthResponse{
				Profile: &gmodels.Profile{Id: id.String()},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				uc: mock.NewMockAuthUsecase(ctrl),
			}

			if tt.prepare != nil {
				tt.prepare(&f)
			}
			tt.h = NewGrpcAuthHandler(f.uc, logger.Set("local", os.Stdout))

			got, err := tt.h.CheckAuth(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("GrpcAuthHandler.CheckAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GrpcAuthHandler.CheckAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
