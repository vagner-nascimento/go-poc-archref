package data

import (
	"github.com/streadway/amqp"
)

type userPubInfo struct {
	queueInfo
	messageInfo
	data []byte
}

func (o userPubInfo) QueueInfo() queueInfo {
	return o.queueInfo
}

func (o userPubInfo) MessageInfo() messageInfo {
	return o.messageInfo
}

func newUserPub(data []byte) userPubInfo {
	return userPubInfo{
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
