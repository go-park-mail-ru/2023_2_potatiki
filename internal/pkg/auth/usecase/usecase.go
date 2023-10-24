package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/google/uuid"
)

type AuthUsecase struct {
	repo   auth.AuthRepo
	Auther auth.AuthAuther
}

func NewAuthUsecase(repo auth.AuthRepo, cfg auth.AuthConfig) *AuthUsecase {
	return &AuthUsecase{
		repo:   repo,
		Auther: NewAuther(cfg),
	}
}

func (uc *AuthUsecase) CheckToken(ctx context.Context, tokenStr string) (uuid.UUID, error) {
	claims, err := uc.Auther.GetClaims(tokenStr)
	if err != nil {
		err = fmt.Errorf("error happened in Auther.GetClaims: %w", err)

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
		err = fmt.Errorf("error happened in repo.CheckUser: %w", err)

		return models.Profile{}, "", time.Now(), err
	}

	token, exp, err := uc.Auther.GenerateToken(&profile)
	if err != nil {
		err = fmt.Errorf("error happened in Auther.GenerateToken: %w", err)

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
		err = fmt.Errorf("error happened in repo.CreateUser: %w", err)

		return models.Profile{}, "", time.Now(), err
	}

	token, exp, err := uc.Auther.GenerateToken(&profile)
	if err != nil {
		err = fmt.Errorf("error happened in Auther.GenerateToken: %w", err)

		return models.Profile{}, "", time.Now(), err
	}

	return profile, token, exp, nil
}

func (uc *AuthUsecase) GetProfile(ctx context.Context, userID uuid.UUID) (models.Profile, error) {
	profile, err := uc.repo.ReadProfile(ctx, userID)
	if err != nil {
		err = fmt.Errorf("error happened in repo.ReadProfile: %w", err)

		return models.Profile{}, err
	}

	return profile, nil
}
