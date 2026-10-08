package cache

import (
	goCache "github.com/patrickmn/go-cache"
	"github.com/vitordm/go-boilerplate-webapi/internal/shared/utils"
	"time"
)

func Build() *Cache {
	if utils.IsDev() {
		return goCache.New(goCache.NoExpiration, goCache.NoExpiration)
	}
	return goCache.New(3*time.Hour, 3*time.Hour)
}
