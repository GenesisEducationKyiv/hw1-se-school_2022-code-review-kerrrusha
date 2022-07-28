package service

import (
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("homePage endpoint"))
}
