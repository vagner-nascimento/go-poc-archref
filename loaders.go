package main

import (
	"github.com/vagner-nascimento/go-poc-archref/infra"
	"github.com/vagner-nascimento/go-poc-archref/infra/repository"
)

func loadSubscribers() {
	err := repository.SubscribeAllConsumers()
	if err != nil {
		infra.LogError("error on load Customer subscriber", err)
	} else {
		infra.LogInfo("customer subscriber loaded")
	}
}
