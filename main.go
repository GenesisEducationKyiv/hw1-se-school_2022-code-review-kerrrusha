package main

import (
	"net/http"
	"os"

	"github.com/kerrrusha/btc-api/api/presentation/routes"
	"github.com/kerrrusha/btc-api/logger"
)

func getPort() string {
	PORT, presented := os.LookupEnv("PORT")
	if !presented {
		PORT = "8000"
	}
	return PORT
}
func handleRequests() {
	PORT := getPort()
	log := logger.CreateRabbitMQLogger()
	log.Info("Started server at " + PORT + " port")

	http.HandleFunc("/rate/", routes.ProvideRate)
	http.HandleFunc("/subscribe/", routes.Subscribe)
	http.HandleFunc("/sendEmails/", routes.SendRateEmails)
	log.Error(http.ListenAndServe("0.0.0.0:"+PORT, nil).Error())
}

func main() {
	handleRequests()
}
