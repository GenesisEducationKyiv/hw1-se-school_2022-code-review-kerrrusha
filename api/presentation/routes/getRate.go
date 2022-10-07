package routes

import (
	"github.com/kerrrusha/btc-api/logger"
	"net/http"

	"github.com/kerrrusha/btc-api/api/application/service"
)

func ProvideRate(w http.ResponseWriter, r *http.Request) {
	log := logger.CreateRabbitMQLogger()
	log.Debug("rate endpoint")

	s := &service.RateService{}
	s.ProvideRateJson(w)
}
