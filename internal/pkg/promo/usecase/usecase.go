package usecase

import (
	"context"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/promo"
)

type PromoUsecase struct {
	repo promo.PromoRepo
}

func NewPromoUsecase(repo promo.PromoRepo) *PromoUsecase {
	return &PromoUsecase{
		repo: repo,
	}
}

func (uc *PromoUsecase) CheckPromocode(ctx context.Context, name string) (*models.Promocode, error) {
	promocode, err := uc.repo.ReadPromocode(ctx, name)
	if err != nil {
		return &models.Promocode{}, err
	}

	if time.Now().After(promocode.Deadline) || promocode.Leftover < 1 {
		return &models.Promocode{}, promo.ErrPromocodeNotFound
	}

	return promocode, nil
}

func (uc *PromoUsecase) UsePromocode(ctx context.Context, name string) (*models.Promocode, error) {
	promocode, err := uc.repo.UsePromocode(ctx, name)
	if err != nil {
		return &models.Promocode{}, err
	}
	if time.Now().After(promocode.Deadline) || promocode.Leftover < 1 {
		return &models.Promocode{}, promo.ErrPromocodeNotFound
	}

	return promocode, nil
}
