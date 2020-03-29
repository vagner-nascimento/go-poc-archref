package integration

import (
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"github.com/vagner-nascimento/go-poc-archref/src/provider"
)

type enterpriseSub struct {
	topic    string
	consumer string
	handler  func(data []byte)
}

func (es *enterpriseSub) getTopic() string {
	return es.topic
}

func (es *enterpriseSub) getConsumer() string {
	return es.consumer
}

func (es *enterpriseSub) getHandler() func([]byte) {
	return es.handler
}

func newEnterpriseSub() subscription {
	entConf := config.Get().Integration.Amqp.Subs.Enterprise
	return &enterpriseSub{
		topic:    entConf.Topic,
		consumer: entConf.Consumer,
		handler: func(data []byte) {
			if ent, err := model.NewEnterpriseFromJsonBytes(data); err == nil {
				if customerUs, err := provider.SupplierUseCase(); err == nil {
					customerUs.UpdateFromEnterprise(ent)
				}
			} else {
				infra.LogError("error on update a customer", err)
			}
		},
	}
}
