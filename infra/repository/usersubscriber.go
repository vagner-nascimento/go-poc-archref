package repository

import (
	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

const (
	userQueue    = "q-user"
	userConsumer = "poc-golang"
)

//TODO realise why it is closing after do operation
func newUserSubscriber() data.AmqSubscriber {
	messageHandler := func(data []byte) {
		if _, err := app.UpdateCustomerFromUser(data, &CustomerRepository{}); err != nil {
			infra.LogError("error on update a customer", err)
		}
	}

	return data.NewAmqpSubscriber(userQueue, userConsumer, messageHandler)
}
