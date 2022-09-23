package utils

import (
	"encoding/json"

	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

func GetJsonStringValueByKey(jsonBytes []byte, key string) (string, *customErrors.JsonUnmarshalError) {
	const ErrorMessage = "Unsuccessful to unmarshal json."
	const InvalidReturnValue = ""

	jsonMap := make(map[string]json.RawMessage)
	e := json.Unmarshal(jsonBytes, &jsonMap)
	if e != nil {
		return InvalidReturnValue, customErrors.CreateJsonUnmarshalError(ErrorMessage)
	}

	return RemoveRedundantGaps(string(jsonMap[key])), nil
}
