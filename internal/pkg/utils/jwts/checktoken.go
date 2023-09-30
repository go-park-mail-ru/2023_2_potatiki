package jwts

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func CheckToken(r *http.Request) (bool, error) {
	cookie, err := r.Cookie("Default")
	if err != nil {
		return false, err
	}
	_, err = jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
