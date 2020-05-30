package graphql

import (
	"github.com/gabivlj/chat-it/internals/repository"
	"github.com/gabivlj/chat-it/internals/services"
)

// Resolver stores the dependencies and resolves the Graphql dependencies
type Resolver struct {
	userRepo services.UserService
	postRepo services.PostService
}

// New returns a new resolver
func New(u services.UserService, p *repository.PostRepository) *Resolver {
	return &Resolver{userRepo: u, postRepo: p}
}
