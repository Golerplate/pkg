package redis

import (
	"github.com/golerplate/pkg/cache"
	"github.com/redis/go-redis/v9"
)

type cacheClient struct {
	rdb redis.UniversalClient
}

func NewDragonflyCache(rdb redis.UniversalClient) cache.Cache {
	return &cacheClient{rdb: rdb}
}
