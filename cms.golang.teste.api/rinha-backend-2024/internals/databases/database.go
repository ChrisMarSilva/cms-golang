package databases

import (
	"log"
	"strconv"
	"time"

	"github.com/chrismarsilva/rinha-backend-2024/internals/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	Conn *sqlx.DB // *sql.DB // *sqlx.DB
	// ConnWriter *sqlx.DB
	// ConnReader *sqlx.DB
)

// type IDatabaseWriter interface { Writer }
// type IDatabaseRead interface { Reader }

type IDatabase interface {
	StartDbConn()
	GetDatabaseConn() interface{}
	CloseDatabaseConn()

	// 	StartDbWriter()
	// 	GetDatabaseWriter() interface{}
	// 	CloseDatabaseWriter()

	// 	StartDbReader()
	// 	GetDatabaseReader() interface{}
	// 	CloseDatabaseReader()
}

type DatabasePostgres struct{}

func (DatabasePostgres) startDb(cfg *utils.Config) *sqlx.DB {
	database, err := sqlx.Connect(cfg.DbDriver, cfg.DbUri)
	if err != nil {
		log.Fatalf("Error connecting to database : error=%v", err)
	}

	if err = database.Ping(); err != nil {
		log.Println(err)
	} else {
		log.Println("Connected")
	}

	maxConnections, _ := strconv.Atoi(cfg.DbDriver)
	database.SetMaxOpenConns(maxConnections) // SetMaxOpenConns define o número máximo de conexões abertas com o banco de dados.
	database.SetMaxIdleConns(maxConnections) // SetMaxIdleConns define o número máximo de conexões no pool de conexão ociosa.
	database.SetConnMaxIdleTime(time.Minute * 1)
	database.SetConnMaxLifetime(time.Minute * 1) // SetConnMaxLifetime define a quantidade máxima de tempo que uma conexão pode ser reutilizada.

	return database
}

func (data *DatabasePostgres) StartDbConn(cfg *utils.Config) {
	Conn = data.startDb(cfg)
}

// func (data *DatabasePostgres) StartDbWriter(cfg *utils.Config) {
// 	ConnWriter = data.startDb(cfg)
// }

// func (data *DatabasePostgres) StartDbReader(cfg *utils.Config) {
// 	ConnReader = data.startDb(cfg)
// }

func (DatabasePostgres) GetDatabaseConn() *sqlx.DB {
	return Conn
}

// func (DatabasePostgres) GetDatabaseWriter() *sqlx.DB {
// 	return ConnWriter
// }

// func (DatabasePostgres) GetDatabaseReader() *sqlx.DB {
// 	return ConnReader
// }

func (DatabasePostgres) CloseDatabaseConn() {
	if Conn == nil {
		return
	}
	Conn.Close()
}

// func (DatabasePostgres) CloseDatabaseWriter() {
// 	if ConnWriter == nil {
// 		return
// 	}
// 	ConnWriter.Close()
// }

// func (DatabasePostgres) CloseDatabaseReader() {
// 	if ConnReader == nil {
// 		return
// 	}
// 	ConnReader.Close()
// }
