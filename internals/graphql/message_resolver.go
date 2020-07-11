package graphql

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/middleware"
)

func (r *messageResolver) User(ctx context.Context, obj *domain.Message) (*domain.User, error) {
	if obj.User == nil {
		return middleware.DataLoaderUser(ctx).Load(obj.UserID)
	}
	return obj.User, nil
}

// Warning, we shouldn't really use this
func (r *messageResolver) Post(ctx context.Context, obj *domain.Message) (*domain.Post, error) {
	// if obj.Post != nil {
	// 	return obj.Post, nil
	// }
	// elements, err := middleware.DataLoaderPost(ctx).Load(obj.UserID)
	// if err != nil {
	// 	return nil, err
	// }
	// var postToChoose *domain.Post = nil
	// for i := range elements {
	// 	if elements[i].ID == obj.PostID {
	// 		postToChoose = elements[i]
	// 		break
	// 	}
	// }

	return middleware.DataLoaderSinglePost(ctx).Load(obj.PostID)
}
