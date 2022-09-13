package responseUtils

import (
	"encoding/json"
	"net/http"

	"github.com/kerrrusha/BTC-API/api/internal/errorUtils"
)

func SendResponse(w http.ResponseWriter, response interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(response)
	errorUtils.CheckForError(err)
}
