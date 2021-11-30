package migration

import (
	"fmt"

	"github.com/anonychun/go-blog-api/internal/config"
	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func load() (*migrate.Migrate, error) {
	sourceURL := "file://migrations"
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Cfg().PostgresUser,
		config.Cfg().PostgresPassword,
		config.Cfg().PostgresHost,
		config.Cfg().PostgresPort,
		config.Cfg().PostgresDatabase,
	)
	return migrate.New(sourceURL, databaseURL)
}

func Up() error {
	m, err := load()
	if err != nil {
		return err
	}
	err = m.Up()
	return ignoreErrNoChange(err)
}

func Down() error {
	m, err := load()
	if err != nil {
		return err
	}
	err = m.Down()
	return ignoreErrNoChange(err)
}

func Steps(n int) error {
	m, err := load()
	if err != nil {
		return err
	}
	err = m.Steps(n)
	return ignoreErrNoChange(err)
}

func Drop() error {
	m, err := load()
	if err != nil {
		return err
	}
	err = m.Drop()
	return ignoreErrNoChange(err)
}

func ignoreErrNoChange(err error) error {
	if err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
