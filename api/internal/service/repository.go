package service

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/kerrrusha/BTC-API/api/internal/config"
	"github.com/kerrrusha/BTC-API/api/internal/model"
	"github.com/kerrrusha/BTC-API/api/internal/requestUtils"
)

type repository struct {
	baseUrl string
}

var lock = &sync.Mutex{}

var repo *repository

func GetRepository(repoUrl string) *repository {
	if repo != nil {
		return repo
	}

	TryInitRepositorySingleton(repoUrl)

	return repo
}

func TryInitRepositorySingleton(repoUrl string) {
	lock.Lock()
	defer lock.Unlock()
	if repo == nil {
		repo = createRepository(repoUrl)
	}
}

func createRepository(repoUrl string) *repository {
	return &repository{
		baseUrl: repoUrl,
	}
}

func (repo *repository) GetCurrencyRate(baseCurrenct string, quoteCurrency string) (int, *model.RequestFailureError) {
	requestUrl := repo.configureUrl()
	jsonResponse := requestUtils.RequestJson(requestUrl)
	rate, err := repo.castResponse(jsonResponse)

	return int(rate), err
}

func (repo *repository) configureUrl() string {
	cfg := config.Get()
	result := repo.baseUrl

	result = strings.ReplaceAll(result, cfg.BaseCurrencyMark, cfg.BaseCurrency)
	result = strings.ReplaceAll(result, cfg.QuoteCurrencyMark, cfg.QuoteCurrency)

	return result
}

func (repo *repository) castResponse(jsonBytes []byte) (float64, *model.RequestFailureError) {
	const INVALID_RETURN_VALUE = -1
	const CAST_ERROR_MESSAGE = "Unsuccessful to unmarshal json Response"
	const THIRD_PARTY_ERROR_MESSAGE = "Third-party side API caused error"

	var rateResponse model.RateResponse
	err := json.Unmarshal(jsonBytes, &rateResponse)
	if err != nil {
		return INVALID_RETURN_VALUE, model.CreateRequestFailureError(CAST_ERROR_MESSAGE)
	}

	var errorResponse model.ErrorResponse
	err = json.Unmarshal(jsonBytes, &errorResponse)
	if err != nil {
		return INVALID_RETURN_VALUE, model.CreateRequestFailureError(THIRD_PARTY_ERROR_MESSAGE)
	}

	if len(errorResponse.Error) > 0 {
		return INVALID_RETURN_VALUE, model.CreateRequestFailureError(errorResponse.Error)
	}

	return rateResponse.Rate, nil
}
