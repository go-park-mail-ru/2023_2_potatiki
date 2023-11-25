package http

import (
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/survey"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"
	uuid "github.com/satori/go.uuid"
	"log/slog"
	"net/http"
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

//func (h *SurveyHandler) SaveSurvey(w http.ResponseWriter, r *http.Request) {
//	h.log = h.log.With(
//		slog.String("op", sl.GFN()),
//		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
//	)
//	userID, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
//	if !ok {
//		h.log.Error("failed cast uuid from context value")
//		resp.JSONStatus(w, http.StatusUnauthorized)
//
//		return
//	}
//
//	body, err := io.ReadAll(r.Body)
//	if resp.BodyErr(err, h.log, w) {
//		return
//	}
//	defer r.Body.Close()
//	h.log.Debug("request body decoded", slog.Any("request", r))
//
//	surveyInfo := models.Survey{
//		UserID: userID,
//	}
//	err = json.Unmarshal(body, &surveyInfo)
//	if err != nil {
//		h.log.Error("failed to unmarshal request body", sl.Err(err))
//		resp.JSONStatus(w, http.StatusTooManyRequests)
//
//		return
//	}
//
//	err = h.uc.SaveSurvey(r.Context(), surveyInfo)
//	if err != nil {
//		h.log.Error("failed to save survey", sl.Err(err))
//		resp.JSONStatus(w, http.StatusTooManyRequests)
//
//		return
//	}
//
//	h.log.Debug("h.uc.SaveSurvey", "survey", surveyInfo)
//	resp.JSONStatus(w, http.StatusOK)
//}

func (h *SurveyHandler) GetSurvey(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)
	// count - обязателен

	surveyIDString := r.URL.Query().Get("surveyID")
	surveyID, err := uuid.FromString(surveyIDString)
	if err != nil {
		h.log.Error("surveyID is invalid", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	survey, err := h.uc.GetSurvey(r.Context(), surveyID)
	if err != nil {
		h.log.Error("failed to get survey", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("got survey", "len", len(survey.Questions))
	resp.JSON(w, http.StatusOK, survey)
}
