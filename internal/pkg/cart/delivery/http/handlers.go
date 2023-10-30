package http

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
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
// @Param @Param input body models.Cart true "cart info"
// @Success	200	{object} models.Cart "cart info"
// @Failure	400	{object} response.Response	"invalid request"
// @Failure	429
// @Router	/api/cart/update [post]
func (h *CartHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
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
// @Failure	401
// @Failure	429
// @Router	/api/cart/summary [get]
func (h *CartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
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

	cart, err := h.uc.GetCart(r.Context(), userID)
	if err != nil {
		h.log.Error("failed to get cart", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.GetCart", "cart", cart)
	resp.JSON(w, http.StatusOK, cart)
}

// @Summary	AddProduct
// @Tags Cart
// @Description	add product to cart or change its number
// @Accept json
// @Produce json
// @Param input body models.CartProduct true "product info"
// @Success	200	{object} models.Cart "cart info"
// @Failure	400	{object} response.Response	"error message"
// @Failure	401
// @Failure	429
// @Router	/api/cart/add_product [post]
func (h *CartHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
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

	p := models.CartProduct{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}
	c := models.Cart{}
	c.ProfileId = userID

	cart, err := h.uc.AddProduct(r.Context(), c, p)
	if err != nil {
		h.log.Error("failed to add product to cart", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("add product to cart")
	resp.JSON(w, http.StatusOK, cart)
}

// @Summary	DeleteProduct
// @Tags Cart
// @Description	delete product from cart
// @Accept json
// @Produce json
// @Param input body models.CartProduct true "product info"
// @Success	200	{object} models.Cart "cart info"
// @Failure	400	{object} response.Response	"error message"
// @Failure	401
// @Failure	429
// @Router	/api/cart/delete_product [post]
func (h *CartHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
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

	p := models.CartProduct{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}
	c := models.Cart{}
	c.ProfileId = userID

	cart, err := h.uc.DeleteProduct(r.Context(), c, p)
	if err != nil {
		h.log.Error("failed to delete product from cart", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("delete product from cart")
	resp.JSON(w, http.StatusOK, cart)
}
