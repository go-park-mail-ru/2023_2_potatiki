package http

import (
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/user"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	log *slog.Logger
	uc  user.UserUsecase
}

func NewUserHandler(log *slog.Logger, uc user.UserUsecase) *UserHandler {
	return &UserHandler{
		log: log,
		uc:  uc,
	}
}

func (h *UserHandler) UpdatePhoto(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) UpdateDescription(w http.ResponseWriter, r *http.Request) {

}
