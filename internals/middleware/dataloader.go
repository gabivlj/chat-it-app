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

// DataloaderMiddleware is the middleware for adding all the dataloader middleware
func DataloaderMiddleware(next http.Handler, userRepository services.UserService, postRepository services.PostService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		userLoader := dataloader.NewUserLoader(userLoaderConfig)
		postLoader := dataloader.NewPostLoader(postLoaderConfig)
		tx := context.WithValue(r.Context(), userDataLoaderKey, userLoader)
		tx = context.WithValue(tx, postDataLoaderKey, postLoader)
		next.ServeHTTP(w, r.WithContext(tx))
	})
}

// DataLoaderUser returns a userloader from the context
func DataLoaderUser(ctx context.Context) *dataloader.UserLoader {
	return ctx.Value(userDataLoaderKey).(*dataloader.UserLoader)
}

// DataLoaderPost returns the postloader from the request's context
func DataLoaderPost(ctx context.Context) *dataloader.PostLoader {
	return ctx.Value(postDataLoaderKey).(*dataloader.PostLoader)
}
