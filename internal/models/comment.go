package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//easyjson:json
type Comment struct {
	ID           uuid.UUID `json:"id"`
	UserName     string    `json:"userName"`
	CreationDate time.Time `json:"creationDate"`
	ProductID    uuid.UUID `json:"productId"`
	Pros         string    `json:"pros"`
	Cons         string    `json:"cons"`
	Comment      string    `json:"comment"`
	Rating       int       `json:"rating"`
}

//easyjson:json
type CommentSlice []Comment
