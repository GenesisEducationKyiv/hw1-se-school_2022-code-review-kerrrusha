package service

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/kerrrusha/BTC-API/api/internal/config"
	"github.com/kerrrusha/BTC-API/api/internal/model"
	"github.com/kerrrusha/BTC-API/api/internal/requestUtils"
	"github.com/kerrrusha/BTC-API/api/internal/service/currencyResponse"
)

type currencyProvider struct {
	baseUrl  string
	response currencyResponse.CurrencyProviderResponse
}

func CreateCurrencyProvider(providerUrl string,
	rateResponse currencyResponse.CurrencyProviderResponse) *currencyProvider {
	return &currencyProvider{
		baseUrl:  providerUrl,
		response: rateResponse,
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

	jsonMap := make(map[string]json.RawMessage)
	e := json.Unmarshal(jsonBytes, &jsonMap)

	if e == nil {
		log.Println(jsonMap)
		log.Println("now values:")
		for key := range jsonMap {
			log.Println(string(jsonMap[key]))
		}
	} else {
		panic(e)
	}

	var rateResponse currencyResponse.BinanceResponse
	log.Println(rateResponse)
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

	return rateResponse.GetRate(), nil
}
