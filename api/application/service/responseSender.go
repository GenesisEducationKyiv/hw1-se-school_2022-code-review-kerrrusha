package service

import (
	"net/http"

	"github.com/kerrrusha/btc-api/api/application"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
	"github.com/kerrrusha/btc-api/logger"
)

type ResponseSender struct{}

func (e *FatalErrorHandler) sendSuccessResponse(w http.ResponseWriter, message string) {
	var presenter = application.CreateJsonSuccessPresenter()
	presenter.PresentSuccess(w, message)
}

func (e *FatalErrorHandler) sendErrorResponse(w http.ResponseWriter, err *customErrors.CustomError, statusCode int) {
	log := logger.CreateRabbitMQLogger()
	log.Error(err.GetMessage())
	var presenter = application.CreateJsonErrorPresenter()
	presenter.PresentError(w, err.GetMessage(), statusCode)
}
