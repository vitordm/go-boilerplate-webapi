package helpers

import (
	"time"

	"github.com/patrickmn/go-cache"
	coreCache "github.com/vitordm/go-boilerplate-webapi/internal/core/cache"
	"github.com/vitordm/go-boilerplate-webapi/internal/core/utils"
)

func BuildCache() *coreCache.Cache {

	if utils.IsDev() {
		return cache.New(cache.NoExpiration, cache.NoExpiration)
	}

	return cache.New(3*time.Hour, 3*time.Hour)
}
