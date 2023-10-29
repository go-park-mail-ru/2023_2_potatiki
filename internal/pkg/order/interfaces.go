package order

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/cart_mock.go -package mock

type OrderUsecase interface {
	GetCart(context.Context, uuid.UUID) (models.Cart, error)
	AddProduct(context.Context, models.Cart, models.CartProduct) (models.Cart, error)
	DeleteProduct(context.Context, models.Cart, models.CartProduct) (models.Cart, error)
}

type OrderRepo interface {
	CreateCart(context.Context, uuid.UUID) (models.Cart, error)
	CheckCart(context.Context, uuid.UUID) (models.Cart, error)
	ReadCart(context.Context, uuid.UUID) (models.Cart, error)
	ReadCartProducts(context.Context, models.Cart) (models.Cart, error)
	AddProduct(context.Context, models.Cart, models.CartProduct) (models.Cart, error)
	DeleteProduct(context.Context, models.Cart, models.CartProduct) (models.Cart, error)
	CheckProduct(context.Context, uuid.UUID) error
}
