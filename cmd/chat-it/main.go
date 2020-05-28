package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gabivlj/chat-it/internals/graphql"
	"github.com/gabivlj/chat-it/internals/graphql/generated"
	"github.com/gorilla/websocket"
)

const defaultPort = "8080"

func main() {
	// user := repository.NewRepository()
	// fmt.Println(user.SaveUser(context.TODO(), &domain.User{Username: "GABI", Password: "mhm?"}))

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	graphqlHandler := generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}})
	srv := handler.NewDefaultServer(graphqlHandler)
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
