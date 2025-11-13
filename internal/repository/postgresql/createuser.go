package postgresql

import (
	"database/sql"
	"fmt"
)

func (d *DBrepo) CreareUser(login string, password string) error {
	q := `
	INSERT INTO users (login, password_hash)
	VALUES ($1, $2)
	RETURING id;
	`

	var id int

	err := d.db.QueryRow(q, login, password).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("login '%s' is already exists", login)
		}
		return fmt.Errorf("db error: %w", err)
	}
	return nil
}
