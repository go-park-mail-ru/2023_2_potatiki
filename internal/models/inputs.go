package models

import "html"

type SignUpPayload struct {
	Login    string `json:"login" validate:"required,min=6,max=30"`
	Phone    string `json:"phone" validate:"required,e164"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type SignInPayload struct {
	Login    string `json:"login" validate:"required,min=6,max=30"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type UpdateProfileDataPayload struct {
	Password string `json:"password" validate:"required,min=8,max=32"`
	Phone    string `json:"phone" validate:"required,e164"`
}

func (p *SignUpPayload) Sanitize() {
	p.Login = html.EscapeString(p.Login)
	p.Phone = html.EscapeString(p.Phone)
	p.Password = html.EscapeString(p.Password)
}

func (p *SignInPayload) Sanitize() {
	p.Login = html.EscapeString(p.Login)
	p.Password = html.EscapeString(p.Password)

}
func (p *UpdateProfileDataPayload) Sanitize() {
	p.Phone = html.EscapeString(p.Phone)
	p.Password = html.EscapeString(p.Password)
}
