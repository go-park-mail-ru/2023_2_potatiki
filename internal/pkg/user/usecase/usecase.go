package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/user"
	"github.com/google/uuid"
)

type UserUsecase struct {
	repo user.UserRepo
}

func NewUserUsecase(repo user.UserRepo) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (uc *UserUsecase) GetProfile(context.Context, uuid.UUID) (models.Profile, error) {
	panic("unimplemented")
}
