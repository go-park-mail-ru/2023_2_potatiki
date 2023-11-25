package grpc

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/delivery/grpc/gen"
	"github.com/go-park-mail-ru/2023_2_potatiki/proto/gmodels"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

type serverAPI struct {
	log *slog.Logger
	uc  order.OrderUsecase

	//gen.ProductsServer
	gen.UnimplementedOrderServer
}

func Register(gRPCServer *grpc.Server, log *slog.Logger, uc order.OrderUsecase) {
	gen.RegisterOrderServer(gRPCServer, &serverAPI{log: log, uc: uc})
}

func (h serverAPI) CreateOrder(ctx context.Context, in *gen.CreateOrderRequest) (*gen.CreateOrderResponse, error) {
	userId, err := uuid.FromString(in.Id)
	if err != nil {

		return &gen.CreateOrderResponse{Error: err.Error()}, nil
	}
	order, err := h.uc.CreateOrder(ctx, userId)
	if err != nil {

		return &gen.CreateOrderResponse{Error: err.Error()}, nil
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

func (h serverAPI) GetOrders(ctx context.Context, in *gen.OrdersRequest) (*gen.OrdersResponse, error) {
	userId, err := uuid.FromString(in.Id)
	if err != nil {

		return &gen.OrdersResponse{Error: err.Error()}, nil
	}
	orders, err := h.uc.GetOrders(ctx, userId)
	if err != nil {

		return &gen.OrdersResponse{Error: err.Error()}, nil
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
