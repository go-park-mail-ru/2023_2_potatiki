package http

import (
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/recommendations"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"
	uuid "github.com/satori/go.uuid"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

type RecommendationsHandler struct {
	log *slog.Logger
	uc  recommendations.RecommendationsUsecase
}

func NewRecommendationsHandler(log *slog.Logger, uc recommendations.RecommendationsUsecase) RecommendationsHandler {
	return RecommendationsHandler{
		log: log,
		uc:  uc,
	}
}

// @Summary	AnonRecommendations
// @Tags Recommendations
// @Description	Get recommendations for product
// @Accept json
// @Produce json
// @Param id query string true "Product UUID"
// @Param category_id query int true "Category id"
// @Success	200	{object} []models.ProductSlice "Products Slice"
// @Failure	400	{object} responser.response	"error message"
// @Failure	429
// @Router	/api/recommendations/get_anon [get]
func (h *RecommendationsHandler) AnonRecommendations(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)

	idStr := r.URL.Query().Get("id")
	if len(idStr) == 0 {
		h.log.Error("id is empty")
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	productID, err := uuid.FromString(idStr)
	if err != nil {
		h.log.Error("id is invalid", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	idStr = r.URL.Query().Get("category_id")
	if len(idStr) == 0 {
		h.log.Error("id is empty")
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	categoryID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.log.Error("category_id is invalid", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	recommendations, err := h.uc.AnonRecommendations(r.Context(), productID, categoryID)
	if err != nil {
		h.log.Error("failed to get recommendations", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.AnonRecommendations", "recommendations", "recommendations")
	resp.JSON(w, http.StatusOK, &recommendations)
}

// @Summary	Recommendations
// @Tags Recommendations
// @Description	Get recommendations
// @Accept json
// @Produce json
// @Param id query string true "Product UUID"
// @Param category_id query int true "Category id"
// @Success	200	{object} []models.ProductSlice "Products Slice"
// @Failure	400	{object} responser.response	"error message"
// @Failure	429
// @Router	/api/recommendations/get_all [get]
func (h *RecommendationsHandler) Recommendations(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)

	idStr := r.URL.Query().Get("id")
	if len(idStr) == 0 {
		h.log.Error("id is empty")
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	productID, err := uuid.FromString(idStr)
	if err != nil {
		h.log.Error("id is invalid", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	idStr = r.URL.Query().Get("category_id")
	if len(idStr) == 0 {
		h.log.Error("id is empty")
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	categoryID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.log.Error("category_id is invalid", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	userID, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
	if !ok {
		h.log.Error("failed cast uuid from context value")
		resp.JSONStatus(w, http.StatusUnauthorized)

		return
	}

	recommendations, err := h.uc.Recommendations(r.Context(), userID, productID, categoryID)
	if err != nil {
		h.log.Error("failed to get recommendations", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.Recommendations", "recommendations", "recommendations")
	resp.JSON(w, http.StatusOK, &recommendations)
}

// @Summary	UpdateUserActivity
// @Tags Recommendations
// @Description	Update user activities
// @Accept json
// @Produce json
// @Param input body models.UserActivity true "User activity info"
// @Success	200
// @Failure	401	"User unauthorized"
// @Failure	400	{object} responser.response	"error message"
// @Failure	429
// @Router	/api/recommendations/update [post]
func (h *RecommendationsHandler) UpdateUserActivity(w http.ResponseWriter, r *http.Request) {
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

	userActivity := models.UserActivity{}

	err = userActivity.UnmarshalJSON(body)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusBadRequest)

		return
	}

	err = h.uc.UpdateUserActivity(r.Context(), userID, userActivity)

	if err != nil {
		h.log.Error("failed to update user activity", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.UpdateUserActivity", "update success", "recommendations")
	resp.JSONStatus(w, http.StatusOK)
}
