package application

import (
	"github.com/kerrrusha/btc-api/api/dataAccess/emailsSource"
	"github.com/kerrrusha/btc-api/api/domain"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

type EmailsGetter interface {
	GetEmails() (*domain.Emails, *customErrors.CustomError)
}

func CreateJsonEmailsGetter(jsonpath string) EmailsGetter {
	return emailsSource.CreateJsonEmailsReader(jsonpath)
}
