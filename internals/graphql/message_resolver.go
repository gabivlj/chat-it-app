package graphql

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
)

func (r *messageResolver) User(ctx context.Context, obj *domain.Message) (*domain.User, error) {
	return obj.User, nil
}

func (r *messageResolver) Post(ctx context.Context, obj *domain.Message) (*domain.Post, error) {
	return obj.Post, nil
}
