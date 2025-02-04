package kzkredis

import (
	"context"
)

func (c *defaultClient) HKeys(ctx context.Context, key string) ([]string, error) {
	vals, err := c.rc.HKeys(ctx, key).Result()
	if err != nil {
		return []string{}, NewRedisError(err)
	}

	return vals, nil
}
