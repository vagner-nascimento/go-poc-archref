package repository

import (
	"encoding/json"
	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

type userRepository struct {
}

func (o *userRepository) Save(u *app.User) error {
	uBytes, err := json.Marshal(u)
	if err != nil {
		return typeConversionError("user", "bytes array")
	}

	pub, err := data.NewAmqpPublisher("q-user")
	if err != nil {
		return operationError("save", "user")
	}

	err = pub.Publish(uBytes)
	if err != nil {
		return operationError("save", "user")
	}

	return nil
}

func (o *userRepository) Update(c *app.User) error {
	return notImplementedError("user repository")
}

func (o *userRepository) Get(id string) (app.User, error) {
	return app.User{}, notImplementedError("user repository")
}

func (o *userRepository) GetMany(params ...interface{}) ([]app.User, error) {
	return []app.User{}, notImplementedError("user repository")
}
