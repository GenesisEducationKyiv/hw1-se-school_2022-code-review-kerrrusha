package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kerrrusha/BTC-API/api/internal/config"
	"github.com/kerrrusha/BTC-API/api/internal/errorUtils"
	"github.com/kerrrusha/BTC-API/api/internal/model"
	"github.com/kerrrusha/BTC-API/api/internal/responseUtils"
	"github.com/kerrrusha/BTC-API/api/internal/service"
)

func Rate(w http.ResponseWriter, r *http.Request) {
	log.Println("rate endpoint")

	cfg := config.Get()
	repo := service.GetRepository(cfg.CoinApiUrl)

	rate, err := repo.GetCurrencyRate(cfg.BaseCurrency, cfg.QuoteCurrency)

	if err != nil {
		responseUtils.SendResponse(w, model.ErrorResponse{Error: err.GetMessage()}, http.StatusBadRequest)
		return
	}

	response := model.RateValue{Rate: uint32(rate)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	errorUtils.CheckForError(json.NewEncoder(w).Encode(response))
}
