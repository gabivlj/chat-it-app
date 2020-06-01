package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gabivlj/chat-it/internals/graphql/generated"
)

// NewDirectives returns the directives of this application
func NewDirectives() *generated.DirectiveRoot {
	conf := &generated.DirectiveRoot{User: func(ctx context.Context, obj interface{}, next graphql.Resolver, id string) (res interface{}, err error) {
		return next(ctx)
	}}
	return conf
}
