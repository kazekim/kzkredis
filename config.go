package kzkredis

import "time"

type Config struct {
	Hostname                 string        `json:"hostname" mapstructure:"hostname"`
	Port                     string        `json:"port" mapstructure:"port"`
	Password                 string        `json:"password" mapstructure:"password"`
	DB                       *int          `json:"db" mapstructure:"db"`
	DefaultExpirationSeconds time.Duration `json:"default_expiration_seconds" mapstructure:"default_expiration_seconds"`
}
