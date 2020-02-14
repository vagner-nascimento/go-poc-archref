package data

import (
	"sync"

	"github.com/streadway/amqp"

	"github.com/vagner-nascimento/go-poc-archref/environment"
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
	queuePublisher interface {
		QueueInfo() queueInfo
		MessageInfo() messageInfo
	}
	amqpConfigTp struct {
		once       sync.Once
		localConn  string
		dockerConn string
	}
)

var (
	// TODO: Amqp - realise how put it on app config
	singletonAmqp struct {
		amqoConn    *amqp.Connection
		amqpChannel *amqp.Channel
	}
	amqpConfig = amqpConfigTp{
		localConn:  "amqp://guest:guest@localhost:5672",
		dockerConn: "amqp://guest:guest@go-rabbit-mq:5672",
	}
)

func PublishMessage(p queuePublisher) error {
	ch, err := amqpConnect()
	if err != nil {
		return handleAmqConnectionError(err)
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

	infra.LogInfo("message published into", p.QueueInfo().Name)
	return nil
}

func SubscribeConsumers() error {
	ch, err := amqpConnect()
	if err != nil {
		return handleAmqConnectionError(err)
	}

	customerSub := NewCustomerSub()
	customerReader, err := messageReader(ch, customerSub)
	if err != nil {
		return handleAmqConnectionError(err)
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

func amqpConnect() (*amqp.Channel, error) {
	var err error
	amqpConfig.once.Do(func() {
		if environment.GetEnv() == "docker" {
			singletonAmqp.amqoConn, err = amqp.Dial(amqpConfig.dockerConn)
		} else {
			singletonAmqp.amqoConn, err = amqp.Dial(amqpConfig.localConn)
		}

		if err == nil {
			infra.LogInfo("successfully connected into AMQP server")

			singletonAmqp.amqpChannel, err = singletonAmqp.amqoConn.Channel()
			if err == nil {
				infra.LogInfo("successfully created AMQP channel")
			}
		}
	})

	return singletonAmqp.amqpChannel, err
}

func handleAmqConnectionError(err error) error {
	infra.LogError("error on try to get amqp channel", err)
	return ConnectionError("amqp server")
}
