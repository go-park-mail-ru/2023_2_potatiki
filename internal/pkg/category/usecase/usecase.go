package usecase

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/category"
)

type CategoryUsecase struct {
	repo category.CategoryRepo
}

func NewCategoryUsecase(repo category.CategoryRepo) *CategoryUsecase {
	return &CategoryUsecase{
		repo: repo,
	}
}

func (uc *CategoryUsecase) Categories(ctx context.Context) ([]models.Category, error) {
	categorySlice, err := uc.repo.ReadCategories(ctx)
	if err != nil {
		err = fmt.Errorf("error happened in repo.ReadCategories: %w", err)

		return []models.Category{}, err
	}

	return categorySlice, nil
}
