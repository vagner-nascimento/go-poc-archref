package model

import "encoding/json"

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	EMail      string `json:"eMail"`
	BirthYear  int    `json:"birthYear"`
	BirthDay   int    `json:"birthDay"`
	BirthMonth int    `json:"birthMonth"`
}

func NewUserFromJsonBytes(data []byte) (user User, err error) {
	err = json.Unmarshal(data, &user)
	return user, err
}
