package data

import (
	"encoding/json"

	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

type customerSub struct {
	queueInfo
	messageInfo
	handler func(data []byte)
}

func (o customerSub) QueueInfo() queueInfo {
	return o.queueInfo
}

func (o customerSub) MessageInfo() messageInfo {
	return o.messageInfo
}

func (o customerSub) MessageHandler() func([]byte) {
	return o.handler
}

func newCustomerSub() *customerSub {
	return &customerSub{
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
			c := app.NewCustomer(&customerDb{db: mongo{}})
			err := json.Unmarshal(data, &c)
			if err == nil {
				err = app.CreateCustomer(c)
				if err != nil {
					infra.LogError("error on create a customer", err)
				} else {
					infra.LogInfo("customer created")
				}
			} else {
				infra.LogError("error on convert message's body into a customer", err)
			}
		},
	}
}
