package services

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
)

// ConnectionService handles connection logic
type ConnectionService interface {
	SendMessage(ctx context.Context, postID string, userFrom string, text string) (*domain.Message, error)
	NewUser(ctx context.Context, user *domain.User, postID string) (<-chan *domain.Message, error)
}
