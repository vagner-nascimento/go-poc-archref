package data

import (
	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

type customerSubscriber struct {
	queueInfo   queueInfo
	messageInfo messageInfo
	handler     func(data []byte)
}

func (o *customerSubscriber) QueueInfo() queueInfo {
	return o.queueInfo
}

func (o *customerSubscriber) MessageInfo() messageInfo {
	return o.messageInfo
}

func (o *customerSubscriber) MessageHandler() func([]byte) {
	return o.handler
}

func NewCustomerSub() *customerSubscriber {
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
			c, err := app.NewCustomerFromBytes(data)
			if err == nil { // TODO: realise how to do it without cycle dependency (data cannot depends on repo)
				err = app.AddCustomer(&c, &customerRepository{})
				if err != nil {
					infra.LogError("error on create a customer", err)
				} else {
					infra.LogInfo("customer created")

					u := app.NewUserFromCustomer(c)
					go app.AddUser(&u, &userRepository{})
				}
			} else {
				infra.LogError("error on convert message's body into a Customer", err)
			}
		},
	}
}
