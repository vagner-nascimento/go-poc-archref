package environment

import (
	"os"
	"sync"
)

var (
	env          string
	envOnce      sync.Once
	httpPort     string
	httpPortOnce sync.Once
)

func GetEnv() string {
	envOnce.Do(func() {
		if env == "" {
			if env = os.Getenv("GO_ENV"); env == "" {
				env = "DEV"
			}
		}
	})

	return env
}

func GetHttpPort(prefix string) string {
	httpPortOnce.Do(func() {
		if httpPort == "" {
			httpPort = os.Getenv("HTTP_PORT")
			if httpPort == "" {
				httpPort = "3000"
			}
		}
	})

	return prefix + httpPort
}
