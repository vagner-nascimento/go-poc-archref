package data

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"sync"

	"github.com/vagner-nascimento/go-poc-archref/src/infra"
)

type (
	queueInfo struct {
		Name         string
		Durable      bool
		DeleteUnused bool
		AutoDelete   bool
		Exclusive    bool
		NoWait       bool
		Args         amqp.Table
	}
	messageInfo struct {
		Consumer  string
		AutoAct   bool
		Exclusive bool
		Local     bool
		NoWait    bool
		Exchange  string
		Mandatory bool
		Immediate bool
		Args      amqp.Table
	}
)

var (
	singletonAmqp struct {
		amqoConn    *amqp.Connection
		amqpChannel *amqp.Channel
	}
)

type AmqPublisher struct {
	queue   queueInfo
	message messageInfo
}

func (o *AmqPublisher) Publish(data []byte) error {
	ch, err := amqpConnect()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		o.queue.Name,
		o.queue.Durable,
		o.queue.AutoDelete,
		o.queue.Exclusive,
		o.queue.NoWait,
		o.queue.Args,
	)

	if err != nil {
		return err
	}

	err = ch.Publish(
		o.message.Exchange,
		q.Name,
		o.message.Mandatory,
		o.message.Immediate,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func NewAmqpPublisher(queueName string) (*AmqPublisher, error) {
	return &AmqPublisher{
		queue: queueInfo{
			Name:       queueName,
			Durable:    false,
			AutoDelete: false,
			Exclusive:  false,
			NoWait:     false,
			Args:       nil,
		},
		message: messageInfo{
			Exchange:  "",
			Mandatory: false,
			Immediate: false,
			Args:      nil,
		},
	}, nil
}

type AmqSubscriber struct {
	queue   queueInfo
	message messageInfo
	handler func([]byte)
}

func NewAmqpSubscriber(queueName string, consumerName string, handler func([]byte)) AmqSubscriber {
	return AmqSubscriber{
		queue: queueInfo{
			Name:         queueName,
			Durable:      false,
			DeleteUnused: false,
			Exclusive:    false,
			NoWait:       false,
		},
		message: messageInfo{
			Consumer:  consumerName,
			AutoAct:   true,
			Exclusive: false,
			Local:     false,
			NoWait:    false,
		},
		handler: handler,
	}
}

func SubscribeConsumers(subscribers []AmqSubscriber) error {
	ch, err := amqpConnect()
	if err != nil {
		return err
	}

	var qNames string
	for i := 0; i < len(subscribers); i = i + 1 {
		c := subscribers[i]
		handler, err := messageHandlers(ch, c)
		if err != nil {
			continue
		}

		go handler()

		if qNames == "" {
			qNames = qNames + c.queue.Name
		} else {
			qNames = fmt.Sprintf("%s, %s", qNames, c.queue.Name)
		}
	}

	if qNames != "" {
		infra.LogInfo("listening to the queues: " + qNames)

		keepListening := make(chan bool)
		<-keepListening
	}

	return errors.New("none queue can be listened")
}

func messageHandlers(ch *amqp.Channel, sub AmqSubscriber) (func(), error) {
	q, err := ch.QueueDeclare(
		sub.queue.Name,
		sub.queue.Durable,
		sub.queue.DeleteUnused,
		sub.queue.Exclusive,
		sub.queue.NoWait,
		sub.queue.Args, // Queue Table Args
	)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name,
		sub.message.Consumer,
		sub.message.AutoAct,
		sub.message.Exclusive,
		sub.message.Local,
		sub.message.NoWait,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return func() {
		for msg := range msgs {
			infra.LogInfo("message received, body: ", string(msg.Body))
			sub.handler(msg.Body)
		}
	}, nil
}

var amqpOnce sync.Once

func amqpConnect() (*amqp.Channel, error) {
	var err error
	amqpOnce.Do(func() {
		if singletonAmqp.amqoConn, err = amqp.Dial(config.Get().Data.Amqp.ConnStr); err == nil {
			infra.LogInfo("successfully connected into AMQP server")

			if singletonAmqp.amqpChannel, err = singletonAmqp.amqoConn.Channel(); err == nil {
				infra.LogInfo("successfully created AMQP channel")
			}
		}
	})

	return singletonAmqp.amqpChannel, err
}
