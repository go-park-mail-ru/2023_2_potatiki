package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/repo"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products"
	uuid "github.com/satori/go.uuid"
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
		if errors.Is(err, repo.ErrProductNotFound) {
			// TODO: implement
		}
		err = fmt.Errorf("error happened in repo.ReadProduct: %w", err)

		return models.Product{}, err
	}

	return product, nil
}

func (uc *ProductsUsecase) GetProducts(ctx context.Context, paging, count int64, ratingBy, sortingBy string) (
	[]models.Product, error) {

	productsSlice, err := uc.repo.ReadProducts(ctx, paging, count, ratingBy, sortingBy)
	if err != nil {
		err = fmt.Errorf("error happened in repo.ReadProducts: %w", err)

		return []models.Product{}, err
	}

	return productsSlice, nil
}

func (uc *ProductsUsecase) GetCategory(ctx context.Context, id int, paging, count int64, ratingBy, sortingBy string) (
	[]models.Product, error) {

	productsSlice, err := uc.repo.ReadCategory(ctx, id, paging, count, ratingBy, sortingBy)
	if err != nil {
		err = fmt.Errorf("error happened in repo.GetCategory: %w", err)

		return []models.Product{}, err
	}

	return productsSlice, nil
}
