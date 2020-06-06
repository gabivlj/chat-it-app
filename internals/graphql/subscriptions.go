package graphql

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/middleware"
)

func (r *subscriptionResolver) NewMessage(ctx context.Context, postID string) (<-chan *domain.Message, error) {
	user, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return r.connectionsPosts.NewUser(ctx, user, postID)
}
