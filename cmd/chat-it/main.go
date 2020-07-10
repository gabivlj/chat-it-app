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
	middlewareDataloadenHTTP, middlewareDataloadenWebSockets := middleware.DataloaderMiddleware(srv, userRepo, postRepo, msgRepo)
	middlewareHTTP, middlewareSessionsWebsockets := middleware.SessionMiddleware(userRepo.Sessions, middlewareDataloadenHTTP.ServeHTTP)
	addTransports(srv, middlewareSessionsWebsockets, middlewareDataloadenWebSockets)

	// Handle everything
	http.Handle("/", middleware.Cors(playground.Handler("GraphQL playground", "/query")))
	// http.Serve("/*filepath", http.Dir("/"))
	// handler := http.FileServer(http.Dir("/"))
	http.Handle("/query", middleware.Cors(middlewareHTTP))
	// http.Handle("/", middleware.Cors(handler.ServeHTTP))
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func addTransports(srv *handler.Server, middlewareWebsocketsSession func(ctx context.Context, token string) context.Context, middlewareWebsocketsDataLoader func(context.Context) context.Context) {
	srv.AddTransport(transport.Websocket{
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			fmt.Println(initPayload.GetString("Authorization"), initPayload)
			token := initPayload.GetString("Authorization")
			// If the authorization param is empty
			if token == "" {
				headers := initPayload["headers"]
				if headers == nil {
					return nil, fmt.Errorf("Unauthorized")
				}
				headersMap, k := headers.(map[string]interface{})
				if !k {
					return nil, fmt.Errorf("Unauthorized")
				}
				header, k := headersMap["Authorization"]
				if !k {
					return nil, fmt.Errorf("Unauthorized")
				}
				headerStr, k := header.(string)
				if k {
					token = headerStr
				}
			}
			tx := middlewareWebsocketsSession(ctx, token)
			tx = middlewareWebsocketsDataLoader(tx)
			user, err := middleware.GetUser(tx)
			fmt.Println(user, err)
			if err != nil {
				return nil, err
			}
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
