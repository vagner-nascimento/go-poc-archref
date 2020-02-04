package loader

import (
	"github.com/vagner-nascimento/go-poc-archref/infra"
	dataamqp "github.com/vagner-nascimento/go-poc-archref/infra/data/amqp"
)

func LoadSubscribers() {
	err := dataamqp.SubscribeConsumers()
	if err != nil {
		infra.LogError("Error on load Customer subscriber", err)
	} else {
		infra.LogInfo("Customer subscriber loaded")
	}
}
