package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/logger"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: log and send handled error to repo pkg
type mongoDb struct {
	collection *mongo.Collection
}

func (o *mongoDb) InsertOne(entity interface{}) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), mongoConnection.insertTimeout*time.Second)
	if _, err := o.collection.InsertOne(ctx, entity); err != nil {
		return nil, err
	}
	return entity, nil
}

func (o *mongoDb) FindOne(id interface{}) ([]byte, error) {
	filter := bson.D{{"id", id}}
	ctx, _ := context.WithTimeout(context.Background(), mongoConnection.findTimeout*time.Second)
	raw, err := o.collection.FindOne(ctx, filter).DecodeBytes()
	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			return nil, nil
		}
		return nil, err
	}

	result, err := bson.Marshal(raw)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o *mongoDb) Find(filters bson.D, maxDocs int64, page int64, results chan interface{}, total *int64) {
	if page < 0 {
		results <- errors.New(fmt.Sprintf("invalid parameters, maxDocs: %s, page: %s", maxDocs, page))
		close(results)
		return
	}
	if maxDocs <= 0 {
		maxDocs = config.Get().Data.NoSql.Mongo.MaxPaginatedSearch
	}

	shouldExit := func(err error) bool {
		if err != nil {
			results <- err
			close(results)
			return true
		}
		return false
	}

	var err error
	ctx, _ := context.WithTimeout(context.Background(), mongoConnection.findTimeout*time.Second)
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

func (o *mongoDb) ReplaceOne(filter bson.M, newData interface{}) (int64, error) {
	ctx, _ := context.WithTimeout(context.Background(), mongoConnection.findTimeout*time.Second)
	res, err := o.collection.ReplaceOne(ctx, filter, newData)
	if err != nil {
		return 0, err
	}

	return res.ModifiedCount, nil
}

func NewNoSqlDb(collectionName string) (NoSqlHandler, error) {
	if err := mongoDbConnect(); err != nil {
		return nil, err
	}

	return &mongoDb{
		collection: mongoConnection.database.Collection(collectionName),
	}, nil
}

var (
	mongoConnection struct {
		once               sync.Once
		database           *mongo.Database
		clientTimeOut      time.Duration
		insertTimeout      time.Duration
		findTimeout        time.Duration
		maxPaginatedSearch int64
	}
)

func mongoDbConnect() (err error) {
	first := false
	mongoConnection.once.Do(func() {
		first = true
		mongoConf := config.Get().Data.NoSql.Mongo
		ctx, _ := context.WithTimeout(context.Background(), mongoConf.ClientTimeOut*time.Second)
		cliOpts := options.Client().ApplyURI(mongoConf.ConnStr)
		if client, err := mongo.Connect(ctx, cliOpts); err == nil {
			if err = client.Ping(context.TODO(), nil); err == nil { // TODO: MONGO PING realise why is crashing app when mongo is down
				mongoConnection.database = client.Database(mongoConf.Database)
				logger.Info(fmt.Sprintf("successfully connected into mongo database %s", mongoConf.Database))
				setMongoConfigs(mongoConf)
			}
		}
	})

	if !first {
		err = mongoConnection.database.Client().Ping(context.TODO(), nil)
	}

	return err
}

func setMongoConfigs(mongoConf config.MongoDataConfig) {
	mongoConnection.findTimeout = mongoConf.FindTimeout
	mongoConnection.insertTimeout = mongoConf.InsertTimeout
	mongoConnection.maxPaginatedSearch = mongoConf.MaxPaginatedSearch
}
