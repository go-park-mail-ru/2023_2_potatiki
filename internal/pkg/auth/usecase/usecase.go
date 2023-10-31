package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/user"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/hasher"
	"github.com/google/uuid"
)

type AuthUsecase struct {
	repo   user.UserRepo
	Auther auth.AuthAuther
}

func NewAuthUsecase(repo user.UserRepo, cfg auth.AuthConfig) *AuthUsecase {
	return &AuthUsecase{
		repo:   repo,
		Auther: NewAuther(cfg),
	}
}

var (
	ErrPassMismatch = errors.New("password does not match")
	ErrInvalidUser  = errors.New("user is not valid")
)

func (uc *AuthUsecase) SignIn(ctx context.Context, user *models.User) (*models.Profile, string, time.Time, error) {
	if !user.IsValid() {
		return &models.Profile{}, "", time.Now(), ErrInvalidUser
	}

	Id, err := uc.repo.GetProfileIdByUser(ctx, user)
	if err != nil {
		err = fmt.Errorf("error happened in repo.GetProfileIdByUser: %w", err)

		return &models.Profile{}, "", time.Now(), err
	}

	profile, err := uc.repo.ReadProfile(ctx, Id)
	if err != nil {
		err = fmt.Errorf("error happened in repo.ReadProfile: %w", err)

		return &models.Profile{}, "", time.Now(), err
	}

	if !hasher.CheckPass(profile.PasswordHash, user.Password) {
		return &models.Profile{}, "", time.Now(), ErrPassMismatch
	}

	token, exp, err := uc.Auther.GenerateToken(profile)
	if err != nil {
		err = fmt.Errorf("error happened in Auther.GenerateToken: %w", err)

		return &models.Profile{}, "", time.Now(), err
	}
	profile.HidePass()

	return profile, token, exp, nil
}

func (uc *AuthUsecase) SignUp(ctx context.Context, user *models.User) (*models.Profile, string, time.Time, error) {
	if !user.IsValid() {
		return &models.Profile{}, "", time.Now(), ErrInvalidUser
	}

	profile := &models.Profile{
		Id:           uuid.New(),
		Login:        user.Login,
		Description:  "",
		ImgSrc:       "default.png",
		PasswordHash: hasher.HashPass(user.Password),
	}

	err := uc.repo.CreateProfile(ctx, profile)
	if err != nil {
		err = fmt.Errorf("error happened in repo.CreateUser: %w", err)

		return &models.Profile{}, "", time.Now(), err
	}

	token, exp, err := uc.Auther.GenerateToken(profile)
	if err != nil {
		err = fmt.Errorf("error happened in Auther.GenerateToken: %w", err)

		return &models.Profile{}, "", time.Now(), err
	}

	profile.HidePass()
	return profile, token, exp, nil
}
