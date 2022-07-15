package db

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
)

func NewDB() (*pg.DB, error) {
	// connect to db
	db := pg.Connect(&pg.Options{
		Addr:     "db:5432", // "localhost:5432"
		User:     "postgres",
		Password: "admin",
		// Database: "pg_migrations_example",
	})

	// run migrations
	collection := migrations.NewCollection()

	err := collection.DiscoverSQLMigrations("migrations")
	if err != nil {
		return nil, err
	}

	_, _, err = collection.Run(db, "init")
	if err != nil {
		return nil, err
	}

	oldVersion, newVersion, err := collection.Run(db, "up")
	if err != nil {
		return nil, err
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}

	// return the db connection
	return db, nil
}
