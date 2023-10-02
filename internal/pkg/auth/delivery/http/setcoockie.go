package http

import (
	"net/http"
	"time"
)

const (
	AccessTokenCookieName = "zuzu-t"
)

func getTokenCookie(name, token string, expiration time.Time) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    token,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // false ?
	}
}
