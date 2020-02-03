package app

import (
	"fmt"
)

func addCustomer(data []byte) {
	c := makeCustomer(data)
	fmt.Println(c)
}
