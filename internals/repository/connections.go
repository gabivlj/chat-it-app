package repository

import (
	"context"
	"errors"
	"log"
	"runtime/debug"
	"sync"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gofrs/uuid"
)

// ConnectionsRepository handles all the connections for the app.
type ConnectionsRepository struct {
	connections map[string]PostConnections
	messageRepo *MessageRepository
	userRepo    *UserRepository
	postRepo    *PostRepository
	mu          *sync.Mutex
}

// newConnectionsRepository returns a new repository of connections
func newConnectionsRepository(p *PostRepository, m *MessageRepository, u *UserRepository) *ConnectionsRepository {
	// 1024 * 1024 => single MB
	debug.SetGCPercent(20)
	return &ConnectionsRepository{connections: make(map[string]PostConnections, 1000), postRepo: p, mu: &sync.Mutex{}, messageRepo: m, userRepo: u}
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
	// Goroutine that will wait until the user disconnects
	go func() {
		<-ctx.Done()
		conn.mu.Lock()
		log.Println("User", user.Username, "disconnected from post", postID)
		delete(conn.Observers, user.ID)
		conn.mu.Unlock()
	}()
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
	msg, err := c.messageRepo.SaveMessage(ctx, postID, userFrom, text)
	if err != nil {
		return nil, err
	}
	c.mu.Lock()
	postConn, k := c.connections[postID]
	if !k {
		post, err := c.postRepo.GetPost(ctx, postID)
		if err != nil {
			c.mu.Unlock()
			return nil, err
		}
		c.connections[postID] = PostConnections{Post: post, Observers: map[string]Observer{}, mu: &sync.Mutex{}}
		postConn = c.connections[postID]
	}
	c.mu.Unlock()
	postConn.mu.Lock()
	defer postConn.mu.Unlock()
	user, k := postConn.Observers[userFrom]
	if !k {
		userModel, err := c.userRepo.FindByID(ctx, userFrom)
		if err != nil {
			return nil, errors.New("Unexpected error")
		}
		msg.User = userModel
	} else {
		msg.User = user.User
	}
	for _, observer := range postConn.Observers {
		observer.NewMessage <- msg
	}
	return msg, nil
}
