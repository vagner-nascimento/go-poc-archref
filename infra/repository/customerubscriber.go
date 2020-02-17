package repository

import (
	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

type customerSubscriber struct {
	queueInfo   data.QueueInfo
	messageInfo data.MessageInfo
	handler     func(data []byte)
}

func (o *customerSubscriber) QueueInfo() data.QueueInfo {
	return o.queueInfo
}

func (o *customerSubscriber) MessageInfo() data.MessageInfo {
	return o.messageInfo
}

func (o *customerSubscriber) MessageHandler() func([]byte) {
	return o.handler
}

func newCustomerSub() *customerSubscriber {
	return &customerSubscriber{
		queueInfo: data.QueueInfo{
			Name:         "q-customer",
			Durable:      false,
			DeleteUnused: false,
			Exclusive:    false,
			NoWait:       false,
		},
		messageInfo: data.MessageInfo{
			Consumer:  "go-poc-archref",
			AutoAct:   true,
			Exclusive: false,
			Local:     false,
			NoWait:    false,
		},
		handler: func(data []byte) {
			c, err := app.NewCustomerFromBytes(data)
			if err == nil {
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
