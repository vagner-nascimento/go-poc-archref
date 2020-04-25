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
	conn            *amqp.Connection
	connect         sync.Once
	statusListeners []*chan bool
	lostForever     bool
}

var singletonRabbitConn rabbitConnection

func ListenToRabbitConnectionStatus(connStatus *chan bool) error {
	if singletonRabbitConn.lostForever {
		close(*connStatus)
		return errors.New("rabbit connection is closed forever")
	}

	singletonRabbitConn.statusListeners = append(singletonRabbitConn.statusListeners, connStatus)
	go func(chSt *chan bool) {
		for {
			if singletonRabbitConn.conn == nil || singletonRabbitConn.conn.IsClosed() {
				// TODO: make a func to do it receiving boolean value
				select {
				case *chSt <- false:
					continue
				default:
					return
				}
			} else {
				select {
				case *chSt <- true:
					continue
				default:
					return
				}
			}
		}
	}(connStatus)

	return nil
}

func getRabbitChannel() (*amqp.Channel, error) {
	var err error
	singletonRabbitConn.connect.Do(func() {
		err = connectIntoRabbit()
	})

	if err != nil {
		return nil, err
	} else if singletonRabbitConn.conn != nil && !singletonRabbitConn.conn.IsClosed() {
		return singletonRabbitConn.conn.Channel()
	} else {
		err = errors.New("rabbit connection is closed")
	}

	return nil, err
}

func connectIntoRabbit() (err error) {
	sleep := config.Get().Data.Amqp.ConnRetry.Sleep
	maxTries := 1

	if pMaxTries := config.Get().Data.Amqp.ConnRetry.MaxTries; pMaxTries != nil {
		maxTries = *pMaxTries
	}

	for currentTry := 1; currentTry <= maxTries; currentTry++ {
		if singletonRabbitConn.conn, err = amqp.Dial(config.Get().Data.Amqp.ConnStr); err != nil {
			if maxTries > 1 {
				fmt.Println(fmt.Sprintf("waiting %d seconds before try to reconnect %d of %d tries", sleep, currentTry, maxTries))
				time.Sleep(sleep * time.Second)
			}
		} else {
			errs := make(chan *amqp.Error)
			singletonRabbitConn.conn.NotifyClose(errs)
			go reconnectRabbitOnClose(errs)
			logger.Info("successfully connected into rabbit mq")
			break
		}
	}

	if err != nil {
		logger.Error("error on connect into rabbit mq", err)
		err = errors.New("an error occurred on try to connect into rabbit mq")

		go func(connection *rabbitConnection) {
			connection.lostForever = true
			if connection.statusListeners != nil {
				for _, listener := range connection.statusListeners {
					close(*listener)
				}
			}

			return
		}(&singletonRabbitConn)
	}

	return err
}

func reconnectRabbitOnClose(errs chan *amqp.Error) {
	var closeErr *amqp.Error
	for {
		if closeErr = <-errs; closeErr != nil {
			fmt.Println("rabbit mq connection was closed, error:", closeErr)
			fmt.Println("trying to reconnecting into rabbit mq server...")
			connectIntoRabbit()
		}
	}

	fmt.Println("error on reconnect into rabbit mq sever", errors.New(closeErr.Error()))
	close(errs)
}
