package app

import (
	"errors"
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

func simpleError(msg string) error {
	return errors.New(msg)
}

func operationError(err error, operation string) error {
	msg := fmt.Sprintf("cannot %s", operation)
	infra.LogError(msg, err)
	return errors.New(msg)
}
