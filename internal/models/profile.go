package models

import (
	"log/slog"

	"github.com/google/uuid"
)

type Profile struct {
	Id           uuid.UUID `json:"id"` //nolint:stylecheck
	Login        string    `json:"login"`
	Description  string    `json:"description,omitempty"`
	ImgSrc       string    `json:"img"`
	Phone        string    `json:"phone"`
	PasswordHash []byte    `json:"password"`
}

func (p *Profile) HidePass() {
	p.PasswordHash = []byte("lolkek")
}

type UserInfo struct {
	NewPassword    string `json:"newPassword"`
	NewDescription string `json:"newDescription"`
	Description    string `json:"description"`
}

type ProfileInfo struct {
	SignUpPayload
	UserInfo
}

func (p *Profile) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("id", p.Id.String()),
		slog.String("login", p.Login),
	)
}
