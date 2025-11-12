package repository

import (
	"database/sql"
	"fmt"
)

func (d *DBrepo) GetUserPasswordHash(login string) (string, error) {
	q := `
	SELECT password_hash 
	FROM users
	WHERE login = $1;
	`
	var password_hash string

	err := d.db.QueryRow(q, login).Scan(&password_hash)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrUserNotFound
		}
		return "", fmt.Errorf("db error %w:", err)
	}
	return password_hash, nil

}
