package model

import "encoding/json"

type Enterprise struct {
	CompanyName string `json:"companyName" bson:"companyName"`
	Document    string `json:"document" bson:"document"`
	Active      bool   `json:"active" bson:"active"`
}

func NewEnterpriseFromJsonBytes(data []byte) (ent Enterprise, err error) {
	err = json.Unmarshal(data, &ent)
	return ent, err
}
