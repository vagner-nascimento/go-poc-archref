package loader

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/logger"
	"github.com/vagner-nascimento/go-poc-archref/src/integration"
	"github.com/vagner-nascimento/go-poc-archref/src/presentation"
	"os"
)

func LoadApplication() error {
	loadConfiguration()
	httpSuccess := loadPresentation()
	subSuccess := loadIntegration()
	if !subSuccess && !httpSuccess {
		return errors.New("cannot load application")
	}
	return nil
}

func loadConfiguration() {
	logger.Info("loading configurations")
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "DEV"
	}
	if err := config.Load(env); err != nil {
		logger.Error("cannot load configurations", err)
		panic(err)
	}
	conf, _ := json.Marshal(config.Get())
	logger.Info(fmt.Sprintf("configurations loaded %s", string(conf)))
}

func loadIntegration() bool {
	logger.Info("loading subscribers")
	if err := integration.SubscribeConsumers(); err != nil {
		logger.Error("cannot subscribe consumers", err)
		return false
	}
	return true
}

func loadPresentation() bool {
	logger.Info("loading http presentation")
	if err := presentation.StartHttpServer(); err != nil {
		logger.Error("cannot load http presentation", err)
		return false
	}
	return true
}
