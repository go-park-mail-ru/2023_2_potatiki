package http

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/coockie"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/errcheck"
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

// @Summary	SignIn
// @Tags Auth
// @Description	LogIn to Account
// @Accept json
// @Produce json
// @Param input body models.User true "user info"
// @Success	200	{object} models.Profile "User profile"
// @Failure	400	{object} response.Response	"request body is empty"
// @Failure	429
// @Router	/api/auth/signin [post]
func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	body, err := io.ReadAll(r.Body)
	if errcheck.BodyErr(err, h.log, w) {
		return
	}
	h.log.Debug("request body decoded", slog.Any("request", r))

	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	profile, token, exp, err := h.uc.SignIn(r.Context(), *u)

	if err != nil {
		h.log.Error("failed to signin", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid login or password"))

		return
	}

	h.log.Debug("got profile", slog.Any("profile", profile.Id))

	http.SetCookie(w, coockie.GetTokenCookie(coockie.AccessTokenCookieName, token, exp))
	resp.JSON(w, http.StatusOK, profile)
}

// @Summary	SignUp
// @Tags Auth
// @Description	Create Account
// @Accept json
// @Produce json
// @Param input body models.User true "user info"
// @Success	200 {object} models.Profile "User profile"
// @Failure	400	{object} response.Response	"request body is empty"
// @Failure	429
// @Router	/api/auth/signup [post]
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	body, err := io.ReadAll(r.Body)
	if errcheck.BodyErr(err, h.log, w) {
		return
	}
	h.log.Debug("request body decoded", slog.Any("request", r))

	u := &models.User{}
	err = json.Unmarshal(body, u)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	profile, token, exp, err := h.uc.SignUp(r.Context(), *u)
	if err != nil {
		h.log.Error("failed to signup", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid login or password"))

		return
	}

	http.SetCookie(w, coockie.GetTokenCookie(coockie.AccessTokenCookieName, token, exp))
	h.log.Debug("got profile", slog.Any("profile", profile.Id))
	resp.JSON(w, http.StatusOK, profile)
}

// @Summary	Logout
// @Tags Auth
// @Description	Logout from Account
// @Accept json
// @Produce json
// @Success	200
// @Router	/api/auth/logout [get]
func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, coockie.GetTokenCookie(coockie.AccessTokenCookieName, "", time.Now().UTC().AddDate(0, 0, -1)))
	h.log.Info("logout")
	resp.JSONStatus(w, http.StatusOK)
}

// @Summary	CheckAuth
// @Tags Auth
// @Description	Check user auth
// @Accept json
// @Produce json
// @Success	200
// @Failure	401
// @security AuthKey
// @Param Cookie header string  false "Token" default(zuzu-t=xxx)
// @Router	/api/auth/check_auth [get]
func (h *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	tokenCookie, err := r.Cookie(coockie.AccessTokenCookieName)
	if errcheck.TokenCookieErr(err, h.log, w) {
		return
	}

	id, err := h.uc.CheckToken(r.Context(), tokenCookie.Value)
	if err != nil {
		h.log.Error("jws token is invalid", sl.Err(err))
		resp.JSONStatus(w, http.StatusUnauthorized)

		return
	}
	h.log.Info("got profile id", slog.Any("profile id", id))

	resp.JSONStatus(w, http.StatusOK)
}

// @Summary	GetProfile
// @Tags Auth
// @Description	Get user profile
// @Accept json
// @Produce json
// @Param id path string true "Profile UUID"
// @Success	200	{object} models.Profile "User profile"
// @Failure	400	{object} response.Response	"invalid request"
// @Failure	429
// @Router	/api/auth/{id} [get]
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
