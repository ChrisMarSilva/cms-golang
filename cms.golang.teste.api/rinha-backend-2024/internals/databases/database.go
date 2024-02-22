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
	ConnWriter *sqlx.DB // *sql.DB // *sqlx.DB
	ConnReader *sqlx.DB // *sql.DB // *sqlx.DB
)

// type IDatabaseWriter interface {
//     Writer
// }

// type IDatabaseRead interface {
//     Reader
// }

type IDatabase interface {
	StartDbWriter()
	GetDatabaseWriter() interface{}
	CloseDatabaseWriter()

	StartDbReader()
	GetDatabaseReader() interface{}
	CloseDatabaseReader()
}

type DatabasePostgres struct{}

func (DatabasePostgres) StartDb() *sqlx.DB {
	cfg := utils.NewConfig()

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

func (data *DatabasePostgres) StartDbWriter() {

	// if ConnWriter != nil {
	// 	log.Println("getting existent connection")
	// 	return
	// }

	// log.Println("opening connections")

	// database, err := sqlx.Connect(viper.GetString("DATABASE_DRIVER"), viper.GetString("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatalf("Error connecting to database : error=%v", err)
	// }

	// err = database.Ping()
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	log.Println("Connected")
	// }

	// maxConnectionsStr := viper.GetString("DATABASE_MAX_CONNECTIONS")
	// if maxConnectionsStr == "" {
	// 	maxConnectionsStr = "50"
	// }
	// maxConnections, _ := strconv.Atoi(maxConnectionsStr)

	// database.SetMaxOpenConns(maxConnections) // SetMaxOpenConns define o número máximo de conexões abertas com o banco de dados.
	// database.SetMaxIdleConns(100)            // SetMaxIdleConns define o número máximo de conexões no pool de conexão ociosa.
	// database.SetConnMaxIdleTime(time.Hour)
	// database.SetConnMaxLifetime(time.Minute * 5) // SetConnMaxLifetime define a quantidade máxima de tempo que uma conexão pode ser reutilizada.

	// ConnWriter = database
	ConnWriter = data.StartDb()
}

func (data *DatabasePostgres) StartDbReader() {

	// if ConnReader != nil {
	// 	log.Println("getting existent connection")
	// 	return
	// }

	// log.Println("opening connections")

	// database, err := sqlx.Connect(viper.GetString("DATABASE_DRIVER"), viper.GetString("DATABASE_URL"))
	// if err != nil {
	// 	log.Fatalf("Error connecting to database : error=%v", err)
	// }

	// err = database.Ping()
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	log.Println("Connected")
	// }

	// maxConnectionsStr := viper.GetString("DATABASE_MAX_CONNECTIONS")
	// if maxConnectionsStr == "" {
	// 	maxConnectionsStr = "50"
	// }
	// maxConnections, _ := strconv.Atoi(maxConnectionsStr)

	// database.SetMaxOpenConns(maxConnections) // SetMaxOpenConns define o número máximo de conexões abertas com o banco de dados.
	// database.SetMaxIdleConns(100)            // SetMaxIdleConns define o número máximo de conexões no pool de conexão ociosa.
	// database.SetConnMaxIdleTime(time.Hour)
	// database.SetConnMaxLifetime(time.Minute * 5) // SetConnMaxLifetime define a quantidade máxima de tempo que uma conexão pode ser reutilizada.

	// ConnReader = database
	ConnReader = data.StartDb()
}

func (DatabasePostgres) GetDatabaseWriter() *sqlx.DB {
	return ConnWriter
}

func (DatabasePostgres) GetDatabaseReader() *sqlx.DB {
	return ConnReader
}

func (DatabasePostgres) CloseDatabaseWriter() {
	if ConnWriter == nil {
		return
	}
	ConnWriter.Close()
}

func (DatabasePostgres) CloseDatabaseReader() {
	if ConnReader == nil {
		return
	}
	ConnReader.Close()
}
