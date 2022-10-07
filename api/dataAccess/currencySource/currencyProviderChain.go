package currencySource

import (
	"github.com/kerrrusha/btc-api/api/domain"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

type currencyProviderChain struct {
	provider *currencyProvider
	next     *currencyProviderChain
}

func CreateCurrencyProviderChain(pr *currencyProvider) *currencyProviderChain {
	return &currencyProviderChain{provider: pr}
}

func (c *currencyProviderChain) GetCurrencyRate(baseCurrency string, quoteCurrency string) (*domain.Rate, *customErrors.RequestFailureError) {
	return c.provider.GetCurrencyRate(baseCurrency, quoteCurrency)
}
func (c *currencyProviderChain) GetCurrencyProvider() *currencyProvider {
	return c.provider
}

func (c *currencyProviderChain) SetNext(nextChain *currencyProviderChain) {
	c.next = nextChain
}

func (c *currencyProviderChain) IsEmpty() bool {
	return c.provider == nil
}
