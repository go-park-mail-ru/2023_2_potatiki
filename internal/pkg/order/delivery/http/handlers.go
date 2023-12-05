package http

import (
	"errors"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/hub"
	"github.com/gorilla/websocket"

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
	hub    *hub.Hub
}

func NewOrderHandler(cl gen.OrderClient, log *slog.Logger, uc order.OrderUsecase, hub *hub.Hub) *OrderHandler {
	return &OrderHandler{
		client: cl,
		log:    log,
		uc:     uc,
		hub:    hub,
	}
}

var upgrader = websocket.Upgrader{} // use default options

func (h *OrderHandler) GetNotifications(w http.ResponseWriter, r *http.Request) {
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

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	h.hub.AddClient(userID, c)

}

// @Summary	CreateOrder
// @Tags Order
// @Description	Create Order using profile ID from cookies
// @Accept json
// @Produce json
// @Param input body models.OrderInfo true "UpdateProfileDataPayload"
// @Success	200	{object} models.Order "New order info"
// @Failure	401	"User unauthorized"
// @Failure	404	{object} responser.Response	"something not found error message"
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

	//order, err := h.uc.CreateOrder(r.Context(), userID)
	orderResponse, err := h.client.CreateOrder(r.Context(), &gen.CreateOrderRequest{
		Id:           userID.String(),
		DeliveryDate: payload.DeliveryAtDate,
		DeliveryTime: payload.DeliveryAtTime,
	})
	if err != nil {
		h.log.Error("failed to get something", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	if orderResponse.Error != "" {
		h.log.Error("failed to get something", sl.Err(errors.New(orderResponse.Error)))
		resp.JSONStatus(w, http.StatusNotFound)

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
		return
	}
	orderProto := orderResponse.Order
	addressProto := orderProto.Address
	productsProto := orderProto.Products
	orderId, err := uuid.FromString(orderProto.Id)
	addressId, err := uuid.FromString(orderProto.Address.Id)
	profileId, err := uuid.FromString(orderProto.Address.ProfileId)
	parsedTime, err := time.Parse(time.RFC3339, orderProto.CreationAt)
	if err != nil {
		h.log.Error("failed to get from string", err)
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}
	orderModel := models.Order{
		Id:           orderId,
		Status:       orderProto.Status,
		CreationAt:   parsedTime,
		DeliveryTime: orderProto.DeliveryTime,
		DeliveryDate: orderProto.DeliveryDate,
		Address: models.Address{
			Id:        addressId,
			ProfileId: profileId,
			City:      addressProto.City,
			Street:    addressProto.Street,
			House:     addressProto.House,
			Flat:      addressProto.Flat,
			IsCurrent: addressProto.IsCurrent,
		},
	}

	var productsSlice []models.OrderProduct
	for _, orderProduct := range productsProto {
		product := orderProduct.Product
		productId, _ := uuid.FromString(product.Id)
		productsSlice = append(productsSlice, models.OrderProduct{
			Quantity: orderProduct.Quantity,
			Product: models.Product{
				Id:          productId,
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
				ImgSrc:      product.ImgSrc,
				Rating:      product.Rating,
				Category: models.Category{
					Id:     product.Category.Id,
					Name:   product.Category.Name,
					Parent: product.Category.Parent,
				},
			},
		})
	}

	orderModel.Products = productsSlice

	h.log.Debug("h.uc.CreateOrder", "order", orderModel)
	resp.JSON(w, http.StatusOK, &orderModel)
}

// @Summary	GetCurrentOrder
// @Tags Order
// @Description	Get Current Order using profile ID from cookies
// @Accept json
// @Produce json
// @Success	200	{object} models.Order "Current order info"
// @Failure	401	"User unauthorized"
// @Failure	404	{object} responser.Response	"something not found error message"
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

// @Summary	GetOrders
// @Tags Order
// @Description	Get all Orders using profile ID from cookies
// @Accept json
// @Produce json
// @Success	200	{array} models.Order "All orders info"
// @Failure	401	"User unauthorized"
// @Failure 404	{object} responser.Response	"something not found error message"
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

	ordersResponse, err := h.client.GetOrders(r.Context(), &gen.OrdersRequest{
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
		h.log.Error("failed to get order", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	if ordersResponse.Error != "" {
		h.log.Error("failed to get order", sl.Err(errors.New(ordersResponse.Error)))
		resp.JSONStatus(w, http.StatusNotFound)
		//
		//if errors.Is(err, orderRepo.ErrOrdersNotFound) {
		//	resp.JSON(w, http.StatusNotFound, resp.Err("orders not found"))
		//
		//	return
		//}
		//if errors.Is(err, orderRepo.ErrPoductsInOrderNotFound) {
		//	resp.JSON(w, http.StatusNotFound, resp.Err("products in order not found"))
		//
		//	return
		//}
		return
	}

	var (
		ordersSlice []models.Order
	)

	for _, orderProto := range ordersResponse.Orders {
		addressProto := orderProto.Address
		productsProto := orderProto.Products
		orderId, err := uuid.FromString(orderProto.Id)
		addressId, err := uuid.FromString(orderProto.Address.Id)
		profileId, err := uuid.FromString(orderProto.Address.ProfileId)
		parsedTime, err := time.Parse(time.RFC3339, orderProto.CreationAt)
		if err != nil {
			h.log.Error("failed to get from string", err)
			resp.JSONStatus(w, http.StatusTooManyRequests)

			return
		}
		orderModel := models.Order{
			Id:           orderId,
			Status:       orderProto.Status,
			CreationAt:   parsedTime,
			DeliveryTime: orderProto.DeliveryTime,
			DeliveryDate: orderProto.DeliveryDate,
			Address: models.Address{
				Id:        addressId,
				ProfileId: profileId,
				City:      addressProto.City,
				Street:    addressProto.Street,
				House:     addressProto.House,
				Flat:      addressProto.Flat,
				IsCurrent: addressProto.IsCurrent,
			},
		}

		var productsSlice []models.OrderProduct
		for _, orderProduct := range productsProto {
			product := orderProduct.Product
			productId, _ := uuid.FromString(product.Id)
			productsSlice = append(productsSlice, models.OrderProduct{
				Quantity: orderProduct.Quantity,
				Product: models.Product{
					Id:          productId,
					Name:        product.Name,
					Description: product.Description,
					Price:       product.Price,
					ImgSrc:      product.ImgSrc,
					Rating:      product.Rating,
					Category: models.Category{
						Id:     product.Category.Id,
						Name:   product.Category.Name,
						Parent: product.Category.Parent,
					},
				},
			})
		}

		orderModel.Products = productsSlice

		ordersSlice = append(ordersSlice, orderModel)
	}

	h.log.Debug("h.uc.GetOrders", "orders", ordersSlice)
	resp.JSON(w, http.StatusOK, (*models.OrderSlice)(&ordersSlice))
}
