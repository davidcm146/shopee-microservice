package common

import "go.mongodb.org/mongo-driver/bson"

func NotDeletedFilter() bson.M {
	return bson.M{"isDeleted": false}
}
