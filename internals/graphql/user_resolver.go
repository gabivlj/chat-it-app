package graphql

import (
	"context"
	"fmt"

	"github.com/gabivlj/chat-it/internals/domain"
)

func (r *userResolver) Posts(ctx context.Context, obj *domain.User) ([]*domain.Post, error) {
	return []*domain.Post{{Text: "DAAAMN", Title: "xDDD", ID: "XDDDD"}}, nil
}

func (r *userResolver) ProfileImage(ctx context.Context, obj *domain.User) (*domain.Image, error) {
	panic(fmt.Errorf("not implemented"))
}
