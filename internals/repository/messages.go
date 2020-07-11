package repository

import (
	"context"
	"fmt"
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

func mapMessages(m []messageMongo) []*domain.Message {
	messages := make([]*domain.Message, 0, len(m))
	for idx := range m {
		messages = append(messages, m[idx].Domain())
	}
	return messages
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
	return mapMessages(mongoMessages), nil
}

// GetMessagesUsers loads the messages of an user in a paginated way (No dataloaden)
func (m *MessageRepository) GetMessagesUsers(ctx context.Context, params *model.Params, users ...string) ([]*domain.Message, error) {
	options, query, err := parsePagination(params)
	if err != nil {
		return nil, err
	}
	query["userId"] = bson.M{"$in": users}
	options.Sort = constants.SortDescendingCreatedAt
	results, err := m.messagesCollection.Find(ctx, query, options)
	mongoMessages := []messageMongo{}
	err = results.All(ctx, &mongoMessages)
	if err != nil {
		return nil, err
	}
	return mapMessages(mongoMessages), nil
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

// CountMessagesPost count messages
func (m *MessageRepository) CountMessagesPost(ctx context.Context, postID string) (int, error) {

	return 0, nil
}

type countMessagesPost struct {
	Total uint64 `bson:"total"`
	ID    string `bson:"_id"`
}

type arrayCounts []countMessagesPost

func (a arrayCounts) reorder(ids []string) ([]*domain.MessageCount, error) {
	zip := map[string]uint64{}
	for idx := range a {
		element := a[idx]
		zip[element.ID] = element.Total
	}
	elements := make([]*domain.MessageCount, 0, len(ids))
	for _, id := range ids {
		elements = append(elements, &domain.MessageCount{Total: zip[id], PostID: id})
	}
	return elements, nil
}

// CountMessagesPosts with dataloaden
func (m *MessageRepository) CountMessagesPosts(ctx context.Context, postIDs []string) ([]*domain.MessageCount, error) {
	firstQuery := bson.M{"$match": bson.M{"postId": bson.M{"$in": postIDs}}}
	secondQuery := bson.M{"$group": bson.M{"_id": "$postId", "total": bson.M{"$sum": 1}}}
	res, err := m.messagesCollection.Aggregate(ctx, bson.A{firstQuery, secondQuery})
	if err != nil {
		return nil, err
	}
	messagesPost := []countMessagesPost{}
	err = res.All(ctx, &messagesPost)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return arrayCounts(messagesPost).reorder(postIDs)
}

// CountMessagesUser returns the number of messages sent by an user (No dataloaden)
func (m *MessageRepository) CountMessagesUser(ctx context.Context, userID string) (int, error) {
	// firstQuery := bson.M{"$match": bson.M{"userId": bson.M{"$in": []string{userID}}}}
	// secondQuery := bson.M{"$group": bson.M{"_id": "$postId", "total": bson.M{"$sum": 1}}}
	// secondQuery := bson.M{"$group": bson.M{"_id": "$postId", "total": bson.M{"$sum": 1}}}
	query := bson.M{"userId": userID}
	nOfMsgs, _ := m.messagesCollection.CountDocuments(ctx, query)
	// todo : check integer overflow
	return int(nOfMsgs), nil
}
