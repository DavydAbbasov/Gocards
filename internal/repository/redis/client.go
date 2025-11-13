package redis

import (
	"gocarts/internal/repository"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	cli *redis.Client
}

func NewRedisStore(cli *redis.Client) (*RedisStore, error) {
	if cli == nil {
		return nil, repository.RedisRequired
	}
	return &RedisStore{
		cli: cli,
	}, nil
}
