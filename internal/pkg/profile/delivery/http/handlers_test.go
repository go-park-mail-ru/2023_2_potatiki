package http

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/profile/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestProfileHandler_UpdateProfileData(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockProfileUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
		jsonString     string
	}{
		{
			name: "SuccessfulUpdateData",
			mockUsecaseFn: func(mockUsecase *mock.MockProfileUsecase) {
				mockUsecase.EXPECT().UpdateData(gomock.Any(), gomock.Any(), gomock.Any()).Return(&models.Profile{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"passwords": {
						"oldPass": "*********",
						"newPass": "**********"
					},
				 "phone": "+123456789"
			}`,
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockProfileUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
			jsonString: `{
				"passwords": {
						"oldPass": "*********",
						"newPass": "**********"
					},
				 "phone": "+123456789"
			}`,
		},
		{
			name:           "EmptyBody",
			mockUsecaseFn:  func(mockUsecase *mock.MockProfileUsecase) {},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: ``,
		},
		{
			name: "UnsuccessfulAddAddress",
			mockUsecaseFn: func(mockUsecase *mock.MockProfileUsecase) {
				mockUsecase.EXPECT().UpdateData(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&models.Profile{}, errors.New("error in add profile"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			jsonString: `{
				"passwords": {
						"oldPass": "*********",
						"newPass": "**********"
					},
				 "phone": "+123456789"
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockProfileUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodPost, "http://zuzu-market.ru/api/profile/update-data",
				strings.NewReader(tc.jsonString))
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewProfileHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.UpdateProfileData(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestProfileHandler_UpdatePhoto(t *testing.T) {
	testCases := []struct {
		name           string
		mockUsecaseFn  func(*mock.MockProfileUsecase)
		expectedStatus int
		funcCtxUser    func(context.Context) context.Context
		imageBytes     []byte
	}{
		{
			name: "SuccessfulUpdateData",
			mockUsecaseFn: func(mockUsecase *mock.MockProfileUsecase) {
				mockUsecase.EXPECT().UpdatePhoto(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&models.Profile{}, nil)
			},
			expectedStatus: http.StatusOK,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			imageBytes: []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46},
		},
		{
			name:           "UnauthorizedUser",
			mockUsecaseFn:  func(mockUsecase *mock.MockProfileUsecase) {},
			expectedStatus: http.StatusUnauthorized,
			funcCtxUser: func(ctx context.Context) context.Context {
				return context.Background()
			},
			imageBytes: []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46},
		},
		{
			name: "UnsuccessfulUpdatePhoto",
			mockUsecaseFn: func(mockUsecase *mock.MockProfileUsecase) {
				mockUsecase.EXPECT().UpdatePhoto(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Return(&models.Profile{}, errors.New("error in add photo"))
			},
			expectedStatus: http.StatusTooManyRequests,
			funcCtxUser: func(ctx context.Context) context.Context {
				id := uuid.NewV4()
				return context.WithValue(ctx, "zuzu-t", id)
			},
			imageBytes: []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUsecase := mock.NewMockProfileUsecase(ctrl)
			tc.mockUsecaseFn(mockUsecase)

			req := httptest.NewRequest(http.MethodPost, "http://zuzu-market.ru/api/profile/update-photo",
				bytes.NewReader(tc.imageBytes))
			w := httptest.NewRecorder()
			ctx := tc.funcCtxUser(req.Context())

			req = req.WithContext(ctx)
			addressHandler := NewProfileHandler(logger.Set("local", os.Stdout), mockUsecase)
			addressHandler.UpdatePhoto(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

func TestGetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockProfileUsecase(ctrl)
	idProfile := uuid.NewV4()

	uc.EXPECT().GetProfile(gomock.Any(), idProfile).Return(&models.Profile{}, nil)

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
	req = mux.SetURLVars(req, map[string]string{"id": idProfile.String()})
	w := httptest.NewRecorder()
	ProfileHandler := NewProfileHandler(logger.Set("prod", os.Stdout), uc)
	ProfileHandler.GetProfile(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProfileBad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("EmptyID", func(t *testing.T) {
		uc := mock.NewMockProfileUsecase(ctrl)
		idProfile := uuid.NewV4()
		uc.EXPECT().GetProfile(gomock.Any(), idProfile).Return(&models.Profile{}, errors.New("invalidProfile"))

		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
		req = mux.SetURLVars(req, map[string]string{"id": idProfile.String()})
		w := httptest.NewRecorder()
		ProfileHandler := NewProfileHandler(logger.Set("prod", os.Stdout), uc)
		ProfileHandler.GetProfile(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("InvalidID", func(t *testing.T) {
		uc := mock.NewMockProfileUsecase(ctrl)

		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "invalidID"})
		w := httptest.NewRecorder()
		UserHandler := NewProfileHandler(logger.Set("prod", os.Stdout), uc)
		UserHandler.GetProfile(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
