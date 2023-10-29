package order

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/cart_mock.go -package mock

type OrderUsecase interface {
	CreateOrder(ctx context.Context, id uuid.UUID) (models.Order, error)
}

type OrderRepo interface {
	CreateOrder(ctx context.Context, cart models.Cart, userID uuid.UUID) (models.Order, error)
}
