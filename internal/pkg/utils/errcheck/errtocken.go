package errcheck

import (
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"log/slog"
	"net/http"
)

func TokenCookieErr(err error, log *slog.Logger, w http.ResponseWriter) bool {
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			log.Error("token cookie not found", sl.Err(err))
			resp.JSONStatus(w, http.StatusUnauthorized)

			return true
		default:
			log.Error("faild to get token cookie", sl.Err(err))
			resp.JSONStatus(w, http.StatusUnauthorized)

			return true
		}
	}

	return false
}
