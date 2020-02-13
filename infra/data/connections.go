package data

import (
	"context"
	"os"
	"sync"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/vagner-nascimento/go-poc-archref/infra"
)

var (
	amqpOnce    sync.Once
	mongoOne    sync.Once
	amqpLocal   = "amqp://guest:guest@localhost:5672"
	mongoLocal  = "mongodb://admin:admin@localhost:27017"
	amqpDocker  = "amqp://guest:guest@go-rabbit-mq:5672"
	mongoDocker = "mongodb://admin:admin@go-mongodb:27017"
	singletons  struct {
		amqoConn    *amqp.Connection
		amqpChannel *amqp.Channel
		mongoClient *mongo.Client
	}
	goEnv = os.Getenv("GO_ENV")
)

func amqpChannel() (*amqp.Channel, error) {
	var err error
	amqpOnce.Do(func() {
		if goEnv == "docker" {
			singletons.amqoConn, err = amqp.Dial(amqpDocker)
		} else {
			singletons.amqoConn, err = amqp.Dial(amqpLocal)
		}

		if err == nil {
			infra.LogInfo("successfully connected into AMQP server")

			singletons.amqpChannel, err = singletons.amqoConn.Channel()
			if err == nil {
				infra.LogInfo("successfully created AMQP channel")
			}
		}
	})

	return singletons.amqpChannel, err
}

func mongoDbClient() (*mongo.Client, error) {
	var err error
	mongoOne.Do(func() {
		var cliOpts *options.ClientOptions
		if goEnv == "docker" {
			cliOpts = options.Client().ApplyURI(mongoDocker)
		} else {
			cliOpts = options.Client().ApplyURI(mongoLocal)
		}

		singletons.mongoClient, err = mongo.Connect(context.TODO(), cliOpts)
		if err != nil {
			return
		}

		err = singletons.mongoClient.Ping(context.TODO(), nil)
		if err == nil {
			infra.LogInfo("successfully connected into MongoDb server")
		}
	})

	return singletons.mongoClient, err
}
