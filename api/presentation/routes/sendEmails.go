package routes

import (
	"log"
	"net/http"

	"github.com/kerrrusha/btc-api/api/application/service"
)

func SendRateEmails(w http.ResponseWriter, r *http.Request) {
	log.Println("sendEmails endpoint")

	ses := service.CreateSendEmailsService()
	ses.SendEmails(w)
}
