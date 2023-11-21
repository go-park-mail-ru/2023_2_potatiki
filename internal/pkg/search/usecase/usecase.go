package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/search"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/search/repo"
)

type SearchUsecase struct {
	repoSearch   search.SearchRepo
	repoProducts products.ProductsRepo
}

func NewSearchUsecase(repoSearch search.SearchRepo, repoProducts products.ProductsRepo) *SearchUsecase {
	return &SearchUsecase{
		repoSearch:   repoSearch,
		repoProducts: repoProducts,
	}
}

func (uc *SearchUsecase) SearchProducts(ctx context.Context, productName string) ([]models.Product, error) {
	var err error
	paging := int64(0)
	count := int64(10)
	var productsSlice []models.Product
	if productName == "" {
		productsSlice, err = uc.repoProducts.ReadProducts(ctx, paging, count)
		if err != nil {
			if errors.Is(err, repo.ErrProductNotFound) {
				return []models.Product{}, err
			}
			err = fmt.Errorf("error happened in repoProducts.ReadProducts: %w", err)

			return []models.Product{}, err
		}
	} else {
		productsSlice, err = uc.repoSearch.ReadProductsByName(ctx, productName)
		if err != nil {
			if errors.Is(err, repo.ErrProductNotFound) {
				return []models.Product{}, err
			}
			err = fmt.Errorf("error happened in repoSearch.SearchProducts: %w", err)

			return []models.Product{}, err
		}
	}

	return productsSlice, nil
}
