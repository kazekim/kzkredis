package kzkredis

import (
	"context"
	"time"
)

type Client interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	SetWithDefaultExpiration(ctx context.Context, key string, value interface{}) error
	SetNXWithDefaultExpiration(ctx context.Context, key string, value interface{}) error

	Get(ctx context.Context, key string) (string, error)
	GetWithParse(ctx context.Context, key string, output interface{}) error

	Delete(ctx context.Context, key ...string) error

	HSet(ctx context.Context, key, field string, value interface{}) error
	HDel(ctx context.Context, key, field string) error
	HGet(ctx context.Context, key, field string) (string, error)
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	HExists(ctx context.Context, key, field string) (bool, error)
	HKeys(ctx context.Context, key string) ([]string, error)
}
