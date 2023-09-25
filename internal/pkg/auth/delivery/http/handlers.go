package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
)

type AuthHandler struct {
	usecase *auth.AuthUsecase
}

func New(usecase auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		usecase: &usecase,
	}
}

func (h *AuthHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		return
	}
}

func (h *AuthHandler) SingUp(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		return
	}
}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {

}
