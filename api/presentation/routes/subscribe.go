package routes

import (
	"log"
	"net/http"

	"github.com/kerrrusha/btc-api/api/application/service"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {
	log.Println("subscribe endpoint")

	s := service.SubscribeService{}
	s.SubscribeEmail(w, r)
}
