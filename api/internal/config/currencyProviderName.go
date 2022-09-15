package config

const (
	BINANCE = "BINANCE"
	COINAPI = "COINAPI"
)

func GetCurrencyProviderNameArray() []string {
	return []string{
		BINANCE,
		COINAPI,
	}
}
func GetDefaultCurrencyProviderName() string {
	return BINANCE
}
func CurrencyProviderNameExists(name string) bool {
	names := GetCurrencyProviderNameArray()
	for _, s := range names {
		if name == s {
			return true
		}
	}
	return false
}
