package main

import (
	routes2 "github.com/kerrrusha/btc-api/api/presentation/routes"
	"log"
	"net/http"
	"os"
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
	log.Println("Started server at " + PORT + " port")

	http.HandleFunc("/rate/", routes2.ProvideRate)
	http.HandleFunc("/subscribe/", routes2.Subscribe)
	http.HandleFunc("/sendEmails/", routes2.SendRateEmails)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+PORT, nil))
}

func main() {
	handleRequests()
}
