package repository

import (
	"encoding/json"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/data"
)

type supplierPublisher struct {
	handler data.AmqpHandler
	topic   string
}

func (sp *supplierPublisher) publish(data interface{}) (err error) {
	if bytes, err := json.Marshal(data); err == nil {
		err = sp.handler.Publish(bytes, sp.topic)
	}
	return err
}

func newSupplierPublisher() (supplierPub amqpPublishHandler, err error) {
	if handler, err := data.NewAmqpHandler(); err == nil {
		supplierPub = &supplierPublisher{handler: handler, topic: config.Get().Integration.Amqp.Pubs.Supplier.Topic}
	}
	return supplierPub, err
}
