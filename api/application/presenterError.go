package application

import (
	"net/http"

	"github.com/kerrrusha/btc-api/api/presentation/presenters"
)

type PresenterError interface {
	PresentError(w http.ResponseWriter, message string, statusCode int)
}

func CreateJsonErrorPresenter() PresenterError {
	return presenters.CreateJsonErrorPresenter()
}
