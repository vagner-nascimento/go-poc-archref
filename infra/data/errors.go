package data

import (
	"errors"
)

func NotImplementedError() error {
	return errors.New("function not implemented")
}

func ConnectionError(resource string) error {
	return errors.New("cannot connect into " + resource)
}

func ExecError(operation string, entity string) error {
	return errors.New("cannot " + operation + " " + entity + "'s data")
}

func Error(msg string) error {
	return errors.New(msg)
}
