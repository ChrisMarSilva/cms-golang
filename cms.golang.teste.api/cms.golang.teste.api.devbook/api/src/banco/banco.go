package banco

import (
	"database/sql"
	"time"

	//"os"

	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/config"
	_ "github.com/go-sql-driver/mysql"
	// sqldblogger "github.com/simukti/sqldb-logger"
	// "github.com/simukti/sqldb-logger/logadapter/zerologadapter"
	// "github.com/rs/zerolog"
)

func Conectar() (*sql.DB, error) {

	db, err := sql.Open("mysql", config.StringConexaoBanco)
	if err != nil {
		return nil, err
	}

	// loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	// db = sqldblogger.OpenDriver(config.StringConexaoBanco, &mysql.MySQLDriver{}, loggerAdapter)

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil

}
