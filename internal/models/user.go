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
	md5Regex := regexp.MustCompile("^[A-Za-z0-9!@#$%^&*()-_+=<>?]+$")
	return len(user.Login) > 6 && len(user.Login) < 30 && md5Regex.MatchString(user.Login)
}

type Profile struct {
	Id          uuid.UUID `json:"id"`
	Login       string    `json:"login"`
	Description string    `json:"description,omitempty"`
	ImgSrc      string    `json:"img"`
}
