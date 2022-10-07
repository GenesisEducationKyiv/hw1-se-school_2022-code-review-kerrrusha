package typecasting

import (
	"encoding/json"
	"github.com/kerrrusha/btc-api/logger"
	"net/http"

	"github.com/kerrrusha/btc-api/api/domain"
)

type HttpCaster struct{}

func (c *HttpCaster) RequestToEmail(r *http.Request) domain.Email {
	decoder := json.NewDecoder(r.Body)

	var email domain.Email
	err := decoder.Decode(&email)
	if err != nil {
		log := logger.CreateRabbitMQLogger()
		log.Error(err.Error())
	}

	return email
}

func CreateHttpCaster() *HttpCaster {
	return &HttpCaster{}
}
