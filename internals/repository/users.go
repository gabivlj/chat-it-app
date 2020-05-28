package repository

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/coocood/freecache"
	"github.com/dgryski/trifles/uuid"
	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository implements business logic for the users
type UserRepository struct {
	db             *mongo.Database
	userCollection *mongo.Collection
	sessions       *freecache.Cache
	// (NOTE) (GABI) : Do pagination with { _id : { $gt: otherid }}
}

type userMongo struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Password string             `bson:"password,omitempty"`
}

// Mongo returns the user in our package
func mongoUser(u *domain.User) *userMongo {
	var id primitive.ObjectID = primitive.NilObjectID
	var err error
	if len(u.ID) > 0 {
		var bytes []byte
		bytes = make([]byte, len(u.ID)*16)
		hex.Encode(bytes, []byte(u.ID))
		b := string(bytes)
		id, err = primitive.ObjectIDFromHex(b)
		if err != nil {
			id = primitive.NilObjectID
		}
	}
	return &userMongo{Username: u.Username, ID: id}
}

func (u *userMongo) Domain() *domain.User {
	return &domain.User{Username: u.Username, ID: u.ID.String()}
}

// NewRepository returns users repo
func NewRepository() *UserRepository {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	mongoURI, _ := os.LookupEnv("MONGO_URI")
	if mongoURI == "" {
		panic(fmt.Errorf("Mongo URI is empty"))
	}
	repo := &UserRepository{db: nil, sessions: freecache.NewCache(100)}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// todo Dont push this
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://gabivlj:%s", mongoURI),
	))
	if err != nil {
		log.Fatal(err)
	}
	repo.db = client.Database("chat-it")
	repo.userCollection = repo.db.Collection("users")
	return repo
}

// SaveUser saves a user into mongo db
func (ur *UserRepository) SaveUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	// todo Handle error
	encryptedPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	mongoU := mongoUser(user)
	mongoU.Password = string(encryptedPass)
	result, err := ur.userCollection.InsertOne(ctx, mongoU)
	if err != nil {
		return nil, err
	}
	id := result.InsertedID.(primitive.ObjectID)
	mongoU.ID = id
	return mongoU.Domain(), nil
}

// LogUser logs an user from the db and creates a session
func (ur *UserRepository) LogUser(ctx context.Context, user *domain.User) (*domain.User, string, error) {
	userMongo := mongoUser(user)
	res := ur.userCollection.FindOne(ctx, userMongo)
	err := res.Decode(&userMongo)
	if err != nil {
		return nil, "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(userMongo.Password), []byte(user.Password)) != nil {
		return nil, "", errors.New("Incorrect password")
	}
	user = userMongo.Domain()
	bytes, _ := json.Marshal(&user)
	s := uuid.UUIDv4()
	ur.sessions.Set([]byte(s), bytes, int(time.Hour)*1000)
	return userMongo.Domain(), s, nil
}

// FindByID returns a user by id
func (ur *UserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	u := domain.User{ID: id}
	userMongo := mongoUser(&u)
	now := time.Now().Unix()
	// Mongo atlas free tier wastes like 2 seconds for finding an user. UNREAL.
	res := ur.userCollection.FindOne(ctx, userMongo)
	fmt.Println("Wasted ", time.Now().Unix()-now, " for getting a user")
	err := res.Decode(&userMongo)
	if err != nil {
		return nil, err
	}
	return userMongo.Domain(), nil
}

// FindByUsername returns a user by id
func (ur *UserRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	u := domain.User{Username: username}
	userMongo := mongoUser(&u)
	// Mongo atlas free tier wastes like 2 seconds for finding an user. UNREAL.
	res := ur.userCollection.FindOne(ctx, userMongo)
	err := res.Decode(&userMongo)
	if err != nil {
		return nil, err
	}
	return userMongo.Domain(), nil
}

// VerifySession returns if the passed session is correct
func (ur *UserRepository) VerifySession(session string) (*domain.User, error) {
	user, err := ur.sessions.Get([]byte(session))
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
