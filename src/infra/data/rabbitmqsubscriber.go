package data

import (
	"fmt"
	"github.com/streadway/amqp"
)

type rabbitSubInfo struct {
	queue   rabbitQueueInfo
	message rabbitMessageInfo
	handler func([]byte)
}

func SubscribeConsumer(queueName string, consumerName string, handler func([]byte)) error {
	rbChan, err := getRabbitChannel()
	if err != nil {
		return err
	}

	sub := newRabbitSubInfo(queueName, consumerName, handler)
	if err = processMessages(rbChan, sub); err != nil {
		return err
	}

	return nil
}

// newRabbitSubInfo - initialize rabbit queue and message with variables and default data
func newRabbitSubInfo(queueName string, consumerName string, handler func([]byte)) rabbitSubInfo {
	return rabbitSubInfo{
		queue: rabbitQueueInfo{
			Name:         queueName,
			Durable:      false,
			DeleteUnused: false,
			Exclusive:    false,
			NoWait:       false,
		},
		message: rabbitMessageInfo{
			Consumer:  consumerName,
			AutoAct:   true,
			Exclusive: false,
			Local:     false,
			NoWait:    false,
		},
		handler: handler,
	}
}

func processMessages(ch *amqp.Channel, sub rabbitSubInfo) error {
	q, err := ch.QueueDeclare(
		sub.queue.Name,
		sub.queue.Durable,
		sub.queue.DeleteUnused,
		sub.queue.Exclusive,
		sub.queue.NoWait,
		sub.queue.Args,
	)
	if err != nil {
		return err
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
		return err
	}

	go func(queueName string, msgs <-chan amqp.Delivery) {
		for msg := range msgs {
			fmt.Println(fmt.Sprintf("message received from %s. body:\r\n %s", q.Name, string(msg.Body)))
			sub.handler(msg.Body)
		}
	}(q.Name, msgs)

	return nil
}
