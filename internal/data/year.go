package data

import (
	"fmt"
	"strconv"
)

type Year int32

func (y Year) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d year", y)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}
