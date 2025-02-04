package kzkredis

import (
	"context"
)

func (c *defaultClient) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rc.Get(ctx, key).Result()
	if err != nil {
		return "", NewRedisError(err)
	}

	return val, nil
}
