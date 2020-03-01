package repository

import (
	"encoding/json"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

const customerQueue = "q-customer"

func publishCustomer(c interface{}) error {
	uBytes, err := json.Marshal(c)
	if err != nil {
		return conversionError(err, "customer", "bytes array")
	}

	pub, err := data.NewAmqpPublisher(customerQueue)
	if err != nil {
		return operationError("publish", "customer")
	}

	err = pub.Publish(uBytes)
	if err != nil {
		return operationError("publish", "customer")
	}

	return nil
}
