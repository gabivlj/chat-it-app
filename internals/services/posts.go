package services

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/graphql/model"
)

// PostService is the service that handles post logic and needs to be implemented
type PostService interface {
	NewPost(ctx context.Context, input *model.PostForm, user *domain.User) (*domain.Post, error)
}
