package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gabivlj/chat-it/internals/dataloader"
	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/services"
)

type dataLoaderKey uint16

const userDataLoaderKey = dataLoaderKey(1)
const postDataLoaderKey = dataLoaderKey(2)
const messageCountDataLoaderKey = dataLoaderKey(3)

func dataloaders(userRepository services.UserService, postRepository services.PostService, messageRepository services.MessageService) (*dataloader.UserLoader, *dataloader.PostLoader, *dataloader.MessageCountLoader) {
	userLoaderConfig := dataloader.UserLoaderConfig{
		MaxBatch: 100,
		Wait:     1 * time.Millisecond,
		Fetch: func(ids []string) ([]*domain.User, []error) {
			var users []*domain.User
			// * NOTE (GABI): Maybe change this context.Background() to another thing
			users, err := userRepository.FindByIDs(context.Background(), ids)
			return users, []error{err}
		},
	}
	postLoaderConfig := dataloader.PostLoaderConfig{
		MaxBatch: 100,
		Wait:     1 * time.Millisecond,
		Fetch: func(ids []string) ([][]*domain.Post, []error) {
			// * NOTE (GABI): Maybe change this context.Background() to another thing
			posts, err := postRepository.GetPostsFromUsers(context.Background(), ids)
			return posts, []error{err}
		},
	}
	messageTotalLoaderConfig := dataloader.MessageCountLoaderConfig{
		MaxBatch: 100,
		Wait:     1 * time.Millisecond,
		Fetch: func(keys []string) ([]*domain.MessageCount, []error) {
			total, err := messageRepository.CountMessagesPosts(context.Background(), keys)
			if err != nil {
				return nil, []error{err}
			}
			return total, nil
		},
	}
	userLoader := dataloader.NewUserLoader(userLoaderConfig)
	postLoader := dataloader.NewPostLoader(postLoaderConfig)
	messageTotalLoader := dataloader.NewMessageCountLoader(messageTotalLoaderConfig)
	return userLoader, postLoader, messageTotalLoader
}

// DataloaderMiddleware is the middleware for adding all the dataloader middleware,
// returns the http and websocket middleware
func DataloaderMiddleware(next http.Handler, userRepository services.UserService, postRepository services.PostService, messageRepository services.MessageService) (http.Handler, func(ctx context.Context) context.Context) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userLoader, postLoader, messageCountLoader := dataloaders(userRepository, postRepository, messageRepository)
			tx := context.WithValue(r.Context(), userDataLoaderKey, userLoader)
			tx = context.WithValue(tx, postDataLoaderKey, postLoader)
			tx = context.WithValue(tx, messageCountDataLoaderKey, messageCountLoader)
			next.ServeHTTP(w, r.WithContext(tx))
		}), func(ctx context.Context) context.Context {
			userLoader, postLoader, messageCountLoader := dataloaders(userRepository, postRepository, messageRepository)
			tx := context.WithValue(ctx, userDataLoaderKey, userLoader)
			tx = context.WithValue(tx, postDataLoaderKey, postLoader)
			tx = context.WithValue(tx, messageCountDataLoaderKey, messageCountLoader)
			return tx
		}
}

// DataLoaderUser returns a userloader from the context
func DataLoaderUser(ctx context.Context) *dataloader.UserLoader {
	return ctx.Value(userDataLoaderKey).(*dataloader.UserLoader)
}

// DataLoaderPost returns the postloader from the request's context
func DataLoaderPost(ctx context.Context) *dataloader.PostLoader {
	return ctx.Value(postDataLoaderKey).(*dataloader.PostLoader)
}

// DataLoaderMessageCount returns the messageCountLoader from the request's context
func DataLoaderMessageCount(ctx context.Context) *dataloader.MessageCountLoader {
	return ctx.Value(messageCountDataLoaderKey).(*dataloader.MessageCountLoader)
}
