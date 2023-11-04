package order

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/cart_mock.go -package mock

type OrderUsecase interface {
	CreateOrder(ctx context.Context, id uuid.UUID) (models.Order, error)
	GetCurrentOrder(ctx context.Context, userID uuid.UUID) (models.Order, error)
	GetOrders(ctx context.Context, userID uuid.UUID) ([]models.Order, error)
}

type OrderRepo interface {
	CreateOrder(ctx context.Context, cart models.Cart, userID uuid.UUID, statusID int) (models.Order, error)
	ReadOrderID(ctx context.Context, userID uuid.UUID) (uuid.UUID, error)
	ReadOrder(ctx context.Context, orderID uuid.UUID) (models.Order, error)
	ReadOrdersID(ctx context.Context, userID uuid.UUID) ([]uuid.UUID, error)
}
