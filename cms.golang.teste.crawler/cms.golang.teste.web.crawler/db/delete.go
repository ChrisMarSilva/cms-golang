package db

import (
	_ "github.com/mattn/go-sqlite3"
)

func Delete() error {

	db, err := getConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	
	_, err = db.Exec("delete from links")
	if err != nil {
		return err
	}

	return nil

}
