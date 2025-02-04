package kzkredis

import "time"

type RedisConfig struct {
	Hostname                 string `json:"redis_hostname" mapstructure:"redis_hostname" env:"REDIS_HOSTNAME"`
	Port                     string `json:"redis_port" mapstructure:"redis_port" env:"REDIS_PORT"`
	Password                 string `json:"redis_password" mapstructure:"redis_password" env:"REDIS_PASSWORD"`
	DB                       *int   `json:"redis_db" mapstructure:"redis_db" env:"REDIS_DB"`
	DefaultExpirationSeconds int64  `json:"redis_default_expiration_seconds" mapstructure:"redis_default_expiration_seconds" env:"REDIS_DEFAULT_EXPIRATION_SECONDS"`
}

func (c *RedisConfig) DefaultExpirationSecondsDuration() time.Duration {
	return time.Duration(c.DefaultExpirationSeconds) * time.Second
}
