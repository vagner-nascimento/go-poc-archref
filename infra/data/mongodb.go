package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/vagner-nascimento/go-poc-archref/environment"
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

type mongoConfigTp struct {
	clientTimeout time.Duration
	insertTimeout time.Duration
	findTimeout   time.Duration
	localConn     string
	dockerConn    string
	once          sync.Once
}

var (
	singletonMongo struct {
		client *mongo.Client
	}
	mongoConfig = mongoConfigTp{
		clientTimeout: 15,
		insertTimeout: 5,
		findTimeout:   8,
		localConn:     "mongodb://admin:admin@localhost:27017",
		dockerConn:    "mongodb://admin:admin@go-mongodb:27017",
	} // TODO: Mongo - realise how put connection into app config
)

type MongoDb struct {
	collection *mongo.Collection
}

func (o *MongoDb) Insert(entity interface{}) (interface{}, error) {

	ctx, _ := context.WithTimeout(context.Background(), mongoConfig.insertTimeout*time.Second)
	if _, err := o.collection.InsertOne(ctx, entity); err != nil {
		return nil, execError(err, "insertOne", "mongodb server")
	}

	return entity, nil
}

func (o *MongoDb) Find(filters bson.D, limit int64) ([]interface{}, error) {
	if limit <= 0 {
		limit = 1
	}

	options := options.Find()
	options.SetLimit(limit)

	ctx, _ := context.WithTimeout(context.Background(), mongoConfig.findTimeout*time.Second)
	cur, err := o.collection.Find(ctx, filters, options)
	if err != nil {
		return nil, execError(err, "find", "mongodb server")
	}

	// TODO: Realise how to send the result
	var results []interface{}
	for cur.Next(ctx) {
		var r interface{}

		err := cur.Decode(&r)
		if err != nil {
			return nil, execError(err, "find", "mongodb server")
		}

		results = append(results, r)
	}

	if err := cur.Err(); err != nil {
		return nil, execError(err, "find", "mongodb server")
	}

	cur.Close(ctx)

	return results, nil
}

func NewMongoDb(collectionName string) (*MongoDb, error) {
	if err := mongoDbConnect(); err != nil {
		return nil, connectionError(err, "mongodb server")
	}

	return &MongoDb{
		collection: singletonMongo.client.Database("golang").Collection(collectionName),
	}, nil
}

func mongoDbConnect() error {
	var err error
	mongoConfig.once.Do(func() {
		var cliOpts *options.ClientOptions
		if environment.GetEnv() == "docker" {
			cliOpts = options.Client().ApplyURI(mongoConfig.dockerConn)
		} else {
			cliOpts = options.Client().ApplyURI(mongoConfig.localConn)
		}

		ctx, _ := context.WithTimeout(context.Background(), mongoConfig.clientTimeout*time.Second)
		singletonMongo.client, err = mongo.Connect(ctx, cliOpts)
		if err != nil {
			return
		}

		if err = singletonMongo.client.Ping(context.TODO(), nil); err == nil {
			infra.LogInfo("successfully connected into MongoDb server")
		}
	})

	if err != nil {
		return connectionError(err, "mongodb server")
	}

	return nil
}

// TODO: Realise how to set mongo db indexes and make customers email unique
// func SetMongoDbIndexes() error {
// 	return nil
// }
