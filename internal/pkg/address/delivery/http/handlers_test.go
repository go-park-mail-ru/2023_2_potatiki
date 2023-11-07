package http

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestAddressHandler_AddAddress(t *testing.T) {
	testCases := []struct {
		name           string
		userID         uuid.UUID
		addressInfo    models.AddressInfo
		mockUsecaseFn  func(*mock.MockAddressUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
	}{
		{
			name:        "SuccessfulAddAddress",
			userID:      uuid.NewV4(),
			addressInfo: models.AddressInfo{},
			mockUsecaseFn: func(mockUsecase *mock.MockAddressUsecase) {
				mockUsecase.EXPECT().AddAddress(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Address{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
		{
			name:        "UnauthorizedUser",
			userID:      uuid.NewV4(),
			addressInfo: models.AddressInfo{},
			mockUsecaseFn: func(mockUsecase *mock.MockAddressUsecase) {
				mockUsecase.EXPECT().AddAddress(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Address{}, errors.New("internal server error"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockAddressUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodPost, "http://example.com/foo",
				strings.NewReader("{ \"login\": \"User\", \"password\": \"Dima@gmail.com\" }"))
			//url := "/api/address/add"
			//req := httptest.NewRequest("POST", url, nil)
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewAddressHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.AddAddress(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
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
