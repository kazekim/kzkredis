package kzkredis

import (
	"context"
	"time"
)

// SetNX set if not exist
func (c *defaultClient) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := c.rc.SetNX(ctx, key, value, expiration).Err()
	if err != nil {
		return NewRedisError(err)
	}
	return nil
}
