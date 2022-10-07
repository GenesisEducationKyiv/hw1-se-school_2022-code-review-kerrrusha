package currencySource

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/kerrrusha/btc-api/api/domain"
	"github.com/kerrrusha/btc-api/api/internal/config"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
	"github.com/kerrrusha/btc-api/api/internal/utils"
	"github.com/kerrrusha/btc-api/api/presentation/model"
)

type currencyProvider struct {
	baseUrl string
	rateKey string
}

func CreateCurrencyProvider(providerUrl string, currencyRateKey string) *currencyProvider {
	return &currencyProvider{
		baseUrl: providerUrl,
		rateKey: currencyRateKey,
	}
}

func (provider *currencyProvider) RequestJson(baseCurrency string, quoteCurrency string) []byte {
	requestUrl := provider.configureUrl(baseCurrency, quoteCurrency)
	return utils.RequestJson(requestUrl)
}
func (provider *currencyProvider) GetCurrencyRate(baseCurrency string, quoteCurrency string) (*domain.Rate, *customErrors.RequestFailureError) {
	jsonResponse := provider.RequestJson(baseCurrency, quoteCurrency)
	rate, err := provider.castResponse(jsonResponse)

	return rate, err
}
func (provider *currencyProvider) GetDomain() string {
	return utils.SubstringBetween(provider.baseUrl, "https://", "/")
}

func (provider *currencyProvider) configureUrl(baseCurrency string, quoteCurrency string) string {
	result := provider.baseUrl
	cfg := config.GetConfig()

	result = strings.ReplaceAll(result, cfg.GetBaseCurrencyMark(), baseCurrency)
	result = strings.ReplaceAll(result, cfg.GetQuoteCurrencyMark(), quoteCurrency)

	return result
}

func (provider *currencyProvider) castResponse(jsonBytes []byte) (*domain.Rate, *customErrors.RequestFailureError) {
	const CastErrorMessage = "Unsuccessful to unmarshal json Response"
	const StringToFloatErrorMessage = "Unsuccessful parsing string to float64"
	const ThirdPartyErrorMessage = "Third-party side API caused error"

	rateStr, unmarshalErr := utils.GetJsonStringValueByKey(jsonBytes, provider.rateKey)
	if unmarshalErr != nil {
		return nil, customErrors.CreateRequestFailureError(CastErrorMessage)
	}

	rateFloat, ParseFloatErr := strconv.ParseFloat(rateStr, 64)
	if ParseFloatErr != nil {
		return nil, customErrors.CreateRequestFailureError(StringToFloatErrorMessage)
	}

	var errorResponse model.ErrorResponse
	err := json.Unmarshal(jsonBytes, &errorResponse)
	if err != nil {
		return nil, customErrors.CreateRequestFailureError(ThirdPartyErrorMessage)
	}

	if len(errorResponse.Error) > 0 {
		return nil, customErrors.CreateRequestFailureError(errorResponse.Error)
	}

	return domain.CreateRate(rateFloat), nil
}
