package kzkredis

import (
	"context"
)

func (c *defaultClient) Delete(ctx context.Context, key ...string) error {

	_, err := c.rc.Del(ctx, key...).Result()
	if err != nil {
		return NewRedisError(err)
	}

	return nil
}
