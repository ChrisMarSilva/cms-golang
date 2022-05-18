package database

// import (
// 	"context"
// 	"database/sql"
// 	"log"
// )

// type Dbtx interface {
// 	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
// 	PrepareContext(context.Context, string) (*sql.Stmt, error)
// 	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
// 	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
// }

// var dbSqlServerSQLC *Queries

// type DatabaseSQLServerSQLC struct {
// }

// func (DatabaseSQLServerSQLC) StartDB() {

// 	driverName := "mssql"
// 	dataSourceName := "sqlserver://sa:sa@127.0.0.1:5401?database=CMS_TESTE_JDSPB" // docker

// 	db, err := sql.Open(driverName, dataSourceName)
// 	if err != nil {
// 		log.Fatalln("Error creating connection pool: ", err.Error())
// 	}

// 	ctx := context.Background()

// 	err = db.PingContext(ctx)
// 	if err != nil {
// 		log.Fatalln("Error ping database: ", err.Error())
// 	}

// 	dbSqlServerSQLC = db
// }

// func (DatabaseSQLServerSQLC) GetDatabase() *Queries {
// 	return dbSqlServerSQLC
// }

// func (DatabaseSQLServerSQLC) Close() {
// 	// dbSqlServerSQLC.Close()
// }
