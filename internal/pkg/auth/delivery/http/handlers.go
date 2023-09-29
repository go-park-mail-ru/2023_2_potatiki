package http

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"

	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
)

type AuthHandler struct {
	usecase auth.AuthUsecase
}

func NewAuthHandler(usecase auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		usecase: usecase,
	}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if errors.Is(err, io.EOF) {
		//log.Error("request body is empty")
		resp.JSON(w, http.StatusBadRequest, resp.Err("request body is empty"))
		return
	}
	if err != nil {
		//log.Error("failed to decode request body", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
		return
	}
	//log.Info("request body decoded", slog.Any("request", req))

	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
		return
	}

	profile, err := h.usecase.SignIn(r.Context(), *u)
	if err != nil {
		resp.JSON(w, http.StatusBadRequest, resp.Err("internal error"))
		return
	}
	//log.Info("got profile", slog.String("profile", profile.Id))
	resp.JSON(w, http.StatusOK, profile)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		resp.JSON(w, http.StatusBadRequest, resp.Err("failed to io.ReadAll(r.Body)"))
		return
	}
	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		resp.JSON(w, http.StatusBadRequest, resp.Err("failed to json.Unmarshal(body, u)"))
		return
	}

	profile, err := h.usecase.SignUp(r.Context(), *u)
	if err != nil {
		resp.JSON(w, http.StatusBadRequest, resp.Err("failed in SignUp"))
		return
	}
	resp.JSON(w, http.StatusOK, profile)
}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	resp.JSON(w, http.StatusOK, resp.OK())
}
