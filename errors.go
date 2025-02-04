package kzkredis

import (
	"errors"
	"strings"
)

func NewRedisError(err error) error {
	if strings.Contains(err.Error(), "redis: nil") {
		return errors.New("redis data not found")
	}
	return err
}
