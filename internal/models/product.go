package models

import "github.com/satori/go.uuid"

type Product struct {
	Id          uuid.UUID `json:"id"` //nolint:stylecheck
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Price       int64     `json:"price"`
	ImgSrc      string    `json:"img"`
	Rating      float64   `json:"rating"` // TODO: category add
}
