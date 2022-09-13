package rest

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"github.com/kerrrusha/BTC-API/api/internal/config"
	"github.com/kerrrusha/BTC-API/api/internal/errorUtils"
	"github.com/kerrrusha/BTC-API/api/internal/model"
	"github.com/kerrrusha/BTC-API/api/internal/responseUtils"
	"github.com/kerrrusha/BTC-API/api/internal/service"
)

func SendEmails(to []string, subject string, body string) {
	const (
		FROM     = "smtp8317@gmail.com"
		USERNAME = "smtp8317@gmail.com"
		PASSWORD = "khtihhqqywqrryan"
		PORT     = "587"
		HOST     = "smtp.gmail.com"
		ADDRESS  = HOST + ":" + PORT
	)

	if len(to) <= 0 {
		return
	}

	auth := smtp.PlainAuth("", USERNAME, PASSWORD, HOST)
	log.Println("SMTP was authorized successfully.")

	msg := []byte("From: " + FROM + "\r\n" +
		"To: you\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body + "\r\n")
	err := smtp.SendMail(ADDRESS, auth, FROM, to, msg)
	errorUtils.CheckForError(err)

	log.Println("Emails was sent successfully via SMTP '" + HOST + "' host.")
}

func SendRateEmails(w http.ResponseWriter, r *http.Request) {
	log.Println("sendEmails endpoint")

	provider, requestFailure := service.GetCurrencyRepository().GetCurrencyProvider()
	if requestFailure != nil {
		responseUtils.SendResponse(w, model.ErrorResponse{Error: requestFailure.GetMessage()}, http.StatusBadRequest)
		return
	}

	cfg := config.Get()
	rate, err := provider.GetCurrencyRate(cfg.BaseCurrency, cfg.QuoteCurrency)

	if err != nil {
		responseUtils.SendResponse(w, model.ErrorResponse{Error: err.GetMessage()}, http.StatusBadRequest)
		return
	}

	emails := ReadEmails(cfg.Filepath)

	subject := "BTC/UAH"
	body := fmt.Sprintf("%d", rate)

	for _, element := range emails.Emails {
		SendEmails([]string{element}, subject, body)
	}

	responseUtils.SendResponse(w, model.SuccessResponse{Success: "Emails was sent successfully!"}, http.StatusOK)
}
