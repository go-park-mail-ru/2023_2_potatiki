package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products"
	"github.com/google/uuid"
)

type ProductsUsecase struct {
	repo products.ProductsRepo
}

func NewProductsUsecase(repo products.ProductsRepo) *ProductsUsecase {
	return &ProductsUsecase{
		repo: repo,
	}
}

func (uc *ProductsUsecase) GetProduct(context.Context, uuid.UUID) (models.Product, error) {
	panic("unimplemented")
}

func (uc *ProductsUsecase) GetProducts(context.Context, int64, int64) ([]models.Product, error) {
	panic("unimplemented")
}
