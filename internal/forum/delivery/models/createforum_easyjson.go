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

func easyjson369c8e19DecodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels(in *jlexer.Lexer, out *CreateForumResponse) {
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
		case "title":
			out.Title = string(in.String())
		case "user":
			out.User = string(in.String())
		case "slug":
			out.Slug = string(in.String())
		case "posts":
			out.Posts = int(in.Int())
		case "threads":
			out.Threads = int(in.Int())
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
func easyjson369c8e19EncodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels(out *jwriter.Writer, in CreateForumResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Title != "" {
		const prefix string = ",\"title\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	if in.User != "" {
		const prefix string = ",\"user\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.User))
	}
	if in.Slug != "" {
		const prefix string = ",\"slug\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Slug))
	}
	if in.Posts != 0 {
		const prefix string = ",\"posts\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Posts))
	}
	if in.Threads != 0 {
		const prefix string = ",\"threads\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Threads))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateForumResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson369c8e19EncodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateForumResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson369c8e19EncodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateForumResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson369c8e19DecodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateForumResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson369c8e19DecodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels(l, v)
}
func easyjson369c8e19DecodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels1(in *jlexer.Lexer, out *CreateForumRequest) {
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
		case "title":
			out.Title = string(in.String())
		case "user":
			out.User = string(in.String())
		case "slug":
			out.Slug = string(in.String())
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
func easyjson369c8e19EncodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels1(out *jwriter.Writer, in CreateForumRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Title != "" {
		const prefix string = ",\"title\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	if in.User != "" {
		const prefix string = ",\"user\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.User))
	}
	if in.Slug != "" {
		const prefix string = ",\"slug\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Slug))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CreateForumRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson369c8e19EncodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CreateForumRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson369c8e19EncodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CreateForumRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson369c8e19DecodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CreateForumRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson369c8e19DecodeLonkidelyTechnoparkDbmsForumInternalForumDeliveryModels1(l, v)
}