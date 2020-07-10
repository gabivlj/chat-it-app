package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewRepository returns all the repositories of the application
func NewRepository() (*UserRepository, *PostRepository, *MessageRepository, *ConnectionsRepository) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	mongoURI, _ := os.LookupEnv("MONGO_URI_LOCAL")
	if mongoURI == "" {
		panic(fmt.Errorf("Mongo URI is empty"))
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	options := options.Client()
	options.SetMaxPoolSize(100)
	options.ApplyURI(fmt.Sprintf("%s", mongoURI))
	defer cancel()
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("chat-it")
	fileUpl := NewCloudStorage()
	pRepo := newPostRepository(db, client, fileUpl)
	mRepo := newMessageRepository(db, client)
	uRepo := newUsersRepo(db, client, fileUpl)
	// mRepo.CountMessagesPosts(context.TODO(), []string{"5ef0166f1173b7f88b007a96", "5ed1b3ab0f94ac59a503574a"})
	return uRepo, pRepo, mRepo, newConnectionsRepository(pRepo, mRepo, uRepo)
}
