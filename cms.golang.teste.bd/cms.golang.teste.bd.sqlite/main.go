package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.bd.sqlite
// go get github.com/mattn/go-sqlite3
// go mod tidy

// go sqlite3 banco.db

// go run main.go

func main() {

	log.Println("sqlite3.ini")

	os.Remove("./banco.db")

	db, err := sql.Open("sqlite3", "./banco.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `create table foo (id integer not null primary key, name text)`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	sqlStmt = `delete from foo`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("Nome-%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}

	tx.Commit()

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("update foo set name=? where id=?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec("Nome-003.4", 3)
	if err != nil {
		log.Fatal(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(affect)

	log.Println("sqlite3.fim")
}
