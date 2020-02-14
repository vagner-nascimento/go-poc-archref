package environment

import (
	"os"
)

var env string

func GetEnv() string {
	if env == "" {
		env = os.Getenv("GO_ENV")
		if env == "" {
			env = "DEV"
		}
	}
	return env
}
