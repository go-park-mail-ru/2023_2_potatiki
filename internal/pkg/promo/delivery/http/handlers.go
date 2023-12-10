package http

import (
	"errors"
	"net/http"

	"log/slog"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/promo"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"
)

type PromoHandler struct {
	log *slog.Logger
	uc  promo.PromoUsecase
}

func NewPromoHandler(log *slog.Logger, uc promo.PromoUsecase) PromoHandler {
	return PromoHandler{
		log: log,
		uc:  uc,
	}
}

// @Summary	Promo
// @Tags Promo
// @Description	Check Promocode
// @Accept json
// @Produce json
// @Param name query string true "promocode name, example: SALE23"
// @Success	200	{object} models.Promocode "Promocode model"
// @Failure	400	{object} responser.response	"error message"
// @Failure	404	"something not found error message"
// @Failure	429
// @Router	/api/promo/check [get]
func (h *PromoHandler) CheckPromocode(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)

	name := r.URL.Query().Get("name")
	if name == "" {
		h.log.Warn("name is invalid", "error", "invalid request")
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
	}

	promocode, err := h.uc.CheckPromocode(r.Context(), name)
	if err != nil {
		if errors.Is(err, promo.ErrPromocodeNotFound) {
			h.log.Debug("promocode not found", sl.Err(err))
			resp.JSONStatus(w, http.StatusNotFound)
			return
		}

		h.log.Error("failed to check promocode", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)
		return
	}

	h.log.Debug("got promocode", "promocode", promocode)
	resp.JSON(w, http.StatusOK, promocode)
}

// @Summary	Promo
// @Tags Promo
// @Description	Use Promocode
// @Accept json
// @Produce json
// @Param name query string true "promocode name, example: SALE23"
// @Success	200	{object} models.Promocode "Promocode model"
// @Failure	400	{object} responser.response	"error message"
// @Failure	404	"something not found error message"
// @Failure	429
// @Router	/api/promo/use [get]
func (h *PromoHandler) UsePromocode(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)

	name := r.URL.Query().Get("name")
	if name == "" {
		h.log.Warn("name is invalid", "error", "invalid request")
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
	}

	promocode, err := h.uc.UsePromocode(r.Context(), name)
	if err != nil {
		if errors.Is(err, promo.ErrPromocodeNotFound) {
			h.log.Debug("promocode not found", sl.Err(err))
			resp.JSONStatus(w, http.StatusNotFound)
			return
		}

		h.log.Error("failed to use promocode", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)
		return
	}

	h.log.Debug("got promocode", "promocode", promocode)
	resp.JSON(w, http.StatusOK, promocode)
}
