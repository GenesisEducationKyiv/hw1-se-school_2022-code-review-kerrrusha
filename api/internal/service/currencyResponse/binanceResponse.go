package currencyResponse

type BinanceResponse struct {
	Symbol string
	Price  float64
}

func (b *BinanceResponse) GetRate() float64 {
	return b.Price
}

func CreateBinanceResponse() CurrencyProviderResponse {
	return &BinanceResponse{}
}
