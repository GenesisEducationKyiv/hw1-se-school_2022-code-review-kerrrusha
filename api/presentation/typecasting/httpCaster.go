package typecasting

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kerrrusha/btc-api/api/domain"
)

type HttpCaster struct{}

func (c *HttpCaster) ToEmail(r *http.Request) domain.Email {
	decoder := json.NewDecoder(r.Body)

	var email domain.Email
	err := decoder.Decode(&email)
	if err != nil {
		log.Fatal(err)
	}

	return email
}
