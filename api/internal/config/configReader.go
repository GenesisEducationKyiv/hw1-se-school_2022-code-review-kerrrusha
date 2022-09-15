package config

import (
	"encoding/json"
	"sync"

	"github.com/kerrrusha/BTC-API/api/model/dataStorage/fileStorage"
)

type config struct {
	Filepath, BaseCurrency, QuoteCurrency, ProjectName             string
	CoinApiUrl, BinanceApiUrl, BaseCurrencyMark, QuoteCurrencyMark string
}

var lock = &sync.Mutex{}

var cfg *config

func Get() *config {
	if cfg != nil {
		return cfg
	}

	TryInitConfigSingleton()

	return cfg
}

func TryInitConfigSingleton() {
	lock.Lock()
	defer lock.Unlock()
	if cfg == nil {
		cfg = createConfig()
	}
}

func createConfig() *config {
	fileReader := fileStorage.CreateFileReader()
	jsonBytes := 1
	jsonMap := make(map[string]json.RawMessage)
	return &config{
		Filepath:          "emails.json",
		ProjectName:       "BTC-API",
		BaseCurrency:      "BTC",
		BaseCurrencyMark:  BASE,
		QuoteCurrency:     "UAH",
		QuoteCurrencyMark: QUOTE,
		CoinApiUrl: "https://rest.coinapi.io/v1/exchangerate/" +
			BASE + "/" + QUOTE +
			"?apikey=735B916A-29E3-49D7-BB21-5142DF49DAAC",
		BinanceApiUrl: "https://api.binance.com/api/v3/ticker/price?symbol=" +
			BASE + QUOTE,
	}
}
