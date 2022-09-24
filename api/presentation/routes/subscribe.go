package routes

import (
	"github.com/kerrrusha/btc-api/api/presentation/typecasting"
	"log"
	"net/http"

	"github.com/kerrrusha/btc-api/api/application/service"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {
	log.Println("subscribe endpoint")

	s := service.SubscribeService{}
	newEmail := typecasting.CreateHttpCaster().RequestToEmail(r)
	s.SubscribeEmail(w, newEmail)
}
