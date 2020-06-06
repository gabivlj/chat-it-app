package services

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/graphql/model"
)

// MessageService handles logic of the messages repos
type MessageService interface {
	GetMessages(ctx context.Context, postID string, params *model.Params) ([]*domain.Message, error)
}
