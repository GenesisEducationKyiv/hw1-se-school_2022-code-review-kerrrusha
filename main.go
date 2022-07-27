package main

import (
	"log"
	"net/http"

	"github.com/kerrrusha/BTC-API/service"
)

func handleRequests() {
	http.HandleFunc("/", service.HomePage)
	http.HandleFunc("/rate/", service.Rate)
	http.HandleFunc("/subscribe/", service.SubscribeNewEmail)
	http.HandleFunc("/sendEmails/", service.SendBTCRateMails)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func main() {
	handleRequests()
}
