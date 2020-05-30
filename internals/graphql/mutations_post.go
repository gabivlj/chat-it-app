package graphql

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/graphql/model"
	"github.com/gabivlj/chat-it/internals/middleware"
)

func (r *mutationResolver) NewPost(ctx context.Context, form model.PostForm) (*domain.Post, error) {
	user, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	post, err := r.postRepo.NewPost(ctx, &form, user)
	return post, err
}
