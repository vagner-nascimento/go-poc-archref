package data

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/logger"
	"sync"
	"time"
)

type rabbitConnection struct {
	conn    *amqp.Connection
	connect sync.Once
	isAlive bool
}

func (rbConn *rabbitConnection) isConnected() bool {
	return singletonRabbitConn.conn != nil && !singletonRabbitConn.conn.IsClosed()
}

var singletonRabbitConn rabbitConnection

func ListenToRabbitConnected(connStatus *chan bool) {
	go func(chSt *chan bool) {
		for singletonRabbitConn.isAlive {
			*chSt <- singletonRabbitConn.isConnected()
		}
		close(*chSt)
	}(connStatus)
}

func getRabbitChannel() (*amqp.Channel, error) {
	var err error
	singletonRabbitConn.connect.Do(func() {
		singletonRabbitConn.isAlive = true
		err = connectIntoRabbit()
	})

	if err != nil {
		return nil, err
	} else if singletonRabbitConn.isConnected() {
		return singletonRabbitConn.conn.Channel()
	} else {
		err = errors.New("rabbit connection is closed")
	}

	return nil, err
}

func connectIntoRabbit() (err error) {
	sleep := config.Get().Data.Amqp.ConnRetry.Sleep
	maxTries := 1
	if config.Get().Data.Amqp.ConnRetry.MaxTries != nil {
		maxTries = *config.Get().Data.Amqp.ConnRetry.MaxTries
	}

	for currentTry := 1; currentTry <= maxTries; currentTry++ {
		if singletonRabbitConn.conn, err = amqp.Dial(config.Get().Data.Amqp.ConnStr); err != nil {
			if maxTries > 1 {
				logger.Info(fmt.Sprintf("waiting %d seconds before try to reconnect %d of %d tries", sleep, currentTry, maxTries))
				time.Sleep(sleep * time.Second)
			}
		} else {
			logger.Info("successfully connected into rabbit mq")

			errs := make(chan *amqp.Error)
			singletonRabbitConn.conn.NotifyClose(errs)
			go reconnectRabbitOnClose(errs)
			break
		}
	}

	if err != nil {
		logger.Error("error on connect into rabbit mq", err)

		singletonRabbitConn.isAlive = false
		err = errors.New("an error occurred on try to connect into rabbit mq")
	}

	return err
}

func reconnectRabbitOnClose(errs chan *amqp.Error) {
	for closeErr := range errs {
		if closeErr != nil {
			fmt.Println("rabbit mq connection was closed, error:", closeErr)
			fmt.Println("trying to reconnecting into rabbit mq server...")
			go connectIntoRabbit()
			return
		}
	}
}
