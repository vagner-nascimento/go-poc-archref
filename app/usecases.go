package app

import (
	"fmt"
	"math/rand"
	"strings"
)

func addCustomer(c *Customer) error {
	c.CreditCardHash = "fake_"
	for i := 0; i < 5; i = i + 1 {
		c.CreditCardHash += strings.Split(fmt.Sprintf("%f", rand.Float64()), ".")[1]
	}

	err := c.save()

	if err != nil {
		return err
	}

	return nil
}
