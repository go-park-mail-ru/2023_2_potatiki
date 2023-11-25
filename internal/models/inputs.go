package models

import (
	"html"
)

type SignUpPayload struct {
	Login    string `json:"login" validate:"required,min=6,max=30"`
	Phone    string `json:"phone" validate:"required,e164"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (p *SignUpPayload) Sanitize() {
	p.Login = html.EscapeString(p.Login)
	p.Phone = html.EscapeString(p.Phone)
	p.Password = html.EscapeString(p.Password)
}

type SignInPayload struct {
	Login    string `json:"login" validate:"required,min=6,max=30"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (p *SignInPayload) Sanitize() {
	p.Login = html.EscapeString(p.Login)
	p.Password = html.EscapeString(p.Password)

}

type UpdateProfileDataPayload struct {
	Passwords struct {
		OldPass string `json:"oldPass" validate:"omitempty,min=8,max=32"`
		NewPass string `json:"newPass" validate:"omitempty,min=8,max=32"`
	} `json:"passwords"`
	Phone string `json:"phone" validate:"omitempty,e164"`
}

func (p *UpdateProfileDataPayload) Sanitize() {
	p.Phone = html.EscapeString(p.Phone)
	p.Passwords.OldPass = html.EscapeString(p.Passwords.OldPass)
	p.Passwords.NewPass = html.EscapeString(p.Passwords.NewPass)
}

type AddressPayload struct {
	City   string `json:"city" validate:"omitempty,max=32"`
	Street string `json:"street" validate:"omitempty,max=32"`
	House  string `json:"house" validate:"omitempty,max=32"`
	Flat   string `json:"flat" validate:"omitempty,max=32"`
}

func (b *AddressPayload) Sanitize() {
	b.City = html.EscapeString(b.City)
	b.Street = html.EscapeString(b.Street)
	b.House = html.EscapeString(b.House)
	b.Flat = html.EscapeString(b.Flat)
}
