package services

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/graphql/model"
)

// PostService is the service that handles post logic and needs to be implemented
type PostService interface {
	NewPost(ctx context.Context, input *model.PostForm, user *domain.User) (*domain.Post, error)
	GetPost(ctx context.Context, id string) (*domain.Post, error)
	GetPosts(ctx context.Context, pagination *model.Params, users ...string) ([]*domain.Post, error)
	GetPostsByIDs(ctx context.Context, ids ...string) ([]*domain.Post, error)
	GetPostsFromUsers(ctx context.Context, userIDs []string) ([][]*domain.Post, error)
	CountPosts(ctx context.Context, userID string) (int, error)
}
