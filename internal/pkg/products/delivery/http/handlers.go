package http

import (
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
)

type AuthHandler struct {
	useCase auth.AuthUsecase
}

func NewAuthHandler(newUseCase auth.AuthUsecase) AuthHandler {
	return AuthHandler{
		useCase: newUseCase,
	}
}
