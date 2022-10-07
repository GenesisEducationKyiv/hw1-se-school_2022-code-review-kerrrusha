package routes

import (
	"net/http"

	"github.com/kerrrusha/btc-api/api/application/service"
	"github.com/kerrrusha/btc-api/api/presentation/typecasting"
	"github.com/kerrrusha/btc-api/logger"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {
	log := logger.CreateRabbitMQLogger()
	log.Debug("subscribe endpoint")

	s := service.SubscribeService{}
	newEmail := typecasting.CreateHttpCaster().RequestToEmail(r)
	s.SubscribeEmail(w, newEmail)
}
