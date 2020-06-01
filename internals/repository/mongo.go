package repository

import (
	"github.com/gabivlj/chat-it/internals/graphql/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getObjectIDs(s []string) []primitive.ObjectID {
	ids := make([]primitive.ObjectID, 0, len(s))
	for _, str := range s {
		id, _ := primitive.ObjectIDFromHex(str)
		ids = append(ids, id)
	}
	return ids
}

func unwrapPointerObjectID(s *string) (primitive.ObjectID, error) {
	if s == nil {
		return primitive.NilObjectID, nil
	}
	return primitive.ObjectIDFromHex(*s)
}

func parsePagination(pagination *model.Params, key ...string) (*options.FindOptions, bson.M, error) {
	if key == nil || len(key) == 0 {
		key = []string{"_id"}
	}
	keyFirst := key[0]
	after, err := unwrapPointerObjectID(pagination.After)
	if err != nil {
		return nil, nil, err
	}
	before, err := unwrapPointerObjectID(pagination.Before)
	if err != nil {
		return nil, nil, err
	}
	options := options.Find()
	l := int64(pagination.Limit)
	options.Limit = &l
	var query bson.M
	if after == primitive.NilObjectID && before == primitive.NilObjectID {
		query = bson.M{}
	} else if after != primitive.NilObjectID {
		query = bson.M{keyFirst: bson.M{"$gt": after}}
	} else {
		query = bson.M{keyFirst: bson.M{"$lt": before}}
	}
	return options, query, nil
}
