package graphql

import (
	"context"
	"fmt"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/middleware"
)

func (r *userResolver) Posts(ctx context.Context, obj *domain.User) ([]*domain.Post, error) {
	return middleware.DataLoaderPost(ctx).Load(obj.ID)
}

func (r *userResolver) ProfileImage(ctx context.Context, obj *domain.User) (*domain.Image, error) {
	panic(fmt.Errorf("not implemented"))
}
