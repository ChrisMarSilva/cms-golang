package db

import (
	"time"
l"
	_ "github.com/mattn/go-sqite3l"
	_ "github.com/mattn/go-sqite3
)

func Insert(data interface{}) error {

	db, err := getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("insert into links(website, link, data) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	switch t := data.(type) {
    case []string:
        for _, v := range t {
			_, err = stmt.Exec(v, v, time.Now().Format("2006-01-02 15:04:05.000")) // 20060102150405
			if err != nil {
				return err
			}
        }
    case string:
        _, err = stmt.Exec(t, t, time.Now().Format("2006-01-02 15:04:05.000")) // 20060102150405
		if err != nil {
			return err
		}
    case model.VisitedLink:
        _, err = stmt.Exec(t.WebSite, t.Link, t.VisitedDate)
		if err != nil {
			return err
		}
    }

	tx.Commit()

	return nil

}
