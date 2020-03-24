package data

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra"
	"strings"
	"sync"
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

var singletonAmqp struct {
	amqoConn    *amqp.Connection
	amqpChannel *amqp.Channel
}

// TODO: Make interfaces of data to interact with repository
type RabbitPublisher struct {
	queue   queueInfo
	message messageInfo
}

func (o *RabbitPublisher) Publish(data []byte) error {
	ch, err := amqpChannel()
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

func NewRabbitPublisher(queueName string) (*RabbitPublisher, error) {
	return &RabbitPublisher{
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

type RabbitSubscriber struct {
	queue   queueInfo
	message messageInfo
	handler func([]byte)
}

func NewRabbitSubscriber(queueName string, consumerName string, handler func([]byte)) RabbitSubscriber {
	return RabbitSubscriber{
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

func SubscribeRabbitConsumers(subscribers []RabbitSubscriber) error {
	ch, err := amqpChannel()
	if err != nil {
		return err
	}

	var qNames []string
	for i := 0; i < len(subscribers); i = i + 1 {
		c := subscribers[i]
		processMsgs, err := processMessages(ch, c)
		if err != nil {
			continue // TODO: realise how to properly handler this kind of error
		}
		go processMsgs()
		qNames = append(qNames, c.queue.Name)
	}

	if len(qNames) <= 0 {
		return errors.New("none queue to be listened")
	}

	infra.LogInfo("listening to the queues: " + strings.Join(qNames, ","))
	return nil
}

func processMessages(ch *amqp.Channel, sub RabbitSubscriber) (func(), error) {
	q, err := ch.QueueDeclare(
		sub.queue.Name,
		sub.queue.Durable,
		sub.queue.DeleteUnused,
		sub.queue.Exclusive,
		sub.queue.NoWait,
		sub.queue.Args,
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
			infra.LogInfo(fmt.Sprintf("message received from %s. body:\r\n %s", q.Name, string(msg.Body)))
			sub.handler(msg.Body)
		}
	}, nil
}

var amqpOnce sync.Once

func amqpChannel() (*amqp.Channel, error) {
	var err error
	amqpOnce.Do(func() {
		if singletonAmqp.amqoConn, err = amqp.Dial(config.Get().Data.Amqp.ConnStr); err == nil {
			infra.LogInfo("successfully connected into AMQP server")

			if singletonAmqp.amqpChannel, err = singletonAmqp.amqoConn.Channel(); err == nil {
				infra.LogInfo("successfully created AMQP channel")
			}
		}
	})

	if singletonAmqp.amqpChannel == nil && err == nil {
		err = errors.New("cannot open channel into amqp sever")
	}

	return singletonAmqp.amqpChannel, err
}
