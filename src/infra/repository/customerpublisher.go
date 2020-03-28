package repository

import (
	"encoding/json"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/data"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
)

type customerPublisher struct {
	handler data.AmqpHandler
	topic   string
}

func newCustomerPublisher() (customerPub *customerPublisher, err error) {
	if handler, err := data.NewAmqpHandler(); err == nil {
		customerPub = &customerPublisher{handler: handler, topic: config.Get().Integration.Amqp.Pubs.Customer.Topic}
	}
	return customerPub, err
}

func (cp *customerPublisher) publish(customer model.Customer) error {
	data, err := json.Marshal(customer)
	if err != nil {
		return err
	}

	return cp.handler.Publish(data, cp.topic)
}
