package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kerrrusha/btc-api/api/internal/config"
	"github.com/kerrrusha/btc-api/api/internal/model"
	"github.com/kerrrusha/btc-api/api/internal/service"
	"github.com/kerrrusha/btc-api/api/internal/utils"
)

func Rate(w http.ResponseWriter, r *http.Request) {
	log.Println("rate endpoint")

	provider, requestFailure := service.GetProviderRepository().GetCurrencyProvider()
	if requestFailure != nil {
		utils.SendResponse(w, model.ErrorResponse{Error: requestFailure.GetMessage()}, http.StatusBadRequest)
		return
	}

	cfg := config.GetConfig()
	rate, err := provider.GetCurrencyRate(cfg.GetBaseCurrency(), cfg.GetQuoteCurrency())

	if err != nil {
		utils.SendResponse(w, model.ErrorResponse{Error: err.GetMessage()}, http.StatusBadRequest)
		return
	}

	response := model.RateValue{Rate: uint32(rate)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	utils.CheckForError(json.NewEncoder(w).Encode(response))
}
