package order

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/cart_mock.go -package mock

type OrderUsecase interface {
	GetCart(context.Context, uuid.UUID) (models.Cart, error)
	UpdateCart(context.Context, models.Cart) (models.Cart, error)
}

type OrderRepo interface {
	CreateOrder(context.Context, uuid.UUID) error
	ReadCart(context.Context, uuid.UUID) (models.Cart, error)
	UpdateCart(context.Context, models.Cart) (models.Cart, error)
	ReadCartProducts(context.Context, models.Cart) (models.Cart, error)
}
