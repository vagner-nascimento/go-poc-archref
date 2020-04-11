package provider

import (
	"github.com/vagner-nascimento/go-poc-archref/src/app"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/repository"
	"sync"
)

var resources struct {
	customerUseCase app.CustomerUseCase
	supplierUseCase app.SupplierUseCase
	amqSub          repository.AmqpSubscriptionHandler
}

var once struct {
	customerUc sync.Once
	amqSub     sync.Once
	providerUc sync.Once
}

func CustomerUseCase() (app.CustomerUseCase, error) {
	var err error
	once.customerUc.Do(func() {
		if repo, err := repository.NewCustomerRepository(); err == nil {
			resources.customerUseCase = app.NewCustomerUseCase(repo)
		}
	})
	return resources.customerUseCase, err
}

func SupplierUseCase() (app.SupplierUseCase, error) {
	var err error
	once.providerUc.Do(func() {
		if repo, err := repository.NewSupplierRepository(); err == nil {
			resources.supplierUseCase = app.NewSupplierUseCase(repo)
		}
	})
	return resources.supplierUseCase, err
}

func AmqpSubscription() (repository.AmqpSubscriptionHandler, error) {
	var err error
	once.amqSub.Do(func() {
		resources.amqSub, err = repository.NewAmqpSubscription()
	})

	return resources.amqSub, err
}
