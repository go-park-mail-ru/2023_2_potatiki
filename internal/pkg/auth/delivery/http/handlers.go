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
		// Так нужно обрабатывать ошибки в handler-aх
		// TODO : переписать resp.JSON на свою функцию
		resp.JSON(w, r, resp.Error("failed to io.ReadAll(r.Body)"))

		//Так нужно отвечать что всё хорошо
		//resp.JSON(w, r, resp.OK())
		return
	}
	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		resp.JSON(w, r, resp.Error("failed to decode request"))
		return
	}

	profile, err := h.usecase.SignUp(r.Context(), u)
	if err != nil {
		resp.JSON(w, r, resp.Error("PLOHO"))
		return
	}
	fmt.Println(profile)

}

func (h *AuthHandler) Kek(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Single page:", r.URL.String())
	resp.JSON(w, r, resp.OK())
}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {

}
