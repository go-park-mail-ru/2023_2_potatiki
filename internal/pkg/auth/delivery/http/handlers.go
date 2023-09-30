package http

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"log/slog"
	"net/http"
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
	if err != nil {
		if errors.Is(err, io.EOF) {
			h.log.Error("request body is empty")
			resp.JSON(w, http.StatusBadRequest, resp.Err("request body is empty"))
			return
		}
		h.log.Error("failed to decode request body", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, nil)
		return
	}
	h.log.Debug("request body decoded", slog.Any("request", r))

	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, nil)
		return
	}

	profile, err := h.uc.SignIn(r.Context(), *u)

	if err != nil {
		h.log.Error("failed to signin", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid login or password"))
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
		resp.JSON(w, http.StatusBadRequest, nil)
		return
	}
	h.log.Info("request body decoded", slog.Any("request", r))

	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, nil)
		return
	}

	profile, err := h.uc.SignUp(r.Context(), *u)
	if err != nil {
		h.log.Error("failed to signup", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid login or password"))
		return
	}
	resp.JSON(w, http.StatusOK, profile)
}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	resp.JSON(w, http.StatusOK, nil)
}

func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok || idStr == "" {
		h.log.Error("id is empty")
		resp.JSON(w, http.StatusAccepted, resp.Err("invalid request"))
		return
	}
	idProfile, err := uuid.Parse(idStr)
	if err != nil {
		h.log.Error("id is invalid", sl.Err(err))
		resp.JSON(w, http.StatusAccepted, resp.Err("invalid request"))
		return
	}

	profile, err := h.uc.GetProfile(r.Context(), idProfile)

	if err != nil {
		h.log.Error("failed to signup", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid uuid"))
		return
	}

	h.log.Debug("got profile", slog.Any("profile", profile.Id))
	resp.JSON(w, http.StatusOK, profile)
}
