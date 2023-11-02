package http

import (
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"
	uuid "github.com/satori/go.uuid"
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

// @Summary	CreateOrder
// @Tags Order
// @Description	Create Order using profile ID from cookies
// @Accept json
// @Produce json
// @Success	200	{object} models.Order "New order info"
// @Failure	401	"User unauthorized"
// @Failure	429
// @Router	/api/order/create [post]
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

// @Summary	GetCurrentOrder
// @Tags Order
// @Description	Get Current Order using profile ID from cookies
// @Accept json
// @Produce json
// @Success	200	{object} models.Order "Current order info"
// @Failure	401	"User unauthorized"
// @Failure	429
// @Router	/api/order/get_current [get]
func (h *OrderHandler) GetCurrentOrder(w http.ResponseWriter, r *http.Request) {
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

	order, err := h.uc.GetCurrentOrder(r.Context(), userID)
	if err != nil {
		h.log.Error("failed to get cart", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.GetCurrentOrder", "order", order)
	resp.JSON(w, http.StatusOK, order)
}

// @Summary	GetOrders
// @Tags Order
// @Description	Get all Orders using profile ID from cookies
// @Accept json
// @Produce json
// @Success	200	{array} models.Order "All orders info"
// @Failure	401	"User unauthorized"
// @Failure	429
// @Router	/api/order/get_all [get]
func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
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

	orders, err := h.uc.GetOrders(r.Context(), userID)
	if err != nil {
		h.log.Error("failed to get cart", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.GetOrders", "orders", orders)
	resp.JSON(w, http.StatusOK, orders)
}
