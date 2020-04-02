package repository

import (
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

// TODO: move it to DATA, repository shouldn't know mongo
func getBsonFilters(params []model.SearchParameter) bson.D {
	convertValue := func(val interface{}) (res interface{}) {
		if res, err := strconv.ParseInt(val.(string), 0, 64); err == nil {
			return res
		}
		if res, err := strconv.ParseFloat(val.(string), 64); err == nil {
			return res
		}
		if res, err := strconv.ParseBool(val.(string)); err == nil {
			return res
		}

		return val.(string)
	}
	convertValues := func(vals []interface{}) (res []interface{}) {
		for _, val := range vals {
			res = append(res, convertValue(val))
		}
		return res
	}

	getBsonD := func(param model.SearchParameter) bson.D {
		bsonD := bson.D{{}}
		if len(param.Values) > 0 {
			if len(param.Values) == 1 {
				bsonD = bson.D{{param.Field, convertValue(param.Values[0])}}
			} else {
				bsonD = bson.D{{param.Field, bson.M{"$in": convertValues(param.Values)}}}
			}
		}
		return bsonD
	}

	var filters bson.D
	switch len(params) {
	case 0:
		filters = bson.D{{}}
	case 1:
		filters = getBsonD(params[0])
	default: //
		var andFilter []bson.D
		for _, param := range params {
			andFilter = append(andFilter, getBsonD(param))
		}
		filters = bson.D{{"$and", andFilter}}
	}

	return filters
}

func unmarshalBsonData(data []byte, target interface{}) error {
	return bson.Unmarshal(data, target)
}
