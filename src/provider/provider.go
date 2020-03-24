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

func CustomerUseCase() app.CustomerUseCase {
	once.customerUs.Do(func() {
		resources.customerUseCase = app.NewCustomerUseCase(repository.NewCustomerRepository())
	})
	return resources.customerUseCase
}

func AmqpSubscription() repository.AmqpSubscriptionHandler {
	once.amqSub.Do(func() {
		resources.amqSub = repository.NewAmqpSubscription()
	})
	return resources.amqSub
}
