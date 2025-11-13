package box

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"gocarts/internal/config"
	redisrepo "gocarts/internal/repository/redis"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type Env struct {
	Config      *config.Config
	RedisClient *redis.Client
	RedisStore  *redisrepo.RedisStore
	Postgres    *sql.DB
}

func New() (*Env, error) {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	postgre, err := provideDB(
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.Database.UserName,
			cfg.Database.Password,
			cfg.Database.Addr,
			cfg.Database.Port,
			cfg.Database.DBName,
			cfg.Database.SSLMODE,
		),
	)
	if err != nil {
		return nil, err
	}

	rdb, err := initRedisClient(cfg.Redis)
	if err != nil {
		return nil, err
	}

	store, err := redisrepo.NewRedisStore(rdb)
	if err != nil {
		return nil, err
	}
	return &Env{
		Config:      cfg,
		RedisStore:  store,
		RedisClient: rdb,
		Postgres:    postgre,
	}, nil
}
func initRedisClient(cfg config.RedisConfig) (*redis.Client, error) {
	if cfg.Addr == "" {
		return nil, errors.New("redis addr required")
	}
	opt := &redis.Options{
		Addr:        cfg.Addr,
		Password:    cfg.Password,
		DB:          cfg.DB,
		DialTimeout: cfg.DialTimeout,
	}
	if cfg.UseTLS {
		opt.TLSConfig = cfg.TLSConfig
	}

	rdb := redis.NewClient(opt)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping: %w", err)
	}

	return rdb, nil
}
func provideDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("can't open connection: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("can't ping database: %w", err)
	}
	return db, nil
}
func (e *Env) Close() error {
	var firstErr error
	if e.Postgres != nil {
		if err := e.Postgres.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	if e.RedisStore != nil {
		if err := e.RedisStore.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}

	return firstErr
}
