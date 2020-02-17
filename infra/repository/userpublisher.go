package repository

import (
	"github.com/streadway/amqp"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

// TODO make it a class that uses amqp
type userPubInfo struct {
	queueInfo   data.QueueInfo
	messageInfo data.MessageInfo
	data        []byte
}

func (o userPubInfo) QueueInfo() data.QueueInfo {
	return o.queueInfo
}

func (o userPubInfo) MessageInfo() data.MessageInfo {
	return o.messageInfo
}

func NewUserPub(dataBytes []byte) userPubInfo {
	return userPubInfo{
		queueInfo: data.QueueInfo{
			Name:       "q-user",
			Durable:    false,
			AutoDelete: false,
			Exclusive:  false,
			NoWait:     false,
			Args:       nil,
		},
		messageInfo: data.MessageInfo{
			Exchange:  "",
			Mandatory: false,
			Immediate: false,
			Publishing: amqp.Publishing{
				ContentType: "application/json",
				Body:        dataBytes,
			},
			Args: nil,
		},
	}
}
