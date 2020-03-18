package data

import (
	"context"
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/vagner-nascimento/go-poc-archref/src/infra"
)

var (
	singletonMongo struct {
		client *mongo.Client
	}
)

type MongoDb struct {
	collection *mongo.Collection
}

func (o *MongoDb) Insert(entity interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), config.Get().Data.NoSql.Mongo.InsertTimeout*time.Second)
	if _, err := o.collection.InsertOne(ctx, entity); err != nil {
		return nil, execError(err, "insertOne", "mongodb server")
	}

	return entity, nil
}

func (o *MongoDb) FindOne(filters bson.D) ([]byte, error) {
	ctx, _ := context.WithTimeout(context.Background(), config.Get().Data.NoSql.Mongo.FindTimeout*time.Second)
	raw, err := o.collection.FindOne(ctx, filters).DecodeBytes()
	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			return nil, nil
		}
		return nil, execError(err, "find one", "mongodb server")
	}

	result, err := bson.Marshal(raw)
	if err != nil {
		return nil, execError(err, "find one", "mongodb server")
	}
	return result, nil
}

func (o *MongoDb) Find(filters bson.D, maxDocs int64, page int64, results chan interface{}, total *int64) {
	if maxDocs <= 0 || page < 0 {
		results <- simpleError(fmt.Sprintf("invalid parameters, maxDocs: %s, page: %s", maxDocs, page))
		close(results)
		return
	}

	shouldExit := func(err error) (exit bool) {
		if err != nil {
			results <- err
			close(results)
			return true
		}
		return false
	}

	var err error

	ctx, _ := context.WithTimeout(context.Background(), config.Get().Data.NoSql.Mongo.FindTimeout*time.Second)
	*total, err = o.collection.CountDocuments(ctx, filters)
	if shouldExit(err) {
		return
	}

	opts := options.Find().SetLimit(maxDocs).SetSkip(maxDocs * page)
	cur, err := o.collection.Find(ctx, filters, opts)
	if shouldExit(err) {
		return
	}

	for cur.Next(ctx) {
		item, err := bson.Marshal(cur.Current)
		if shouldExit(err) {
			break
		}
		results <- item
	}

	if err := cur.Err(); err == nil {
		cur.Close(ctx)
	}

	close(results)
}

func (o *MongoDb) ReplaceOne(filter bson.M, newData interface{}) (int64, error) {
	ctx, _ := context.WithTimeout(context.Background(), config.Get().Data.NoSql.Mongo.FindTimeout*time.Second)

	res, err := o.collection.ReplaceOne(ctx, filter, newData)
	if err != nil {
		return 0, err
	}

	return res.ModifiedCount, nil
}

func NewMongoDb(collectionName string) (*MongoDb, error) {
	if err := mongoDbConnect(); err != nil {
		return nil, connectionError(err, "mongodb server")
	}

	return &MongoDb{
		collection: singletonMongo.client.Database("golang").Collection(collectionName),
	}, nil
}

var mongoOnce sync.Once

func mongoDbConnect() (err error) {
	mongoOnce.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), config.Get().Data.NoSql.Mongo.ClientTimeOut*time.Second)
		cliOpts := options.Client().ApplyURI(config.Get().Data.NoSql.Mongo.ConnStr)
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

	return err
}
