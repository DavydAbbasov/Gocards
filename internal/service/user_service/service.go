package userservice

import (
	"context"
	"gocarts/internal/lib/models"
)

type Service interface {
	CreateUser(ctx context.Context, login string, password string) (*models.User, error)
	GetUser(ctx context.Context, id int64) (*models.User, error)
	
}
