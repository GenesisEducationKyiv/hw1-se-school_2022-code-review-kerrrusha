package presenters

import (
	"encoding/json"
	"net/http"
)

type jsonSuccessPresenter struct {
	*jsonPresenter
}

func CreateJsonSuccessPresenter() *jsonSuccessPresenter {
	return &jsonSuccessPresenter{}
}

func (p *jsonSuccessPresenter) PresentSuccess(w http.ResponseWriter, message string) {
	p.setHeader(w)
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(message)
	p.checkForEncodeError(err)
}
