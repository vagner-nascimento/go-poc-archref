package dataamqp

import (
	"errors"

	"github.com/vagner-nascimento/go-poc-archref/app"
)

type customerPublisher struct {
}

func (cp *customerPublisher) Save(user *app.User) (app.User, error) {
	// TODO: User publisher send implementation
	return *user, nil
}

func (cp *customerPublisher) Get(id string) (app.User, error) {
	u := app.User{}
	return u, errors.New("Not implemented")
}

func (cp *customerPublisher) GetMany(params ...interface{}) ([]app.User, error) {
	u := []app.User{}
	return u, errors.New("Not implemented")
}

func (cp *customerPublisher) Update(user *app.User) (app.User, error) {
	return *user, errors.New("Not implemented")
}
