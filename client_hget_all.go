package kzkredis

import (
	"context"
)

func (c *defaultClient) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	val, err := c.rc.HGetAll(ctx, key).Result()
	if err != nil {
		return map[string]string{}, NewRedisError(err)
	}

	return val, nil
}
