package kzkredis

import (
	"context"
	"time"
)

func (c *defaultClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {

	err := c.rc.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return NewRedisError(err)
	}
	return nil
}
