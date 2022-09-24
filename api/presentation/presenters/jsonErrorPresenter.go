package presenters

import (
	"encoding/json"
	"net/http"
)

type jsonErrorPresenter struct {
	*jsonPresenter
}

func CreateJsonErrorPresenter() *jsonErrorPresenter {
	return &jsonErrorPresenter{}
}

func (p *jsonErrorPresenter) PresentError(w http.ResponseWriter, message string, statusCode int) {
	p.setHeader(w)
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(message)
	p.checkForEncodeError(err)
}
