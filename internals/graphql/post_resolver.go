package graphql

import (
	"context"
	"fmt"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/middleware"
)

func (r *postResolver) User(ctx context.Context, obj *domain.Post) (*domain.User, error) {
	return middleware.DataLoaderUser(ctx).Load(obj.UserID)
}

func (r *postResolver) Image(ctx context.Context, obj *domain.Post) (*domain.Image, error) {
	return &domain.Image{URLMD: obj.URLImage, URLXL: obj.URLImage, URLSM: obj.URLImage}, nil
}

// TODO dataLoader for []postIDS
func (r *postResolver) Chat(ctx context.Context, obj *domain.Post) ([]*domain.Message, error) {
	panic(fmt.Errorf("not implemented"))
}
