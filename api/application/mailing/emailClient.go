package mailing

import "github.com/kerrrusha/btc-api/api/internal/customErrors"

type EmailClient interface {
	SendEmails(to []string, subject string, body string) *customErrors.SendMailError
}
