package http

import (
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	mock "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)

	uc := mock.NewMockAuthUsecase(ctrl)
	uc.EXPECT().SignUp(gomock.Any(), models.User{
		Login:        "User",
		PasswordHash: "Dima@gmail.com",
	}).Return(models.Profile{}, nil)

	req := httptest.NewRequest("POST", "http://example.com/foo",
		strings.NewReader("{ \"login\": \"User\", \"password\": \"Dima@gmail.com\" }"))
	w := httptest.NewRecorder()
	AuthHandler := NewAuthHandler(uc)
	AuthHandler.SignUp(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
