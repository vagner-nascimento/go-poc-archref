package integration

import (
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"github.com/vagner-nascimento/go-poc-archref/src/provider"
)

type userSub struct {
	topic    string
	consumer string
	handler  func(data []byte)
}

func (es *userSub) getTopic() string {
	return es.topic
}

func (es *userSub) getConsumer() string {
	return es.consumer
}

func (es *userSub) getHandler() func([]byte) {
	return es.handler
}

func newUserSub() subscription {
	userConf := config.Get().Integration.Amqp.Subs.User
	return &userSub{
		topic:    userConf.Topic,
		consumer: userConf.Consumer,
		handler: func(data []byte) {
			if user, err := model.NewUserFromJsonBytes(data); err == nil {
				if customerUs, err := provider.CustomerUseCase(); err == nil {
					_, err = customerUs.UpdateFromUser(user)
				}
			} else {
				infra.LogError("error on update a customer", err)
			}
		},
	}
}
