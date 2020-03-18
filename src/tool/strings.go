package tool

import (
	"strconv"
	"strings"
)

func IsStringArray(str string) bool {
	return strings.HasPrefix(str, "[") && strings.HasSuffix(str, "]")
}

func SafeParseInt(str string) (res int64, err error) {
	if len(str) > 0 {
		res, err = strconv.ParseInt(str, 0, 64)
	}
	return res, err
}
