package data

import (
	"errors"
)

func notImplementedError() error {
	return errors.New("function not implemented")
}
