package main

import (
	"encoding/json"
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/repository"
	"github.com/vagner-nascimento/go-poc-archref/src/presentation"
	"os"
)

func loadConfiguration() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "DEV"
	}
	if err := config.Load(env); err != nil {
		infra.LogInfo("cannot load http presentation")
		panic(err)
	}
	conf, _ := json.Marshal(config.Get())
	fmt.Println("configuration loaded", string(conf))
}

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
