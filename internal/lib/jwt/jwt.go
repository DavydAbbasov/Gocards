package lib

import (
	"context"
	"time"
)

type Claims struct {
	ID        int64
	Login     string
	Email     string
	ExpiresAt time.Time
}

type Driver interface {
	GenerateAccessToken(ctx context.Context, id int64, login, email string) (string, error)
	GenerateRefreshToken(login string) (string, error)

	ExtractAccessToken(ctx context.Context, token string) (*Claims, error)
	ExtractRefreshToken(ctx context.Context, token string) (*Claims, error)

	CheckAccessTokenBlacklist(ctx context.Context, token string) (bool, error)
	CheckRefreshTokenBlacklist(ctx context.Context, token string) (bool, error)

	AddAccessTokenBlacklist(ctx context.Context, token string, ttl time.Duration) error
	AddRefreshTokenBlacklist(ctx context.Context, token string, ttl time.Duration) error
}
