package http

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/delivery/grpc/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	orderRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/repo"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"
	uuid "github.com/satori/go.uuid"
)

type OrderHandler struct {
	client gen.OrderClient
	log    *slog.Logger
	uc     order.OrderUsecase
}

func NewOrderHandler(cl gen.OrderClient, log *slog.Logger, uc order.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		client: cl,
		log:    log,
		uc:     uc,
	}
}

// @Summary	CreateOrder
// @Tags Order
// @Description	Create Order using profile ID from cookies
// @Accept json
// @Produce json
// @Param input body models.OrderInfo true "DeliveryDate and DeliveryTime"
// @Success	200	{object} models.Order "New order info"
// @Failure	401	"User unauthorized"
// @Failure	404	"Promocode not found"
// @Failure	403	"Promocode leftout"
// @Failure	419	"Promocode expired"
// @Failure	406	{object} responser.response	"something not found error message"
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

	body, err := io.ReadAll(r.Body)
	if resp.BodyErr(err, h.log, w) {
		return
	}
	defer r.Body.Close()
	h.log.Debug("got file from r.Body", slog.Any("request", r))

	payload := &models.OrderInfo{}
	err = payload.UnmarshalJSON(body)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	gorder, err := h.client.CreateOrder(r.Context(), &gen.CreateOrderRequest{
		Id:            userID.String(),
		DeliveryDate:  payload.DeliveryAtDate,
		DeliveryTime:  payload.DeliveryAtTime,
		PromocodeName: payload.PromocodeName,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			h.log.Error("failed cast grpc error", sl.Err(err))
			resp.JSONStatus(w, http.StatusTooManyRequests)
			return
		}
		switch st.Code() {
		case codes.NotFound:
			h.log.Warn("failed to CreateOrder", sl.Err(st.Err()))
			resp.JSONStatus(w, http.StatusNotFound)
		case codes.OutOfRange:
			h.log.Warn("failed to CreateOrder", sl.Err(st.Err()))
			resp.JSONStatus(w, http.StatusForbidden)
		case codes.DeadlineExceeded:
			h.log.Warn("failed to CreateOrder", sl.Err(st.Err()))
			resp.JSONStatus(w, 419)
		case codes.Unavailable: //orderRepo.ErrPoductNotFound:
			h.log.Warn("failed to CreateOrder", sl.Err(st.Err()))
			resp.JSON(w, http.StatusNotAcceptable, resp.Err(st.Message()))
		case codes.Aborted: //cartRepo.ErrCartNotFound
			h.log.Warn("failed to CreateOrder", sl.Err(st.Err()))
			resp.JSON(w, http.StatusNotAcceptable, resp.Err(st.Message()))
		case codes.ResourceExhausted: //addressRepo.ErrAddressNotFound
			h.log.Warn("failed to CreateOrder", sl.Err(st.Err()))
			resp.JSON(w, http.StatusNotAcceptable, resp.Err(st.Message()))
		default:
			h.log.Error("failed to CreateOrder", sl.Err(st.Err()))
			resp.JSONStatus(w, http.StatusTooManyRequests)
		}

		return

	}
	if err != nil {
		h.log.Error("failed to get CreateOrder", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		//if errors.Is(err, cartRepo.ErrCartNotFound) {
		//	resp.JSON(w, http.StatusNotFound, resp.Err("cart not found"))
		//
		//	return
		//}
		//
		//if errors.Is(err, orderRepo.ErrPoductNotFound) {
		//	resp.JSON(w, http.StatusNotFound, resp.Err("product not found"))
		//
		//	return
		//}
		//
		//if errors.Is(err, addressRepo.ErrAddressNotFound) {
		//	resp.JSON(w, http.StatusNotFound, resp.Err("address not found"))
		//
		//	return
		//}

	}

	orderId, err := uuid.FromString(gorder.Order.Id)
	addressId, err := uuid.FromString(gorder.Order.Address.Id)
	profileId, err := uuid.FromString(gorder.Order.Address.ProfileId)
	parsedTime, err := time.Parse(time.RFC3339, gorder.Order.CreationAt)
	if err != nil {
		h.log.Error("failed to parse order response field", err)
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}
	order := models.Order{
		Id:           orderId,
		Status:       gorder.Order.Status,
		CreationAt:   parsedTime,
		DeliveryTime: gorder.Order.DeliveryTime,
		DeliveryDate: gorder.Order.DeliveryDate,
		PomocodeName: gorder.Order.PromocodeName,
		Address: models.Address{
			Id:        addressId,
			ProfileId: profileId,
			City:      gorder.Order.Address.City,
			Street:    gorder.Order.Address.Street,
			House:     gorder.Order.Address.House,
			Flat:      gorder.Order.Address.Flat,
			IsCurrent: gorder.Order.Address.IsCurrent,
		},
		Products: make([]models.OrderProduct, len(gorder.Order.Products)),
	}

	for i, gproduct := range gorder.Order.Products {
		productId, err := uuid.FromString(gproduct.Product.Id)
		if err != nil {
			h.log.Error("failed to cast id", sl.Err(err))
			resp.JSONStatus(w, http.StatusTooManyRequests)
			return
		}
		order.Products[i] = models.OrderProduct{
			Quantity: gproduct.Quantity,
			Product: models.Product{
				Id:          productId,
				Name:        gproduct.Product.Name,
				Description: gproduct.Product.Description,
				Price:       gproduct.Product.Price,
				ImgSrc:      gproduct.Product.ImgSrc,
				Rating:      gproduct.Product.Rating,
				Category: models.Category{
					Id:     gproduct.Product.Category.Id,
					Name:   gproduct.Product.Category.Name,
					Parent: gproduct.Product.Category.Parent,
				},
			},
		}
	}

	h.log.Debug("h.uc.CreateOrder", "order", order)
	resp.JSON(w, http.StatusOK, &order)
}

// @Summary	GetOrders
// @Tags Order
// @Description	Get all Orders using profile ID from cookies
// @Accept json
// @Produce json
// @Success	200	{array} models.Order "All orders info"
// @Failure	401	"User unauthorized"
// @Failure 404	{object} responser.response	"something not found error message"
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

	gorders, err := h.client.GetOrders(r.Context(), &gen.OrdersRequest{
		Id: userID.String(),
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			h.log.Error("failed cast grpc error", sl.Err(err))
			resp.JSONStatus(w, http.StatusTooManyRequests)
			return
		}
		if st.Code() == codes.NotFound {
			h.log.Warn("orders not found", slog.Any("grpc status", st))
			resp.JSONStatus(w, http.StatusNotFound)
			return
		}
		//if errors.Is(err, orderRepo.ErrPoductsInOrderNotFound) {
		//	resp.JSON(w, http.StatusNotFound, resp.Err("products in order not found"))
		//
		//	return
		h.log.Error("failed to get order", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	orders := make([]models.Order, len(gorders.Orders))

	for i, gorder := range gorders.Orders {

		orderId, err := uuid.FromString(gorder.Id)
		addressId, err := uuid.FromString(gorder.Address.Id)
		profileId, err := uuid.FromString(gorder.Address.ProfileId)
		parsedTime, err := time.Parse(time.RFC3339, gorder.CreationAt)
		if err != nil {
			h.log.Error("failed to cast uuid to string", sl.Err(err))
			resp.JSONStatus(w, http.StatusTooManyRequests)
			return
		}
		orders[i] = models.Order{
			Id:           orderId,
			Status:       gorder.Status,
			CreationAt:   parsedTime,
			DeliveryTime: gorder.DeliveryTime,
			DeliveryDate: gorder.DeliveryDate,
			PomocodeName: gorder.PromocodeName,
			Address: models.Address{
				Id:        addressId,
				ProfileId: profileId,
				City:      gorder.Address.City,
				Street:    gorder.Address.Street,
				House:     gorder.Address.House,
				Flat:      gorder.Address.Flat,
				IsCurrent: gorder.Address.IsCurrent,
			},
			Products: make([]models.OrderProduct, len(gorder.Products)),
		}

		for j, product := range gorder.Products {
			productId, err := uuid.FromString(product.Product.Id)
			if err != nil {
				h.log.Error("failed to cast uuid to string", sl.Err(err))
				resp.JSONStatus(w, http.StatusTooManyRequests)
				return
			}
			orders[i].Products[j] = models.OrderProduct{
				Quantity: product.Quantity,
				Product: models.Product{
					Id:          productId,
					Name:        product.Product.Name,
					Description: product.Product.Description,
					Price:       product.Product.Price,
					ImgSrc:      product.Product.ImgSrc,
					Rating:      product.Product.Rating,
					Category: models.Category{
						Id:     product.Product.Category.Id,
						Name:   product.Product.Category.Name,
						Parent: product.Product.Category.Parent,
					},
				},
			}
		}
	}

	h.log.Debug("h.uc.GetOrders", "orders", orders)
	resp.JSON(w, http.StatusOK, (*models.OrderSlice)(&orders))
}

// @Summary	GetCurrentOrder
// @Tags Order
// @Description	Get Current Order using profile ID from cookies
// @Accept json
// @Produce json
// @Success	200	{object} models.Order "Current order info"
// @Failure	401	"User unauthorized"
// @Failure	404	{object} responser.response	"something not found error message"
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
		h.log.Error("failed to get order", sl.Err(err))
		if errors.Is(err, orderRepo.ErrOrderNotFound) {
			resp.JSON(w, http.StatusNotFound, resp.Err("order not found"))

			return
		}
		if errors.Is(err, orderRepo.ErrPoductsInOrderNotFound) {
			resp.JSON(w, http.StatusNotFound, resp.Err("products in order not found"))

			return
		}
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.GetCurrentOrder", "order", order)
	resp.JSON(w, http.StatusOK, &order)
}
