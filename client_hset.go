package kzkredis

import (
	"context"
)

func (c *defaultClient) HSet(ctx context.Context, key, field string, value interface{}) error {
	err := c.rc.HSet(ctx, key, field, value).Err()
	if err != nil {
		return NewRedisError(err)
	}

	return nil
}
