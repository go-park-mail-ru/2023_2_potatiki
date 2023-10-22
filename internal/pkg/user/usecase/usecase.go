package usecase

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/usecase"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/user"
	"github.com/google/uuid"
)

type UserUsecase struct {
	repo   user.UserRepo
	auther auth.AuthAuther
}

func NewUserUsecase(repo user.UserRepo, cfg auth.AuthConfig) *UserUsecase {
	return &UserUsecase{
		repo:   repo,
		auther: usecase.NewAuther(cfg),
	}
}

func (uc *UserUsecase) UpdatePhoto(ctx context.Context, userID uuid.UUID) error {
	return nil
}

func (uc *UserUsecase) UpdatePassword(ctx context.Context, userID uuid.UUID, password string) error {
	return nil
}

func (uc *UserUsecase) UpdateDescription(ctx context.Context, userID uuid.UUID, description string) error {
	return nil
}
