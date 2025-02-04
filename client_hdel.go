package kzkredis

import (
	"context"
)

func (c *defaultClient) HDel(ctx context.Context, key, field string) error {
	err := c.rc.HDel(ctx, key, field).Err()
	if err != nil {
		return NewRedisError(err)
	}

	return nil
}
