package constants

import "go.mongodb.org/mongo-driver/bson"

// SortDescendingCreatedAt returns the configuration for having descending elements
var SortDescendingCreatedAt = bson.M{"createdAt": -1}
