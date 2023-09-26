package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"

	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
)

type AuthHandler struct {
	usecase auth.AuthUsecase
}

func New(usecase auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		usecase: usecase,
	}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		resp.Status(w, http.StatusBadRequest, resp.Err("failed to io.ReadAll(r.Body)"))
		return
	}

	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		resp.Status(w, http.StatusBadRequest, resp.Err("failed to json.Unmarshal(body, u)"))
		return
	}

	profile, err := h.usecase.SignIn(r.Context(), *u)
	if err != nil {
		resp.Status(w, http.StatusBadRequest, resp.Err("failed in SignUp"))
		return
	}
	fmt.Println(profile)

	resp.Status(w, http.StatusOK, resp.OK())
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		resp.Status(w, http.StatusBadRequest, resp.Err("failed to io.ReadAll(r.Body)"))
		return
	}
	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		resp.Status(w, http.StatusBadRequest, resp.Err("failed to json.Unmarshal(body, u)"))
		return
	}

	profile, err := h.usecase.SignUp(r.Context(), *u)
	if err != nil {
		resp.Status(w, http.StatusBadRequest, resp.Err("failed in SignUp"))
		return
	}
	fmt.Println(profile)
	resp.Status(w, http.StatusOK, resp.OK())
}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	resp.Status(w, http.StatusOK, resp.OK())
}
