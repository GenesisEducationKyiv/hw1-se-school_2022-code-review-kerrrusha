package application

import (
	"net/http"

	"github.com/kerrrusha/btc-api/api/domain"
	"github.com/kerrrusha/btc-api/api/presentation/presenters"
)

type PresenterRate interface {
	PresentRate(w http.ResponseWriter, rate *domain.Rate)
}

func CreateJsonRatePresenter() PresenterRate {
	var presenter PresenterRate

	presenter = presenters.CreateJsonRatePresenter()
	return presenter
}
