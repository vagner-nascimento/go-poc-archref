package dataamqp

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/streadway/amqp"

	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

type connSingleton struct {
	AmqpConn *amqp.Connection
}

const (
	qName         = "q-customer"
	qDurable      = false
	qDeleteUnused = false
	qExclusive    = false
	qNoWait       = false
	mConsumer     = "go-poc-archref"
	mAutoAct      = true
	mExclusive    = false
	mNoLocal      = false
	mNoWait       = false
)

var (
	amqpUrl   = "localhost"
	amqpPort  = "5672"
	amqpUser  = "guest"
	amqpPass  = "guest"
	connError error
	singleton connSingleton
	once      sync.Once
)

func getConnection() (*amqp.Connection, error) {
	once.Do(func() {
		singleton.AmqpConn, connError = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", amqpUser, amqpPass, amqpUrl, amqpPort))
		if connError == nil {
			infra.LogInfo("Successfully connected in AMQP server")
		}
	})

	return singleton.AmqpConn, connError
}

func SubscribeConsumers() error {
	conn, err := getConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		qName,
		qDurable,
		qDeleteUnused,
		qExclusive,
		qNoWait,
		nil, // Queue Table Args
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		mConsumer,
		mAutoAct,
		mExclusive,
		mNoLocal,
		mNoWait,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			infra.LogInfo(fmt.Sprintf("[customer subscriber] Message body %s", msg.Body))
			var c app.Customer
			err := json.Unmarshal(msg.Body, &c)
			if err == nil {
				app.CreateCustomer(c)
			} else {
				infra.LogInfo("Invalid data type from Customer queue:\n" + string(msg.Body))
			}
		}
	}()

	keepListening := make(chan bool)
	infra.LogInfo("Listening to the queues: " + qName)
	<-keepListening

	return nil
}
