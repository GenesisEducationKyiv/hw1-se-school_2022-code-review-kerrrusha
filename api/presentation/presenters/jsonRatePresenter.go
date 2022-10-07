package presenters

import (
	"encoding/json"
	"net/http"

	"github.com/kerrrusha/btc-api/api/domain"
)

const (
	EncodeResponseErrorMessage = "Failed to encode json response."
)

type jsonRatePresenter struct {
	*jsonPresenter
}

func CreateJsonRatePresenter() *jsonRatePresenter {
	return &jsonRatePresenter{}
}

func (p *jsonRatePresenter) PresentRate(w http.ResponseWriter, rate *domain.Rate) {
	p.setHeader(w)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(rate.GetValue())
	p.checkForEncodeError(err)
}
