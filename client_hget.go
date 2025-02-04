package kzkredis

import (
	"context"
)

func (c *defaultClient) HGet(ctx context.Context, key, field string) (string, error) {
	val, err := c.rc.HGet(ctx, key, field).Result()
	if err != nil {
		return "", NewRedisError(err)
	}

	return val, nil
}
