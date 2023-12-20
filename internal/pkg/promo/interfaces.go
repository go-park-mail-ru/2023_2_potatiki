package promo

import (
	"context"
	"errors"

	uuid "github.com/satori/go.uuid"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
)

//go:generate mockgen -source interfaces.go -destination ./mocks/promo_mock.go -package mock

var (
	ErrPromocodeNotFound = errors.New("promocode not found")    //404
	ErrPromocodeLeftout  = errors.New("promocode leftout")      //403
	ErrPromocodeExpired  = errors.New("promocode expired")      //419
	ErrAlreadyUsed       = errors.New("promocode already used") //410
)

type PromoUsecase interface {
	CheckPromocode(context.Context, uuid.UUID, string) (*models.Promocode, error)
	UsePromocode(context.Context, string) (*models.Promocode, error)
}

type PromoRepo interface {
	ReadPromocode(context.Context, string) (*models.Promocode, error)
	CheckUniq(context.Context, uuid.UUID, int) error
	UsePromocode(context.Context, string) (*models.Promocode, error)
}
