package models

import (
	"encoding/json"
	"html"

	uuid "github.com/satori/go.uuid"
)

type Address struct {
	Id        uuid.UUID `json:"addressId"`
	ProfileId uuid.UUID `json:"-"`
	City      string    `json:"city"`
	Street    string    `json:"street"`
	House     string    `json:"house"`
	Flat      string    `json:"flat"`
	IsCurrent bool      `json:"addressIsCurrent"`
}

func (b *Address) Sanitize() {
	b.City = html.EscapeString(b.City)
	b.Street = html.EscapeString(b.Street)
	b.House = html.EscapeString(b.House)
	b.Flat = html.EscapeString(b.Flat)
}

func (u *Address) MarshalJSON() ([]byte, error) {
	type address *Address
	b := address(u)
	b.City = html.EscapeString(b.City)
	b.Street = html.EscapeString(b.Street)
	b.House = html.EscapeString(b.House)
	b.Flat = html.EscapeString(b.Flat)
	return json.Marshal(b)
}

type AddressDelete struct {
	ProfileId uuid.UUID `json:"-"`
	Id        uuid.UUID `json:"addressId"`
}

type AddressMakeCurrent struct {
	ProfileId uuid.UUID `json:"-"`
	Id        uuid.UUID `json:"addressId"`
}
