package main

import (
	"database/sql"
	"fmt"
	"runtime"

	_ "github.com/godror/godror"
)

// go mod init github.com/chrismarsilva/cms-golang-teste-bd-oracle
// go mod tidy

// go get github.com/godror/godror

// go run main.go
// go build

func main() {

	//db, err := sql.Open("godror", `user="CRIS_JDSPB" password="CRIS_JDSPB" connectString="JDDSVDB"`)
	db, err := sql.Open("godror", "CRIS_JDSPB/CRIS_JDSPB@JDDSVDB")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select sysdate from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {
		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)

	runtime.Goexit()

	fmt.Print("ok")
}
