package data

import (
	"encoding/json"
	"github.com/vagner-nascimento/go-poc-archref/app"
)

type userRepository struct {
}

func (o *userRepository) Save(u *app.User) error {
	uBytes, _ := json.Marshal(u)
	uPub := newUserPub(uBytes)
	return publish(uPub)
}

func (o *userRepository) Update(c *app.User) error {
	return notImplementedError()
}

func (o *userRepository) Get(id string) (app.User, error) {
	return app.User{}, notImplementedError()
}

func (o *userRepository) GetMany(params ...interface{}) ([]app.User, error) {
	return []app.User{}, notImplementedError()
}
