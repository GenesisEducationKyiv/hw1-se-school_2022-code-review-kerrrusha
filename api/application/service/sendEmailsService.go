package service

import (
	"fmt"
	"net/http"

	"github.com/kerrrusha/btc-api/api/application/mailing"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

type sendEmailsService struct {
	*FatalErrorHandler
	*ResponseSender

	client mailing.EmailClient
}

func CreateSendEmailsService() *sendEmailsService {
	return &sendEmailsService{client: mailing.CreateSmtpEmailClient()}
}

func (s *sendEmailsService) SendEmails(w http.ResponseWriter) {
	rs := &RateService{}
	rate, rateErr := rs.GetRate()
	s.ifErrorPanic(rateErr)

	ss := &SubscribeService{}
	emails, emailsErr := ss.GetSubscribedEmails()
	s.ifErrorPanic(emailsErr)

	if s.client == nil {
		panic("Mailing client not initialised.")
	}

	subject := "BTC/UAH"
	body := fmt.Sprintf("%d", rate.GetValue())

	for _, emailStr := range emails.ToList() {
		sendErr := s.client.SendEmails([]string{emailStr}, subject, body)
		if sendErr != nil {
			s.sendErrorResponse(
				w,
				customErrors.CreateCustomError("Email '"+emailStr+"' was not sent: "+sendErr.GetMessage()),
				http.StatusBadRequest,
			)
			return
		}
	}
	s.sendSuccessResponse(w, "Emails was sent successfully!")
}
