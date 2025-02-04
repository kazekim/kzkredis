package kzkredis

import (
	"context"
)

func (c *defaultClient) HExists(ctx context.Context, key, field string) (bool, error) {
	isExists, err := c.rc.HExists(ctx, key, field).Result()
	if err != nil {
		return false, NewRedisError(err)
	}

	return isExists, nil
}
