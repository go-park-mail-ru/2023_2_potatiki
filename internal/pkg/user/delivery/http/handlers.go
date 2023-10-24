package http

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/user"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/coockie"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/errcheck"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"io"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	log      *slog.Logger
	ucUser   user.UserUsecase
	ucAuther auth.AuthUsecase
}

func NewUserHandler(log *slog.Logger, ucUser user.UserUsecase, ucAuther auth.AuthUsecase) *UserHandler {
	return &UserHandler{
		log:      log,
		ucUser:   ucUser,
		ucAuther: ucAuther,
	}
}

func (h *UserHandler) UpdatePhoto(w http.ResponseWriter, r *http.Request) {
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
