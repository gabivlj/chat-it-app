package repository

import (
	"context"
	"log"
	"time"

	"github.com/gabivlj/chat-it/internals/constants"
	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/graphql/model"
	"go.mongodb.org/mongo-driver/bson"
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

func newMessageRepository(db *mongo.Database, client *mongo.Client) *MessageRepository {
	coll := db.Collection("messages")
	name, err := coll.Indexes().CreateOne(context.Background(), mongo.IndexModel{Keys: bson.M{"userId": 1}})
	log.Printf("repository - messages.go: tried creating indexes on userId, name: %s, err: %v", name, err)
	name, err = coll.Indexes().CreateOne(context.Background(), mongo.IndexModel{Keys: bson.M{"postId": 1}})
	log.Printf("repository - messages.go: tried creating indexes on postId, name: %s, err: %v", name, err)
	return &MessageRepository{db: db, client: client, messagesCollection: coll}
}

// GetMessages returns the messages of the chat
func (m *MessageRepository) GetMessages(ctx context.Context, postID string, params *model.Params) ([]*domain.Message, error) {

	options, query, err := parsePagination(params)
	query["postId"] = postID
	if err != nil {
		return nil, err
	}
	options.Sort = constants.SortDescendingCreatedAt
	results, err := m.messagesCollection.Find(ctx, query, options)
	if err != nil {
		return nil, err
	}
	mongoMessages := []messageMongo{}
	err = results.All(ctx, &mongoMessages)
	if err != nil {
		return nil, err
	}
	messages := make([]*domain.Message, 0, len(mongoMessages))
	for idx := range mongoMessages {
		messages = append(messages, mongoMessages[idx].Domain())
	}
	return messages, nil
}

// SaveMessage saves the message in the database
func (m *MessageRepository) SaveMessage(ctx context.Context, postID, userID, text string) (*domain.Message, error) {
	msg := messageMongo{PostID: postID, UserID: userID, Text: text, CreatedAt: time.Now().UnixNano()}
	res, err := m.messagesCollection.InsertOne(ctx, msg)
	if err != nil {
		return nil, err
	}
	msg.ID = res.InsertedID.(primitive.ObjectID)
	return msg.Domain(), nil
}
