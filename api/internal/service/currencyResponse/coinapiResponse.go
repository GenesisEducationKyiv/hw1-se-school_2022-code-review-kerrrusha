package currencyResponse

type CoinapiResponse struct {
	Time           string
	Asset_id_base  string
	Asset_id_quote string
	Rate           float64
}

func (c *CoinapiResponse) GetRate() float64 {
	return c.Rate
}

func CreateCoinapiResponse() CurrencyProviderResponse {
	return &CoinapiResponse{}
}
