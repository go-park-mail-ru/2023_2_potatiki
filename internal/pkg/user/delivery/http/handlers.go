package http

import (
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/user"
	"log/slog"
	"net/http"
)

type UserHandler struct {
	log      *slog.Logger
	ucUser   user.UserUsecase
	ucAuther auth.AuthUsecase
}

func NewUserHandler(log *slog.Logger, uc user.UserUsecase) *UserHandler {
	return &UserHandler{
		log:    log,
		ucUser: uc,
		// ucAuther:
	}
}

func (h *UserHandler) UpdatePhoto(w http.ResponseWriter, r *http.Request) {

}

func (h *UserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {

}

// TODO: одна ручка для всех текстовых полей
func (h *UserHandler) UpdateDescription(w http.ResponseWriter, r *http.Request) {

}
