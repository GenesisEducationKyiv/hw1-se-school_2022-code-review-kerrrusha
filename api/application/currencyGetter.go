package application

import (
	"github.com/kerrrusha/btc-api/api/dataAccess/currencySource"
	"github.com/kerrrusha/btc-api/api/domain"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

type CurrencyGetter interface {
	GetCurrencyRate(baseCurrency string, quoteCurrency string) (*domain.Rate, *customErrors.RequestFailureError)
}

func CreateCurrencyProviderFacadeGetter() (CurrencyGetter, *customErrors.CurrencyProviderChainAreOverError) {
	var getter CurrencyGetter
	var emptyRepoErr *customErrors.CurrencyProviderChainAreOverError

	getter, emptyRepoErr = currencySource.CreateCurrencyProviderFacade()
	return getter, emptyRepoErr
}
