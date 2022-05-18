package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func getConnection() (*sql.DB, error) {

	// os.Remove("./banco.db")

	db, err := sql.Open("sqlite3", "./banco.db")
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	sqlStmt := `
		create table IF NOT EXISTS links (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, 
			website text NOT NULL,
			link text NOT NULL,
			data text NOT NULL
		)
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	return db, nil

}
