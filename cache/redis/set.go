package redis

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (c *cacheClient) SAdd(ctx context.Context, key string, value interface{}) (int64, error) {
	val, err := c.rdb.SAdd(ctx, key, value).Result()
	if err != nil {
		log.
			Error().
			Err(err).
			Str("key", key).
			Msg("unable to sadd key in the cache")
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) SAddAll(ctx context.Context, key string, values ...interface{}) (int64, error) {
	val, err := c.rdb.SAdd(ctx, key, values...).Result()
	if err != nil {
		log.
			Error().
			Err(err).
			Str("key", key).
			Msg("unable to saddAll key in the cache")
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) SRem(ctx context.Context, key string, value interface{}) (int64, error) {

	val, err := c.rdb.SRem(ctx, key, value).Result()
	if err != nil {
		log.
			Error().
			Err(err).
			Str("key", key).
			Msg("unable to SRem key in the cache")
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) SCard(ctx context.Context, key string) (int64, error) {
	val, err := c.rdb.SCard(ctx, key).Result()
	if err != nil {
		log.
			Error().
			Err(err).
			Str("key", key).
			Msg("unable to scard key in the cache")
		return 0, err
	}

	return val, nil
}

func (c *cacheClient) SIsMember(ctx context.Context, key string, value interface{}) (bool, error) {
	val, err := c.rdb.SIsMember(ctx, key, value).Result()
	if err != nil {
		log.
			Error().
			Err(err).
			Str("key", key).
			Msg("unable to SIsMember key in the cache")
		return false, err
	}

	return val, nil
}

func (c *cacheClient) SMembers(ctx context.Context, key string) ([]string, error) {
	values, err := c.rdb.SMembers(ctx, key).Result()
	if err != nil {
		log.
			Error().
			Err(err).
			Str("key", key).
			Msg("unable to SMembers key in the cache")
		return nil, err
	}

	return values, nil
}
