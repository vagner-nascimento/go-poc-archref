package main

import (
	"github.com/vagner-nascimento/go-poc-archref/infra"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

func loadSubscribers() {
	err := data.SubscribeConsumers()
	if err != nil {
		infra.LogError("error on load Customer subscriber", err)
	} else {
		infra.LogInfo("customer subscriber loaded")
	}
}
