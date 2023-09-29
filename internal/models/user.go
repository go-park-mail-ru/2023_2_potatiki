package models

import "github.com/google/uuid"

type User struct {
	Login        string `json:"login"`
	PasswordHash string `json:"password"`
}

type Profile struct {
	Id          uuid.UUID `json:"id"`
	Login       string    `json:"login"`
	Description string    `json:"description,omitempty"`
	ImgSrc      string    `json:"img"`
}
