package databases

import (
	"log"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type IDatabase interface {
	StartDB()
	GetDatabase() interface{}
	Close()
}

var (
	Conn *sqlx.DB // *sql.DB // *sqlx.DB
)

type DatabasePostgres struct {
}

func (data *DatabasePostgres) StartDB() {
	if Conn != nil {
		log.Println("getting existent connection")
		return
	}

	log.Println("opening connections")

	database, err := sqlx.Connect(viper.GetString("DATABASE_DRIVER"), viper.GetString("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error connecting to database : error=%v", err)
	}

	err = database.Ping()
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Connected")
	}

	maxConnectionsS := viper.GetString("DATABASE_MAX_CONNECTIONS")
	if maxConnectionsS == "" {
		maxConnectionsS = "50"
	}
	maxConnections, _ := strconv.Atoi(maxConnectionsS)

	database.SetMaxOpenConns(maxConnections) // SetMaxOpenConns define o número máximo de conexões abertas com o banco de dados.
	database.SetMaxIdleConns(100)            // SetMaxIdleConns define o número máximo de conexões no pool de conexão ociosa.
	database.SetConnMaxIdleTime(time.Hour)
	database.SetConnMaxLifetime(time.Minute * 5) // 24 *time.Hour // SetConnMaxLifetime define a quantidade máxima de tempo que uma conexão pode ser reutilizada.

	// database.MustExec("TRUNCATE TABLE cliente_transacao")
	// database.MustExec("UPDATE cliente_saldo SET total = 0")
	// database.MustExec("UPDATE cliente SET saldo = 0")

	Conn = database
}

func (DatabasePostgres) GetDatabase() *sqlx.DB {
	return Conn
}

func (DatabasePostgres) Close() {
}
