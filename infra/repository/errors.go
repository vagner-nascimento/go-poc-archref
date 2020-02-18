package repository

import (
	"errors"
	"fmt"
)

func notImplementedError(typeNma string) error {
	return errors.New("function not implemented on " + typeNma)
}

func typeConversionError(originType string, destinyType string) error {
	return errors.New(fmt.Sprintf("cannot convert %s into %s", originType, destinyType))
}

func operationError(operation string, entity string) error {
	return errors.New(fmt.Sprintf("cannot %s %s", operation, entity))
}
