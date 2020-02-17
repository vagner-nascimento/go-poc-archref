package repository

import "errors"

func notImplementedError(typeNma string) error {
	return errors.New("function not implemented on " + typeNma)
}
