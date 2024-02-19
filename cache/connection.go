package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func GetConnection(cfg Config) redis.UniversalClient {
	return redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)},
		DB:    0,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			log.Info().Msg("connected to the cache")
			return nil
		},
		ContextTimeoutEnabled: true,
		MinIdleConns:          10,
	})
}
