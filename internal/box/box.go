package box

import (
	"gocarts/internal/config"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type Env struct {
	Config      *config.Config
	MySQLRead   *sqlx.DB
	MySQLWrite  *sqlx.DB
	RedisClient *redis.Client
}

func New() (*Env, error) {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	return &Env{
		Config: cfg,
	}, nil
}

func (e *Env) Close() error {
	var firstErr error

	return firstErr
}
