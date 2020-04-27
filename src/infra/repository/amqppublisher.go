package repository

import (
	"github.com/vagner-nascimento/go-poc-archref/src/infra/data"
	"github.com/vagner-nascimento/go-poc-archref/src/infra/logger"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
)

func publishMessage(entity model.BytesCaster, topic string) {
	err := data.PublishIntoRabbit(entity.GetBytes(), topic)
	if err != nil {
		logger.Error("error on publish data", err)
	}
}
