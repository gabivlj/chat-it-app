package services

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/graphql/model"
)

// MessageService handles logic of the messages repos
type MessageService interface {
	GetMessages(ctx context.Context, postID string, params *model.Params) ([]*domain.Message, error)
	CountMessagesPost(ctx context.Context, postID string) (int, error)
	CountMessagesUser(ctx context.Context, userID string) (int, error)
	CountMessagesPosts(ctx context.Context, postIDs []string) ([]*domain.MessageCount, error)
}
