package routes

import (
	"log"
	"net/http"

	"github.com/kerrrusha/btc-api/api/application/service"
)

func ProvideRate(w http.ResponseWriter, r *http.Request) {
	log.Println("rate endpoint")

	s := &service.RateService{}
	s.ProvideRateJson(w)
}
