package integration

import (
	"github.com/vagner-nascimento/go-poc-archref/src/infra/repository"
)

func SubscribeConsumers() error {
	return repository.SubscribeConsumers(getSubscriptions())
}

func getSubscriptions() (subs []repository.Subscription) {
	return append(subs,
		newUserSub(),
		newEnterpriseSub())
}
