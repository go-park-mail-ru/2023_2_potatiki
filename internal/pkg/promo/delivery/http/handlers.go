package http

import (
	"errors"
	"net/http"

	"log/slog"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/promo"
	uuid "github.com/satori/go.uuid"

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
// @Failure	401	"user Unauthorized"
// @Failure	404	"promocode not found"
// @Failure	403	"promocode not leftout"
// @Failure	419	"promocode expired"
// @Failure	410	"promocode already used"
// @Failure	429 "internal error"
// @Router	/api/promo/check [get]
func (h *PromoHandler) CheckPromocode(w http.ResponseWriter, r *http.Request) {
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

	name := r.URL.Query().Get("name")
	if name == "" {
		h.log.Warn("name is invalid", "error", "invalid request")
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
	}

	promocode, err := h.uc.CheckPromocode(r.Context(), userID, name)
	if err != nil {
		switch {
		case errors.Is(err, promo.ErrPromocodeNotFound):
			h.log.Debug("promocode not found", sl.Err(err))
			resp.JSONStatus(w, http.StatusNotFound)
		case errors.Is(err, promo.ErrPromocodeExpired):
			h.log.Debug("promocode expired", sl.Err(err))
			resp.JSONStatus(w, 419) //http code doesn't implement in lib
		case errors.Is(err, promo.ErrPromocodeLeftout):
			h.log.Debug("promocode not leftout", sl.Err(err))
			resp.JSONStatus(w, http.StatusForbidden)
		case errors.Is(err, promo.ErrAlreadyUsed):
			h.log.Debug("promocode already used", sl.Err(err))
			resp.JSONStatus(w, http.StatusGone)
		default:
			h.log.Error("failed to check promocode", sl.Err(err))
			resp.JSONStatus(w, http.StatusTooManyRequests)
		}

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
