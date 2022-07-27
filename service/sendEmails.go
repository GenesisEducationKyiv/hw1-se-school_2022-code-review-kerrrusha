package service

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"github.com/kerrrusha/BTC-API/config"
	"github.com/kerrrusha/BTC-API/error"
	"github.com/kerrrusha/BTC-API/response"
)

func sendEmails(to []string, subject string, body string) {
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
	error.CheckForError(err)

	log.Println("Emails was sent successfully via SMTP '" + HOST + "' host.")
}

func SendBTCRateMails(w http.ResponseWriter, r *http.Request) {
	emails := readEmails(config.FILENAME)

	result, errorMsg := GetBitcoinPriceUAH()

	if len(errorMsg) > 0 {
		response.SendErrorResponse(w, errorMsg, http.StatusBadRequest)
		return
	}

	subject := "BTC/UAH"
	body := fmt.Sprintf("%d", result)

	for _, element := range emails.Emails {
		sendEmails([]string{element}, subject, body)
	}

	response.SendSuccessResponse(w, "Emails was sent successfully!", http.StatusOK)
}
