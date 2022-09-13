package service

import (
	"sync"

	"github.com/kerrrusha/BTC-API/api/internal/config"
	"github.com/kerrrusha/BTC-API/api/internal/model"
)

type currencyRepository struct {
	current *currencyProviderChain
}

var lock = &sync.Mutex{}

var repo *currencyRepository

func GetCurrencyRepository() *currencyRepository {
	if repo != nil {
		return repo
	}

	TryInitCurrencyRepositorySingleton()

	return repo
}

func TryInitCurrencyRepositorySingleton() {
	lock.Lock()
	defer lock.Unlock()
	if repo != nil {
		return
	}

	cfg := config.Get()

	coinapiProvider := CreateCurrencyProvider(cfg.CoinApiUrl)
	startChain := CreateCurrencyProviderChain(coinapiProvider)

	repo = createCurrencyRepository(startChain)
}

func createCurrencyRepository(currentChain *currencyProviderChain) *currencyRepository {
	return &currencyRepository{current: currentChain}
}

func (c *currencyRepository) GetCurrencyProvider() (*currencyProvider, *model.CurrencyProviderChainAreOverError) {
	if c.current.IsLast() {
		return nil, model.CreateCurrencyProviderChainAreOverError("")
	}
	return c.current.GetCurrencyProvider(), nil
}
