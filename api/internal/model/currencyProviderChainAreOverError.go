package model

type CurrencyProviderChainAreOverError struct {
	*CustomError
}

func CreateCurrencyProviderChainAreOverError(message string) *CurrencyProviderChainAreOverError {
	return &CurrencyProviderChainAreOverError{&CustomError{errorMessage: message}}
}
