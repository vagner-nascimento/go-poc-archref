package data

import (
	"github.com/streadway/amqp"
)

type (
	rabbitQueueInfo struct {
		Name         string
		Durable      bool
		DeleteUnused bool
		AutoDelete   bool
		Exclusive    bool
		NoWait       bool
		Args         amqp.Table
	}
	rabbitMessageInfo struct {
		Consumer  string
		AutoAct   bool
		Exclusive bool
		Local     bool
		NoWait    bool
		Exchange  string
		Mandatory bool
		Immediate bool
		Args      amqp.Table
	}
)
