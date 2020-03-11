package presentation

import (
	"errors"
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/src/infra"
)

func simpleError(err error, msg string) error {
	infra.LogError(msg, err)
	return errors.New(msg)
}

func castError(err error, from string, to string) error {
	msg := fmt.Sprintf("erro on cast %s into %s", from, to)
	infra.LogError(msg, err)

	return errors.New(msg)
}
