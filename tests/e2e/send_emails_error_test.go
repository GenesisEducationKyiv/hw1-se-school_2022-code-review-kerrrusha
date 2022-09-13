package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const emailsErrorMsg = "An exception was thrown while sending emails"

func TestSendEmailsError(t *testing.T) {
	to := []string{
		"bdfj@gjfd.com",
		"fkdg_kfdgnk@kdf.net",
		"kjdfbgk123@company.ua",
	}
	subject := "test"
	body := "test"
	assert.NotPanics(t, func() { service.SendEmails(to, subject, body) })
}
