package graphql

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/middleware"
)

func (r *messageResolver) User(ctx context.Context, obj *domain.Message) (*domain.User, error) {
	if obj.User == nil {
		return middleware.DataLoaderUser(ctx).Load(obj.UserID)
	}
	return obj.User, nil
}

// Warning, we shouldn't really use this
func (r *messageResolver) Post(ctx context.Context, obj *domain.Message) (*domain.Post, error) {
	if obj.Post == nil {
		return r.postRepo.GetPost(ctx, obj.PostID)
	}

	return obj.Post, nil
}
