package models

import (
	"log/slog"

	"github.com/google/uuid"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (u *User) IsValid() bool {
	// strings.Contains()
	return len(u.Login) >= 6 && len(u.Login) <= 30
}

type Profile struct {
	Id           uuid.UUID `json:"id"` //nolint:stylecheck
	Login        string    `json:"login"`
	Description  string    `json:"description,omitempty"`
	ImgSrc       string    `json:"img"`
	PasswordHash []byte    `json:"password"`
}

type UserInfo struct {
	NewPassword    string `json:"newPassword"`
	NewDescription string `json:"newDescription"`
	Description    string `json:"description"`
}

type ProfileInfo struct {
	User
	UserInfo
}

func (p *Profile) LogValue() slog.Value {
	//nolint:lll
	// check https://betterstack.com/community/guides/logging/logging-in-go/#hiding-sensitive-fields-with-the-logvaluer-interface
	return slog.GroupValue(
		slog.String("id", p.Id.String()),
		slog.String("login", p.Login),
	)
}
