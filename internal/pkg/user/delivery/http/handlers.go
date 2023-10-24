package http

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/user"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/cookie"
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

	body, err := io.ReadAll(r.Body)
	if resp.BodyErr(err, h.log, w) {
		return
	}
	h.log.Debug("request body decoded", slog.Any("request", r))

	ID, ok := r.Context().Value(cookie.AccessTokenCookieName).(uuid.UUID)
	if !ok {
		h.log.Error("failed cast uuid from context value")
		resp.JSONStatus(w, http.StatusUnauthorized)
	}

	h.log.Info("UpdatePhoto", "body", body, "ID", ID)
}

func (h *UserHandler) UpdateInfo(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	id, ok := r.Context().Value(cookie.AccessTokenCookieName).(uuid.UUID)
	if !ok {
		h.log.Error("failed cast uuid from context value")
		resp.JSONStatus(w, http.StatusUnauthorized)
	}

	body, err := io.ReadAll(r.Body)
	if resp.BodyErr(err, h.log, w) {
		return
	}

	profileInfo := &models.ProfileInfo{}
	err = json.Unmarshal(body, profileInfo)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	err = h.uc.UpdateInfo(r.Context(), id, *profileInfo)
	if err != nil {
		h.log.Error("failed to update user info", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}
	h.log.Info("updated user info")
	resp.JSONStatus(w, http.StatusOK)
}
