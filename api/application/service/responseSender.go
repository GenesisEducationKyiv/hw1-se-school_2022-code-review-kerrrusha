package service

import (
	"net/http"

	"github.com/kerrrusha/btc-api/api/application"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

type ResponseSender struct{}

func (e *FatalErrorHandler) sendSuccessResponse(w http.ResponseWriter, message string) {
	var errPresenter application.PresenterSuccess
	errPresenter = application.CreateJsonSuccessPresenter()
	errPresenter.PresentSuccess(w, message)
}

func (e *FatalErrorHandler) sendErrorResponse(w http.ResponseWriter, err *customErrors.CustomError, statusCode int) {
	var errPresenter application.PresenterError
	errPresenter = application.CreateJsonErrorPresenter()
	errPresenter.PresentError(w, err.GetMessage(), statusCode)
}
