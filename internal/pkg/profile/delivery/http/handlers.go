package http

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/profile"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

const maxRequestBodySize = 1024 * 1024 * 5 // 5 MB

type ProfileHandler struct {
	log *slog.Logger
	uc  profile.ProfileUsecase
}

func NewProfileHandler(log *slog.Logger, uc profile.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{
		log: log,
		uc:  uc,
	}
}

// @Summary	GetProfile
// @Tags Profile
// @Description	Get user profile
// @Accept json
// @Produce json
// @Param id path string true "Profile UUID"
// @Success	200	{object} models.Profile "User profile"
// @Failure	400	{object} response.Response	"invalid request"
// @Failure	429
// @Router	/api/profile/{id} [get]
func (h *ProfileHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
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

// @Summary	UpdatePhoto
// @Tags Profile
// @Description	Update user photo
// @Accept json
// @Produce json
// @Param id path string true "Profile UUID"
// @Success	200	{object} models.Profile "User profile"
// @Failure	401
// @Failure 413
// @Failure	429
// @Router	/api/profile/update-photo/{id} [post]
func (h *ProfileHandler) UpdatePhoto(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)
	ID, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
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

	if err != nil && !errors.Is(err, io.EOF) {
		if errors.As(err, new(*http.MaxBytesError)) {
			resp.JSONStatus(w, http.StatusRequestEntityTooLarge)
		} else {
			resp.JSONStatus(w, http.StatusTooManyRequests)
		}

		return
	}

	h.log.Info("UpdatePhoto", "ID", ID)
	err = h.uc.UpdatePhoto(r.Context(), ID, bodyContent, fileFormat)
	if err != nil {
		h.log.Error("failed in uc.UpdatePhoto", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)
	}

	h.log.Info("updated profile info")
	resp.JSONStatus(w, http.StatusOK)
}

// @Summary	UpdateInfo
// @Tags Profile
// @Description	Update profile data
// @Accept json
// @Produce json
// @Param id path string true "Profile UUID"
// @Success	200	{object} models.Profile "User profile"
// @Failure	400	{object} response.Response	"error messege"
// @Failure	401
// @Failure	429
// @Router	/api/profile/update-info/{id} [get]
func (h *ProfileHandler) UpdateInfo(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)

	id, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
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

	profileInfo := &models.ProfileInfo{}
	err = json.Unmarshal(body, profileInfo)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	err = h.uc.UpdateInfo(r.Context(), id, profileInfo)
	if err != nil {
		h.log.Error("failed to update profile info", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Info("updated profile info")
	resp.JSONStatus(w, http.StatusOK)
}
