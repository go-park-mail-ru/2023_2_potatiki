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

func (uc *ProductsUsecase) GetProducts(ctx context.Context, paging, count int64, ratingBy, priceBy string) (
	[]models.Product, error) {

	var err error
	productsSlice := []models.Product{}
	if (ratingBy == "DESC" || ratingBy == "ASC") && (priceBy == "DESC" || priceBy == "ASC") {
		productsSlice, err = uc.repo.ReadProductsByRatingPrice(ctx, paging, count, ratingBy, priceBy)
	} else if ratingBy == "DESC" || ratingBy == "ASC" {
		productsSlice, err = uc.repo.ReadProductsByRating(ctx, paging, count, ratingBy)
	} else if priceBy == "DESC" || priceBy == "ASC" {
		productsSlice, err = uc.repo.ReadProductsByPrice(ctx, paging, count, priceBy)
	} else {
		productsSlice, err = uc.repo.ReadProducts(ctx, paging, count)
	}

	if err != nil {
		err = fmt.Errorf("error happened in repo.ReadProducts: %w", err)

		return []models.Product{}, err
	}

	return productsSlice, nil
}

func (uc *ProductsUsecase) GetCategory(ctx context.Context, id int, paging, count int64, ratingBy, priceBy string) (
	[]models.Product, error) {

	var err error
	productsSlice := []models.Product{}
	if (ratingBy == "DESC" || ratingBy == "ASC") && (priceBy == "DESC" || priceBy == "ASC") {
		productsSlice, err = uc.repo.ReadProductsCategoryByRatingPrice(ctx, id, paging, count, ratingBy, priceBy)
	} else if ratingBy == "DESC" || ratingBy == "ASC" {
		productsSlice, err = uc.repo.ReadProductsCategoryByRating(ctx, id, paging, count, ratingBy)
	} else if priceBy == "DESC" || priceBy == "ASC" {
		productsSlice, err = uc.repo.ReadProductsCategoryByPrice(ctx, id, paging, count, priceBy)
	} else {
		productsSlice, err = uc.repo.ReadProductsCategory(ctx, id, paging, count)
	}

	if err != nil {
		err = fmt.Errorf("error happened in repo.GetCategory: %w", err)

		return []models.Product{}, err
	}

	return productsSlice, nil
}
