package service

import (
	"net/http"

	"github.com/kerrrusha/btc-api/api/application"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
	"github.com/kerrrusha/btc-api/logger"
)

type FatalErrorHandler struct{}

func (e *FatalErrorHandler) handleFatalError(w http.ResponseWriter, err *customErrors.CustomError, statusCode int) {
	if err == nil {
		return
	}

	application.CreateJsonErrorPresenter().PresentError(w, err.GetMessage(), statusCode)
	log := logger.CreateRabbitMQLogger()
	log.Error(err.GetMessage())
}
func (e *FatalErrorHandler) ifErrorPanic(err *customErrors.CustomError) {
	if err != nil {
		panic(err)
	}
}
