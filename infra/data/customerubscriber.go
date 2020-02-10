package data

import (
	"encoding/json"

	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

type customerSubscriber struct {
	queueInfo
	messageInfo
	handler func(data []byte)
}

func (o customerSubscriber) QueueInfo() queueInfo {
	return o.queueInfo
}

func (o customerSubscriber) MessageInfo() messageInfo {
	return o.messageInfo
}

func (o customerSubscriber) MessageHandler() func([]byte) {
	return o.handler
}

func newCustomerSub() *customerSubscriber {
	return &customerSubscriber{
		queueInfo: queueInfo{
			Name:         "q-customer",
			Durable:      false,
			DeleteUnused: false,
			Exclusive:    false,
			NoWait:       false,
		},
		messageInfo: messageInfo{
			Consumer:  "go-poc-archref",
			AutoAct:   true,
			Exclusive: false,
			Local:     false,
			NoWait:    false,
		},
		handler: func(data []byte) {
			c := app.NewCustomer(&customerDb{})
			err := json.Unmarshal(data, &c)
			if err == nil {
				err = app.CreateCustomer(c)
				if err != nil {
					infra.LogError("error on create a customer", err)
				} else {
					infra.LogInfo("customer created")
					// TODO save user and send to queue with its ID
					go publish(newUserPub(data))
				}
			} else {
				infra.LogError("error on convert message's body into a customer", err)
			}
		},
	}
}
