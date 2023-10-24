package http

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/user"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/coockie"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/cookie"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/errcheck"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"github.com/google/uuid"
)

type UserHandler struct {
	log *slog.Logger
	uc  user.UserUsecase
}

func NewUserHandler(log *slog.Logger, uc user.UserUsecase) *UserHandler {
	return &UserHandler{
		log: log,
		uc:  uc,
	}
}

func (h *UserHandler) UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)
	val := r.Context().Value(cookie.AccessTokenCookieName)
	if ID, err := val.(uuid.UUID); !ok {
		h.log.Error("failed cast uuid from context value")
	}
	_ = ID

	_, err = io.ReadAll(r.Body)
	if errcheck.BodyErr(err, h.log, w) {
		return
	}
	h.log.Debug("request body decoded", slog.Any("request", r))
}

func (h *UserHandler) UpdateInfo(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	tokenCookie, err := r.Cookie(coockie.AccessTokenCookieName)
	if errcheck.TokenCookieErr(err, h.log, w) {
		return
	}

	id, err := h.ucAuther.CheckToken(r.Context(), tokenCookie.Value)
	if err != nil {
		h.log.Error("jws token is invalid", sl.Err(err))
		resp.JSONStatus(w, http.StatusUnauthorized)

		return
	}
	h.log.Info("got profile id", slog.Any("profile id", id))

	body, err := io.ReadAll(r.Body)
	if errcheck.BodyErr(err, h.log, w) {
		return
	}
	h.log.Debug("request body decoded", slog.Any("request", r))

	profileInfo := &models.ProfileInfo{}
	err = json.Unmarshal(body, profileInfo)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	err = h.ucUser.UpdateInfo(r.Context(), id, *profileInfo)
	if err != nil {
		h.log.Error("failed to update user info", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}
	h.log.Info("updated user info")

	resp.JSONStatus(w, http.StatusOK)
}
