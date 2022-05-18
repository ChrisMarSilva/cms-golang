package database

import (
	"context"
	"fmt"
	"log"

	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

var dbSqlServerSQL *sql.DB

type DatabaseSQLServerSQL struct {
}

func (DatabaseSQLServerSQL) StartDB() {

	driverName := "mssql"
	dataSourceName := "sqlserver://sa:sa@127.0.0.1:5401?database=CMS_TESTE_JDSPB" // docker

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalln("Error creating connection pool: ", err.Error())
	}

	ctx := context.Background()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln("Error ping database: ", err.Error())
	}

	dbSqlServerSQL = db
}

func (DatabaseSQLServerSQL) GetDatabase() *sql.DB {
	return dbSqlServerSQL
}

func (DatabaseSQLServerSQL) Close() {
	dbSqlServerSQL.Close()
}

func (DatabaseSQLServerSQL) SelectVersion() {

	ctx := context.Background()

	err := dbSqlServerSQL.PingContext(ctx)
	if err != nil {
		log.Fatal("Error pinging database: " + err.Error())
	}

	var result string

	err = dbSqlServerSQL.QueryRowContext(ctx, "SELECT @@version").Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}

	fmt.Printf("%s\n", result)
}
