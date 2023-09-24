package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/google/uuid"
)

type AuthUsecase struct {
	repo auth.AuthRepo
}

func NewAuthUsecase(NewRepo auth.AuthRepo) AuthUsecase {
	return AuthUsecase{
		repo: NewRepo,
	}
}

func (uc *AuthUsecase) Login(context.Context, models.User) (models.Profile, error) {
	return models.Profile{}, nil
}

func (uc *AuthUsecase) SingUp(context.Context, models.User) (models.Profile, error) {
	return models.Profile{}, nil
}

func (uc *AuthUsecase) GetProfile(context.Context, uuid.UUID) (models.Profile, error) {
	return models.Profile{}, nil
}
