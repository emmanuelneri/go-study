package config

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type DBConfig struct {
	Host         string
	Port         string
	Name         string
	User         string
	Password     string
	Driver       string
	MaxOpenConns int
	MigrationDir string
}

func (d DBConfig) Url() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		d.Host, d.Port, d.User, d.Password, d.Name)
}

func StartDB(config DBConfig) *sql.DB {
	db, err := Connect(config)
	if err != nil {
		panic(err)
	}

	return db
}

func Connect(config DBConfig) (*sql.DB, error) {
	fmt.Println("DB starting... " + config.Url())

	db, err := sql.Open(config.Driver, config.Url())
	if err != nil {
		return nil, errors.Wrap(err, "connection fail")
	}

	db.SetMaxOpenConns(config.MaxOpenConns)

	if err = db.Ping(); err != nil {
		return nil, errors.Wrap(err, "ping fail")
	}

	if err = migration(config, db); err != nil {
		return nil, errors.Wrap(err, "migration fail")
	}

	return db, nil
}

func migration(config DBConfig, db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "driver start fail")
	}

	m, err := migrate.NewWithDatabaseInstance(
		config.MigrationDir,
		config.Driver, driver)

	if err != nil {
		return errors.Wrap(err, "migration start fail")
	}

	defer m.Close()
	if err = m.Up(); err != migrate.ErrNoChange {
		return err
	}

	return nil
}
