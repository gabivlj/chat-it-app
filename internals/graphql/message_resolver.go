package graphql

import (
	"context"
	"fmt"

	"github.com/gabivlj/chat-it/internals/domain"
)

func (r *messageResolver) User(ctx context.Context, obj *domain.Message) (*domain.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *messageResolver) Post(ctx context.Context, obj *domain.Message) (*domain.Post, error) {
	panic(fmt.Errorf("not implemented"))
}
