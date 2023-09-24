package models

import "github.com/google/uuid"

type Product struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Price       int64     `json:"price"`
}
