package application

import (
	"github.com/kerrrusha/btc-api/api/dataAccess/emailsSource"
	"github.com/kerrrusha/btc-api/api/domain"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

type EmailsAppender interface {
	Append(email domain.Email) *customErrors.CustomError
}

func CreateJsonEmailsAppender(jsonpath string) EmailsAppender {
	return emailsSource.CreateJsonEmailsAppender(jsonpath)
}
