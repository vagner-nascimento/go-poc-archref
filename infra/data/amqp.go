package data

import (
	"github.com/streadway/amqp"

	"github.com/vagner-nascimento/go-poc-archref/infra"
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
	amqpUrl  = "localhost"
	amqpPort = "5672"
	amqpUser = "guest"
	amqpPass = "guest"
)

func handleAmqError(err error) error {
	infra.LogError("error on try to get amqp channel", err)
	return connectionError("amqp server")
}

func publish(p queuePublishHandler) error {
	ch, err := amqpChannel()
	if err != nil {
		return handleAmqError(err)
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

	infra.LogInfo("message published into", q.Name)
	return nil
}

func SubscribeConsumers() error {
	ch, err := amqpChannel()
	if err != nil {
		return handleAmqError(err)
	}

	customerSub := newCustomerSub()
	customerReader, err := messageReader(ch, customerSub)
	if err != nil {
		return handleAmqError(err)
	}

	go customerReader()

	keepListening := make(chan bool)
	infra.LogInfo("listening to the queues: " + customerSub.QueueInfo().Name)
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
			infra.LogInfo("message received", string(msg.Body))
			handleMessage(msg.Body)
		}
	}, nil
}
