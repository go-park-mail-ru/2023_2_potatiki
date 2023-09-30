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

func (uc *ProductsUsecase) GetProduct(ctx context.Context, id uuid.UUID) (models.Product, error) {
	product, err := uc.repo.ReadProduct(ctx, id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (uc *ProductsUsecase) GetProducts(ctx context.Context, paging int64, count int64) ([]models.Product, error) {
	productsSlice, err := uc.repo.ReadProducts(ctx, paging, count)
	if err != nil {
		return []models.Product{}, err
	}
	return productsSlice, nil
}
