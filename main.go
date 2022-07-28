package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kerrrusha/BTC-API/service"
)

func handleRequests() {
	PORT, presented := os.LookupEnv("PORT")
	if !presented {
		PORT = "8000"
	}
	log.Println("Started server at " + PORT + " port")

	http.HandleFunc("/rate/", service.Rate)
	http.HandleFunc("/subscribe/", service.SubscribeNewEmail)
	http.HandleFunc("/sendEmails/", service.SendBTCRateMails)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+PORT, nil))
}

func main() {
	handleRequests()
}
