package models

import (
	"github.com/google/uuid"
	"regexp"
)

type User struct {
	Login        string `json:"login"`
	PasswordHash string `json:"password"`
}

func (user User) IsValid() bool {
	md5Regex := regexp.MustCompile("^[0-9a-fA-F]{32}$")
	return len(user.Login) > 7 && len(user.Login) < 50 && md5Regex.MatchString(user.PasswordHash)
}

type Profile struct {
	Id          uuid.UUID `json:"id"`
	Login       string    `json:"login"`
	Description string    `json:"description,omitempty"`
	ImgSrc      string    `json:"img"`
}

type UserId struct {
	Id uuid.UUID `json:"id"`
}
