package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/recommendations"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/recommendations/repo"
	uuid "github.com/satori/go.uuid"
	"sort"
)

type RecommendationsUsecase struct {
	repo recommendations.RecommendationsRepo
}

func NewRecommendationsUsecase(repo recommendations.RecommendationsRepo) *RecommendationsUsecase {
	return &RecommendationsUsecase{
		repo: repo,
	}
}

func (uc *RecommendationsUsecase) AnonRecommendations(
	ctx context.Context,
	productID uuid.UUID,
	categoryID int64,
) (models.ProductSlice, error) {

	products, err := uc.repo.ReadProductsFromParentCategory(ctx, categoryID)
	if err != nil {
		err = fmt.Errorf("error happened in repo.ReadProductsFromParentCategory: %w", err)

		return models.ProductSlice{}, err
	}

	return products, nil
}

func (uc *RecommendationsUsecase) Recommendations(
	ctx context.Context,
	userID uuid.UUID,
	productID uuid.UUID,
	categoryID int64,
) (models.ProductSlice, error) {

	oldActivity, err := uc.repo.ReadUserActivity(ctx, userID)
	if err != nil {
		if errors.Is(err, repo.ErrActivityNotFound) {
			products, err := uc.repo.ReadProductsFromParentCategory(ctx, categoryID)
			if err != nil {
				err = fmt.Errorf("error happened in repo.ReadProductsFromParentCategory: %w", err)

				return models.ProductSlice{}, err
			}
			return products, nil
		} else {
			err = fmt.Errorf("error happened in repo.ReadUserActivity: %w", err)

			return models.ProductSlice{}, err
		}
	}

	if oldActivity.Product == nil || oldActivity.Category == nil {
		products, err := uc.repo.ReadProductsFromParentCategory(ctx, categoryID)
		if err != nil {
			err = fmt.Errorf("error happened in repo.ReadProductsFromParentCategory: %w", err)

			return models.ProductSlice{}, err
		}

		return products, nil
	}

	productStatisticSlice := make([]models.ProductStatistic, 0, len(oldActivity.Product))
	for _, product := range oldActivity.Product {
		productStatisticSlice = append(productStatisticSlice, product)
	}

	sort.Slice(productStatisticSlice, func(i, j int) bool {
		return productStatisticSlice[i].ActivityPoints > productStatisticSlice[j].ActivityPoints
	})

	productIds := make(models.ProductIDs, 0, models.MinProductsCount/2)
	for index, product := range productStatisticSlice {
		productIds = append(productIds, product.ProductID)
		if index == models.MinProductsCount-1 {
			break
		}
	}

	products, err := uc.repo.ReadRecommendations(ctx, productIds)
	if err != nil {
		err = fmt.Errorf("error happened in repo.ReadRecommendations: %w", err)

		return models.ProductSlice{}, err
	}

	categoryStatisticSlice := make([]models.CategoryStatistic, 0, len(oldActivity.Category))
	for _, category := range oldActivity.Category {
		categoryStatisticSlice = append(categoryStatisticSlice, category)
	}
	sort.Slice(categoryStatisticSlice, func(i, j int) bool {
		return categoryStatisticSlice[i].ActivityPoints > categoryStatisticSlice[j].ActivityPoints
	})

	categoryIDs := make(models.CategoryIDs, 0, models.MinCateggoriesCount)
	for index, category := range categoryStatisticSlice {
		categoryIDs = append(categoryIDs, category.CategoryID)
		if index == models.MinCateggoriesCount-1 {
			break
		}
	}

	productsFromCategories, err := uc.repo.ReadProductsFromCategories(ctx, categoryIDs)
	products = append(products, productsFromCategories...)
	productsMap := make(map[uuid.UUID]models.Product, len(products))
	for _, product := range products {
		productsMap[product.Id] = product
	}

	if len(products) < models.MinProductsCount {
		productsExtra, err := uc.repo.ReadProductsFromParentCategory(ctx, categoryID)
		if err != nil {
			err = fmt.Errorf("error happened in repo.ReadProductsFromParentCategory: %w", err)

			return models.ProductSlice{}, err
		}
		for _, product := range productsExtra {
			productsMap[product.Id] = product
		}
	}

	products = make(models.ProductSlice, 0, len(productsMap))
	for _, product := range productsMap {
		products = append(products, product)
	}

	return products, nil
}

func (uc *RecommendationsUsecase) UpdateUserActivity(
	ctx context.Context,
	userID uuid.UUID,
	newActivity models.UserActivity,
) error {

	oldActivity, err := uc.repo.ReadUserActivity(ctx, userID)
	if err != nil {
		if errors.Is(err, repo.ErrActivityNotFound) {
			err = uc.repo.CreateUserActivity(ctx, userID)

			if err != nil {
				err = fmt.Errorf("error happened in repo.CreateUserActivity: %w", err)

				return err
			}
		} else {
			err = fmt.Errorf("error happened in repo.ReadUserActivity: %w", err)

			return err
		}
	}

	if oldActivity.Product == nil {
		oldActivity.Product = make(map[uuid.UUID]models.ProductStatistic, len(newActivity.Product))
	}

	if oldActivity.Category == nil {
		oldActivity.Category = make(map[int64]models.CategoryStatistic, len(newActivity.Category))
	}

	for _, productInfo := range newActivity.Product {
		oldProductInfo, ok := oldActivity.Product[productInfo.ProductID]
		if !ok {
			oldActivity.Product[productInfo.ProductID] = productInfo
		} else {
			oldProductInfo.ActivityPoints += productInfo.ActivityPoints
			oldProductInfo.IsBought = oldProductInfo.IsBought || productInfo.IsBought
			oldProductInfo.IsReviewed = oldProductInfo.IsReviewed || productInfo.IsReviewed
			oldActivity.Product[productInfo.ProductID] = oldProductInfo
		}
	}

	for _, categoryInfo := range newActivity.Category {
		oldCategoryInfo, ok := oldActivity.Category[categoryInfo.CategoryID]
		if !ok {
			oldActivity.Category[categoryInfo.CategoryID] = categoryInfo
		} else {
			oldCategoryInfo.ActivityPoints += categoryInfo.ActivityPoints
			oldActivity.Category[categoryInfo.CategoryID] = oldCategoryInfo
		}
	}

	err = uc.repo.UpdateUserActivity(ctx, userID, oldActivity)
	if err != nil {
		err = fmt.Errorf("error happened in repo.UpdateUserActivity: %w", err)

		return err
	}

	return nil
}
