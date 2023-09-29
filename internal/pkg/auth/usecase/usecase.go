package usecase

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/google/uuid"
)

type AuthUsecase struct {
	repo auth.AuthRepo
}

func NewAuthUsecase(repo auth.AuthRepo) *AuthUsecase {
	return &AuthUsecase{
		repo: repo,
	}
}

func (uc *AuthUsecase) SignIn(ctx context.Context, user models.User) (models.Profile, error) {
	if !user.IsValid() {
		err := errors.New("user is not valid")
		return models.Profile{}, err
	}
	profile, err := uc.repo.CheckUser(ctx, user)
	if err != nil {
		return models.Profile{}, err
	}
	return profile, nil
}

func (uc *AuthUsecase) SignUp(ctx context.Context, user models.User) (models.Profile, error) {
	if !user.IsValid() {
		err := errors.New("user is not valid")
		return models.Profile{}, err
	}
	profile, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return models.Profile{}, err
	}
	return profile, nil
}

func (uc *AuthUsecase) GetProfile(ctx context.Context, userId uuid.UUID) (models.Profile, error) {
	profile, err := uc.repo.ReadProfile(ctx, userId)
	if err != nil {
		return models.Profile{}, err
	}
	return profile, nil
}
