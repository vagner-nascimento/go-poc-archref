package app

import (
	"errors"
	"fmt"
	"strings"
)

func simpleError(msg string) error {
	return errors.New(msg)
}

func conversionError(originType string, destinyType string) error {
	msg := fmt.Sprintf("cannot convert %s into %s", originType, destinyType)
	return errors.New(msg)
}

func validationError(msgs []string) error {
	return errors.New(strings.Join(msgs, ","))
}

func notFoundError(entity string) error {
	return errors.New(fmt.Sprintf("%s not found", entity))
}
