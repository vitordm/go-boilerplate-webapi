package helpers

import (
	"github.com/patrickmn/go-cache"
	coreCache "github.concur.com/I573758/example-golang-webapi/internal/core/cache"
	"github.concur.com/I573758/example-golang-webapi/internal/core/utils"
	"time"
)

func BuildCache() *coreCache.Cache {

	if utils.IsDev() {
		return cache.New(cache.NoExpiration, cache.NoExpiration)
	}

	return cache.New(3*time.Hour, 3*time.Hour)
}
