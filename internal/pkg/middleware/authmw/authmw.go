package authmw

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"github.com/gorilla/mux"
)

const (
	AccessTokenCookieName = "zuzu-t"
)

func MakeTokenCookie(token string, expiration time.Time) *http.Cookie {
	return &http.Cookie{
		Name:     AccessTokenCookieName,
		Value:    token,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: true,
	}
}

func New(log *slog.Logger, auther auth.AuthAuther) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler { // TODO: del
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenCookie, err := r.Cookie(AccessTokenCookieName)
			if err != nil {
				switch {
				case errors.Is(err, http.ErrNoCookie):
					log.Debug("token cookie not found", sl.Err(err))
					resp.JSONStatus(w, http.StatusUnauthorized)

					return
				default:
					log.Error("faild to get token cookie", sl.Err(err))
					resp.JSONStatus(w, http.StatusUnauthorized)

					return
				}
			}

			claims, err := auther.GetClaims(tokenCookie.Value)
			if err != nil {
				log.Error("jws token is invalid", sl.Err(err))
				resp.JSONStatus(w, http.StatusUnauthorized)

				return
			}

			log.Info("got profile id", slog.Any("profile id", claims.ID))

			ctx := context.WithValue(r.Context(), AccessTokenCookieName, claims.ID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
