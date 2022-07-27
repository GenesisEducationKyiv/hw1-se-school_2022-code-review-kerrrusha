package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/kerrrusha/BTC-API/config"
	"github.com/kerrrusha/BTC-API/error"
	"github.com/kerrrusha/BTC-API/model"
	"github.com/kerrrusha/BTC-API/response"
)

func getJson(url string) []byte {
	client := http.Client{Timeout: time.Second * 2}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	error.CheckForError(err)

	res, err := client.Do(req)
	error.CheckForError(err)

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	error.CheckForError(readErr)

	return body
}

func GetBitcoinPriceUAH() (int, string) {
	jsonAnswer := getJson(config.URL_GET_BTC)

	var rateAnswer model.RateAnswer
	err := json.Unmarshal(jsonAnswer, &rateAnswer)
	error.CheckForError(err)

	var errorAnswer model.ErrorAnswer
	err = json.Unmarshal(jsonAnswer, &errorAnswer)
	error.CheckForError(err)

	if len(errorAnswer.Error) > 0 {
		return -1, errorAnswer.Error
	}

	rateUAH := int(rateAnswer.Rate)

	return rateUAH, ""
}

func Rate(w http.ResponseWriter, r *http.Request) {
	log.Println("rate endpoint")

	result, errorMsg := GetBitcoinPriceUAH()

	if len(errorMsg) > 0 {
		response.SendErrorResponse(w, errorMsg, http.StatusBadRequest)
		return
	}

	rateUAH := model.RateValue{Rate: uint32(result)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rateUAH)
}
