package data

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/logger"
	"strings"
	"sync"
	"time"
)

type (
	queueInfo struct {
		Name         string
		Durable      bool
		DeleteUnused bool
		AutoDelete   bool
		Exclusive    bool
		NoWait       bool
		Args         amqp.Table
	}
	messageInfo struct {
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
	pubDefaultValues struct {
		queue   queueInfo
		message messageInfo
	}
	rabbitAmqpHandler struct {
		subscribers []rabbitSubscriber
		pubValues   pubDefaultValues
	}
	rabbitSubscriber struct {
		queue   queueInfo
		message messageInfo
		handler func([]byte)
	}
)

type rabbitConnection struct {
	conn *amqp.Connection
}

func (rc *rabbitConnection) getChannel() (*amqp.Channel, error) {
	if rc.conn == nil || rc.conn.IsClosed() {
		return nil, errors.New("rabbit connection is closed")
	}

	return rc.conn.Channel()
}

var (
	rabbitConn     rabbitConnection
	onceRabbitConn sync.Once
)

func setRabbitConnection() (err error) {
	onceRabbitConn.Do(func() {
		errs := make(chan *amqp.Error)
		go connectIntoRabbit(errs)
		go reconnectRabbitOnErrors(errs)
	})

	return err // TODO REMOVE THIS ERR
}

// TODO: Realise how to subscribe consumers if amqp gets down
func reconnectRabbitOnErrors(errs chan *amqp.Error) {
	var closeErr *amqp.Error
	for {
		if closeErr = <-errs; closeErr != nil {
			logger.Error("an error occurred on rabbit mq connection", closeErr)
			logger.Info("trying to reconnecting into rabbit mq server...")
			connectIntoRabbit(errs)
		}
	}

	logger.Error("error on connect into rabbit mq sever", errors.New(closeErr.Error()))
	close(errs)
}

func connectIntoRabbit(errs chan *amqp.Error) {
	maxTries := config.Get().Data.Amqp.ConnRetry.MaxTries
	sleep := config.Get().Data.Amqp.ConnRetry.Sleep

	var err error
	for currentTry := -1; maxTries == nil || currentTry < *maxTries; currentTry++ {
		if rabbitConn.conn, err = amqp.Dial(config.Get().Data.Amqp.ConnStr); err != nil {
			if maxTries != nil {
				logger.Info(fmt.Sprintf("waiting %d seconds until try to reconnect %d of %d tries", sleep, currentTry+1, maxTries))
			} else {
				logger.Info(fmt.Sprintf("waiting %d seconds until try to reconnect %d try of infinite", sleep, currentTry+1))
			}

			time.Sleep(sleep * time.Second)
		} else {
			rabbitConn.conn.NotifyClose(errs)
			break
		}
	}

	if err != nil {
		logger.Info("cannot connect into amq server")
		close(errs)
	}
}

func subscribeRabbitConsumers(subscribers []rabbitSubscriber) error {
	ch, err := rabbitConn.getChannel()
	if err != nil {
		return err
	}

	var qNames []string
	for i := 0; i < len(subscribers); i = i + 1 {
		c := subscribers[i]

		processMsg, err := processMessages(ch, c)
		if err != nil {
			logger.Error(fmt.Sprintf("error on try subbscribe consumer %s", c.message.Consumer), err)
			continue
		}

		go processMsg()
		qNames = append(qNames, c.queue.Name)
	}

	if len(qNames) <= 0 {
		return errors.New("none queue to be listened")
	}

	logger.Info("listening to the queues: " + strings.Join(qNames, ","))
	return nil
}

// TODO if AMQP connection is lost, it keep connected when it reconnect?
func processMessages(ch *amqp.Channel, sub rabbitSubscriber) (func(), error) {
	q, err := ch.QueueDeclare(
		sub.queue.Name,
		sub.queue.Durable,
		sub.queue.DeleteUnused,
		sub.queue.Exclusive,
		sub.queue.NoWait,
		sub.queue.Args,
	)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name,
		sub.message.Consumer,
		sub.message.AutoAct,
		sub.message.Exclusive,
		sub.message.Local,
		sub.message.NoWait,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return func() {
		for msg := range msgs {
			logger.Info(fmt.Sprintf("message received from %s. body:\r\n %s", q.Name, string(msg.Body)))
			sub.handler(msg.Body)
		}
	}, nil
}

func (rh *rabbitAmqpHandler) AddSubscriber(topicName string, consumerName string, handler func([]byte)) error {
	if err := validateSub(topicName, consumerName, handler); err != nil {
		return err
	}

	rh.subscribers = append(rh.subscribers, newRabbitSubscriber(topicName, consumerName, handler))
	return nil
}

func (rh *rabbitAmqpHandler) SubscribeAll() (err error) {
	if len(rh.subscribers) > 0 {
		err = subscribeRabbitConsumers(rh.subscribers)
	} else {
		err = errors.New("there are no subscribers to consume topics")
	}

	return err
}

func (rh *rabbitAmqpHandler) Publish(data []byte, topicName string) (err error) {
	ch, err := rabbitConn.getChannel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		topicName,
		rh.pubValues.queue.Durable,
		rh.pubValues.queue.AutoDelete,
		rh.pubValues.queue.Exclusive,
		rh.pubValues.queue.NoWait,
		rh.pubValues.queue.Args,
	)
	if err == nil {
		err = ch.Publish(
			rh.pubValues.message.Exchange,
			q.Name,
			rh.pubValues.message.Mandatory,
			rh.pubValues.message.Immediate,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        data,
			},
		)
	}
	return err
}

func NewAmqpHandler() (AmqpHandler, error) {
	if err := setRabbitConnection(); err != nil {
		return nil, err
	}

	return &rabbitAmqpHandler{
		pubValues: pubDefaultValues{
			queue: queueInfo{
				Durable:    false,
				AutoDelete: false,
				Exclusive:  false,
				NoWait:     false,
				Args:       nil,
			},
			message: messageInfo{
				Exchange:  "",
				Mandatory: false,
				Immediate: false,
				Args:      nil,
			},
		},
	}, nil
}

func validateSub(topicName string, consumerName string, messageHandler func(data []byte)) error {
	if topicName == "" || consumerName == "" || messageHandler == nil {
		return errors.New(fmt.Sprintf("invalid subscriber's data: topic: %s, consumer: %s, handler is nil: %t",
			topicName,
			consumerName,
			messageHandler == nil))
	}
	return nil
}

func newRabbitSubscriber(queueName string, consumerName string, handler func([]byte)) rabbitSubscriber {
	return rabbitSubscriber{
		queue: queueInfo{
			Name:         queueName,
			Durable:      false,
			DeleteUnused: false,
			Exclusive:    false,
			NoWait:       false,
		},
		message: messageInfo{
			Consumer:  consumerName,
			AutoAct:   true,
			Exclusive: false,
			Local:     false,
			NoWait:    false,
		},
		handler: handler,
	}
}
