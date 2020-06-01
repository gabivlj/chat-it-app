package graphql

import (
	"github.com/gabivlj/chat-it/internals/services"
)

// Resolver stores the dependencies and resolves the Graphql dependencies
type Resolver struct {
	userRepo services.UserService
	postRepo services.PostService

	connectionsPosts  services.ConnectionService
	messageRepository services.MessageService
}

// New returns a new resolver todo: Use services
func New(u services.UserService, p services.PostService, c services.ConnectionService, m services.MessageService) *Resolver {
	return &Resolver{userRepo: u, postRepo: p, connectionsPosts: c, messageRepository: m}
}
