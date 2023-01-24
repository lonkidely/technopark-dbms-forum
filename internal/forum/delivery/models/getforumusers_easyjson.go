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

func easyjson5e434b3DecodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels(in *jlexer.Lexer, out *GetForumUsersResponse) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(GetForumUsersResponse, 0, 1)
			} else {
				*out = GetForumUsersResponse{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 GetForumUserResponse
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
func easyjson5e434b3EncodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels(out *jwriter.Writer, in GetForumUsersResponse) {
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
func (v GetForumUsersResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5e434b3EncodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetForumUsersResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5e434b3EncodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetForumUsersResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5e434b3DecodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetForumUsersResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5e434b3DecodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels(l, v)
}
func easyjson5e434b3DecodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels1(in *jlexer.Lexer, out *GetForumUserResponse) {
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
		case "nickname":
			out.Nickname = string(in.String())
		case "fullname":
			out.FullName = string(in.String())
		case "about":
			out.About = string(in.String())
		case "email":
			out.Email = string(in.String())
		default:
			in.AddError(&jlexer.LexerError{
				Offset: in.GetPos(),
				Reason: "unknown field",
				Data:   key,
			})
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5e434b3EncodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels1(out *jwriter.Writer, in GetForumUserResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Nickname != "" {
		const prefix string = ",\"nickname\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Nickname))
	}
	if in.FullName != "" {
		const prefix string = ",\"fullname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FullName))
	}
	if in.About != "" {
		const prefix string = ",\"about\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.About))
	}
	if in.Email != "" {
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetForumUserResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5e434b3EncodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetForumUserResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5e434b3EncodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetForumUserResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5e434b3DecodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetForumUserResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5e434b3DecodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels1(l, v)
}