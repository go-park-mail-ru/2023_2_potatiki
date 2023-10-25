package http

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/user"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/cookie"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"github.com/google/uuid"
	"io"
	"log/slog"
	"net/http"
)

const maxRequestBodySize = 1024 * 1024 * 5 // 5 MB

type UserHandler struct {
	log *slog.Logger
	uc  user.UserUsecase
}

func NewUserHandler(uc user.UserUsecase, log *slog.Logger) *UserHandler {
	return &UserHandler{
		log: log,
		uc:  uc,
	}
}

func (h *UserHandler) UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)
	ID, ok := r.Context().Value(cookie.AccessTokenCookieName).(uuid.UUID)
	if !ok {
		h.log.Error("failed cast uuid from context value")
		resp.JSONStatus(w, http.StatusUnauthorized)
		return
	}

	limitedReader := http.MaxBytesReader(w, r.Body, maxRequestBodySize)
	defer r.Body.Close()

	bodyContent, err := io.ReadAll(limitedReader)
	fileFormat := http.DetectContentType(bodyContent)
	h.log.Debug("got []byte file", slog.Any("request", r))

	if err != nil && err != io.EOF {
		if errors.As(err, new(*http.MaxBytesError)) {
			resp.JSONStatus(w, http.StatusRequestEntityTooLarge)
		} else {
			resp.JSONStatus(w, http.StatusTooManyRequests)
		}

		return
	}

	h.log.Info("UpdatePhoto", "ID", ID)
	_ = h.uc.UpdatePhoto(r.Context(), ID, bodyContent, fileFormat)

}

func (h *UserHandler) UpdateInfo(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	id, ok := r.Context().Value(cookie.AccessTokenCookieName).(uuid.UUID)
	if !ok {
		h.log.Error("failed cast uuid from context value")
		resp.JSONStatus(w, http.StatusUnauthorized)

		return
	}

	body, err := io.ReadAll(r.Body)
	if resp.BodyErr(err, h.log, w) {
		return
	}
	defer r.Body.Close()
	h.log.Debug("got file from r.Body", slog.Any("request", r))

	profileInfo := models.ProfileInfo{}
	err = json.Unmarshal(body, &profileInfo)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	err = h.uc.UpdateInfo(r.Context(), id, profileInfo)
	if err != nil {
		h.log.Error("failed to update user info", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}
	h.log.Info("updated user info")
	resp.JSONStatus(w, http.StatusOK)
}
