package service

import (
	"log"
	"net/http"

	"github.com/kerrrusha/btc-api/api/application"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

type FatalErrorHandler struct{}

func (e *FatalErrorHandler) handleFatalError(w http.ResponseWriter, err *customErrors.CustomError, statusCode int) {
	if err == nil {
		return
	}

	var errPresenter application.PresenterError
	errPresenter = application.CreateJsonErrorPresenter()
	errPresenter.PresentError(w, err.GetMessage(), statusCode)
	log.Fatal(err.GetMessage())
}
func (e *FatalErrorHandler) ifErrorPanic(err *customErrors.CustomError) {
	if err != nil {
		panic(err)
	}
}
