package tests

import (
	"testing"

	"github.com/kerrrusha/BTC-API/fileManager"
	"github.com/stretchr/testify/assert"
)

const readErrorMsg = "File contains is incorrect"

func TestReadEmailsFromCorrectNonEmptyJson(t *testing.T) {
	filename := "correctNonEmpty.json"
	emails := `{"emails":["test_email@gmail.com","james2394cahfd@eod.net","dgfb234894____fjkd______________________@a.a","a@a.a"]}`
	databyte := fileManager.ReadFile(filename)
	assert.Equal(t, emails, string(databyte), readErrorMsg)
}

func TestReadEmailsFromCorrectEmptyJson(t *testing.T) {
	filename := "correctNonEmpty.json"
	emails := `{"emails":["test_email@gmail.com","james2394cahfd@eod.net","dgfb234894____fjkd______________________@a.a","a@a.a"]}`
	databyte := fileManager.ReadFile(filename)
	assert.Equal(t, emails, string(databyte), readErrorMsg)
}

func TestReadEmailsError(t *testing.T) {
	filename := "sfbdjbsdkf.json"
	assert.Panics(t, func() { fileManager.ReadFile(filename) })
}
