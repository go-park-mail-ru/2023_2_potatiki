package grpc

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/metricsmw"

	generatedOrder "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/delivery/grpc/gen"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/delivery/grpc/gen"
	"github.com/go-park-mail-ru/2023_2_potatiki/proto/gmodels"
	uuid "github.com/satori/go.uuid"
)

//go:generate mockgen -source=./gen/order_grpc.pb.go -destination=../../mocks/order_grpc.go -package=mock

type GrpcOrderHandler struct {
	log *slog.Logger
	uc  order.OrderUsecase

	generatedOrder.OrderServer
}

func NewGrpcOrderHandler(uc order.OrderUsecase, log *slog.Logger) *GrpcOrderHandler {
	return &GrpcOrderHandler{
		uc:  uc,
		log: log,
	}
}

func (h GrpcOrderHandler) CreateOrder(ctx context.Context, in *gen.CreateOrderRequest) (*gen.CreateOrderResponse, error) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	userId, err := uuid.FromString(in.Id)
	if err != nil {
		h.log.Error("failed to get uuid from string", sl.Err(err))
		return &gen.CreateOrderResponse{Error: err.Error()}, metricsmw.ClientError
	}
	order, err := h.uc.CreateOrder(ctx, userId)
	if err != nil {
		h.log.Error("failed in h.uc.CreateOrder", sl.Err(err))
		return &gen.CreateOrderResponse{Error: err.Error()}, metricsmw.ServerError
	}

	orderResponse := gen.CreateOrderResponse{
		Order: &gmodels.Order{
			Id:     order.Id.String(),
			Status: order.Status,
			Address: &gmodels.Address{
				Id:        order.Address.Id.String(),
				ProfileId: order.Address.ProfileId.String(),
				City:      order.Address.City,
				Street:    order.Address.Street,
				House:     order.Address.House,
				Flat:      order.Address.Flat,
				IsCurrent: order.Address.IsCurrent,
			},
		},
	}

	for _, p := range order.Products {
		orderResponse.Order.Products = append(orderResponse.Order.Products,
			&gmodels.OrderProduct{
				Quantity: p.Quantity,
				Product: &gmodels.Product{
					Id:          p.Id.String(),
					Name:        p.Name,
					Description: p.Description,
					Price:       p.Price,
					ImgSrc:      p.ImgSrc,
					Rating:      p.Rating,
					Category: &gmodels.Category{
						Id:     p.Category.Id,
						Name:   p.Category.Name,
						Parent: p.Category.Parent,
					},
				},
			})
	}

	orderResponse.Error = ""

	return &orderResponse, nil
}

func (h GrpcOrderHandler) GetOrders(ctx context.Context, in *gen.OrdersRequest) (*gen.OrdersResponse, error) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	userId, err := uuid.FromString(in.Id)
	if err != nil {
		h.log.Error("failed to get uuid from string", sl.Err(err))
		return &gen.OrdersResponse{Error: err.Error()}, metricsmw.ClientError
	}
	orders, err := h.uc.GetOrders(ctx, userId)
	if err != nil {
		h.log.Error("failed in h.uc.GetOrders", sl.Err(err))
		return &gen.OrdersResponse{Error: err.Error()}, metricsmw.ServerError
	}

	var ordersResponse gen.OrdersResponse
	for _, o := range orders {
		order := gmodels.Order{
			Id:     o.Id.String(),
			Status: o.Status,
			Address: &gmodels.Address{
				Id:        o.Address.Id.String(),
				ProfileId: o.Address.ProfileId.String(),
				City:      o.Address.City,
				Street:    o.Address.Street,
				House:     o.Address.House,
				Flat:      o.Address.Flat,
				IsCurrent: o.Address.IsCurrent,
			},
		}

		for _, p := range o.Products {
			order.Products = append(order.Products, &gmodels.OrderProduct{
				Quantity: p.Quantity,
				Product: &gmodels.Product{
					Id:          p.Id.String(),
					Name:        p.Name,
					Description: p.Description,
					Price:       p.Price,
					ImgSrc:      p.ImgSrc,
					Rating:      p.Rating,
					Category: &gmodels.Category{
						Id:     p.Category.Id,
						Name:   p.Category.Name,
						Parent: p.Category.Parent,
					},
				},
			})
		}

		ordersResponse.Orders = append(ordersResponse.Orders, &order)
	}

	ordersResponse.Error = ""

	return &ordersResponse, nil
}
