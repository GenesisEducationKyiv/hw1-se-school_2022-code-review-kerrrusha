package service

import (
	"net/http"

	"github.com/kerrrusha/btc-api/api/application"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

type ResponseSender struct{}

func (e *FatalErrorHandler) sendSuccessResponse(w http.ResponseWriter, message string) {
	var presenter = application.CreateJsonSuccessPresenter()
	presenter.PresentSuccess(w, message)
}

func (e *FatalErrorHandler) sendErrorResponse(w http.ResponseWriter, err *customErrors.CustomError, statusCode int) {
	var presenter = application.CreateJsonErrorPresenter()
	presenter.PresentError(w, err.GetMessage(), statusCode)
}
