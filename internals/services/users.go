package services

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gabivlj/chat-it/internals/domain"
)

// UserService handles user bussiness
type UserService interface {
	// Needs to encrypt password and save in a secure way
	SaveUser(ctx context.Context, user *domain.User) (*domain.User, string, error)
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
	LogUser(ctx context.Context, user *domain.User) (*domain.User, string, error)
	VerifySession(session string) (*domain.User, error)
	Disconnect(ctx context.Context)
	FindByIDs(ctx context.Context, ids []string) ([]*domain.User, error)
	UpdateProfileImage(ctx context.Context, image graphql.Upload, user *domain.User) (*domain.User, error)
}
