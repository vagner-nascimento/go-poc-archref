package repository

import (
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

const (
	userQueue    = "q-user"
	userConsumer = "poc-golang"
)

func newUserSubscriber() data.AmqSubscriber {
	messageHandler := func(data []byte) {
		//u, err := app.MakeUserFromBytes(data) // TODO: messageHandler not done yet
		//if err == nil {
		//	if _, err = app.UpdateCustomer(u, &CustomerRepository{}); err != nil {
		//		infra.LogError("error on update a customer", err)
		//	}
		//}
	}

	return data.NewAmqpSubscriber(userQueue, userConsumer, messageHandler)
}
