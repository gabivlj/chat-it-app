// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/gabivlj/chat-it/internals/domain"
)

type FormLogInRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Params struct {
	Before *string `json:"before"`
	After  *string `json:"after"`
	Limit  int     `json:"limit"`
}

type PostForm struct {
	Image *graphql.Upload `json:"image"`
	Text  string          `json:"text"`
	Title string          `json:"title"`
}

type UserQuery struct {
	Username *string `json:"username"`
	ID       *string `json:"id"`
}

type UserSession struct {
	User    *domain.User `json:"user"`
	Session string       `json:"session"`
}
