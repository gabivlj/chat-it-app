package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/coocood/freecache"
	"github.com/gabivlj/chat-it/internals/domain"
)

type userCtx uint8

const userKey = userCtx(1)

// sessionMiddleware testing
type sessionMiddleware struct {
	sessions *freecache.Cache
}

// SessionMiddleware .
func SessionMiddleware(sessions *freecache.Cache, graphQLHandle http.HandlerFunc) http.HandlerFunc {
	middle := sessionMiddleware{sessions: sessions}
	return middle.Auth(graphQLHandle)
}

// Auth .
func (s *sessionMiddleware) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			next(w, r)
			return
		}
		userBytes, err := s.sessions.Get([]byte(auth))
		if err != nil {
			next(w, r)
			return
		}
		var user domain.User
		err = json.Unmarshal(userBytes, &user)
		if err != nil {
			next(w, r)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, userKey, &user)
		r = r.WithContext(ctx)
		next(w, r)
	}
}

// GetUser from the context
func GetUser(ctx context.Context) (*domain.User, error) {
	user := ctx.Value(userKey)
	if user == nil {
		return nil, errors.New("Unauthorized")
	}
	userPointer, ok := user.(*domain.User)
	if !ok || userPointer == nil {
		return userPointer, errors.New("Unauthorized")
	}
	return userPointer, nil
}
