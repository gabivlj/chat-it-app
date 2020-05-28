package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

func getObjectIDs(s []string) []primitive.ObjectID {
	ids := make([]primitive.ObjectID, 0, len(s))
	for _, str := range s {
		id, _ := primitive.ObjectIDFromHex(str)
		ids = append(ids, id)
	}
	return ids
}
