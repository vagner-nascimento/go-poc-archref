package app

import (
	"encoding/json"
	"fmt"
)

func addCustomer(c *Customer) error {
	c.repository.Save(c)
	b, _ := json.Marshal(*c)
	fmt.Println("Saved customer", string(b))
	return c.repository.Save(c)
}
