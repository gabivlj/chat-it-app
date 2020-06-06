package graphql

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/middleware"
)

func (r *mutationResolver) SendMessage(ctx context.Context, text string, postID string, userID string) (*domain.Message, error) {
	user, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	return r.connectionsPosts.SendMessage(ctx, postID, user.ID, text)
}
