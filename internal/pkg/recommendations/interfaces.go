package recommendations

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	uuid "github.com/satori/go.uuid"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/recommendations_mock.go -package mock

type RecommendationsUsecase interface {
	UpdateUserActivity(context.Context, uuid.UUID, models.UserActivity) error
	Recommendations(context.Context, uuid.UUID, uuid.UUID, int64) (models.ProductSlice, error)
	AnonRecommendations(context.Context, uuid.UUID, int64) (models.ProductSlice, error)
}

type RecommendationsRepo interface {
	ReadUserActivity(context.Context, uuid.UUID) (models.UserActivityStore, error)
	CreateUserActivity(context.Context, uuid.UUID) error
	UpdateUserActivity(context.Context, uuid.UUID, models.UserActivityStore) error
	ReadProductsFromParentCategory(context.Context, int64) (models.ProductSlice, error)
	ReadProduct(context.Context, uuid.UUID) (models.Product, error)
	ReadProductsFromCategory(context.Context, int64) (models.ProductSlice, error)
}
