package data

import (
	"github.com/streadway/amqp"
)

type userPub struct {
	queueInfo
	messageInfo
	data []byte
}

func (o userPub) QueueInfo() queueInfo {
	return o.queueInfo
}

func (o userPub) MessageInfo() messageInfo {
	return o.messageInfo
}

func newUserPub(data []byte) userPub {
	return userPub{
		queueInfo: queueInfo{
			Name:       "q-user",
			Durable:    false,
			AutoDelete: false,
			Exclusive:  false,
			NoWait:     false,
			Args:       nil,
		},
		messageInfo: messageInfo{
			Exchange:  "",
			Mandatory: false,
			Immediate: false,
			Publishing: amqp.Publishing{
				ContentType: "application/json",
				Body:        data,
			},
			Args: nil,
		},
	}
}
