package main

import (
	"github.com/vagner-nascimento/go-poc-archref/infra"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

func loadSubscribers() {
	err := dataamqp.SubscribeConsumers()
	if err != nil {
		infra.LogError("Error on load Customer subscriber", err)
	} else {
		infra.LogInfo("Customer subscriber loaded")
	}
}
