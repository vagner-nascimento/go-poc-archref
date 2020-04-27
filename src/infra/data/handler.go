package data

import "go.mongodb.org/mongo-driver/bson"

type NoSqlHandler interface {
	InsertOne(entity interface{}) (interface{}, error)
	FindOne(id interface{}) ([]byte, error)
	// TODO: Realise another generic filter to use instead bson obj
	Find(filters bson.D, maxDocs int64, page int64, results chan interface{}, total *int64)
	ReplaceOne(filter bson.M, newData interface{}) (int64, error)
}
