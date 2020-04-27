package loader

import (
	"encoding/json"
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/logger"
	"github.com/vagner-nascimento/go-poc-archref/src/integration"
	"github.com/vagner-nascimento/go-poc-archref/src/presentation"
	"os"
)

func LoadApplication() *chan error {
	loadConfiguration()

	errs := make(chan error)
	loadIntegration(&errs)
	loadPresentation(&errs)

	return &errs
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

func loadPresentation(errs *chan error) {
	logger.Info("loading http presentation asynchronously")
	go presentation.StartHttpPresentation(errs)
}

func loadIntegration(errsCh *chan error) {
	logger.Info("loading subscribers asynchronously")
	go func(errs *chan error) {
		if err := integration.SubscribeConsumers(); err != nil {
			logger.Error("error subscribe consumers", err)
			*errs <- err
		} else {
			logger.Info("consumers successfully subscribed")
		}
	}(errsCh)
}
