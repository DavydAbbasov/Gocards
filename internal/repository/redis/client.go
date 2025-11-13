package redis

import (
	"gocarts/internal/repository"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(cli *redis.Client) (*RedisStore, error) {
	if cli == nil {
		return nil, repository.RedisRequired
	}
	return &RedisStore{
		client: cli,
	}, nil
}
func (r *RedisStore) Close() error {
	if r == nil || r.client == nil {
		return nil
	}
	return r.client.Close()
}
