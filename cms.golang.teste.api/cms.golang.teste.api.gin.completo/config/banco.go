package config

import (
	"database/sql"
	"time"

	//"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const dbDriver string = "mysql"
const dbHost string = "localhost"
const dbUser string = "root"
const dbPass string = "senha"
const dbName string = "database"
const dbSource string = dbUser + ":" + dbPass + "@tcp(" + dbHost + ":3306)/" + dbName

func GetMySQLBD() (db *sql.DB, err error) {

	errEnv := godotenv.Load()
	if errEnv != nil {
		return nil, errEnv
	}

	db, errBD := sql.Open(dbDriver, dbSource)
	if errBD != nil {
		return nil, errBD
	}

	// dbUserName := os.Getenv("DB_USERNAME")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbNAME := os.Getenv("DB_NAME")

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	errPing := db.Ping()
	if errPing != nil {
		// log.Printf("Error making connection to DB. Please check credentials. The error is: " + errPing.Error())
		return nil, errPing
	}

	// defer db.Close() // adiar o fechamento até depois que a função principal terminar

	return db, nil
}
