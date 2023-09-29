package http

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/repo"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
)

type AuthHandler struct {
	log *slog.Logger
	uc  auth.AuthUsecase
}

func NewAuthHandler(log *slog.Logger, uc auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		log: log,
		uc:  uc,
	}
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if errors.Is(err, io.EOF) {
		h.log.Error("request body is empty")
		resp.JSON(w, http.StatusBadRequest, resp.Err("request body is empty"))
		return
	}
	if err != nil {
		h.log.Error("failed to decode request body", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
		return
	}
	h.log.Debug("request body decoded", slog.Any("request", r))

	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
		return
	}

	profile, err := h.uc.SignIn(r.Context(), *u)

	if err != nil {
		if errors.Is(err, repo.ErrInvalidPass) {
			h.log.Error("failed to signin", sl.Err(err))
			resp.JSON(w, http.StatusBadRequest, resp.Err("invalid login or password"))
		}
		h.log.Error("failed to signin", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("internal error"))
		return
	}

	h.log.Debug("got profile", slog.Any("profile", profile.Id))
	resp.JSON(w, http.StatusOK, profile)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		if errors.Is(err, io.EOF) {
			h.log.Error("request body is empty")
			resp.JSON(w, http.StatusBadRequest, resp.Err("request body is empty"))
			return
		}
		h.log.Error("failed to decode request body", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
		return
	}
	h.log.Info("request body decoded", slog.Any("request", r))

	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
		return
	}

	profile, err := h.uc.SignUp(r.Context(), *u)
	if err != nil {
		h.log.Error("failed to signup", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("internal error"))
		return
	}
	resp.JSON(w, http.StatusOK, profile)
}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	resp.JSON(w, http.StatusOK, resp.OK())
}

func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if errors.Is(err, io.EOF) {
		h.log.Error("request body is empty")
		resp.JSON(w, http.StatusBadRequest, resp.Err("request body is empty"))
		return
	}
	if err != nil {
		h.log.Error("failed to decode request body", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
		return
	}
	h.log.Debug("request body decoded", slog.Any("request", r))

	ud := &models.UserId{}
	err = json.Unmarshal(body, ud)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
		return
	}

	profile, err := h.uc.GetProfile(r.Context(), ud.Id)

	if err != nil {
		if errors.Is(err, repo.ErrInvalidPass) {
			h.log.Error("failed to get profile", sl.Err(err))
			resp.JSON(w, http.StatusBadRequest, resp.Err("invalid uuid"))
		}
		h.log.Error("failed to get profile", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("internal error"))
		return
	}

	h.log.Debug("got profile", slog.Any("profile", profile.Id))
	resp.JSON(w, http.StatusOK, profile)
}
