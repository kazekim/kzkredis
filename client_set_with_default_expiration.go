package kzkredis

import (
	"context"
	"time"
)

func (c *defaultClient) SetWithDefaultExpiration(ctx context.Context, key string, value interface{}) error {

	err := c.rc.Set(ctx, key, value, time.Second*c.cfg.DefaultExpirationSecondsDuration()).Err()
	if err != nil {
		return NewRedisError(err)
	}
	return nil
}
