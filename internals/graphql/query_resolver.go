package graphql

import (
	"context"
	"errors"
	"fmt"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/graphql/model"
)

func (r *queryResolver) Users(ctx context.Context) ([]*domain.User, error) {
	return []*domain.User{{Username: "mmm", ID: "SDSD"}}, nil
}

func (r *queryResolver) User(ctx context.Context, id model.UserQuery) (*domain.User, error) {
	if id.ID != nil {
		return r.userRepo.FindByID(ctx, *id.ID)
	}
	if id.Username == nil {
		return nil, errors.New("Username or ID params need to be filled")
	}
	return r.userRepo.FindByUsername(ctx, *id.Username)
}

func (r *queryResolver) Image(ctx context.Context, id string) (*domain.Image, error) {
	return &domain.Image{URLMD: "x", URLSM: "X", URLXL: "damn"}, nil
}

func (r *queryResolver) ImageFromObject(ctx context.Context, objectID string) (*domain.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MessagesPost(ctx context.Context, id string, params *model.Params) ([]*domain.Message, error) {
	return r.messageRepository.GetMessages(ctx, id, params)
}

func (r *queryResolver) PostsUser(ctx context.Context, id string, params *model.Params) ([]*domain.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context, params *model.Params) ([]*domain.Post, error) {
	return r.postRepo.GetPosts(ctx, params)
}

func (r *queryResolver) Post(ctx context.Context, id string) (*domain.Post, error) {
	return r.postRepo.GetPost(ctx, id)
}
