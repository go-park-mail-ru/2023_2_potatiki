package products

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/products_mock.go -package mock

type ProductsUsecase interface {
	GetProduct(context.Context, uuid.UUID) (models.Product, error)
	GetProducts(context.Context, int64, int64, string, string) ([]models.Product, error)
	GetCategory(context.Context, int, int64, int64, string, string) ([]models.Product, error)
}

type ProductsRepo interface {
	ReadProduct(context.Context, uuid.UUID) (models.Product, error)
	ReadProducts(context.Context, int64, int64) ([]models.Product, error)
	ReadProductsByPrice(context.Context, int64, int64, string) ([]models.Product, error)
	ReadProductsByRating(context.Context, int64, int64, string) ([]models.Product, error)
	ReadProductsCategory(context.Context, int, int64, int64) ([]models.Product, error)
	ReadProductsByRatingPrice(context.Context, int64, int64, string, string) ([]models.Product, error)
	ReadProductsCategoryByPrice(context.Context, int, int64, int64, string) ([]models.Product, error)
	ReadProductsCategoryByRating(context.Context, int, int64, int64, string) ([]models.Product, error)
	ReadProductsCategoryByRatingPrice(context.Context, int, int64, int64, string, string) ([]models.Product, error)
}
