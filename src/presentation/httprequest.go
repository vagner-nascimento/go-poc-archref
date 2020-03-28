package presentation

import (
	"encoding/json"
	"github.com/vagner-nascimento/go-poc-archref/src/model"
	"github.com/vagner-nascimento/go-poc-archref/src/tool"
	"net/url"
	"strings"
)

func getDataFromPath(path string, skip int) string {
	params := strings.Split(path, "/")
	return params[len(params)-skip]
}

func getPaginatedParamsFromQuery(values url.Values) (params []model.SearchParameter, page int64, pageSize int64, err error) {
	for key := range values {
		if key == "page" {
			if page, err = tool.SafeParseInt(values.Get(key)); err != nil {
				params = nil
				break
			}
		} else if key == "pageSize" {
			if pageSize, err = tool.SafeParseInt(values.Get(key)); err != nil {
				params = nil
				break
			}
		} else {
			param := values.Get(key)
			var paramValues []interface{}
			if tool.IsStringArray(param) {
				dec := json.NewDecoder(strings.NewReader(param))
				if err = dec.Decode(&paramValues); err != nil {
					params = nil
					break
				}
			} else {
				paramValues = append(paramValues, strings.Replace(param, "\"", "", -1))
			}

			params = append(params, model.SearchParameter{
				Field:  key,
				Values: paramValues,
			})
		}
	}

	return params, page, pageSize, err
}
