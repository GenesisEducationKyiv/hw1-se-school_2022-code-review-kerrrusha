package presenters

import (
	"log"
	"net/http"
)

type jsonPresenter struct{}

func (p *jsonPresenter) setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
func (p *jsonPresenter) checkForEncodeError(err error) {
	if err != nil {
		log.Fatal(EncodeResponseErrorMessage)
	}
}
