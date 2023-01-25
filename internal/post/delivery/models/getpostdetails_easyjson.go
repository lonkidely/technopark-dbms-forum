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

func easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels(in *jlexer.Lexer, out *PostGetDetailsResponse) {
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
		case "post":
			if in.IsNull() {
				in.Skip()
				out.Post = nil
			} else {
				if out.Post == nil {
					out.Post = new(GetPostDetailsPostResponse)
				}
				(*out.Post).UnmarshalEasyJSON(in)
			}
		case "thread":
			if in.IsNull() {
				in.Skip()
				out.Thread = nil
			} else {
				if out.Thread == nil {
					out.Thread = new(GetPostDetailsThreadResponse)
				}
				(*out.Thread).UnmarshalEasyJSON(in)
			}
		case "author":
			if in.IsNull() {
				in.Skip()
				out.Author = nil
			} else {
				if out.Author == nil {
					out.Author = new(GetPostDetailsAuthorResponse)
				}
				(*out.Author).UnmarshalEasyJSON(in)
			}
		case "forum":
			if in.IsNull() {
				in.Skip()
				out.Forum = nil
			} else {
				if out.Forum == nil {
					out.Forum = new(GetPostDetailsForumResponse)
				}
				(*out.Forum).UnmarshalEasyJSON(in)
			}
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
func easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels(out *jwriter.Writer, in PostGetDetailsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Post != nil {
		const prefix string = ",\"post\":"
		first = false
		out.RawString(prefix[1:])
		(*in.Post).MarshalEasyJSON(out)
	}
	if in.Thread != nil {
		const prefix string = ",\"thread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Thread).MarshalEasyJSON(out)
	}
	if in.Author != nil {
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Author).MarshalEasyJSON(out)
	}
	if in.Forum != nil {
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Forum).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PostGetDetailsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PostGetDetailsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PostGetDetailsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PostGetDetailsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels(l, v)
}
func easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels1(in *jlexer.Lexer, out *GetPostDetailsThreadResponse) {
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
			out.ID = int(in.Int())
		case "title":
			out.Title = string(in.String())
		case "author":
			out.Author = string(in.String())
		case "forum":
			out.Forum = string(in.String())
		case "slug":
			out.Slug = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
		case "votes":
			out.Votes = int(in.Int())
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
func easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels1(out *jwriter.Writer, in GetPostDetailsThreadResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	if in.Author != "" {
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Author))
	}
	if in.Forum != "" {
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Forum))
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
	if in.Message != "" {
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	if true {
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Created).MarshalJSON())
	}
	if in.Votes != 0 {
		const prefix string = ",\"votes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Votes))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetPostDetailsThreadResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetPostDetailsThreadResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetPostDetailsThreadResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetPostDetailsThreadResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels1(l, v)
}
func easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels2(in *jlexer.Lexer, out *GetPostDetailsPostResponse) {
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
			out.ID = int(in.Int())
		case "parent":
			out.Parent = int(in.Int())
		case "author":
			out.Author = string(in.String())
		case "message":
			out.Message = string(in.String())
		case "isEdited":
			out.IsEdited = bool(in.Bool())
		case "forum":
			out.Forum = string(in.String())
		case "thread":
			out.Thread = int(in.Int())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
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
func easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels2(out *jwriter.Writer, in GetPostDetailsPostResponse) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int(int(in.ID))
	}
	if in.Parent != 0 {
		const prefix string = ",\"parent\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Parent))
	}
	if in.Author != "" {
		const prefix string = ",\"author\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Author))
	}
	if in.Message != "" {
		const prefix string = ",\"message\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Message))
	}
	if in.IsEdited {
		const prefix string = ",\"isEdited\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.IsEdited))
	}
	if in.Forum != "" {
		const prefix string = ",\"forum\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Forum))
	}
	if in.Thread != 0 {
		const prefix string = ",\"thread\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Thread))
	}
	if true {
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Created).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GetPostDetailsPostResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetPostDetailsPostResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetPostDetailsPostResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetPostDetailsPostResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels2(l, v)
}
func easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels3(in *jlexer.Lexer, out *GetPostDetailsForumResponse) {
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
func easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels3(out *jwriter.Writer, in GetPostDetailsForumResponse) {
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
func (v GetPostDetailsForumResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetPostDetailsForumResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetPostDetailsForumResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetPostDetailsForumResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels3(l, v)
}
func easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels4(in *jlexer.Lexer, out *GetPostDetailsAuthorResponse) {
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
func easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels4(out *jwriter.Writer, in GetPostDetailsAuthorResponse) {
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
func (v GetPostDetailsAuthorResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GetPostDetailsAuthorResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson539848a4EncodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GetPostDetailsAuthorResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GetPostDetailsAuthorResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson539848a4DecodeLonkidelyTechnoparkDbmsForumInternalPostDeliveryModels4(l, v)
}
