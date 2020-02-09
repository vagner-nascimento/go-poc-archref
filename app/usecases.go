package app

import (
	"encoding/json"
	"fmt"
)

func addCustomer(c *Customer) error {
	c.CreditCardHash = "fake_22FB265865D30BE2E9362CFF01D5B95DCDCEC27D06DE4EEB042B67C0C72FE622"
	err := c.save()
	if err != nil {
		return err
	}

	b, _ := json.Marshal(*c)
	fmt.Println("saved customer", string(b))

	return nil
}
