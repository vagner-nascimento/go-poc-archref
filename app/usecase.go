package app

import (
	"encoding/json"
	"fmt"
)

func addCustomer(c *Customer) error {
	c.CreditCardHash = "CardHash"

	err := c.save()
	if err != nil {
		return err
	}

	b, _ := json.Marshal(*c)
	fmt.Println("saved customer", string(b))

	return nil
}
