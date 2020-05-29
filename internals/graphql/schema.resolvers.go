package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gabivlj/chat-it/internals/domain"
	generated1 "github.com/gabivlj/chat-it/internals/graphql/generated"
)

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

// User returns generated1.UserResolver implementation.
func (r *Resolver) User() generated1.UserResolver { return &userResolver{r} }

type imageResolver struct{ *Resolver }
type messageResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *messageResolver) CreatedAt(ctx context.Context, obj *domain.Message) (int, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *imageResolver) URLSm(ctx context.Context, obj *domain.Image) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
