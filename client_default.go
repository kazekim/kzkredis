package kzkredis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type defaultClient struct {
	rc  *redis.Client
	cfg *RedisConfig
}

func NewClient(cfg *RedisConfig) Client {

	addr := fmt.Sprintf("%s:%s", cfg.Hostname, cfg.Port)

	options := &redis.Options{
		Addr:     addr,
		Password: cfg.Password,
	}

	if cfg.DB != nil {
		options.DB = *cfg.DB
	}
	rc := redis.NewClient(options)

	return &defaultClient{
		rc:  rc,
		cfg: cfg,
	}
}
