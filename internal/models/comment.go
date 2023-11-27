package models

import uuid "github.com/satori/go.uuid"

type Comment struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"productId"`
	Pros      string    `json:"pros"`
	Cons      string    `json:"cons"`
	Comment   string    `json:"comment"`
	Rating    int       `json:"rating"`
}
