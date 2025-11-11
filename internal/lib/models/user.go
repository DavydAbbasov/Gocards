package models

import "time"

type User struct {
	ID        int64
	Name      string
	State     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserUpdate struct {
	ID    int64
	Name  *string
	State *string
}
