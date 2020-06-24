package graphql

// generate:go run github.com/99designs/gqlgen

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/middleware"
)

func (r *userResolver) Posts(ctx context.Context, obj *domain.User) ([]*domain.Post, error) {
	return middleware.DataLoaderPost(ctx).Load(obj.ID)
}

func (r *userResolver) ProfileImage(ctx context.Context, obj *domain.User) (*domain.Image, error) {
	return &domain.Image{URLXL: obj.ImageURL, URLMD: obj.ImageURL, URLSM: obj.ImageURL}, nil
}

func (r *userResolver) NumberOfPosts(ctx context.Context, obj *domain.User) (int, error) {
	return r.postRepo.CountPosts(ctx, obj.ID)

}

func (r *userResolver) NumberOfComments(ctx context.Context, obj *domain.User) (int, error) {
	return r.messageRepository.CountMessagesUser(ctx, obj.ID)
}
