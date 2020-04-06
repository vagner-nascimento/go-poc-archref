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

func LoadApplication() <-chan error {
	loadConfiguration()

	return multiplexErrorChannels(
		loadPresentation(),
		loadIntegration(),
	)
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

func loadPresentation() <-chan error {
	logger.Info("loading http presentation")
	return presentation.StartHttpPresentation()
}

func loadIntegration() <-chan error {
	logger.Info("loading subscribers")
	return integration.SubscribeConsumers()
}

// TODO: realise best place to put concurrency resources
func multiplexErrorChannels(errChannels ...<-chan error) <-chan error {
	outChan := make(chan error)
	for _, errCh := range errChannels {
		go forwardErrChannels(errCh, outChan)
	}

	return outChan
}

func forwardErrChannels(from <-chan error, to chan error) {
	for {
		to <- <-from
	}
}
