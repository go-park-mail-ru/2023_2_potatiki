package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/config"
	"github.com/google/uuid"
)

type AuthUsecase struct {
	repo   auth.AuthRepo
	auther auth.AuthAuther
}

func NewAuthUsecase(repo auth.AuthRepo, cfg config.Auther) *AuthUsecase {
	return &AuthUsecase{
		repo:   repo,
		auther: NewAuther(cfg),
	}
}

func (uc *AuthUsecase) CheckToken(ctx context.Context, tokenStr string) (uuid.UUID, error) {
	claims, err := uc.auther.GetClaims(tokenStr)
	if err != nil {
		return uuid.UUID{}, err
	}
	return claims.ID, nil
}

func (uc *AuthUsecase) SignIn(ctx context.Context, user models.User) (models.Profile, string, time.Time, error) {
	if !user.IsValid() {
		err := errors.New("user is not valid")
		return models.Profile{}, "", time.Now(), err
	}
	profile, err := uc.repo.CheckUser(ctx, user)
	if err != nil {
		return models.Profile{}, "", time.Now(), err
	}
	token, exp, err := uc.auther.GenerateToken(&profile)
	if err != nil {
		return models.Profile{}, "", time.Now(), err
	}
	return profile, token, exp, nil
}

func (uc *AuthUsecase) SignUp(ctx context.Context, user models.User) (models.Profile, string, time.Time, error) {
	if !user.IsValid() {
		err := errors.New("user is not valid")
		return models.Profile{}, "", time.Now(), err
	}
	profile, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return models.Profile{}, "", time.Now(), err
	}
	token, exp, err := uc.auther.GenerateToken(&profile)
	if err != nil {
		return models.Profile{}, "", time.Now(), err
	}
	return profile, token, exp, nil
}

func (uc *AuthUsecase) GetProfile(ctx context.Context, userId uuid.UUID) (models.Profile, error) {
	profile, err := uc.repo.ReadProfile(ctx, userId)
	if err != nil {
		return models.Profile{}, err
	}
	return profile, nil
}
