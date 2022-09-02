package tests

import (
	"testing"

	"github.com/kerrrusha/BTC-API/service"
	"github.com/stretchr/testify/assert"
)

const getBtcErrorMsg = "Exception was thrown"

func TestGetBtcError(t *testing.T) {
	assert.NotPanics(t, func() { service.GetBitcoinPriceUAH() })
}
