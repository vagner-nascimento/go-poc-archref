package integration

import (
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/logger"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/repository"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"github.com/vagner-nascimento/go-poc-archref/src/provider"
)

type enterpriseSub struct {
	topic    string
	consumer string
	handler  func(data []byte)
}

func (es *enterpriseSub) GetTopic() string {
	return es.topic
}

func (es *enterpriseSub) GetConsumer() string {
	return es.consumer
}

func (es *enterpriseSub) GetHandler() func([]byte) {
	return es.handler
}

func newEnterpriseSub() repository.Subscription {
	entConf := config.Get().Integration.Amqp.Subs.Enterprise
	return &enterpriseSub{
		topic:    entConf.Topic,
		consumer: entConf.Consumer,
		handler: func(data []byte) {
			var err error
			var ent model.Enterprise
			if ent, err = model.NewEnterpriseFromJsonBytes(data); err == nil {
				if supUc, err := provider.SupplierUseCase(); err == nil {
					go supUc.UpdateFromEnterprise(ent)
				}
			}
			if err != nil {
				logger.Error("error on update a customer", err)
			}
		},
	}
}
