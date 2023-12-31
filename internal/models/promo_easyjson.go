// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson47107b8bDecodeGithubComGoParkMailRu20232PotatikiInternalModels(in *jlexer.Lexer, out *Promocode) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int64(in.Int64())
		case "discount":
			out.Discount = int(in.Int())
		case "name":
			out.Name = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson47107b8bEncodeGithubComGoParkMailRu20232PotatikiInternalModels(out *jwriter.Writer, in Promocode) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Id))
	}
	{
		const prefix string = ",\"discount\":"
		out.RawString(prefix)
		out.Int(int(in.Discount))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Promocode) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson47107b8bEncodeGithubComGoParkMailRu20232PotatikiInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Promocode) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson47107b8bEncodeGithubComGoParkMailRu20232PotatikiInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Promocode) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson47107b8bDecodeGithubComGoParkMailRu20232PotatikiInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Promocode) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson47107b8bDecodeGithubComGoParkMailRu20232PotatikiInternalModels(l, v)
}
