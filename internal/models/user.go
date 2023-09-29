package models

import "github.com/google/uuid"

type User struct {
	Login        string `json:"login"`
	PasswordHash string `json:"password"`
}

func (user User) IsValid() bool {
	return len(user.Login) > 7 && len(user.Login) < 50 //Вопрос про валидацию хэша
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
