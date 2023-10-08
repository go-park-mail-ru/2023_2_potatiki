package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockAuthUsecase(ctrl)
	uc.EXPECT().SignUp(gomock.Any(), models.User{
		Login:        "User",
		PasswordHash: "Dima@gmail.com",
	}).Return(models.Profile{}, "", time.Time{}, nil)

	req := httptest.NewRequest(http.MethodPost, "http://example.com/foo",
		strings.NewReader("{ \"login\": \"User\", \"password\": \"Dima@gmail.com\" }"))
	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(logger.Set("prod"), uc)
	AuthHandler.SignUp(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
func TestSignUpBad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockAuthUsecase(ctrl)

	t.Run("EmptyRequestBody", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "http://example.com/foo", nil)
		w := httptest.NewRecorder()
		AuthHandler := NewAuthHandler(logger.Set("prod"), uc)
		AuthHandler.SignUp(w, req)
		assert.Equal(t, http.StatusTooManyRequests, w.Code)
	})
}

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockAuthUsecase(ctrl)
	uc.EXPECT().SignIn(gomock.Any(), models.User{
		Login:        "User",
		PasswordHash: "Dima@gmail.com",
	}).Return(models.Profile{}, "", time.Time{}, nil)

	req := httptest.NewRequest(http.MethodPost, "http://example.com/foo",
		strings.NewReader("{ \"login\": \"User\", \"password\": \"Dima@gmail.com\" }"))
	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(logger.Set("prod"), uc)
	AuthHandler.SignIn(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSignInBad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockAuthUsecase(ctrl)

	t.Run("EmptyRequestBody", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "http://example.com/foo", nil)
		w := httptest.NewRecorder()
		AuthHandler := NewAuthHandler(logger.Set("prod"), uc)
		AuthHandler.SignIn(w, req)
		assert.Equal(t, http.StatusTooManyRequests, w.Code)
	})
}

func TestLogOut(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockAuthUsecase(ctrl)

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(logger.Set("prod"), uc)
	AuthHandler.LogOut(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockAuthUsecase(ctrl)
	uc.EXPECT().CheckToken(gomock.Any(), gomock.Any()).Return(uuid.New(), nil)

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
	cookie := &http.Cookie{
		Name:  AccessTokenCookieName,
		Value: "token",
	}
	req.AddCookie(cookie)

	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(logger.Set("prod"), uc)
	AuthHandler.CheckAuth(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckAuthBad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NoTokenCookie", func(t *testing.T) {
		uc := mock.NewMockAuthUsecase(ctrl)
		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
		w := httptest.NewRecorder()
		AuthHandler := NewAuthHandler(logger.Set("prod"), uc)
		AuthHandler.CheckAuth(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("InvalidToken", func(t *testing.T) {
		uc := mock.NewMockAuthUsecase(ctrl)
		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
		req.AddCookie(&http.Cookie{
			Name:  AccessTokenCookieName,
			Value: "invalidTokenValue",
		})
		uc.EXPECT().CheckToken(gomock.Any(), gomock.Any()).Return(uuid.UUID{}, errors.New("invalidTokenValue"))
		w := httptest.NewRecorder()
		AuthHandler := NewAuthHandler(logger.Set("prod"), uc)
		AuthHandler.CheckAuth(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestGetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockAuthUsecase(ctrl)
	idProfile := uuid.New()

	uc.EXPECT().GetProfile(gomock.Any(), idProfile).Return(models.Profile{}, nil)

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
	req = mux.SetURLVars(req, map[string]string{"id": idProfile.String()})
	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(logger.Set("prod"), uc)
	AuthHandler.GetProfile(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProfileBad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("EmptyID", func(t *testing.T) {
		uc := mock.NewMockAuthUsecase(ctrl)
		idProfile := uuid.New()
		uc.EXPECT().GetProfile(gomock.Any(), idProfile).Return(models.Profile{}, errors.New("invalidProfile"))

		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
		req = mux.SetURLVars(req, map[string]string{"id": idProfile.String()})
		w := httptest.NewRecorder()
		AuthHandler := NewAuthHandler(logger.Set("prod"), uc)
		AuthHandler.GetProfile(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("InvalidID", func(t *testing.T) {
		uc := mock.NewMockAuthUsecase(ctrl)

		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
		req = mux.SetURLVars(req, map[string]string{"id": "invalidID"})
		w := httptest.NewRecorder()
		AuthHandler := NewAuthHandler(logger.Set("prod"), uc)
		AuthHandler.GetProfile(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
