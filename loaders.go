package main

import (
	"github.com/vagner-nascimento/go-poc-archref/infra/repository"
)

func loadSubscribers() {
	repository.SubscribeAllConsumers()
}
