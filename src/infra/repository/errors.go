package repository

import (
	"errors"
	"fmt"

	"github.com/vagner-nascimento/go-poc-archref/src/infra"
)

func notImplementedError(typeNma string) error {
	return errors.New("function not implemented on " + typeNma)
}

func conversionError(err error, originType string, destinyType string) error {
	msg := fmt.Sprintf("cannot convert %s into %s", originType, destinyType)
	infra.LogError(msg, err)
	return errors.New(msg)
}

func operationError(operation string, entity string) error {
	return errors.New(fmt.Sprintf("cannot %s %s", operation, entity))
}
