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

func easyjson120d1ca2DecodeGithubComGoParkMailRu20232PotatikiInternalModels(in *jlexer.Lexer, out *OrderSlice) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(OrderSlice, 0, 0)
			} else {
				*out = OrderSlice{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Order
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
func easyjson120d1ca2EncodeGithubComGoParkMailRu20232PotatikiInternalModels(out *jwriter.Writer, in OrderSlice) {
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
func (v OrderSlice) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeGithubComGoParkMailRu20232PotatikiInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrderSlice) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeGithubComGoParkMailRu20232PotatikiInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrderSlice) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeGithubComGoParkMailRu20232PotatikiInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrderSlice) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeGithubComGoParkMailRu20232PotatikiInternalModels(l, v)
}
func easyjson120d1ca2DecodeGithubComGoParkMailRu20232PotatikiInternalModels1(in *jlexer.Lexer, out *OrderProduct) {
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
		case "quantity":
			out.Quantity = int64(in.Int64())
		case "productId":
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Id).UnmarshalText(data))
			}
		case "productName":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "price":
			out.Price = int64(in.Int64())
		case "img":
			out.ImgSrc = string(in.String())
		case "rating":
			out.Rating = float32(in.Float32())
		case "countComments":
			out.CountComments = int64(in.Int64())
		case "category":
			(out.Category).UnmarshalEasyJSON(in)
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
func easyjson120d1ca2EncodeGithubComGoParkMailRu20232PotatikiInternalModels1(out *jwriter.Writer, in OrderProduct) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"quantity\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Quantity))
	}
	{
		const prefix string = ",\"productId\":"
		out.RawString(prefix)
		out.RawText((in.Id).MarshalText())
	}
	{
		const prefix string = ",\"productName\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	if in.Description != "" {
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.Int64(int64(in.Price))
	}
	{
		const prefix string = ",\"img\":"
		out.RawString(prefix)
		out.String(string(in.ImgSrc))
	}
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix)
		out.Float32(float32(in.Rating))
	}
	{
		const prefix string = ",\"countComments\":"
		out.RawString(prefix)
		out.Int64(int64(in.CountComments))
	}
	{
		const prefix string = ",\"category\":"
		out.RawString(prefix)
		(in.Category).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OrderProduct) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeGithubComGoParkMailRu20232PotatikiInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrderProduct) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeGithubComGoParkMailRu20232PotatikiInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrderProduct) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeGithubComGoParkMailRu20232PotatikiInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrderProduct) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeGithubComGoParkMailRu20232PotatikiInternalModels1(l, v)
}
func easyjson120d1ca2DecodeGithubComGoParkMailRu20232PotatikiInternalModels2(in *jlexer.Lexer, out *OrderInfo) {
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
		case "deliveryDate":
			out.DeliveryAtDate = string(in.String())
		case "deliveryTime":
			out.DeliveryAtTime = string(in.String())
		case "promocodeName":
			out.PromocodeName = string(in.String())
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
func easyjson120d1ca2EncodeGithubComGoParkMailRu20232PotatikiInternalModels2(out *jwriter.Writer, in OrderInfo) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"deliveryDate\":"
		out.RawString(prefix[1:])
		out.String(string(in.DeliveryAtDate))
	}
	{
		const prefix string = ",\"deliveryTime\":"
		out.RawString(prefix)
		out.String(string(in.DeliveryAtTime))
	}
	if in.PromocodeName != "" {
		const prefix string = ",\"promocodeName\":"
		out.RawString(prefix)
		out.String(string(in.PromocodeName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OrderInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeGithubComGoParkMailRu20232PotatikiInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OrderInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeGithubComGoParkMailRu20232PotatikiInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OrderInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeGithubComGoParkMailRu20232PotatikiInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OrderInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeGithubComGoParkMailRu20232PotatikiInternalModels2(l, v)
}
func easyjson120d1ca2DecodeGithubComGoParkMailRu20232PotatikiInternalModels3(in *jlexer.Lexer, out *Order) {
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
			if data := in.UnsafeBytes(); in.Ok() {
				in.AddError((out.Id).UnmarshalText(data))
			}
		case "status":
			out.Status = string(in.String())
		case "deliveryDate":
			out.DeliveryDate = string(in.String())
		case "deliveryTime":
			out.DeliveryTime = string(in.String())
		case "promocodeName":
			out.PomocodeName = string(in.String())
		case "creationDate":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreationAt).UnmarshalJSON(data))
			}
		case "address":
			(out.Address).UnmarshalEasyJSON(in)
		case "products":
			if in.IsNull() {
				in.Skip()
				out.Products = nil
			} else {
				in.Delim('[')
				if out.Products == nil {
					if !in.IsDelim(']') {
						out.Products = make([]OrderProduct, 0, 0)
					} else {
						out.Products = []OrderProduct{}
					}
				} else {
					out.Products = (out.Products)[:0]
				}
				for !in.IsDelim(']') {
					var v4 OrderProduct
					(v4).UnmarshalEasyJSON(in)
					out.Products = append(out.Products, v4)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson120d1ca2EncodeGithubComGoParkMailRu20232PotatikiInternalModels3(out *jwriter.Writer, in Order) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.RawText((in.Id).MarshalText())
	}
	{
		const prefix string = ",\"status\":"
		out.RawString(prefix)
		out.String(string(in.Status))
	}
	{
		const prefix string = ",\"deliveryDate\":"
		out.RawString(prefix)
		out.String(string(in.DeliveryDate))
	}
	{
		const prefix string = ",\"deliveryTime\":"
		out.RawString(prefix)
		out.String(string(in.DeliveryTime))
	}
	{
		const prefix string = ",\"promocodeName\":"
		out.RawString(prefix)
		out.String(string(in.PomocodeName))
	}
	{
		const prefix string = ",\"creationDate\":"
		out.RawString(prefix)
		out.Raw((in.CreationAt).MarshalJSON())
	}
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix)
		(in.Address).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"products\":"
		out.RawString(prefix)
		if in.Products == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Products {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Order) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson120d1ca2EncodeGithubComGoParkMailRu20232PotatikiInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Order) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson120d1ca2EncodeGithubComGoParkMailRu20232PotatikiInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Order) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson120d1ca2DecodeGithubComGoParkMailRu20232PotatikiInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Order) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson120d1ca2DecodeGithubComGoParkMailRu20232PotatikiInternalModels3(l, v)
}
