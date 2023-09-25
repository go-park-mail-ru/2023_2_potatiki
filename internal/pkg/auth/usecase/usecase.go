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

func New(repo auth.AuthRepo) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
	}
}

func (uc *AuthUsecase) SignIn(context.Context, models.User) (models.Profile, error) {
	panic("unimplemented")
}

func (uc *AuthUsecase) SignUp(context.Context, *models.User) (models.Profile, error) {
	panic("unimplemented")
}

func (uc *AuthUsecase) GetProfile(context.Context, uuid.UUID) (models.Profile, error) {
	panic("unimplemented")
}
