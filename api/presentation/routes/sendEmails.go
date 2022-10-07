package routes

import (
	"github.com/kerrrusha/btc-api/logger"
	"net/http"

	"github.com/kerrrusha/btc-api/api/application/service"
)

func SendRateEmails(w http.ResponseWriter, r *http.Request) {
	log := logger.CreateRabbitMQLogger()
	log.Debug("sendEmails endpoint")

	ses := service.CreateSendEmailsService()
	ses.SendEmails(w)
}
