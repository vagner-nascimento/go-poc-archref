package repository

import (
	"github.com/vagner-nascimento/go-poc-archref/src/app"
	"github.com/vagner-nascimento/go-poc-archref/src/infra"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/data"
)

const (
	userQueue    = "q-user"
	userConsumer = "poc-golang"
)

func newUserSubscriber() data.AmqSubscriber {
	messageHandler := func(data []byte) {
		if _, err := app.UpdateCustomerFromUser(data, &CustomerRepository{}); err != nil {
			infra.LogError("error on update a customer", err)
		}
	}

	return data.NewAmqpSubscriber(userQueue, userConsumer, messageHandler)
}
