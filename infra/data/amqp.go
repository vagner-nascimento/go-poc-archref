package dataamqp

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/streadway/amqp"

	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

type queueInfo struct {
	QName         string
	QDurable      bool
	QDeleteUnused bool
	QExclusive    bool
	QNoWait       bool
}

type messageInfo struct {
	MConsumer  string
	MAutoAct   bool
	MExclusive bool
	MNoLocal   bool
	MNoWait    bool
}

type connSingleton struct {
	AmqpConn *amqp.Connection
}

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

	cosumerFunc, err := customerSubscribe(ch)
	if err != nil {
		return err
	}

	go cosumerFunc()

	keepListening := make(chan bool)
	infra.LogInfo("Listening to the queues")
	<-keepListening

	return nil
}

func customerSubscribe(ch *amqp.Channel) (func(), error) {
	customerInfo := getCustomerInfo()
	q, err := ch.QueueDeclare(
		customerInfo.Queue.QName,
		customerInfo.Queue.QDurable,
		customerInfo.Queue.QDeleteUnused,
		customerInfo.Queue.QExclusive,
		customerInfo.Queue.QNoWait,
		nil, // Queue Table Args
	)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name,
		customerInfo.Message.MConsumer,
		customerInfo.Message.MAutoAct,
		customerInfo.Message.MExclusive,
		customerInfo.Message.MNoLocal,
		customerInfo.Message.MNoWait,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return func() {
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
	}, nil
}
