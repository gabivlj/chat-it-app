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
	coll := db.Collection("posts")
	// GABI TODO: Maybe just check if the index exists, and if it doesn't create it.
	// name, err := coll.Indexes().CreateOne(context.Background(), mongo.IndexModel{Keys: bson.M{"userId": 1}})
	// log.Printf("repository - posts.go: tried creating indexes on userId, name: %s, err: %v", name, err)
	return &PostRepository{db: db, client: client, postCollection: coll, fileUpl: fileUpl}
}

// NewPost returns a new post saved in the database
func (p *PostRepository) NewPost(ctx context.Context, input *model.PostForm, user *domain.User) (*domain.Post, error) {
	var uri string
	var err error
	if input.Image != nil {
		uri, err = p.fileUpl.UploadFile(ctx, input.Image.File, input.Image.ContentType)
		if input.Image.File != nil && err != nil {
			return nil, err
		}
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
	fmt.Println("POST:", idStr)
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

// GetPosts returns the posts of the frontpage
func (p *PostRepository) GetPosts(ctx context.Context, pagination *model.Params) ([]*domain.Post, error) {
	fmt.Println("sdffdsfdssfd")
	options, query, err := parsePagination(pagination)
	options.Sort = constants.SortDescendingCreatedAt
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
func (p *PostRepository) GetPostsFromUsers(ctx context.Context, userIDs []string) ([][]*domain.Post, error) {
	query := bson.M{"userId": bson.M{"$in": userIDs}}
	posts, err := p.postCollection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	var postsMongo []postMongo = []postMongo{}
	err = posts.All(ctx, &postsMongo)
	if err != nil {
		return nil, err
	}
	postHash := map[string][]*domain.Post{}
	for idx := range postsMongo {
		post := postsMongo[idx]
		_, ok := postHash[post.UserID]
		if !ok {
			postHash[post.UserID] = []*domain.Post{post.Domain()}
			continue
		}
		postHash[post.UserID] = append(postHash[post.UserID], post.Domain())
	}
	matrixOfPosts := make([][]*domain.Post, 0, len(userIDs))
	for _, userID := range userIDs {
		matrixOfPosts = append(matrixOfPosts, postHash[userID])
	}
	return matrixOfPosts, nil
}
