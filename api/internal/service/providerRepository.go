package service

import (
	"os"
	"sync"

	"github.com/kerrrusha/BTC-API/api/internal/config"
	"github.com/kerrrusha/BTC-API/api/internal/model"
	"github.com/kerrrusha/BTC-API/api/internal/service/currencyResponse"
)

type providerRepository struct {
	current *currencyProviderChain
}

var lock = &sync.Mutex{}

var repo *providerRepository

func GetProviderRepository() *providerRepository {
	if repo != nil {
		return repo
	}

	TryInitProviderRepositorySingleton()

	return repo
}
func getMainCurrencyProviderName() string {
	providerName, presented := os.LookupEnv("CRYPTO_CURRENCY_PROVIDER")
	if !presented || !config.CurrencyProviderNameExists(providerName) {
		return config.GetDefaultCurrencyProviderName()
	}
	return providerName
}
func TryInitProviderRepositorySingleton() {
	lock.Lock()
	defer lock.Unlock()
	if repo != nil {
		return
	}

	createProviderRepository()

	initialiseProviderRepository()
}
func createProviderRepository() {
	repo = &providerRepository{}
}
func initialiseProviderRepository() {
	cfg := config.Get()

	coinapiProvider := CreateCurrencyProvider(cfg.CoinApiUrl, &currencyResponse.CoinapiResponse{})
	coinapiChain := CreateCurrencyProviderChain(coinapiProvider)

	binanceProvider := CreateCurrencyProvider(cfg.BinanceApiUrl, &currencyResponse.BinanceResponse{})
	binanceChain := CreateCurrencyProviderChain(binanceProvider)

	mainProviderName := getMainCurrencyProviderName()
	if mainProviderName == config.BINANCE {
		repo.addCurrencyProviderChain(binanceChain)
		repo.addCurrencyProviderChain(coinapiChain)
	}
	if mainProviderName == config.COINAPI {
		repo.addCurrencyProviderChain(coinapiChain)
		repo.addCurrencyProviderChain(binanceChain)
	}
}

func (c *providerRepository) GetCurrencyProvider() (*currencyProvider, *model.CurrencyProviderChainAreOverError) {
	if c.current.IsEmpty() {
		return nil, model.CreateCurrencyProviderChainAreOverError("Current provider chain is empty.")
	}
	return c.current.GetCurrencyProvider(), nil
}
func (c *providerRepository) isEmpty() bool {
	return c.current == nil
}
func (c *providerRepository) getLastCurrencyProviderChain() *currencyProviderChain {
	if c.isEmpty() {
		return nil
	}
	lastChain := c.current
	for lastChain.next != nil {
		lastChain = lastChain.next
	}
	return lastChain
}
func (c *providerRepository) addCurrencyProviderChain(nextChain *currencyProviderChain) {
	if c.isEmpty() {
		c.current = nextChain
		return
	}
	c.getLastCurrencyProviderChain().next = nextChain
}
