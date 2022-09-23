package service

import (
	"github.com/kerrrusha/btc-api/api/domain"
	"net/http"

	"github.com/kerrrusha/btc-api/api/application"
	"github.com/kerrrusha/btc-api/api/internal/config"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

type RateService struct {
	*FatalErrorHandler
}

func (r *RateService) ProvideRateJson(w http.ResponseWriter) {
	rate, err := r.GetRate()
	if err != nil {
		r.handleFatalError(w, err, http.StatusBadRequest)
	}

	var ratePresenter application.PresenterRate
	ratePresenter = application.CreateJsonRatePresenter()
	ratePresenter.PresentRate(w, rate)
}

func (r *RateService) GetRate() (*domain.Rate, *customErrors.CustomError) {
	var getter application.CurrencyGetter
	var emptyRepoErr *customErrors.CurrencyProviderChainAreOverError

	getter, emptyRepoErr = application.CreateCurrencyProviderFacadeGetter()
	if emptyRepoErr != nil {
		return nil, emptyRepoErr.CustomError
	}

	cfg := config.GetConfig()

	rate, reqFailErr := getter.GetCurrencyRate(cfg.GetBaseCurrency(), cfg.GetQuoteCurrency())
	if reqFailErr != nil {
		return nil, reqFailErr.CustomError
	}

	return rate, nil
}
