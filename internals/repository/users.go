package repository

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/coocood/freecache"
	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/services"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository implements business logic for the users
type UserRepository struct {
	db             *mongo.Database
	client         *mongo.Client
	userCollection *mongo.Collection
	sessions       *freecache.Cache
	// (NOTE) (GABI) : Do pagination with { _id : { $gt: otherid }}
}

// Disconnect disconnects from the server
func (u *UserRepository) Disconnect(ctx context.Context) {
	u.client.Disconnect(ctx)
}

type userMongo struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
}

// Mongo returns the user in our package
func mongoUser(u *domain.User) *userMongo {
	var id primitive.ObjectID = primitive.NilObjectID
	if len(u.ID) > 0 {
		ids, err := primitive.ObjectIDFromHex(u.ID)
		if err == nil {
			id = ids
		}
	}
	return &userMongo{Username: u.Username, ID: id}
}

func (u *userMongo) Domain() *domain.User {
	return &domain.User{Username: u.Username, ID: u.ID.Hex()}
}

// NewRepository returns users repo
func NewRepository() services.UserService {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	mongoURI, _ := os.LookupEnv("MONGO_URI_LOCAL")
	if mongoURI == "" {
		panic(fmt.Errorf("Mongo URI is empty"))
	}
	repo := &UserRepository{db: nil, sessions: freecache.NewCache(100)}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	options := options.Client()
	options.SetMaxPoolSize(100)
	options.ApplyURI(fmt.Sprintf("%s", mongoURI))
	defer cancel()
	client, err := mongo.Connect(ctx, options)
	repo.client = client
	if err != nil {
		log.Fatal(err)
	}
	repo.db = client.Database("chat-it")
	repo.userCollection = repo.db.Collection("users")
	return repo
}

// SaveUser saves a user into mongo db
func (u *UserRepository) SaveUser(ctx context.Context, user *domain.User) (*domain.User, string, error) {
	mongoU := mongoUser(user)
	// Sanitize inputs
	usr := u.userCollection.FindOne(ctx, mongoU)
	if usr.Err() == nil {
		return nil, "", fmt.Errorf("Error, the user already exists")
	}
	user.Username = strings.TrimSpace(user.Username)
	if len(user.Username) < 2 || len(user.Username) > 16 {
		return nil, "", fmt.Errorf("Username must be more than 1 character or less than 17")
	}
	if len(user.Password) < 6 || len(user.Password) >= 60 {
		return nil, "", fmt.Errorf("Error, password length must be more than 5 characters or less than 60")
	}
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, "", err
	}
	mongoU.Password = string(encryptedPass)
	result, err := u.userCollection.InsertOne(ctx, mongoU)
	if err != nil {
		return nil, "", err
	}
	id := result.InsertedID.(primitive.ObjectID)
	mongoU.ID = id
	userEnd := mongoU.Domain()
	bytes, _ := json.Marshal(userEnd)
	sessionID := newSessionID()
	u.sessions.Set([]byte(sessionID), bytes, 10000*60*60)
	return mongoU.Domain(), sessionID, nil
}

// LogUser logs an user from the db and creates a session
func (u *UserRepository) LogUser(ctx context.Context, user *domain.User) (*domain.User, string, error) {
	userMongo := mongoUser(user)
	res := u.userCollection.FindOne(ctx, userMongo)
	err := res.Decode(&userMongo)
	if err != nil {
		return nil, "", err
	}
	if userMongo.Username != user.Username {
		return nil, "", errors.New("User not found")
	}
	if bcrypt.CompareHashAndPassword([]byte(userMongo.Password), []byte(user.Password)) != nil {
		return nil, "", errors.New("Incorrect password")
	}
	user = userMongo.Domain()
	bytes, _ := json.Marshal(&user)
	s := newSessionID()
	u.sessions.Set([]byte(s), bytes, int(time.Hour)*1000)
	return userMongo.Domain(), s, nil
}

// FindByID returns a user by id
func (u *UserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	domainUser := domain.User{ID: id}
	userMongo := mongoUser(&domainUser)
	// now := time.Now().Unix()

	res := u.userCollection.FindOne(ctx, userMongo)

	err := res.Decode(&userMongo)
	if err != nil {
		return nil, err
	}
	fmt.Println(userMongo.ID.Hex())
	return userMongo.Domain(), nil
}

// FindByUsername returns a user by id
func (u *UserRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	domainUser := domain.User{Username: username}
	userMongo := mongoUser(&domainUser)
	// Mongo atlas free tier wastes like 2 seconds for finding an user. UNREAL.
	res := u.userCollection.FindOne(ctx, userMongo)
	err := res.Decode(&userMongo)
	if err != nil {
		return nil, err
	}
	return userMongo.Domain(), nil
}

// VerifySession returns if the passed session is correct
func (u *UserRepository) VerifySession(session string) (*domain.User, error) {
	user, err := u.sessions.Get([]byte(session))
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("Error getting user")
	}
	var userData *domain.User = &domain.User{}
	err = json.Unmarshal(user, userData)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return userData, err
}

// FindByIDs finds users by ids (we will use this for data loaden)
func (u *UserRepository) FindByIDs(ctx context.Context, ids []string) ([]*domain.User, error) {
	cursor, err := u.userCollection.Find(ctx, bson.M{"_id": bson.M{"$in": getObjectIDs(ids)}})
	if err != nil {
		return nil, err
	}
	users := []userMongo{}
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	usersRef := make([]*domain.User, len(users))
	for i := range users {
		usersRef[i] = users[i].Domain()
	}
	return usersRef, nil
}

func newSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)

}
