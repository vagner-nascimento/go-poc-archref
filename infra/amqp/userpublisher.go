package amqp

import (
	"errors"

	"github.com/vagner-nascimento/go-poc-archref/app"
)

type userPublisher struct {
}

func (cp *userPublisher) Save(user *app.User) error {
	// TODO: User publisher send implementation
	return nil
}

func (cp *userPublisher) Get(id string) (app.User, error) {
	u := app.User{} // Usar NEW
	return u, errors.New("Not implemented")
}

func (cp *userPublisher) GetMany(params ...interface{}) ([]app.User, error) {
	u := []app.User{} //Usar NEW
	return u, errors.New("Not implemented")
}

func (cp *userPublisher) Update(user *app.User) error {
	return errors.New("Not implemented")
}
