package presentation

import (
	"errors"
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

func simpleError(err error, msg string) error {
	infra.LogError(msg, err)
	return errors.New(msg)
}
