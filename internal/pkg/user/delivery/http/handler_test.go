package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/user/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockUserUsecase(ctrl)
	idProfile := uuid.New()

	uc.EXPECT().GetProfile(gomock.Any(), idProfile).Return(models.Profile{}, nil)

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
	req = mux.SetURLVars(req, map[string]string{"id": idProfile.String()})
	w := httptest.NewRecorder()
	UserHandler := NewUserHandler(logger.Set("prod", os.Stdout), uc)
	UserHandler.GetProfile(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProfileBad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("EmptyID", func(t *testing.T) {
		uc := mock.NewMockUserUsecase(ctrl)
		idProfile := uuid.New()
		uc.EXPECT().GetProfile(gomock.Any(), idProfile).Return(models.Profile{}, errors.New("invalidProfile"))

		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
		req = mux.SetURLVars(req, map[string]string{"id": idProfile.String()})
		w := httptest.NewRecorder()
		UserHandler := NewUserHandler(logger.Set("prod", os.Stdout), uc)
		UserHandler.GetProfile(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("InvalidID", func(t *testing.T) {
		uc := mock.NewMockUserUsecase(ctrl)

		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "invalidID"})
		w := httptest.NewRecorder()
		UserHandler := NewUserHandler(logger.Set("prod", os.Stdout), uc)
		UserHandler.GetProfile(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
