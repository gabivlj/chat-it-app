package graphql

import (
	"context"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/graphql/model"
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
