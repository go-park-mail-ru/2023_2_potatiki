package http

import (
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, token string, ttlive time.Time) {
	LoginCookie := &http.Cookie{
		Name:     "Default",
		Value:    token,
		HttpOnly: true,
		Expires:  ttlive,
		Secure:   true,
	}
	http.SetCookie(w, LoginCookie)
}
