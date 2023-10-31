package http

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"github.com/google/uuid"
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
// @Description	Login to Account
// @Accept json
// @Produce json
// @Param input body models.SignInPayload true "SignInPayload"
// @Success	200	{object} models.Profile "Profile"
// @Failure	400	{object} response.Response	"error messege"
// @Failure	429
// @Router	/api/auth/signin [post]
func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)

	body, err := io.ReadAll(r.Body)
	if resp.BodyErr(err, h.log, w) {
		return
	}
	h.log.Debug("request body decoded", slog.Any("request", r))
	defer r.Body.Close()

	u := &models.SignInPayload{}
	err = json.Unmarshal(body, u)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	profile, token, exp, err := h.uc.SignIn(r.Context(), u)

	if err != nil {
		h.log.Error("failed to signin", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid login or password"))

		return
	}

	h.log.Debug("got profile", slog.Any("profile", profile.Id))

	http.SetCookie(w, authmw.MakeTokenCookie(token, exp))
	resp.JSON(w, http.StatusOK, profile)
}

// @Summary	SignUp
// @Tags Auth
// @Description	Create Account
// @Accept json
// @Produce json
// @Param input body models.SignUpPayload true "SignUpPayload"
// @Success	200 {object} models.Profile "Profile"
// @Failure	400	{object} response.Response	"error messege"
// @Failure	429
// @Router	/api/auth/signup [post]
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)

	body, err := io.ReadAll(r.Body)
	if resp.BodyErr(err, h.log, w) {
		return
	}
	h.log.Debug("request body decoded", slog.Any("request", r))

	u := &models.SignUpPayload{}
	err = json.Unmarshal(body, u)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	profile, token, exp, err := h.uc.SignUp(r.Context(), u)
	if err != nil {
		h.log.Error("failed to signup", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid login or password"))

		return
	}

	http.SetCookie(w, authmw.MakeTokenCookie(token, exp))
	h.log.Debug("got profile", slog.Any("profile", profile.Id))
	resp.JSON(w, http.StatusOK, profile)
}

// @Summary	Logout
// @Tags Auth
// @Description	Logout from Account
// @Accept json
// @Produce json
// @Success	200
// @Failure	401
// @Router	/api/auth/logout [get]
func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, authmw.MakeTokenCookie("", time.Now().UTC().AddDate(0, 0, -1)))
	h.log.Info("logout")
	resp.JSONStatus(w, http.StatusOK)
}

// @Summary	CheckAuth
// @Tags Auth
// @Description	Check user is logged in
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
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)

	id, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
	if !ok {
		h.log.Error("failed cast uuid from context value")
		resp.JSONStatus(w, http.StatusUnauthorized)
	}

	h.log.Info("check auth success", "id", id)
	resp.JSONStatus(w, http.StatusOK)
}
