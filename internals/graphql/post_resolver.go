package graphql

import (
	"context"
	"fmt"

	"github.com/gabivlj/chat-it/internals/domain"
)

// TODO dataLoader for userIDs
func (r *postResolver) User(ctx context.Context, obj *domain.Post) (*domain.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// TODO dataLoader for objectIDs
func (r *postResolver) Image(ctx context.Context, obj *domain.Post) (*domain.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

// TODO dataLoader for []postIDS
func (r *postResolver) Chat(ctx context.Context, obj *domain.Post) ([]*domain.Message, error) {
	panic(fmt.Errorf("not implemented"))
}
