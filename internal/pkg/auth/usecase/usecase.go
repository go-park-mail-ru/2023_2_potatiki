package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/profile"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/hasher"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

type AuthUsecase struct {
	repo   profile.ProfileRepo
	Auther auth.AuthAuther
}

func NewAuthUsecase(repo profile.ProfileRepo, cfg auth.AuthConfig) *AuthUsecase {
	return &AuthUsecase{
		repo:   repo,
		Auther: NewAuther(cfg),
	}
}

var (
	ErrPassMismatch = errors.New("password does not match")
)

func (uc *AuthUsecase) SignIn(ctx context.Context, payload *models.SignInPayload) (*models.Profile, string, time.Time, error) {
	if err := validator.New().Struct(payload); err != nil {
		return &models.Profile{}, "", time.Now(), err
	}

	Id, err := uc.repo.GetProfileIdByLogin(ctx, payload.Login)
	if err != nil {
		err = fmt.Errorf("error happened in repo.GetProfileIdByUser: %w", err)

		return &models.Profile{}, "", time.Now(), err
	}

	profile, err := uc.repo.ReadProfile(ctx, Id)
	if err != nil {
		err = fmt.Errorf("error happened in repo.ReadProfile: %w", err)

		return &models.Profile{}, "", time.Now(), err
	}

	if !hasher.CheckPass(profile.PasswordHash, payload.Password) {
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

func (uc *AuthUsecase) SignUp(ctx context.Context, user *models.SignUpPayload) (*models.Profile, string, time.Time, error) {
	if err := validator.New().Struct(user); err != nil {
		return &models.Profile{}, "", time.Now(), err
	}

	profile := &models.Profile{
		Id:           uuid.NewV4(),
		Login:        user.Login,
		Description:  "",
		ImgSrc:       "default.png",
		Phone:        user.Phone,
		PasswordHash: hasher.HashPass(user.Password),
	}

	err := uc.repo.CreateProfile(ctx, profile)
	if err != nil {
		err = fmt.Errorf("error happened in repo.CreateProfile: %w", err)

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

func (uc *AuthUsecase) CheckAuth(ctx context.Context, Id uuid.UUID) (*models.Profile, error) {
	profile, err := uc.repo.ReadProfile(ctx, Id)
	if err != nil {
		err = fmt.Errorf("error happened in repo.ReadProfile: %w", err)

		return &models.Profile{}, err
	}

	profile.HidePass()

	return profile, nil
}
