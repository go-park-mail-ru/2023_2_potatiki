package search

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/search_mock.go -package mock

type SearchUsecase interface {
	SearchProducts(context.Context, string) ([]models.Product, error)
}

type SearchRepo interface {
	ReadProductsByName(context.Context, string) ([]models.Product, error)
}
