package app

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
)

var customerType reflect.Type
var once = sync.Once{}

func getCustomerType() reflect.Type {
	once.Do(func() {
		customerType = reflect.TypeOf(Customer{})
	})

	return customerType
}

func conversionError(originType string, destinyType string) error {
	msg := fmt.Sprintf("cannot convert %s into %s", originType, destinyType)
	return errors.New(msg)
}

func validationError(msgs []string) error {
	return errors.New(strings.Join(msgs, ","))
}

func customerNotFoundError() error {
	return errors.New(fmt.Sprintf("%s not found", getCustomerType().Name()))
}
