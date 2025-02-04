package kzkredis

import (
	"context"
	"time"
)

// SetNX set if not exist
func (c *defaultClient) SetNXWithDefaultExpiration(ctx context.Context, key string, value interface{}) error {
	err := c.rc.SetNX(ctx, key, value, time.Second*c.cfg.DefaultExpirationSeconds).Err()
	if err != nil {
		return NewRedisError(err)
	}
	return nil
}
