package service

import (
	"encoding/json"
	"strings"

	"github.com/kerrrusha/BTC-API/api/internal/config"
	"github.com/kerrrusha/BTC-API/api/internal/model"
	"github.com/kerrrusha/BTC-API/api/internal/requestUtils"
)

type currencyProvider struct {
	baseUrl string
}

func CreateCurrencyProvider(providerUrl string) *currencyProvider {
	return &currencyProvider{
		baseUrl: providerUrl,
	}
}

func (provider *currencyProvider) GetCurrencyRate(baseCurrency string, quoteCurrency string) (int, *model.RequestFailureError) {
	requestUrl := provider.configureUrl(baseCurrency, quoteCurrency)
	jsonResponse := requestUtils.RequestJson(requestUrl)
	rate, err := provider.castResponse(jsonResponse)

	return int(rate), err
}

func (provider *currencyProvider) configureUrl(baseCurrency string, quoteCurrency string) string {
	result := provider.baseUrl
	cfg := config.Get()

	result = strings.ReplaceAll(result, cfg.BaseCurrencyMark, baseCurrency)
	result = strings.ReplaceAll(result, cfg.QuoteCurrencyMark, quoteCurrency)

	return result
}

func (provider *currencyProvider) castResponse(jsonBytes []byte) (float64, *model.RequestFailureError) {
	const INVALID_RETURN_VALUE = -1
	const CAST_ERROR_MESSAGE = "Unsuccessful to unmarshal json Response"
	const THIRD_PARTY_ERROR_MESSAGE = "Third-party side API caused error"

	var rateResponse model.RateResponse
	err := json.Unmarshal(jsonBytes, &rateResponse)
	if err != nil {
		return INVALID_RETURN_VALUE, model.CreateRequestFailureError(CAST_ERROR_MESSAGE)
	}

	var errorResponse model.ErrorResponse
	err = json.Unmarshal(jsonBytes, &errorResponse)
	if err != nil {
		return INVALID_RETURN_VALUE, model.CreateRequestFailureError(THIRD_PARTY_ERROR_MESSAGE)
	}

	if len(errorResponse.Error) > 0 {
		return INVALID_RETURN_VALUE, model.CreateRequestFailureError(errorResponse.Error)
	}

	return rateResponse.Rate, nil
}
