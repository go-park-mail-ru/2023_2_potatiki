package cookie

import (
	"net/http"
	"time"
)

const (
	AccessTokenCookieName = "zuzu-t"
)

func GetTokenCookie(name, token string, expiration time.Time) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    token,
		Expires:  expiration,
		Path:     "/",
		HttpOnly: true,
	}
}
