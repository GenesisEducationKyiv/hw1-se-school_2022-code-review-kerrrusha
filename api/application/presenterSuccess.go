package application

import (
	"net/http"

	"github.com/kerrrusha/btc-api/api/presentation/presenters"
)

type PresenterSuccess interface {
	PresentSuccess(w http.ResponseWriter, message string)
}

func CreateJsonSuccessPresenter() PresenterSuccess {
	var presenter PresenterSuccess

	presenter = presenters.CreateJsonSuccessPresenter()
	return presenter
}
