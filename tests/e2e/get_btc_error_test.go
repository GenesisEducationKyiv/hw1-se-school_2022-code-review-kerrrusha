package e2e

import (
	"testing"

	"github.com/kerrrusha/BTC-API/api/internal/config"
	"github.com/kerrrusha/BTC-API/api/internal/service"
)

const getBtcErrorMsg = "Exception was thrown"

func TestGetRateError(t *testing.T) {
	cfg := config.Get()
	repo := service.GetRepository(cfg.CoinApiUrl)

	rate, err := repo.GetCurrencyRate(cfg.BaseCurrency, cfg.QuoteCurrency)
}
