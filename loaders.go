package main

import (
	"github.com/vagner-nascimento/go-poc-archref/src/infra"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/repository"
	"github.com/vagner-nascimento/go-poc-archref/src/presentation"
)

func loadHttpPresentation() {
	if err := presentation.StartHttpServer(); err != nil {
		infra.LogInfo("cannot load http presentation")
		return
	}

	infra.LogInfo("http presentation loaded")
}

func loadSubscribers() {
	repository.SubscribeAllConsumers()
}
