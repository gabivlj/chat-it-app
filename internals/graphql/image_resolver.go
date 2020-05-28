package graphql

import (
	"context"
	"fmt"

	"github.com/gabivlj/chat-it/internals/domain"
)

func (r *imageResolver) Post(ctx context.Context, obj *domain.Image) (*domain.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *imageResolver) User(ctx context.Context, obj *domain.Image) (*domain.User, error) {
	panic(fmt.Errorf("not implemented"))
}
