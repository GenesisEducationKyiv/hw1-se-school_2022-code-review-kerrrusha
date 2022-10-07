package currencySource

import (
	"github.com/kerrrusha/btc-api/api/domain"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
	"github.com/kerrrusha/btc-api/logger"
)

type CurrencyProviderFacade struct {
	provider *currencyProvider
	cache    *currencyCache
}

func CreateCurrencyProviderFacade() (*CurrencyProviderFacade, *customErrors.CurrencyProviderChainAreOverError) {
	provider, emptyRepoErr := GetProviderRepository().GetCurrencyProvider()
	if emptyRepoErr != nil {
		return nil, emptyRepoErr
	}

	cache := GetCurrencyCache()

	return &CurrencyProviderFacade{provider: provider, cache: cache}, nil
}

func (providerFacade *CurrencyProviderFacade) GetCurrencyRate(baseCurrency string, quoteCurrency string) (*domain.Rate, *customErrors.RequestFailureError) {
	cachedRate, absentErr := providerFacade.cache.Get()
	if absentErr == nil {
		return cachedRate, nil
	}

	jsonResponse := providerFacade.provider.RequestJson(baseCurrency, quoteCurrency)
	providerFacade.logProviderResponse(jsonResponse)

	rate, err := providerFacade.provider.castResponse(jsonResponse)
	if err != nil {
		return rate, err
	}

	cacheSetErr := providerFacade.cache.Set(rate)
	if cacheSetErr != nil {
		return nil, customErrors.CreateRequestFailureError(cacheSetErr.Error())
	}

	return rate, err
}
func (providerFacade *CurrencyProviderFacade) logProviderResponse(response []byte) {
	responseMsg := "Response from " + providerFacade.provider.GetDomain() + ": " + string(response)
	log := logger.CreateRabbitMQLogger()
	log.Info(responseMsg)
}
