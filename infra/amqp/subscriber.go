package amqp

import (
	"encoding/json"
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/infra/nosql"

	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

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
			c, _ := app.NewCustomer(nosql.New())

			err := json.Unmarshal(msg.Body, &c)

			if err == nil {
				app.CreateCustomer(&c)
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
