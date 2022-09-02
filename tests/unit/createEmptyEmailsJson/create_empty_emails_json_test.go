package tests

import (
	"io/ioutil"
	"testing"

	"github.com/kerrrusha/BTC-API/fileManager"
	"github.com/stretchr/testify/assert"
)

const createErrorMsg = "File is not empty."

func TestCreateEmptyJson(t *testing.T) {
	filename := "emailsTest.json"
	fileManager.CreateEmptyEmailsJSON(filename)
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, `{"emails":[]}`, string(databyte), createErrorMsg)
}
