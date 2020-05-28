package graphql

import (
	"context"
	"fmt"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/graphql/model"
)

func (r *queryResolver) Users(ctx context.Context) ([]*domain.User, error) {
	return []*domain.User{{Username: "mmm", ID: "SDSD"}}, nil
}

func (r *queryResolver) User(ctx context.Context, id model.UserQuery) (*domain.User, error) {
	return &domain.User{Username: "mmm", ID: "SDSD"}, nil
}

func (r *queryResolver) Image(ctx context.Context, id string) (*domain.Image, error) {
	return &domain.Image{URLMD: "x", URLSM: "X", URLXL: "damn"}, nil
}

func (r *queryResolver) ImageFromObject(ctx context.Context, objectID string) (*domain.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MessagesPost(ctx context.Context, id string, params *model.Params) ([]*domain.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) PostsUser(ctx context.Context, id string, params *model.Params) ([]*domain.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context, params *model.Params) ([]*domain.Post, error) {
	panic(fmt.Errorf("not implemented"))
}
