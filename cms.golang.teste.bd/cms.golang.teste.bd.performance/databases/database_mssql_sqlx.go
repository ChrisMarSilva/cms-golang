package database

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

var dbSqlServerSQLX *sqlx.DB

type DatabaseSQLServerSQLX struct {
}

func (DatabaseSQLServerSQLX) StartDB() {

	driverName := "mssql"
	dataSourceName := "sqlserver://sa:sa@127.0.0.1:5401?database=CMS_TESTE_JDSPB" // docker

	dbx, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	err = dbx.PingContext(ctx)
	if err != nil {
		log.Fatalln("Error ping database: ", err.Error())
	}

	dbSqlServerSQLX = dbx
}

func (DatabaseSQLServerSQLX) GetDatabase() *sqlx.DB {
	return dbSqlServerSQLX
}

func (DatabaseSQLServerSQLX) Close() {
	dbSqlServerSQLX.Close()
}
