package http

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)

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

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)

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

func TestLogOut(t *testing.T) {
	ctrl := gomock.NewController(t)

	uc := mock.NewMockAuthUsecase(ctrl)

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(logger.Set("prod"), uc)
	AuthHandler.LogOut(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckAuth(t *testing.T) {
	ctrl := gomock.NewController(t)

	uc := mock.NewMockAuthUsecase(ctrl)
	uc.EXPECT().CheckToken(gomock.Any(), gomock.Any()).Return(uuid.New(), nil)

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
	cookie := &http.Cookie{
		Name:  AccessTokenCookieName,
		Value: "your_token_here",
	}
	req.AddCookie(cookie)

	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(logger.Set("prod"), uc)
	AuthHandler.CheckAuth(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)

	uc := mock.NewMockAuthUsecase(ctrl)
	idProfile := uuid.New()

	uc.EXPECT().GetProfile(gomock.Any(), idProfile).Return(models.Profile{}, nil)

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
	req = mux.SetURLVars(req, map[string]string{"id": idProfile.String()}) // Устанавливаем параметр "id" в URL.
	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(logger.Set("prod"), uc)
	AuthHandler.GetProfile(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
