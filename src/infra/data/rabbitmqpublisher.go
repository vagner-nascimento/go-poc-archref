package data

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/logger"
)

type rabbitPubInfo struct {
	queue   rabbitQueueInfo
	message rabbitMessageInfo
	data    amqp.Publishing
}

func PublishIntoRabbit(data []byte, topic string) (err error) {
	pubInfo := newRabbitPubInfo(data, topic)

	if rbCh, err := getRabbitChannel(); err == nil {
		if qPub, err := rbCh.QueueDeclare(
			pubInfo.queue.Name,
			pubInfo.queue.Durable,
			pubInfo.queue.AutoDelete,
			pubInfo.queue.Exclusive,
			pubInfo.queue.NoWait,
			pubInfo.queue.Args,
		); err == nil {
			err = rbCh.Publish(
				pubInfo.message.Exchange,
				qPub.Name,
				pubInfo.message.Mandatory,
				pubInfo.message.Immediate,
				pubInfo.data,
			)
		}
	}

	if err != nil {
		msg := fmt.Sprintf("error on publish data into rabbit queue %s", topic)
		logger.Error(msg, err)
		err = errors.New(msg)
	}

	return err
}

func newRabbitPubInfo(data []byte, topic string) rabbitPubInfo {
	return rabbitPubInfo{
		queue: rabbitQueueInfo{
			Name:       topic,
			Durable:    false,
			AutoDelete: false,
			Exclusive:  false,
			NoWait:     false,
			Args:       nil,
		},
		message: rabbitMessageInfo{
			Exchange:  "",
			Mandatory: false,
			Immediate: false,
		},
		data: amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	}
}
