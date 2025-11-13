package models

import "time"

type UserCredentials struct {
	ID        int64
	Login     string
	Email     string
	Password  string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Jwt struct {
	AcssesToken  string
	RefreshToken string
	ExpiresAt    time.Time
}
