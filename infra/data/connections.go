package data

import (
	"context"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/vagner-nascimento/go-poc-archref/infra"
)

var (
	amqpOnce   sync.Once
	mongoOne   sync.Once
	singletons struct {
		amqoConn    *amqp.Connection
		amqpChannel *amqp.Channel
		mongoClient *mongo.Client
	}
)

func amqpChannel() (*amqp.Channel, error) {
	var err error
	amqpOnce.Do(func() {
		singletons.amqoConn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", amqpUser, amqpPass, amqpUrl, amqpPort))
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
		cliOpts := options.Client().ApplyURI("mongodb://admin:admin@localhost:27017")
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

func CloseConnections() {
	err := singletons.mongoClient.Disconnect(context.TODO())
	if err != nil {
		infra.LogError("error on close mongo db connection", err)
	}

	if !singletons.amqoConn.IsClosed() {
		err := singletons.amqpChannel.Close()
		if err != nil {
			infra.LogError("error on close amqp channel", err)
		}

		err = singletons.amqoConn.Close()
		if err != nil {
			infra.LogError("error on close amqp connection", err)
		}
	}
}
