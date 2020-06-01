package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gabivlj/chat-it/internals/graphql"
	"github.com/gabivlj/chat-it/internals/graphql/generated"
	"github.com/gabivlj/chat-it/internals/middleware"
	"github.com/gabivlj/chat-it/internals/repository"
	"github.com/gorilla/websocket"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// App repositories
	userRepo, postRepo, msgRepo, repoConnections := repository.NewRepository()

	// Schema and resolver loading
	graphqlHandler := generated.NewExecutableSchema(generated.Config{Resolvers: graphql.New(userRepo, postRepo, repoConnections, msgRepo), Directives: *graphql.NewDirectives()})

	// Initialize handler
	srv := handler.New(graphqlHandler)

	// Middleware and transports
	middlewareDataloadenHTTP, middlewareDataloadenWebSockets := middleware.DataloaderMiddleware(srv, userRepo, postRepo)
	middlewareHTTP, middlewareSessionsWebsockets := middleware.SessionMiddleware(userRepo.Sessions, middlewareDataloadenHTTP.ServeHTTP)
	addTransports(srv, middlewareSessionsWebsockets, middlewareDataloadenWebSockets)

	// Handle everything
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middlewareHTTP)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	// !! Test websockets
	// testWs(srv)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func addTransports(srv *handler.Server, middlewareWebsocketsSession func(ctx context.Context, token string) context.Context, middlewareWebsocketsDataLoader func(context.Context) context.Context) {
	srv.AddTransport(transport.Websocket{
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			token := initPayload.Authorization()
			tx := middlewareWebsocketsSession(ctx, token)
			tx = middlewareWebsocketsDataLoader(tx)
			user, err := middleware.GetUser(tx)
			fmt.Println(user, err)
			return tx, nil
		},
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			EnableCompression: true,
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
}
