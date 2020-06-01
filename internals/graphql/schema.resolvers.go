package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
	generated1 "github.com/gabivlj/chat-it/internals/graphql/generated"
	"github.com/gabivlj/chat-it/internals/middleware"
)

func (r *mutationResolver) SendMessage(ctx context.Context, text string, postID string, userID string) (*domain.Message, error) {
	return r.connectionsPosts.SendMessage(ctx, postID, userID, text)
}

func (r *subscriptionResolver) NewMessage(ctx context.Context, postID string) (<-chan *domain.Message, error) {
	user, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	// if err != nil {

	// 	return nil, err
	// }
	return r.connectionsPosts.NewUser(ctx, user, postID)
}

// Image returns generated1.ImageResolver implementation.
func (r *Resolver) Image() generated1.ImageResolver { return &imageResolver{r} }

// Message returns generated1.MessageResolver implementation.
func (r *Resolver) Message() generated1.MessageResolver { return &messageResolver{r} }

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Post returns generated1.PostResolver implementation.
func (r *Resolver) Post() generated1.PostResolver { return &postResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

// Subscription returns generated1.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated1.SubscriptionResolver { return &subscriptionResolver{r} }

// User returns generated1.UserResolver implementation.
func (r *Resolver) User() generated1.UserResolver { return &userResolver{r} }

type imageResolver struct{ *Resolver }
type messageResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
