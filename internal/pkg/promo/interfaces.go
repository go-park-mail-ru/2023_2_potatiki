package promo

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/promo_mock.go -package mock

var (
	ErrPromocodeNotFound = errors.New("promocode not found")
)

type PromoUsecase interface {
	CheckPromocode(context.Context, string) (*models.Promocode, error)
	UsePromocode(context.Context, string) (*models.Promocode, error)
}

type PromoRepo interface {
	ReadPromocode(context.Context, string) (*models.Promocode, error)
	UsePromocode(context.Context, string) (*models.Promocode, error)
}
