package graphql

// generate:go run github.com/99designs/gqlgen

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/graphql/model"
	"github.com/gabivlj/chat-it/internals/middleware"
)

func (r *userResolver) Posts(ctx context.Context, user *domain.User) ([]*domain.Post, error) {
	return middleware.DataLoaderPost(ctx).Load(user.ID)
}

func (r *userResolver) ProfileImage(ctx context.Context, user *domain.User) (*domain.Image, error) {
	return &domain.Image{URLXL: user.ImageURL, URLMD: user.ImageURL, URLSM: user.ImageURL}, nil
}

func (r *userResolver) NumberOfPosts(ctx context.Context, user *domain.User) (int, error) {
	return r.postRepo.CountPosts(ctx, user.ID)

}

func (r *userResolver) NumberOfComments(ctx context.Context, user *domain.User) (int, error) {
	return r.messageRepository.CountMessagesUser(ctx, user.ID)
}

func (r *userResolver) PostsUser(ctx context.Context, user *domain.User, params *model.Params) ([]*domain.Post, error) {
	return r.postRepo.GetPosts(ctx, params, user.ID)
}

func (r *userResolver) CommentsUser(ctx context.Context, obj *domain.User, params *model.Params) ([]*domain.Message, error) {
	return r.messageRepository.GetMessagesUsers(ctx, params, obj.ID)
}
