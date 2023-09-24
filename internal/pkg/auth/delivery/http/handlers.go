package http

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"io"
	"net/http"
)

type AuthHandler struct {
	useCase auth.AuthUsecase
}

func NewAuthHandler(newUseCase auth.AuthUsecase) AuthHandler {
	return AuthHandler{
		useCase: newUseCase,
	}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
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

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
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
