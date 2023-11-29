package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/grpc/gen"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/go-park-mail-ru/2023_2_potatiki/proto/gmodels"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockAuthClient(ctrl)
	client.EXPECT().SignUp(gomock.Any(), &gen.SignUpRequest{
		Login:    "User",
		Password: "Dima@gmail.com",
		Phone:    "+7989123456879",
	}).Return(&gen.SignUpResponse{
		Profile: &gmodels.Profile{Id: "fb11fe90-09bb-4e72-98a5-5ffba93aa39a"},
		Expires: "2006-01-02 15:04:05.999999999 -0700 UTC"}, nil)

	req := httptest.NewRequest(http.MethodPost, "http://example.com/foo",
		strings.NewReader("{ \"login\": \"User\", \"password\": \"Dima@gmail.com\", \"phone\": \"+7989123456879\" }"))
	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(client, logger.Set("local", os.Stdout))
	AuthHandler.SignUp(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSignUpBad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockAuthClient(ctrl)

	t.Run("EmptyRequestBody", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "http://example.com/foo", nil)
		w := httptest.NewRecorder()
		AuthHandler := NewAuthHandler(client, logger.Set("local", os.Stdout))
		AuthHandler.SignUp(w, req)
		assert.Equal(t, http.StatusTooManyRequests, w.Code)
	})
}

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockAuthClient(ctrl)
	client.EXPECT().SignIn(gomock.Any(), &gen.SignInRequest{
		Login:    "User",
		Password: "Dima@gmail.com",
	}).Return(&gen.SignInResponse{
		Profile: &gmodels.Profile{Id: "fb11fe90-09bb-4e72-98a5-5ffba93aa39a"},
		Expires: "2006-01-02 15:04:05.999999999 -0700 UTC"}, nil)

	req := httptest.NewRequest(http.MethodPost, "http://example.com/foo",
		strings.NewReader("{ \"login\": \"User\", \"password\": \"Dima@gmail.com\" }"))
	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(client, logger.Set("local", os.Stdout))
	AuthHandler.SignIn(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSignInBad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockAuthClient(ctrl)

	t.Run("EmptyRequestBody", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "http://example.com/foo", nil)
		w := httptest.NewRecorder()
		AuthHandler := NewAuthHandler(client, logger.Set("local", os.Stdout))
		AuthHandler.SignIn(w, req)
		assert.Equal(t, http.StatusTooManyRequests, w.Code)
	})
}

func TestLogOut(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock.NewMockAuthClient(ctrl)

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(client, logger.Set("local", os.Stdout))
	AuthHandler.LogOut(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := uuid.NewV4()

	client := mock.NewMockAuthClient(ctrl)
	client.EXPECT().CheckAuth(gomock.Any(), &gen.CheckAuthRequst{
		ID: id.String(),
	}).Return(&gen.CheckAuthResponse{
		Profile: &gmodels.Profile{Id: id.String()},
	}, nil)

	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)

	ctx := context.WithValue(req.Context(), authmw.AccessTokenCookieName, id)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(client, logger.Set("local", os.Stdout))
	AuthHandler.CheckAuth(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckAuthBad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//id := uuid.NewV4()

	t.Run("NoTokenCookie", func(t *testing.T) {
		client := mock.NewMockAuthClient(ctrl)

		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)

		w := httptest.NewRecorder()
		AuthHandler := NewAuthHandler(client, logger.Set("local", os.Stdout))
		AuthHandler.CheckAuth(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("InvalidToken", func(t *testing.T) {
		client := mock.NewMockAuthClient(ctrl)

		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
		req.AddCookie(&http.Cookie{
			Name:  authmw.AccessTokenCookieName,
			Value: "invalidTokenValue",
		})
		w := httptest.NewRecorder()
		AuthHandler := NewAuthHandler(client, logger.Set("local", os.Stdout))
		AuthHandler.CheckAuth(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
