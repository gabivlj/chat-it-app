package graphql

import "github.com/gabivlj/chat-it/internals/services"

// Resolver stores the dependencies and resolves the Graphql dependencies
type Resolver struct {
	userRepo services.UserService
}

// New returns a new resolver
func New(u services.UserService) *Resolver {
	return &Resolver{userRepo: u}
}
