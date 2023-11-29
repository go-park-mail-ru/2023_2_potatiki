package grpc

/*
import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/grpc/gen"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/mocks"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/go-playground/assert"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	//"github.com/gofrs/uuid"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockAuthUsecase(ctrl)
	uc.EXPECT().SignUp(gomock.Any(), &models.SignUpPayload{
		Login:    "User",
		Password: "Dima@gmail.com",
	}).Return(&models.Profile{}, "", time.Time{}, nil)

	AuthHandler := NewGrpcAuthHandler(uc, logger.Set("local", os.Stdout))
	AuthHandler.SignUp(gomock.Any(),
		&gen.SignUpRequest{
			Login:    "User",
			Password: "Dima@gmail.com",
			Phone:    "+7989123456879",
		})
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSignUpBad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockAuthUsecase(ctrl)

	t.Run("EmptyRequestBody", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "http://example.com/foo", nil)
		w := httptest.NewRecorder()
		AuthHandler := NewAuthHandler(logger.Set("local", os.Stdout), uc)
		AuthHandler.SignUp(w, req)
		assert.Equal(t, http.StatusTooManyRequests, w.Code)
	})
}

func TestSignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mock.NewMockAuthUsecase(ctrl)
	uc.EXPECT().SignIn(gomock.Any(), &models.SignInPayload{
		Login:    "Userrr",
		Password: "Dima@gmail.com",
	}).Return(&models.Profile{}, "", time.Time{}, nil)

	req := httptest.NewRequest(http.MethodPost, "http://example.com/foo",
		strings.NewReader("{ \"login\": \"Userrr\", \"password\": \"Dima@gmail.com\" }"))
	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(logger.Set("local", os.Stdout), uc)
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
		AuthHandler := NewAuthHandler(logger.Set("local", os.Stdout), uc)
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
	AuthHandler := NewAuthHandler(logger.Set("local", os.Stdout), uc)
	AuthHandler.LogOut(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := uuid.NewV4()

	uc := mock.NewMockAuthUsecase(ctrl)
	uc.EXPECT().CheckAuth(gomock.Any(), id).Return(&models.Profile{}, nil)
	req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)

	ctx := context.WithValue(req.Context(), authmw.AccessTokenCookieName, id)
	req = req.WithContext(ctx)

	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(logger.Set("local", os.Stdout), uc)
	AuthHandler.CheckAuth(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCheckAuthBad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("NoTokenCookie", func(t *testing.T) {
		uc := mock.NewMockAuthUsecase(ctrl)
		//uc.EXPECT().CheckAuth(gomock.Any(), uuid.NewV4()).Return(nil, errors.New("NoTokenCookie"))

		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
		w := httptest.NewRecorder()
		AuthHandler := NewAuthHandler(logger.Set("local", os.Stdout), uc)
		AuthHandler.CheckAuth(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("InvalidToken", func(t *testing.T) {
		uc := mock.NewMockAuthUsecase(ctrl)
		req := httptest.NewRequest(http.MethodGet, "http://example.com/foo", nil)
		req.AddCookie(&http.Cookie{
			Name:  authmw.AccessTokenCookieName,
			Value: "invalidTokenValue",
		})
		w := httptest.NewRecorder()
		AuthHandler := NewAuthHandler(logger.Set("local", os.Stdout), uc)
		AuthHandler.CheckAuth(w, req)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
*/
