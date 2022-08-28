package response

import (
	"encoding/json"
	"net/http"

	"github.com/kerrrusha/BTC-API/error"
	"github.com/kerrrusha/BTC-API/model"
)

func SendErrorResponse(w http.ResponseWriter, msg string, code int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	errorAnswer := model.ErrorAnswer{Error: msg}
	err := json.NewEncoder(w).Encode(errorAnswer)
	error.CheckForError(err)
}

func SendSuccessResponse(w http.ResponseWriter, msg string, code int) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	successAnswer := model.SuccessAnswer{Success: msg}
	err := json.NewEncoder(w).Encode(successAnswer)
	error.CheckForError(err)
}
