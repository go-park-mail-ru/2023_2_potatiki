package http

import (
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/user"
)

type UserHandler struct {
	usecase user.UserUsecase
}

func NewUserHandler(usecase user.UserUsecase) UserHandler {
	return UserHandler{
		usecase: usecase,
	}
}

func (h *UserHandler) Profile(w http.ResponseWriter, r *http.Request) {
	panic("uninplemented")
}
