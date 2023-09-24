package usecase

import "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"

type AuthUsecase struct {
	repo auth.AuthRepo
}

func NewAuthUsecase()
