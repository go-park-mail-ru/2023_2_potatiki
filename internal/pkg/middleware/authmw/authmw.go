package authmw

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/jwter"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"
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

func New(log *slog.Logger, authJWT jwter.JWTer) mux.MiddlewareFunc {
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

			//log.Debug("AUTH MW", "token", tokenCookie.Value)

			ID, err := authJWT.DecodeAuthToken(tokenCookie.Value)
			if err != nil {
				log.Error("jws token is invalid auth", sl.Err(err))
				resp.JSONStatus(w, http.StatusUnauthorized)

				return
			}

			log.Info("got profile id", slog.Any("profile id", ID))

			ctx := context.WithValue(r.Context(), AccessTokenCookieName, ID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
