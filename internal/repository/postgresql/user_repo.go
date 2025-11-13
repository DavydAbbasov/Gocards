package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"gocarts/internal/repository"
)

func (d *DBStorage) CreareUser(ctx context.Context, login string, password string) error {
	q := `
	INSERT INTO users (login, password_hash)
	VALUES ($1, $2)
	RETURING id;
	`

	var id int

	err := d.db.QueryRowContext(ctx, q, login, password).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("login '%s' is already exists", login)
		}
		return fmt.Errorf("db error: %w", err)
	}
	return nil
}

func (d *DBStorage) GetUserPasswordHash(ctx context.Context, login string) (string, error) {
	q := `
	SELECT password_hash 
	FROM users
	WHERE login = $1;
	`
	var password_hash string

	err := d.db.QueryRowContext(ctx, q, login).Scan(&password_hash)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", repository.ErrUserNotFound
		}
		return "", fmt.Errorf("db error %w:", err)
	}
	return password_hash, nil

}
