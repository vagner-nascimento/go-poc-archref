package loader

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra"
	"github.com/vagner-nascimento/go-poc-archref/src/integration"
	"github.com/vagner-nascimento/go-poc-archref/src/presentation"
	"os"
)

func LoadApplication() error {
	loadConfiguration()
	subSuccess := loadSubscribers()
	httpSuccess := loadHttpPresentation()
	if !subSuccess && !httpSuccess {
		return errors.New("cannot load application")
	}
	return nil
}

func loadConfiguration() {
	infra.LogInfo("loading configurations")
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "DEV"
	}
	if err := config.Load(env); err != nil {
		infra.LogInfo("cannot load configurations")
		panic(err)
	}

	conf, _ := json.Marshal(config.Get())
	fmt.Println("configurations loaded", string(conf))
}

func loadSubscribers() bool {
	infra.LogInfo("loading subscribers")
	if err := integration.SubscribeConsumers(); err != nil {
		infra.LogError("cannot subscribe consumers", err)
		return false
	}
	return true
}

func loadHttpPresentation() bool {
	infra.LogInfo("loading http presentation")
	if err := presentation.StartHttpServer(); err != nil {
		infra.LogError("cannot load http presentation", err)
		return false
	}
	return true
}
