package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order"
	generatedOrder "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/delivery/grpc/generated"
	uuid "github.com/satori/go.uuid"
)

type GrpcOrderHandler struct {
	uc order.OrderUsecase
	// TODO: Add logger

	generatedOrder.OrderServiceServer
}

func NewGrpcOrderHandler(uc order.OrderUsecase) *GrpcOrderHandler {
	return &GrpcOrderHandler{
		uc: uc,
	}
}

func (h GrpcOrderHandler) CreateOrder(ctx context.Context, in *generatedOrder.ProfilePayload) (*generatedOrder.OrderResponse, error) {
	userId, err := uuid.FromString(in.Id)
	if err != nil {

		return &generatedOrder.OrderResponse{Error: err.Error()}, nil
	}
	order, err := h.uc.CreateOrder(ctx, userId)
	if err != nil {

		return &generatedOrder.OrderResponse{Error: err.Error()}, nil
	}

	var orderResponse generatedOrder.OrderResponse
	orderResponse.Order = &generatedOrder.Order{
		Id:     order.Id.String(),
		Status: order.Status,
		Address: &generatedOrder.Address{
			Id:        order.Address.Id.String(),
			ProfileId: order.Address.ProfileId.String(),
			City:      order.Address.City,
			Street:    order.Address.Street,
			House:     order.Address.House,
			Flat:      order.Address.Flat,
			IsCurrent: order.Address.IsCurrent,
		},
	}

	for _, p := range order.Products {
		orderResponse.Order.Products = append(orderResponse.Order.Products, &generatedOrder.Product{
			Id:          p.Id.String(),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			ImgSrc:      p.ImgSrc,
			Rating:      p.Rating,
			Category: &generatedOrder.Category{
				Id:     p.Category.Id,
				Name:   p.Category.Name,
				Parent: p.Category.Parent,
			},
		})
	}

	orderResponse.Error = ""

	return &orderResponse, nil
}

func (h GrpcOrderHandler) GetOrders(ctx context.Context, in *generatedOrder.ProfilePayload) (*generatedOrder.OrdersResponse, error) {
	userId, err := uuid.FromString(in.Id)
	if err != nil {

		return &generatedOrder.OrdersResponse{Error: err.Error()}, nil
	}
	orders, err := h.uc.GetOrders(ctx, userId)
	if err != nil {

		return &generatedOrder.OrdersResponse{Error: err.Error()}, nil
	}

	var ordersResponse generatedOrder.OrdersResponse
	for _, o := range orders {
		var order generatedOrder.Order

		order = generatedOrder.Order{
			Id:     o.Id.String(),
			Status: o.Status,
			Address: &generatedOrder.Address{
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
			order.Products = append(order.Products, &generatedOrder.Product{
				Id:          p.Id.String(),
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				ImgSrc:      p.ImgSrc,
				Rating:      p.Rating,
				Category: &generatedOrder.Category{
					Id:     p.Category.Id,
					Name:   p.Category.Name,
					Parent: p.Category.Parent,
				},
			})
		}

		ordersResponse.Orders = append(ordersResponse.Orders, &order)
	}

	ordersResponse.Error = ""

	return &ordersResponse, nil
}
