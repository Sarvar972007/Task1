package repo

import (
	"database/sql"
	"embed"
	"github.com/pressly/goose/v3"
)

var embedMigrations embed.FS

func GoMigrationUp(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
	return nil
}
func GoMigrationDown(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Down(db, "migrations"); err != nil {
		panic(err)
	}
	return nil
}
