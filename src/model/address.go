package model

import "encoding/json"

type Address struct {
	Street       string `json:"street" bson:"address"`
	Number       string `json:"number" bson:"number"`
	Neighborhood string `json:"neighborhood" bson:"neighborhood"`
	PostalCode   string `json:"postalCode" bson:"postalCode"`
	City         string `json:"city" bson:"city"`
	Country      string `json:"country" bson:"country"`
	State        string `json:"state" bson:"state"`
}

func NeeAddressFromJsonBytes(data []byte) (address Address, err error) {
	err = json.Unmarshal(data, &address)
	return address, err
}
