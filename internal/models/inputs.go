package models

import "html"

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

//func (u SignUpPayload) MarshalJSON() ([]byte, error) {
//	type signUpPayload SignUpPayload
//	b := signUpPayload(u)
//	b.Login = html.EscapeString(b.Login)
//	b.Phone = html.EscapeString(b.Phone)
//	b.Password = html.EscapeString(b.Password)
//	return json.Marshal(b)
//}

type SignInPayload struct {
	Login    string `json:"login" validate:"required,min=6,max=30"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (p *SignInPayload) Sanitize() {
	p.Login = html.EscapeString(p.Login)
	p.Password = html.EscapeString(p.Password)

}

//func (u SignInPayload) MarshalJSON() ([]byte, error) {
//	type signInPayload SignInPayload
//	b := signInPayload(u)
//	b.Login = html.EscapeString(b.Login)
//	b.Password = html.EscapeString(b.Password)
//	return json.Marshal(b)
//}

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

//func (u UpdateProfileDataPayload) MarshalJSON() ([]byte, error) {
//	type updateProfileDataPayload UpdateProfileDataPayload
//	b := updateProfileDataPayload(u)
//	b.Phone = html.EscapeString(b.Phone)
//	b.Passwords.OldPass = html.EscapeString(b.Passwords.OldPass)
//	b.Passwords.NewPass = html.EscapeString(b.Passwords.NewPass)
//	return json.Marshal(b)
//}
