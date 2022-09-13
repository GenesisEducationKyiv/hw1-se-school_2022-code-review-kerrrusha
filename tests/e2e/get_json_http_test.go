package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const getErrorMsg = "JSON answer is not correct."

func TestGetJsonCoinapi(t *testing.T) {
	url := "https://rest.coinapi.io/"
	expected := `{
  "message": "This is the base URL of the API. Please read the https://docs.coinapi.io/ for more information on how to use the API."
}`
	actual := string(service.GetJson(url))
	assert.Equal(t, actual, expected, getErrorMsg)
}
