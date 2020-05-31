package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/graphql/model"
	"github.com/gabivlj/chat-it/internals/middleware"
)

func (r *mutationResolver) NewUser(ctx context.Context, parameters *model.FormLogInRegister) (*model.UserSession, error) {
	user, session, err := r.userRepo.SaveUser(ctx, &domain.User{Username: parameters.Username, Password: parameters.Password})
	if err != nil {
		return nil, err
	}
	return &model.UserSession{User: user, Session: session}, nil
}

func (r *mutationResolver) LogUser(ctx context.Context, parameters *model.FormLogInRegister) (*model.UserSession, error) {
	user, session, err := r.userRepo.LogUser(ctx, &domain.User{Username: parameters.Username, Password: parameters.Password})
	if err != nil {
		return nil, err
	}
	return &model.UserSession{User: user, Session: session}, nil
}

func (r *mutationResolver) NewProfileImage(ctx context.Context, image graphql.Upload) (*domain.User, error) {
	user, err := middleware.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return r.userRepo.UpdateProfileImage(ctx, image, user)
}
