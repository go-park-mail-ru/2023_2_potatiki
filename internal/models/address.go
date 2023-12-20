package models

import (
	"html"

	uuid "github.com/satori/go.uuid"
)

//go:generate easyjson -all /home/scremyda/GolandProjects/2023_2_potatiki/internal/models/address.go

//easyjson:json
type Address struct {
	Id        uuid.UUID `json:"addressId"`
	ProfileId uuid.UUID `json:"-"`
	City      string    `json:"city"`
	Street    string    `json:"street"`
	House     string    `json:"house"`
	Flat      string    `json:"flat"`
	IsCurrent bool      `json:"addressIsCurrent"`
}

//easyjson:json
type AddressSlice []Address

func (b *Address) Sanitize() {
	b.City = html.EscapeString(b.City)
	b.Street = html.EscapeString(b.Street)
	b.House = html.EscapeString(b.House)
	b.Flat = html.EscapeString(b.Flat)
}

//easyjson:json
type AddressDelete struct {
	ProfileId uuid.UUID `json:"-"`
	Id        uuid.UUID `json:"addressId"`
}

//easyjson:json
type AddressMakeCurrent struct {
	ProfileId uuid.UUID `json:"-"`
	Id        uuid.UUID `json:"addressId"`
}
