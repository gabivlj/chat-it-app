package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/graphql/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PostRepository uses mongodb to store and handle posts bussiness logic
type PostRepository struct {
	db             *mongo.Database
	client         *mongo.Client
	postCollection *mongo.Collection
	fileUpl        *CloudStorageImages
}

type postMongo struct {
	UserID    string             `bson:"userId,omitempty"`
	Text      string             `bson:"text,omitempty"`
	Title     string             `bson:"title,omitempty"`
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt int64              `bson:"createdAt,omitempty"`
	URLImage  string             `bson:"urlImg,omitempty"`
}

func (p *postMongo) Domain() *domain.Post {
	return &domain.Post{Text: p.Text, UserID: p.UserID, CreatedAt: p.CreatedAt, Title: p.Title, ID: p.ID.Hex(), URLImage: p.URLImage}
}

func newPostRepository(db *mongo.Database, client *mongo.Client, fileUpl *CloudStorageImages) *PostRepository {
	return &PostRepository{db: db, client: client, postCollection: db.Collection("posts"), fileUpl: fileUpl}
}

// NewPost returns a new post saved in the database
func (p *PostRepository) NewPost(ctx context.Context, input *model.PostForm, user *domain.User) (*domain.Post, error) {
	uri, err := p.fileUpl.UploadFile(ctx, input.Image.File, input.Image.ContentType)
	if input.Image.File != nil && err != nil {
		return nil, err
	}
	log.Printf("New Post image: %s \n", uri)
	postInsertion := &postMongo{UserID: user.ID, Text: input.Text, Title: input.Title, CreatedAt: time.Now().Unix(), URLImage: uri}
	id, errorInserting := p.postCollection.InsertOne(ctx, postInsertion)
	if errorInserting != nil {
		return nil, errorInserting
	}
	idObject := id.InsertedID.(primitive.ObjectID)
	postInsertion.ID = idObject
	return postInsertion.Domain(), nil
}

// GetPost returns the post info
func (p *PostRepository) GetPost(ctx context.Context, id string) (*domain.Post, error) {
	idStr, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	fmt.Println(idStr)
	post := &postMongo{ID: idStr}
	res := p.postCollection.FindOne(ctx, post)
	err = res.Err()
	if err != nil {
		return nil, err
	}
	err = res.Decode(&post)
	if err != nil {
		return nil, err
	}
	return post.Domain(), nil
}

var empty = ""

func unwrapPointerObjectID(s *string) (primitive.ObjectID, error) {
	if s == nil {
		return primitive.NilObjectID, nil
	}
	return primitive.ObjectIDFromHex(*s)
}

// GetPosts returns the posts of the frontpage
func (p *PostRepository) GetPosts(ctx context.Context, pagination *model.Params) ([]*domain.Post, error) {
	after, err := unwrapPointerObjectID(pagination.After)
	if err != nil {
		return nil, err
	}
	before, err := unwrapPointerObjectID(pagination.Before)
	if err != nil {
		return nil, err
	}
	options := options.Find()
	l := int64(pagination.Limit)
	options.Limit = &l
	options.Sort = bson.M{"createdAt": -1}
	var query bson.M
	if after == primitive.NilObjectID && before == primitive.NilObjectID {
		query = bson.M{}
	} else if after != primitive.NilObjectID {
		query = bson.M{"_id": bson.M{"$gt": after}}
	} else {
		query = bson.M{"_id": bson.M{"$lt": before}}
	}
	postsResult, err := p.postCollection.Find(ctx, query, options)
	if err != nil {
		return nil, err
	}
	postsMongo := []postMongo{}
	err = postsResult.All(ctx, &postsMongo)
	if err != nil {
		return nil, err
	}
	posts := make([]*domain.Post, len(postsMongo))
	for i := range postsMongo {
		posts[i] = postsMongo[i].Domain()
	}
	return posts, nil
}

// GetPostsFromUsers returns all the posts from users
func (p *PostRepository) GetPostsFromUsers(ctx context.Context, userIds []string) ([][]*domain.Post, error) {
	return nil, nil
}
