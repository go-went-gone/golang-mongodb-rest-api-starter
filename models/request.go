package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"regexp"
)

var passwordRule = []validation.Rule{
	validation.Required,
	validation.Length(8, 32),
	validation.Match(regexp.MustCompile("^\\S+$")).Error("cannot contain whitespaces"),
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a RegisterRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required, validation.Length(3, 64)),
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, passwordRule...),
	)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (a LoginRequest) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Password, passwordRule...),
	)
}

type RefreshRequest struct {
	Token string `json:"token"`
}

func (refreshRequest RefreshRequest) Validate() error {
	return validation.ValidateStruct(&refreshRequest,
		validation.Field(
			&refreshRequest.Token,
			validation.Required,
			validation.Match(regexp.MustCompile("^\\S+$")).Error("cannot contain whitespaces"),
		),
	)
}

type NoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (noteRequest NoteRequest) Validate() error {
	return validation.ValidateStruct(&noteRequest,
		validation.Field(&noteRequest.Title, validation.Required),
		validation.Field(&noteRequest.Content, validation.Required),
	)
}
