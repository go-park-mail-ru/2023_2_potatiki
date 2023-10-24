package middleware

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/coockie"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"github.com/gorilla/mux"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST,PUT,DELETE,GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Authenticate(log *slog.Logger, auther auth.AuthAuther) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenCookie, err := r.Cookie(coockie.AccessTokenCookieName)

			if err != nil {
				switch {
				case errors.Is(err, http.ErrNoCookie):
					log.Error("token cookie not found", sl.Err(err))
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

			ctx := context.WithValue(r.Context(), coockie.AccessTokenCookieName, claims.ID)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
