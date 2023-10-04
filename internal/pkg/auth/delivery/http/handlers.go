package http

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		if errors.Is(err, io.EOF) {
			h.log.Error("request body is empty")
			resp.JSON(w, http.StatusBadRequest, resp.Err("request body is empty"))
			return
		}
		h.log.Error("failed to decode request body", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Nil())
		return
	}
	h.log.Debug("request body decoded", slog.Any("request", r))

	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSON(w, http.StatusTooManyRequests, resp.Nil())
		return
	}

	profile, token, exp, err := h.uc.SignIn(r.Context(), *u)

	if err != nil {
		h.log.Error("failed to signin", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid login or password"))
		return
	}

	h.log.Debug("got profile", slog.Any("profile", profile.Id))

	http.SetCookie(w, getTokenCookie(AccessTokenCookieName, token, exp))
	resp.JSON(w, http.StatusOK, profile)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		if errors.Is(err, io.EOF) {
			h.log.Error("request body is empty")
			resp.JSON(w, http.StatusBadRequest, resp.Err("request body is empty"))
			return
		}
		h.log.Error("failed to decode request body", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Nil())
		return
	}
	h.log.Info("request body decoded", slog.Any("request", r))

	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSON(w, http.StatusTooManyRequests, resp.Nil())
		return
	}

	profile, token, exp, err := h.uc.SignUp(r.Context(), *u)
	if err != nil {
		h.log.Error("failed to signup", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid login or password"))
		return
	}

	http.SetCookie(w, getTokenCookie(AccessTokenCookieName, token, exp))
	h.log.Debug("got profile", slog.Any("profile", profile.Id))
	resp.JSON(w, http.StatusOK, profile)
}

func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, getTokenCookie(AccessTokenCookieName, "", time.Now().UTC().AddDate(0, 0, -1)))
	h.log.Info("logout")
	resp.JSON(w, http.StatusOK, resp.Nil())
}

func (h *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	tokenCookie, err := r.Cookie(AccessTokenCookieName)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			h.log.Error("token cookie not found", sl.Err(err))
			resp.JSON(w, http.StatusUnauthorized, resp.Nil())
			return
		default:
			h.log.Error("faild to get token cookie", sl.Err(err))
			resp.JSON(w, http.StatusUnauthorized, resp.Nil())
			return
		}
	}

	id, err := h.uc.CheckToken(r.Context(), tokenCookie.Value)
	if err != nil {
		h.log.Error("jws token is invalid", sl.Err(err))
		resp.JSON(w, http.StatusUnauthorized, resp.Nil())
		return
	}
	h.log.Info("got profile id", slog.Any("profile id", id))
	resp.JSON(w, http.StatusOK, resp.Nil())
}

func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok || idStr == "" {
		h.log.Error("id is empty")
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
		return
	}
	idProfile, err := uuid.Parse(idStr)
	if err != nil {
		h.log.Error("id is invalid", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
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
