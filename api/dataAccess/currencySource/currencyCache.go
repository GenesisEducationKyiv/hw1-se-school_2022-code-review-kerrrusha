package currencySource

import (
	"sync"
	"time"

	"github.com/bluele/gcache"
	"github.com/kerrrusha/btc-api/api/domain"
	"github.com/kerrrusha/btc-api/api/internal/customErrors"
)

const (
	CacheSize           = 10
	CacheExpirationTime = 300 * time.Second
	InvalidResult       = -1
	DefaultKey          = 0
)

type currencyCache struct {
	rateCache gcache.Cache
}

var lockCacheCreation = &sync.Mutex{}

var cache *currencyCache

func GetCurrencyCache() *currencyCache {
	if cache != nil {
		return cache
	}

	tryCreateCurrencyCacheSingleton()

	return cache
}

func tryCreateCurrencyCacheSingleton() {
	lockCacheCreation.Lock()
	defer lockCacheCreation.Unlock()
	if cache != nil {
		return
	}

	createCurrencyCache()
}
func createCurrencyCache() {
	cache = &currencyCache{
		rateCache: gcache.New(CacheSize).Expiration(CacheExpirationTime).ARC().Build(),
	}
}

func (cache *currencyCache) Set(rate *domain.Rate) error {
	return cache.rateCache.SetWithExpire(DefaultKey, rate, CacheExpirationTime)
}

func (cache *currencyCache) IsEmpty() bool {
	return cache.rateCache.Len(false) == 0
}

func (cache *currencyCache) Get() (*domain.Rate, *customErrors.RateNotInCacheError) {
	ErrorMessage := "Rate not in cache."

	rate, err := cache.rateCache.Get(DefaultKey)
	if err != nil {
		return nil, customErrors.CreateRateNotInCacheError(ErrorMessage)
	}

	return rate.(*domain.Rate), nil
}

func (cache *currencyCache) Clear() {
	cache.rateCache.Remove(DefaultKey)
}
