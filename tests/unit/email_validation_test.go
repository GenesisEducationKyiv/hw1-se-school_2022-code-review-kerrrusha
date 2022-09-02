package tests

import (
	"testing"

	"github.com/kerrrusha/BTC-API/service"
	"github.com/stretchr/testify/assert"
)

const validationErrorMsg = "Email validate invalid"

func TestEmailValidateCorrect(t *testing.T) {
	emails := []string{
		"kirill@gmail.com",
		"a@a.a",
		"a_a@ejbrjk.oacm",
		"22222@gfjd.g",
		"a2_@gf.c",
	}
	for _, email := range emails {
		actual := service.EmailIsValid(email)
		assert.Equal(t, true, actual, validationErrorMsg)
	}
}

func TestEmailValidateIncorrect(t *testing.T) {
	emails := []string{
		"kirill@@gmail.com",
		"dgdfg",
		"gfd@.",
		"@sdf.c",
		"",
		" ",
	}
	for _, email := range emails {
		actual := service.EmailIsValid(email)
		assert.Equal(t, false, actual, validationErrorMsg)
	}
}
