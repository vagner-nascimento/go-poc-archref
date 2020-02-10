package data

import (
	"errors"
)

func notImplementedError() error {
	return errors.New("function not implemented")
}

func connectionError(resource string) error {
	return errors.New("cannot connect into " + resource)
}

func execError(operation string, entity string) error {
	return errors.New("cannot " + operation + " " + entity + "'s data")
}
