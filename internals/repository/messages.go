package repository

import (
	"github.com/gabivlj/chat-it/internals/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MessageRepository uses mongodb to store and handle messages bussiness logic
type MessageRepository struct {
	db                 *mongo.Database
	client             *mongo.Client
	messagesCollection *mongo.Collection
}

type messageMongo struct {
	UserID    string             `bson:"userId,omitempty"`
	Text      string             `bson:"text,omitempty"`
	PostID    string             `bson:"postId,omitempty"`
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt int64              `bson:"createdAt,omitempty"`
}

func (p *messageMongo) Domain() *domain.Message {
	return &domain.Message{Text: p.Text, UserID: p.UserID, CreatedAt: p.CreatedAt, PostID: p.PostID, ID: p.ID.Hex()}
}

func newMessageRepository(db *mongo.Database, client *mongo.Client, fileUpl *CloudStorageImages) *MessageRepository {
	return &MessageRepository{db: db, client: client, messagesCollection: db.Collection("messages")}
}
