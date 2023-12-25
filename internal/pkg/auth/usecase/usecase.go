package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/profile"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/hasher"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/jwter"

	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

type AuthUsecase struct {
	repo    profile.ProfileRepo
	authJWT jwter.JWTer
}

func NewAuthUsecase(repo profile.ProfileRepo, cfg jwter.Configer) *AuthUsecase {
	return &AuthUsecase{
		repo:    repo,
		authJWT: jwter.New(cfg),
	}
}

var (
	ErrPassMismatch = errors.New("password does not match")
)

func (uc *AuthUsecase) SignIn(ctx context.Context, payload *models.SignInPayload) (*models.Profile, string, time.Time, error) {
	if err := validator.New().Struct(payload); err != nil {
		return &models.Profile{}, "", time.Now(), err
	}
	payload.Sanitize()

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

	token, exp, err := uc.authJWT.EncodeAuthToken(profile.Id)
	if err != nil {
		err = fmt.Errorf("error happened in Auther.GenerateToken: %w", err)

		return &models.Profile{}, "", time.Now(), err
	}

	return profile, token, exp, nil
}

func (uc *AuthUsecase) SignUp(ctx context.Context, payload *models.SignUpPayload) (*models.Profile, string, time.Time, error) {
	if err := validator.New().Struct(payload); err != nil {
		return &models.Profile{}, "", time.Now(), err
	}
	payload.Sanitize()

	profile := &models.Profile{
		Id:           uuid.NewV4(),
		Login:        payload.Login,
		Description:  "",
		Phone:        payload.Phone,
		PasswordHash: hasher.HashPass(payload.Password),
	}

	err := uc.repo.CreateProfile(ctx, profile)
	if err != nil {
		err = fmt.Errorf("error happened in repo.CreateProfile: %w", err)

		return &models.Profile{}, "", time.Now(), err
	}

	token, exp, err := uc.authJWT.EncodeAuthToken(profile.Id)
	if err != nil {
		err = fmt.Errorf("error happened in Auther.GenerateToken: %w", err)

		return &models.Profile{}, "", time.Now(), err
	}

	return profile, token, exp, nil
}

func (uc *AuthUsecase) CheckAuth(ctx context.Context, Id uuid.UUID) (*models.Profile, error) {
	profile, err := uc.repo.ReadProfile(ctx, Id)
	if err != nil {
		err = fmt.Errorf("error happened in repo.ReadProfile: %w", err)

		return &models.Profile{}, err
	}

	return profile, nil
}
