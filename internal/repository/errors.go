package repository

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	RedisRequired   = errors.New("redis client required")
)
