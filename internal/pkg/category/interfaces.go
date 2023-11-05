package category

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/category_mock.go -package mock

type CategoryUsecase interface {
	Categories(context.Context) (models.CategoryTree, error)
}

type CategoryRepo interface {
	ReadCategories(context.Context) (models.CategoryTree, error)
}
