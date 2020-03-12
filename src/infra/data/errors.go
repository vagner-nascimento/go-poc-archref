package data

import (
	"errors"
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/src/infra"
)

func connectionError(err error, resource string) error {
	errMsg := fmt.Sprintf("cannot connect into %s", resource)
	infra.LogError(errMsg, err)

	return errors.New(errMsg)
}

func execError(err error, operation string, dataResource string) error {
	errMsg := fmt.Sprintf("cannot %s %s's data", operation, dataResource)
	infra.LogError(errMsg, err)

	return errors.New(errMsg)
}

func simpleError(msg string) error {
	return errors.New(msg)
}