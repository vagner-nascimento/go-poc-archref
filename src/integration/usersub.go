package integration

import (
	"github.com/vagner-nascimento/go-poc-archref/config"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/logger"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/repository"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"github.com/vagner-nascimento/go-poc-archref/src/provider"
)

type userSub struct {
	topic    string
	consumer string
	handler  func(data []byte)
}

func (es *userSub) GetTopic() string {
	return es.topic
}

func (es *userSub) GetConsumer() string {
	return es.consumer
}

func (es *userSub) GetHandler() func([]byte) {
	return es.handler
}

func newUserSub() repository.Subscription {
	userConf := config.Get().Integration.Amqp.Subs.User
	return &userSub{
		topic:    userConf.Topic,
		consumer: userConf.Consumer,
		handler: func(data []byte) {
			var err error
			var user model.User
			if user, err = model.NewUserFromJsonBytes(data); err == nil {
				if customerUc, err := provider.CustomerUseCase(); err == nil {
					go customerUc.UpdateFromUser(user)
				}
			}
			if err != nil {
				logger.Error("error on update a customer", err)
			}
		},
	}
}
