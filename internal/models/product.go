package models

import uuid "github.com/satori/go.uuid"

//easyjson:json
type Product struct {
	Id            uuid.UUID `json:"productId"`
	Name          string    `json:"productName"`
	Description   string    `json:"description,omitempty"`
	Price         int64     `json:"price"`
	ImgSrc        string    `json:"img"`
	Rating        float32   `json:"rating"`
	CountComments int64     `json:"countComments"`
	Category      `json:"category"`
}

//easyjson:json
type ProductSlice []Product
