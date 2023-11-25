package http

import (
	"log/slog"
	"net/http"

	"encoding/json"
	"io"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/survey"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type SurveyHandler struct {
	log *slog.Logger
	uc  survey.SurveyUsecase
}

func NewSurveyHandler(log *slog.Logger, uc survey.SurveyUsecase) *SurveyHandler {
	return &SurveyHandler{
		log: log,
		uc:  uc,
	}
}

func (h *SurveyHandler) SaveResponse(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)
	userID, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
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
	h.log.Debug("request body decoded", slog.Any("request", r))

	surveyInfo := models.SurveyResponse{
		UserID: userID,
	}
	err = json.Unmarshal(body, &surveyInfo)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	err = h.uc.SaveResponse(r.Context(), surveyInfo)
	if err != nil {
		h.log.Error("failed to save survey", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.SaveSurvey", "survey", surveyInfo)
	resp.JSONStatus(w, http.StatusOK)
}

func (h *SurveyHandler) GetSurvey(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)
	// count - обязателен
	userID, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
	if !ok {
		h.log.Error("failed cast uuid from context value")
		resp.JSONStatus(w, http.StatusUnauthorized)

		return
	}

	surveyIDString := r.URL.Query().Get("surveyID")
	surveyID, err := uuid.FromString(surveyIDString)
	if err != nil {
		h.log.Error("surveyID is invalid", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	survey, err := h.uc.GetSurvey(r.Context(), surveyID, userID)
	if err != nil {
		h.log.Error("failed to get survey", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("got survey", "len", len(survey.Questions))
	resp.JSON(w, http.StatusOK, survey)
}

func (h *SurveyHandler) GetStat(w http.ResponseWriter, r *http.Request) {
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
	surveyID, err := uuid.FromString(idStr)
	if err != nil {
		h.log.Error("surveyID is invali", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	statistics, err := h.uc.GetStat(r.Context(), surveyID)
	if err != nil {
		h.log.Error("failed to get answers", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("got survey", "len", len(statistics))
	resp.JSON(w, http.StatusOK, statistics)
}
