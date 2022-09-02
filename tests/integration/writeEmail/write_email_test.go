package tests

import (
	"io/ioutil"
	"testing"

	"github.com/kerrrusha/BTC-API/fileManager"
	"github.com/kerrrusha/BTC-API/model"
	"github.com/kerrrusha/BTC-API/service"
	"github.com/stretchr/testify/assert"
)

const writeErrorMsg = "File contains is incorrect"

func TestWriteEmailToEmptyJson(t *testing.T) {
	filename := "emailsTest.json"
	emails := model.Emails{
		Emails: []string{
			"test_email@gmail.com",
			"james2394cahfd@eod.net",
			"dgfb234894____fjkd______________________@a.a",
			"a@a.a",
		},
	}
	for _, email := range emails.Emails {
		fileManager.CreateEmptyEmailsJSON(filename)
		service.WriteNewEmailToFile(filename, email)
		databyte, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		assert.Equal(t, `{"emails":["`+email+`"]}`, string(databyte), writeErrorMsg)
	}
}

func getSeparator(index int) string {
	if index == 0 {
		return ""
	}
	return ","
}

func addEmailToJson(json *string, email string, sep string) {
	containsEnd := `]}`
	*json = (*json)[:len((*json))-len(containsEnd)] +
		sep + `"` + email + `"` + containsEnd
}

func TestWriteEmailToNonEmptyJson(t *testing.T) {
	filename := "emailsTest.json"
	fileManager.CreateEmptyEmailsJSON(filename)
	emails := model.Emails{
		Emails: []string{
			"test_email@gmail.com",
			"james2394cahfd@eod.net",
			"dgfb234894____fjkd______________________@a.a",
			"a@a.a",
		},
	}

	contains := `{"emails":[]}`
	for index, email := range emails.Emails {
		service.WriteNewEmailToFile(filename, email)
		databyte, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		addEmailToJson(&contains, email, getSeparator(index))
		assert.Equal(t, contains, string(databyte), writeErrorMsg)
	}
}
