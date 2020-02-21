package repository

import (
	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

const userQueue = "q-user"
const userConsumer = "poc-golang"

func newUserSubscriber() data.AmqSubscriber {
	consumerMessageHandler := func(data []byte) {
		c, err := app.BuildCustomerFromBytes(data) // TODO: Create a user model and convert into a customer
		if err == nil {
			err = app.UpdateCustomer(&c, &CustomerRepository{})
			if err != nil {
				infra.LogError("error on create a customer", err)
			}
		}
	}
	return data.NewAmqpSubscriber(userQueue, userConsumer, consumerMessageHandler)
}
