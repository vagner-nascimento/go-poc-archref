package data

import (
	"fmt"
	"github.com/streadway/amqp"
	"sync"

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
	amqpConfigTp struct {
		once       sync.Once
		localConn  string
		dockerConn string
	}
)

var (
	singletonAmqp struct {
		amqoConn    *amqp.Connection
		amqpChannel *amqp.Channel
	}
	// TODO: Amqp - realise how put connection infos it on app config
	amqpConfig = amqpConfigTp{
		localConn:  "amqp://guest:guest@localhost:5672",
		dockerConn: "amqp://guest:guest@go-rabbit-mq:5672",
	}
)

type AmqPublisher struct {
	channel *amqp.Channel
	queue   queueInfo
	message messageInfo
}

func (o *AmqPublisher) Publish(data []byte) error {
	q, err := o.channel.QueueDeclare(
		o.queue.Name,
		o.queue.Durable,
		o.queue.AutoDelete,
		o.queue.Exclusive,
		o.queue.NoWait,
		o.queue.Args,
	)

	if err != nil {
		return execError(err, "declare queue", "amqp server")
	}

	err = o.channel.Publish(
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
		return execError(err, "publish message", "amqp server")
	}

	return nil
}

func NewAmqpPublisher(queueName string) (*AmqPublisher, error) {
	ch, err := amqpConnect()
	if err != nil {
		return nil, connectionError(err, "amqp server")
	}

	return &AmqPublisher{
		channel: ch,
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
		return connectionError(err, "amqp server")
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

	return simpleError("none queue can be listened")
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
		return nil, execError(err, "declare a channel", "amqp server")
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
		return nil, execError(err, "consume a queue", "amqp server")
	}

	return func() {
		for msg := range msgs {
			infra.LogInfo("message received, body: ", string(msg.Body))
			sub.handler(msg.Body)
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
