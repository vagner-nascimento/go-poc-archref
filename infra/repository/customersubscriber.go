package repository

import (
	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

func newCustomerSubscriber() data.AmqSubscriber {
	consumerMessageHandler := func(data []byte) {
		c, err := app.BuildCustomerFromBytes(data)
		if err == nil {
			err = app.AddCustomer(&c, &CustomerRepository{})
			if err != nil {
				infra.LogError("error on create a customer", err)
			} else {
				u := app.BuildUserFromCustomer(c)
				go app.AddUser(&u, &userRepository{})
			}
		}
	}

	return data.NewAmqpSubscriber("q-customer", "poc-goland", consumerMessageHandler)
}