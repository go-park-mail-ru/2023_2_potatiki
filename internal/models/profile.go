package models

import uuid "github.com/satori/go.uuid"

//go:generate easyjson -all /home/scremyda/GolandProjects/2023_2_potatiki/internal/models/profile.go

//easyjson:json
type Profile struct {
	Id           uuid.UUID `json:"id"`
	Login        string    `json:"login"`
	Description  string    `json:"description,omitempty"`
	ImgSrc       string    `json:"img"`
	Phone        string    `json:"phone"`
	PasswordHash []byte    `json:"-"`
}

//
//func (p *Profile) LogValue() slog.Value {
//	return slog.GroupValue(
//		slog.String("id", p.Id.String()),
//		slog.String("login", p.Login),
//	)
//}
