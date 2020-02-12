package docker

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

const (
	maxRetry  = 5
	sleepSec  = 2
	amqpConn  = "amqp://guest:guest@go-rabbit-mq:5672"
	mongoConn = "mongodb://admin:admin@go-mongodb:27017"
)

func waitMongo(ch chan string) {
	opt := options.Client().ApplyURI(mongoConn)
	retry := 1
	client, err := mongo.Connect(context.TODO(), opt)
	for err != nil || retry >= maxRetry {
		ch <- fmt.Sprintf("retry %d, mongo not ready yet", retry)
		time.Sleep(sleepSec * time.Second)
		retry += 1
		client, err = mongo.Connect(context.TODO(), opt)
	}

	if retry <= 0 {
		ch <- "fail to connect with mongo"
		close(ch)
		return
	}

	ch <- "mongo ready to use"
	client.Disconnect(context.TODO())
	close(ch)
	return
}

func waitAmqp(ch chan string) {
	conn, err := amqp.Dial(amqpConn)
	retry := 1

	for (retry > maxRetry) || err != nil { //looping no working
		ch <- fmt.Sprintf("retry %d of %d, amq server not ready yet", retry, maxRetry)
		time.Sleep(sleepSec * time.Second)
		retry = retry + 1
		conn, err = amqp.Dial(amqpConn)
	}

	if retry > maxRetry {
		ch <- "fail to connect with amq"
		close(ch)
		return
	}

	ch <- "amq server ready to use"
	conn.Close()
	close(ch)
	return
}

func WaitForInfra() {
	if os.Getenv("GO_ENV") == "docker" {
		chMongo := make(chan string)
		chAmqp := make(chan string)

		go waitMongo(chMongo)
		go waitAmqp(chAmqp)

		for msg := range chMongo {
			fmt.Println(msg)
		}

		for msg := range chAmqp {
			fmt.Println(msg)
		}
	}
}
