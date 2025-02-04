package kzkredis

import (
	"context"
	"encoding/json"
)

func (c *defaultClient) GetWithParse(ctx context.Context, key string, output interface{}) error {

	val, vErr := c.Get(ctx, key)
	if vErr != nil {
		return vErr
	}

	err := json.Unmarshal([]byte(val), output)
	if err != nil {
		return NewRedisError(err)
	}

	return nil
}
