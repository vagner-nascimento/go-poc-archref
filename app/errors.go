package app

import (
	"errors"
	"fmt"
	"github.com/vagner-nascimento/go-poc-archref/infra"
)

func simpleError(msg string) error {
	return errors.New(msg)
}

func conversionError(err error, originType string, destinyType string) error {
	msg := fmt.Sprintf("cannot convert %s into %s", originType, destinyType)
	infra.LogError(msg, err)
	return errors.New(msg)
}

func requiredDataError(entity string, data string) error {
	return errors.New(fmt.Sprintf("%s.%s is required"))
}
