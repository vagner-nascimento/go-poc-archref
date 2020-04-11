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

func LoadApplication(errs chan error) {
	loadConfiguration()

	if err := loadIntegration(); err != nil {
		errs <- err
	}

	loadPresentation(errs)
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

func loadPresentation(errs chan error) {
	go presentation.StartHttpPresentation(errs)
}

func loadIntegration() error {
	logger.Info("loading subscribers")
	return integration.SubscribeConsumers()
}
