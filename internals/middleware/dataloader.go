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

// DataloaderMiddleware is the middleware for adding all the dataloader middleware
func DataloaderMiddleware(next http.Handler, userRepository services.UserService) http.Handler {
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
		userLoader := dataloader.NewUserLoader(userLoaderConfig)
		tx := context.WithValue(r.Context(), userDataLoaderKey, userLoader)
		next.ServeHTTP(w, r.WithContext(tx))
	})
}

// DataLoaderUser returns a userloader from the context
func DataLoaderUser(ctx context.Context) *dataloader.UserLoader {
	return ctx.Value(userDataLoaderKey).(*dataloader.UserLoader)
}
