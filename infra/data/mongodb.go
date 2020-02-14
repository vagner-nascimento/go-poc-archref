package data

import (
	"context"
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
	localConn     string
	dockerConn    string
	once          sync.Once
}

var (
	singletonMongoClient struct {
		client *mongo.Client
	}
	mongoConfig = mongoConfigTp{
		clientTimeout: 15,
		insertTimeout: 5,
		localConn:     "mongodb://admin:admin@localhost:27017",
		dockerConn:    "mongodb://admin:admin@go-mongodb:27017",
	} // TODO: Mongo - realise how put connection into app config
)

type MongoDb struct {
	collection *mongo.Collection
}

func (o *MongoDb) Insert(entity interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), mongoConfig.insertTimeout*time.Second)
	_, err := o.collection.InsertOne(ctx, entity)
	if err != nil {
		infra.LogError("error on try to insert data into MongoDb", err)
		return nil, ExecError("insert one", o.collection.Name())
	}

	return entity, nil
}

func NewMongoDb(collectionName string) (*MongoDb, error) {
	var db *MongoDb

	err := mongoDbConnect()
	if err != nil {
		return db, err
	}
	return &MongoDb{
		collection: singletonMongoClient.client.Database("golang").Collection(collectionName),
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
		singletonMongoClient.client, err = mongo.Connect(ctx, cliOpts)
		if err != nil {
			return
		}

		err = singletonMongoClient.client.Ping(context.TODO(), nil)
		if err == nil {
			infra.LogInfo("successfully connected into MongoDb server")
		}
	})

	if err != nil {
		infra.LogError("error on try connect into mongo db", err)
		return ConnectionError("database server")
	}

	return nil
}

// TODO: Realise how to set mongo db indexes
// func SetMongoDbIndexes() error {
// 	return nil
// }
