package repository

import "github.com/vagner-nascimento/go-poc-archref/infra/data"

func SubscribeAllConsumers() error {
	var consumers []data.QueueConsumer

	consumers = append(consumers, newCustomerSub())
	err := data.SubscribeConsumers(consumers)
	// TODO: append other consumers to test multiple queue readers

	if err != nil {
		return err
	}

	return nil
}
