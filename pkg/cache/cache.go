package cache

import (
	"time"

	go_cache "github.com/patrickmn/go-cache"
)

var cache *go_cache.Cache

const (
	defaultExpirationTime  = 2 * time.Minute
	defaultCleanupInterval = 5 * time.Minute
)

func init() {
	cache = go_cache.New(defaultExpirationTime, defaultCleanupInterval)
}

func Init(expirationTime, cleanupInterval time.Duration) {
	cache = go_cache.NewFrom(expirationTime, cleanupInterval, cache.Items())
}

func Get(k string) interface{} {
	v, found := cache.Get(k)
	if !found {
		return nil
	}

	return v
}

func Set(k string, v interface{}, d time.Duration) {
	cache.Set(k, v, d)
}

func SetDefault(k string, v interface{}) {
	cache.SetDefault(k, v)
}
