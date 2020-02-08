package amqp

import (
	"fmt"
	"sync"

	"github.com/streadway/amqp"

	"github.com/vagner-nascimento/go-poc-archref/infra"
)

var once sync.Once

type connSingleton struct {
	AmqpConn *amqp.Connection
}

var (
	amqpUrl   = "localhost"
	amqpPort  = "5672"
	amqpUser  = "guest"
	amqpPass  = "guest"
	connError error
	singleton connSingleton
)

func getConnection() (*amqp.Connection, error) {
	once.Do(func() {
		singleton.AmqpConn, connError = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", amqpUser, amqpPass, amqpUrl, amqpPort))
		if connError == nil {
			infra.LogInfo("Successfully connected in AMQP server")
		}
	})

	return singleton.AmqpConn, connError
}
