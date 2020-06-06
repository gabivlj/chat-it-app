package middleware

// ! warning: Delete the "test" user in the future
// ! warning: Delete the "test" user in the future
// ! warning: Delete the "test" user in the future

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

// SessionMiddlewareData testing
type SessionMiddlewareData struct {
	sessions *freecache.Cache
}

// WSSession is a middleware but for websockets
type WSSession func(ctx context.Context, token string) context.Context

// SessionMiddleware .
func SessionMiddleware(sessions *freecache.Cache, graphQLHandle http.HandlerFunc) (http.HandlerFunc, WSSession) {
	middle := SessionMiddlewareData{sessions: sessions}
	return middle.Auth(graphQLHandle)
}

// Auth .
func (s *SessionMiddlewareData) Auth(next http.HandlerFunc) (http.HandlerFunc, WSSession) {
	return func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				next(w, r)
				return
			}
			userBytes, err := s.sessions.Get([]byte(auth))
			if err != nil && auth != "test" {
				next(w, r)
				return
			}
			var user domain.User
			// TODO Remove this
			if auth != "test" {
				err = json.Unmarshal(userBytes, &user)
				if err != nil {
					next(w, r)
					return
				}
			} else {
				user = domain.User{Username: "test", ID: "1", ImageURL: "none"}
			}
			ctx := r.Context()
			ctx = context.WithValue(ctx, userKey, &user)
			r = r.WithContext(ctx)
			next(w, r)
		}, func(ctx context.Context, token string) context.Context { // The same as before but simplified
			// I might reduce code by simplyfing in a single fn but right now for this case works
			userBytes, err := s.sessions.Get([]byte(token))
			var user domain.User
			if token != "test" {
				err = json.Unmarshal(userBytes, &user)
				if err != nil {
					return ctx
				}
			} else {
				user = domain.User{Username: "test", ID: "1", ImageURL: "none"}
			}
			ctx = context.WithValue(ctx, userKey, &user)
			return ctx
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
