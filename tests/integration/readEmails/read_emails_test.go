package tests

import (
	"testing"

	"github.com/kerrrusha/BTC-API/model"
	"github.com/kerrrusha/BTC-API/service"
	"github.com/stretchr/testify/assert"
)

const readErrorMsg = "Emails list is incorrect"

func TestReadEmailsFromEmptyJson(t *testing.T) {
	filename := "empty.json"
	expected := model.Emails{
		Emails: []string{},
	}
	actual := service.ReadEmails(filename)
	assert.Equal(t, expected, actual, readErrorMsg)
}

func TestReadEmailsFromTotalEmptyJson(t *testing.T) {
	filename := "totalEmpty.json"
	expected := model.Emails{
		Emails: []string{},
	}
	actual := service.ReadEmails(filename)
	assert.Equal(t, expected, actual, readErrorMsg)
}

func TestReadEmailsFromNotEmptyJson(t *testing.T) {
	filename := "notEmpty.json"
	expected := model.Emails{
		Emails: []string{
			"test_email@gmail.com",
			"james2394cahfd@eod.net",
			"dgfb234894____fjkd______________________@a.a",
			"a@a.a",
		},
	}
	actual := service.ReadEmails(filename)
	assert.Equal(t, expected, actual, readErrorMsg)
}

func TestReadEmailsFromBrokenJson(t *testing.T) {
	filename := "broken.json"
	assert.Panics(t, func() { service.ReadEmails(filename) })
}

func TestReadEmailsFromNotFoundJson(t *testing.T) {
	filename := "fshfsd.json"
	expected := model.Emails{
		Emails: []string{},
	}
	actual := service.ReadEmails(filename)
	assert.Equal(t, expected, actual, readErrorMsg)
}
