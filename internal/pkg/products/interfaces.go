package products

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/products_mock.go -package mock

type ProductsUsecase interface {
	GetProduct(context.Context, uuid.UUID) (models.Product, error)
	GetProducts(context.Context, int64, int64) ([]models.Product, error)
}

type ProductsRepo interface {
	ReadProduct(context.Context, uuid.UUID) (models.Product, error)
	ReadProducts(context.Context, int64, int64) ([]models.Product, error)
}
