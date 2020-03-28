package provider

import (
	"github.com/vagner-nascimento/go-poc-archref/src/app"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/repository"
	"sync"
)

var resources struct {
	customerUseCase app.CustomerUseCase
	amqSub          repository.AmqpSubscriptionHandler
}

var once struct {
	customerUs sync.Once
	amqSub     sync.Once
}

func CustomerUseCase() (app.CustomerUseCase, error) {
	var err error
	once.customerUs.Do(func() {
		if repo, err := repository.NewCustomerRepository(); err == nil {
			resources.customerUseCase = app.NewCustomerUseCase(repo)
		}
	})
	return resources.customerUseCase, err
}

func AmqpSubscription() (repository.AmqpSubscriptionHandler, error) {
	var err error
	once.amqSub.Do(func() {
		resources.amqSub, err = repository.NewAmqpSubscription()
	})
	return resources.amqSub, err
}
