package http

import (
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"github.com/google/uuid"
)

type OrderHandler struct {
	log *slog.Logger
	uc  order.OrderUsecase
}

func NewOrderHandler(log *slog.Logger, uc order.OrderUsecase) OrderHandler {
	return OrderHandler{
		log: log,
		uc:  uc,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
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

	order, err := h.uc.CreateOrder(r.Context(), userID)
	if err != nil {
		h.log.Error("failed to get cart", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.CreateOrder", "order", order)
	resp.JSON(w, http.StatusOK, order)
}

//func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
//	h.log = h.log.With(
//		slog.String("op", sl.GFN()),
//		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
//	)
//
//	userID, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
//	if !ok {
//		h.log.Error("failed cast uuid from context value")
//		resp.JSONStatus(w, http.StatusUnauthorized)
//
//		return
//	}
//
//	order, err := h.uc.GetOrder(r.Context(), userID)
//	if err != nil {
//		h.log.Error("failed to get cart", sl.Err(err))
//		resp.JSONStatus(w, http.StatusTooManyRequests)
//
//		return
//	}
//
//	h.log.Debug("h.uc.GetOrder", "order", order)
//	resp.JSON(w, http.StatusOK, order)
//}
