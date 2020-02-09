package data

import (
	"fmt"
	"sync"

	"github.com/streadway/amqp"

	"github.com/vagner-nascimento/go-poc-archref/infra"
)

type (
	connSingleton struct {
		AmqpConn *amqp.Connection
	}
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
		Consumer   string
		AutoAct    bool
		Exclusive  bool
		Local      bool
		NoWait     bool
		Exchange   string
		Mandatory  bool
		Immediate  bool
		Publishing amqp.Publishing
		Args       amqp.Table
	}
	queueConsumer interface {
		QueueInfo() queueInfo
		MessageInfo() messageInfo
		MessageHandler() func([]byte)
	}

	queuePublishHandler interface {
		QueueInfo() queueInfo
		MessageInfo() messageInfo
	}
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

func amqpConnection() (*amqp.Connection, error) {
	once.Do(func() {
		singleton.AmqpConn, connError = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", amqpUser, amqpPass, amqpUrl, amqpPort))
		if connError == nil {
			infra.LogInfo("Successfully connected in AMQP server")
		}
	})

	return singleton.AmqpConn, connError
}

func publish(p queuePublishHandler) error {
	conn, err := amqpConnection()
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		p.QueueInfo().Name,
		p.QueueInfo().Durable,
		p.QueueInfo().AutoDelete,
		p.QueueInfo().Exclusive,
		p.QueueInfo().NoWait,
		p.QueueInfo().Args,
	)

	ch.Publish(
		p.MessageInfo().Exchange,
		q.Name,
		p.MessageInfo().Mandatory,
		p.MessageInfo().Immediate,
		p.MessageInfo().Publishing,
	)

	infra.LogInfo("message published on" + q.Name)
	return nil
}

func SubscribeConsumers() error {
	conn, err := amqpConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	customerReader, err := messageReader(ch, newCustomerSub())
	if err != nil {
		return err
	}

	go customerReader()

	keepListening := make(chan bool)
	infra.LogInfo("Listening to the queues")
	<-keepListening

	return nil
}

func messageReader(ch *amqp.Channel, consumer queueConsumer) (func(), error) {
	q, err := ch.QueueDeclare(
		consumer.QueueInfo().Name,
		consumer.QueueInfo().Durable,
		consumer.QueueInfo().DeleteUnused,
		consumer.QueueInfo().Exclusive,
		consumer.QueueInfo().NoWait,
		nil, // Queue Table Args
	)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name,
		consumer.MessageInfo().Consumer,
		consumer.MessageInfo().AutoAct,
		consumer.MessageInfo().Exclusive,
		consumer.MessageInfo().Local,
		consumer.MessageInfo().NoWait,
		nil,
	)
	if err != nil {
		return nil, err
	}

	handleMessage := consumer.MessageHandler()

	return func() {
		for msg := range msgs {
			infra.LogInfo(fmt.Sprintf("Message erecieved:  %s", msg.Body))
			handleMessage(msg.Body)
		}
	}, nil
}
