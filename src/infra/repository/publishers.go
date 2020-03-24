package repository

import (
	"encoding/json"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/data"
)

const customerQueue = "q-customer"

func publishCustomer(c interface{}) error {
	uBytes, err := json.Marshal(c)
	if err != nil {
		return err
	}

	pub, err := data.NewRabbitPublisher(customerQueue)
	if err != nil {
		return err
	}

	err = pub.Publish(uBytes)
	if err != nil {
		return err
	}

	return nil
}
