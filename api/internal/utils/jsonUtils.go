package utils

import (
	"encoding/json"
	"github.com/kerrrusha/btc-api/api/internal/errors"
)

func GetJsonStringValueByKey(jsonBytes []byte, key string) (string, *errors.JsonUnmarshalError) {
	const ERROR_MESSAGE = "Unsuccessful to unmarshal json."
	const INVALID_RETURN_VALUE = ""

	jsonMap := make(map[string]json.RawMessage)
	e := json.Unmarshal(jsonBytes, &jsonMap)
	if e != nil {
		return INVALID_RETURN_VALUE, errors.CreateJsonUnmarshalError(ERROR_MESSAGE)
	}

	return RemoveRedundantGaps(string(jsonMap[key])), nil
}
