package migrations

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
)

func RunMigrations(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	upMigration, err := ioutil.ReadFile("./migrations/up.sql")
	if err != nil {
		return fmt.Errorf("failed to read up.sql: %w", err)
	}

	_, err = db.Exec(string(upMigration))
	if err != nil {
		return fmt.Errorf("failed to execute up.sql: %w", err)
	}

	return nil
}
