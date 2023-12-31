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

func easyjsonF4fdf71eDecodeGithubComGoParkMailRu20232PotatikiInternalModels(in *jlexer.Lexer, out *AddressSlice) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(AddressSlice, 0, 0)
			} else {
				*out = AddressSlice{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Address
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonF4fdf71eEncodeGithubComGoParkMailRu20232PotatikiInternalModels(out *jwriter.Writer, in AddressSlice) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v AddressSlice) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fdf71eEncodeGithubComGoParkMailRu20232PotatikiInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AddressSlice) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fdf71eEncodeGithubComGoParkMailRu20232PotatikiInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AddressSlice) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fdf71eDecodeGithubComGoParkMailRu20232PotatikiInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AddressSlice) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fdf71eDecodeGithubComGoParkMailRu20232PotatikiInternalModels(l, v)
}
func easyjsonF4fdf71eDecodeGithubComGoParkMailRu20232PotatikiInternalModels1(in *jlexer.Lexer, out *AddressMakeCurrent) {
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
		case "addressId":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Id).UnmarshalText(data))
			}
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
func easyjsonF4fdf71eEncodeGithubComGoParkMailRu20232PotatikiInternalModels1(out *jwriter.Writer, in AddressMakeCurrent) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"addressId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.RawText((in.Id).MarshalText())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AddressMakeCurrent) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fdf71eEncodeGithubComGoParkMailRu20232PotatikiInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AddressMakeCurrent) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fdf71eEncodeGithubComGoParkMailRu20232PotatikiInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AddressMakeCurrent) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fdf71eDecodeGithubComGoParkMailRu20232PotatikiInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AddressMakeCurrent) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fdf71eDecodeGithubComGoParkMailRu20232PotatikiInternalModels1(l, v)
}
func easyjsonF4fdf71eDecodeGithubComGoParkMailRu20232PotatikiInternalModels2(in *jlexer.Lexer, out *AddressDelete) {
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
		case "addressId":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Id).UnmarshalText(data))
			}
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
func easyjsonF4fdf71eEncodeGithubComGoParkMailRu20232PotatikiInternalModels2(out *jwriter.Writer, in AddressDelete) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"addressId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.RawText((in.Id).MarshalText())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AddressDelete) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fdf71eEncodeGithubComGoParkMailRu20232PotatikiInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AddressDelete) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fdf71eEncodeGithubComGoParkMailRu20232PotatikiInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AddressDelete) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fdf71eDecodeGithubComGoParkMailRu20232PotatikiInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AddressDelete) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fdf71eDecodeGithubComGoParkMailRu20232PotatikiInternalModels2(l, v)
}
func easyjsonF4fdf71eDecodeGithubComGoParkMailRu20232PotatikiInternalModels3(in *jlexer.Lexer, out *Address) {
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
		case "addressId":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Id).UnmarshalText(data))
			}
		case "city":
			out.City = string(in.String())
		case "street":
			out.Street = string(in.String())
		case "house":
			out.House = string(in.String())
		case "flat":
			out.Flat = string(in.String())
		case "addressIsCurrent":
			out.IsCurrent = bool(in.Bool())
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
func easyjsonF4fdf71eEncodeGithubComGoParkMailRu20232PotatikiInternalModels3(out *jwriter.Writer, in Address) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"addressId\":"
		out.RawString(prefix[1:])
		out.RawText((in.Id).MarshalText())
	}
	{
		const prefix string = ",\"city\":"
		out.RawString(prefix)
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"street\":"
		out.RawString(prefix)
		out.String(string(in.Street))
	}
	{
		const prefix string = ",\"house\":"
		out.RawString(prefix)
		out.String(string(in.House))
	}
	{
		const prefix string = ",\"flat\":"
		out.RawString(prefix)
		out.String(string(in.Flat))
	}
	{
		const prefix string = ",\"addressIsCurrent\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsCurrent))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Address) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonF4fdf71eEncodeGithubComGoParkMailRu20232PotatikiInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Address) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonF4fdf71eEncodeGithubComGoParkMailRu20232PotatikiInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Address) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonF4fdf71eDecodeGithubComGoParkMailRu20232PotatikiInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Address) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonF4fdf71eDecodeGithubComGoParkMailRu20232PotatikiInternalModels3(l, v)
}
