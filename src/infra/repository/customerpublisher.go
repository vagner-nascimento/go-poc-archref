package repository

import (
	"encoding/json"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/data"
)

type customerPublisher struct {
	handler data.AmqpHandler
	topic   string
}

func (cp *customerPublisher) publish(data interface{}) (err error) {
	if bytes, err := json.Marshal(data); err == nil {
		err = cp.handler.Publish(bytes, cp.topic)
	}
	return err
}

func newCustomerPublisher() (customerPub amqpPublishHandler, err error) {
	if handler, err := data.NewAmqpHandler(); err == nil {
		customerPub = &customerPublisher{handler: handler, topic: config.Get().Integration.Amqp.Pubs.Customer.Topic}
	}
	return customerPub, err
}
