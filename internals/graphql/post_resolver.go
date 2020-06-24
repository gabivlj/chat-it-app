package graphql

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/graphql/model"
	"github.com/gabivlj/chat-it/internals/middleware"
)

func (r *postResolver) User(ctx context.Context, obj *domain.Post) (*domain.User, error) {
	return middleware.DataLoaderUser(ctx).Load(obj.UserID)
}

func (r *postResolver) Image(ctx context.Context, obj *domain.Post) (*domain.Image, error) {
	return &domain.Image{URLMD: obj.URLImage, URLXL: obj.URLImage, URLSM: obj.URLImage}, nil
}

// TODO dataLoader
func (r *postResolver) Chat(ctx context.Context, obj *domain.Post) ([]*domain.Message, error) {
	return r.messageRepository.GetMessages(ctx, obj.ID, &model.Params{Limit: 10})
}

func (r *postResolver) NumberOfComments(ctx context.Context, obj *domain.Post) (int, error) {
	postMessageCount, err := middleware.DataLoaderMessageCount(ctx).Load(obj.ID)
	if err != nil {
		return 0, err
	}
	return int(postMessageCount.Total), nil
}
