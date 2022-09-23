package service

import (
	"github.com/kerrrusha/btc-api/api/presentation/typecasting"
	"net/http"

	"github.com/kerrrusha/btc-api/api/application"
	"github.com/kerrrusha/btc-api/api/domain"
	"github.com/kerrrusha/btc-api/api/internal/config"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
	"github.com/kerrrusha/btc-api/api/internal/utils"
)

type SubscribeService struct {
	*FatalErrorHandler
	*ResponseSender
}

func (s *SubscribeService) SubscribeEmail(w http.ResponseWriter, r *http.Request) {
	var email domain.Email
	caster := typecasting.HttpCaster{}
	email = caster.ToEmail(r)

	emails, err := s.GetSubscribedEmails()
	s.handleFatalError(w, err, http.StatusBadRequest)

	if s.emailAlreadyExists(emails, email) {
		s.sendErrorResponse(
			w,
			customErrors.CreateCustomError("Email was not subscribed: it already exists"),
			http.StatusConflict,
		)
		return
	}
	if !email.IsValid() {
		s.sendErrorResponse(
			w,
			customErrors.CreateCustomError("Email is not correct. Please, enter valid email"),
			http.StatusConflict,
		)
		return
	}

	var appender application.EmailsAppender
	cfg := config.GetConfig()
	appender = application.CreateJsonEmailsAppender(cfg.GetEmailsFilepath())
	appender.Append(email)

	s.sendSuccessResponse(w, "Email was subscribed successfully")
}

func (s *SubscribeService) GetSubscribedEmails() (*domain.Emails, *customErrors.CustomError) {
	var getter application.EmailsGetter
	cfg := config.GetConfig()
	getter = application.CreateJsonEmailsGetter(cfg.GetEmailsFilepath())
	return getter.GetEmails()
}

func (s *SubscribeService) emailAlreadyExists(emails *domain.Emails, email domain.Email) bool {
	return utils.StringArraySearch(emails.ToList(), email.Email) != -1
}
