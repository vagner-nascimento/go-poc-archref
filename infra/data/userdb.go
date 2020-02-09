package data

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/vagner-nascimento/go-poc-archref/app"
)

// TODO: Implement userDb to call mongo connection
type userDb struct {
	db mongo
}

func (o *userDb) Save(c *app.User) error {
	c.Id = "fake_" + uuid.New().String()
	fmt.Println("(fake customerDd) customer save")
	// TODO publish user
	return nil
}

func (o *userDb) Get(id string) (app.User, error) {
	return app.User{}, notImplementedError()
}

func (o *userDb) GetMany(params ...interface{}) ([]app.User, error) {
	return []app.User{}, notImplementedError()
}

func (o *userDb) Update(c *app.User) error {
	return notImplementedError()
}
