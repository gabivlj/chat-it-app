package repository

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gofrs/uuid"
)

// ConnectionsRepository handles all the connections for the app.
type ConnectionsRepository struct {
	connections map[string]PostConnections
	postRepo    *PostRepository
	mu          *sync.Mutex
}

// newConnectionsRepository returns a new repository of connections
func newConnectionsRepository(p *PostRepository) *ConnectionsRepository {
	// 1024 * 1024 => single MB
	debug.SetGCPercent(20)
	return &ConnectionsRepository{connections: make(map[string]PostConnections, 1000), postRepo: p, mu: &sync.Mutex{}}
}

// PostConnections .
type PostConnections struct {
	Post      *domain.Post        `json:"post"`
	Observers map[string]Observer `json:"observers"`
	mu        *sync.Mutex
}

// Observer .
type Observer struct {
	User       *domain.User
	NewMessage chan *domain.Message
}

// NewUser connection on a post
func (c *ConnectionsRepository) NewUser(ctx context.Context, user *domain.User, postID string) (<-chan *domain.Message, error) {
	c.mu.Lock()
	conn, ok := c.connections[postID]
	if !ok {
		post, err := c.postRepo.GetPost(ctx, postID)
		if err != nil {
			return nil, err
		}
		c.connections[postID] = PostConnections{Post: post, Observers: map[string]Observer{user.ID: {User: user, NewMessage: make(chan *domain.Message, 1)}}, mu: &sync.Mutex{}}
		conn = c.connections[postID]
	}
	c.mu.Unlock()
	connChann := make(chan *domain.Message, 1)
	conn.mu.Lock()
	conn.Observers[user.ID] = Observer{User: user, NewMessage: connChann}
	conn.mu.Unlock()
	return connChann, nil
}

func tempunwrap(s uuid.UUID, err error) string {
	return s.String()
}

// SendMessage sends a message across the app
func (c *ConnectionsRepository) SendMessage(ctx context.Context, postID string, userFrom string, text string) (*domain.Message, error) {
	msg := &domain.Message{User: nil, Text: text, CreatedAt: time.Now().Unix(), ID: tempunwrap(uuid.NewV4())}
	c.mu.Lock()
	postConn, k := c.connections[postID]
	if !k {
		return nil, fmt.Errorf("No post")
	}
	c.mu.Unlock()
	postConn.mu.Lock()
	user, k := postConn.Observers[userFrom]
	if !k {
		return nil, fmt.Errorf("bad user ID")
	}
	msg.User = user.User
	for _, observer := range postConn.Observers {
		observer.NewMessage <- msg
	}
	postConn.mu.Unlock()
	return msg, nil
}
