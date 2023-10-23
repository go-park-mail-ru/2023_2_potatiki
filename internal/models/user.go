package models

import (
	"github.com/google/uuid"
)

type User struct {
	Login        string `json:"login"`
	PasswordHash string `json:"password"`
}

func (user User) IsValid() bool {
	// strings.Contains()
	return len(user.Login) >= 6 && len(user.Login) <= 30
}

type Profile struct {
	Id          uuid.UUID `json:"id"` //nolint:stylecheck
	Login       string    `json:"login"`
	Description string    `json:"description,omitempty"`
	ImgSrc      string    `json:"img"`
}

type UserPhoto struct {
	ID    uuid.UUID `json:"id"`
	Photo uuid.UUID `json:"photo"`
}

type UserInfo struct {
	ID              uuid.UUID `json:"id"`
	NewPasswordHash string    `json:"newPassword"`
	OldPasswordHash string    `json:"oldPassword"`
	NewDescription  string    `json:"newDescription"`
	OldDescription  string    `json:"oldDescription"`
}
