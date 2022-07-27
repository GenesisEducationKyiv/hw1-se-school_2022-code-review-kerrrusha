package service

import (
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "homePage endpoint")
}
