package repository

import (
	"encoding/json"
	"github.com/vagner-nascimento/go-poc-archref/app"
	"github.com/vagner-nascimento/go-poc-archref/infra/data"
)

type userRepository struct {
}

func (o *userRepository) Save(u *app.User) error {
	uBytes, _ := json.Marshal(u)
	uPub := newUserPub(uBytes)
	return data.PublishMessage(uPub)
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
