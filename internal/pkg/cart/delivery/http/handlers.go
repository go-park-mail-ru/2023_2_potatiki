package http

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/cookie"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"github.com/google/uuid"
)

type CartHandler struct {
	log *slog.Logger
	uc  cart.CartUsecase
}

func NewCartHandler(log *slog.Logger, uc cart.CartUsecase) CartHandler {
	return CartHandler{
		log: log,
		uc:  uc,
	}
}

// @Summary	UpdateCart
// @Tags Cart
// @Description	Update cart
// @Accept json
// @Produce json
// @Param input body models.Cart true "cart info"
// @Success	200	{object} models.Cart "cart info"
// @Failure	400	{object} response.Response	"invalid request"
// @Failure	429
// @Router	/api/cart/update [post]
func (h *CartHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)
	userID, ok := r.Context().Value(cookie.AccessTokenCookieName).(uuid.UUID)
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

	c := models.Cart{}
	err = json.Unmarshal(body, &c)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}
	c.ProfileId = userID

	cart, err := h.uc.UpdateCart(r.Context(), c)
	if err != nil {
		h.log.Error("failed to update cart", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("update cart")
	resp.JSON(w, http.StatusOK, cart)
}

// @Summary	GetCart
// @Tags Cart
// @Description	Get cart
// @Accept json
// @Produce json
// @Success	200	{object} models.Cart "Cart info"
// @Failure	400	{object} response.Response	"invalid request"
// @Failure	429
// @Router	/api/cart/summary [get]
func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	userID, ok := r.Context().Value(cookie.AccessTokenCookieName).(uuid.UUID)
	if !ok {
		h.log.Error("failed cast uuid from context value")
		resp.JSONStatus(w, http.StatusUnauthorized)

		return
	}

	cart, err := h.uc.GetCart(r.Context(), userID)
	if err != nil {
		h.log.Error("failed to get cart", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.GetCart", "cart", cart)
	resp.JSON(w, http.StatusOK, cart)
}
