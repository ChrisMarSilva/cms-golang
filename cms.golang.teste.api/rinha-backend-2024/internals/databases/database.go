package databases

import (
	"context"
	"log"

	//"strconv"

	"github.com/chrismarsilva/rinha-backend-2024/internals/utils"
	// "github.com/jmoiron/sqlx"
	// _ "github.com/lib/pq"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	// Conn *sqlx.DB // *sql.DB // *sqlx.DB
	// ConnWriter *sqlx.DB
	// ConnReader *sqlx.DB
	ConnPgx *pgxpool.Pool
)

// type IDatabaseWriter interface { Writer }
// type IDatabaseRead interface { Reader }

type IDatabase interface {
	// StartDbConn()
	// GetDatabaseConn() interface{}
	// CloseDatabaseConn()

	// 	StartDbWriter()
	// 	GetDatabaseWriter() interface{}
	// 	CloseDatabaseWriter()

	// 	StartDbReader()
	// 	GetDatabaseReader() interface{}
	// 	CloseDatabaseReader()

	StartDbConnPgx()
	GetDatabaseConnPgx() interface{}
	CloseDatabaseConnPgx()
}

type DatabasePostgres struct{}

// func (DatabasePostgres) startDb(cfg *utils.Config) *sqlx.DB {
// 	database, err := sqlx.Connect(cfg.DbDriver, cfg.DbUri)
// 	if err != nil {
// 		log.Fatalf("Error connecting to database : error=%v", err)
// 	}
//
// 	if err = database.Ping(); err != nil {
// 		log.Println(err)
// 	} else {
// 		//log.Println("Connected")
// 	}
//
// 	maxConnections, _ := strconv.Atoi(cfg.DbDriver)
// 	database.SetMaxOpenConns(maxConnections) // SetMaxOpenConns define o número máximo de conexões abertas com o banco de dados.
// 	database.SetMaxIdleConns(maxConnections) // SetMaxIdleConns define o número máximo de conexões no pool de conexão ociosa.
// 	database.SetConnMaxIdleTime(0)           // time.Minute * 1
// 	database.SetConnMaxLifetime(0)           // SetConnMaxLifetime define a quantidade máxima de tempo que uma conexão pode ser reutilizada.
//
// 	return database
// }

// func (data *DatabasePostgres) StartDbConn(cfg *utils.Config) {
// 	Conn = data.startDb(cfg)
// }

func (data *DatabasePostgres) StartDbConnPgx(cfg *utils.Config) {

	/*
	 config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	    if err != nil {
	        return nil, err
	    }
	    config.MaxConns = 10000 // Defina o número máximo de conexões
	    config.MinConns = 100   // Defina o número mínimo de conexões
	    config.MaxConnIdleTime = 0 // Desative o tempo máximo de inatividade da conexão
	    config.AcquireTimeout = 0  // Desative o tempo máximo de espera para adquirir uma conexão

	    pool, err := pgxpool.ConnectConfig(context.Background(), config)
	    if err != nil {
	        return nil, err
	    }
	*/

	database, err := pgxpool.New(context.Background(), cfg.DbUri)
	//database, err := pgxpool.Connect(context.Background(), cfg.DbUri)
	if err != nil {
		log.Fatalf("Error connecting pool to database : error=%v", err)
	}

	if err = database.Ping(context.Background()); err != nil {
		log.Println(err)
	} else {
		//log.Println("Connected")
	}

	database.Config().MaxConns = 1000     // Define o número máximo de conexões abertas// Defina o número máximo de conexões
	database.Config().MinConns = 100      // Defina o número mínimo de conexões
	database.Config().MaxConnIdleTime = 0 // Define o tempo máximo de ociosidade para 5 minutos// Desative o tempo máximo de inatividade da conexão
	database.Config().MaxConnLifetime = 0 // Define o tempo máximo de vida (lifetime) para 30 minutos// Desative o tempo máximo de espera para adquirir uma conexão
	//database.Config().MaxConnIdleTime = time.Minute * 3

	// database.Exec(context.Background(), "PREPARE consulta_cliente_por_id (INTEGER) AS SELECT id, limite, saldo  FROM cliente  WHERE id = $1;")

	ConnPgx = database
}

// func (data *DatabasePostgres) StartDbWriter(cfg *utils.Config) {
// 	ConnWriter = data.startDb(cfg)
// }

// func (data *DatabasePostgres) StartDbReader(cfg *utils.Config) {
// 	ConnReader = data.startDb(cfg)
// }

// func (DatabasePostgres) GetDatabaseConn() *sqlx.DB {
// 	return Conn
// }

func (DatabasePostgres) GetDatabaseConnPgx() *pgxpool.Pool {
	return ConnPgx
}

// func (DatabasePostgres) GetDatabaseWriter() *sqlx.DB {
// 	return ConnWriter
// }

// func (DatabasePostgres) GetDatabaseReader() *sqlx.DB {
// 	return ConnReader
// }

// func (DatabasePostgres) CloseDatabaseConn() {
// 	if Conn == nil {
// 		return
// 	}
// 	Conn.Close()
// }

func (DatabasePostgres) CloseDatabaseConnPgx() {
	if ConnPgx == nil {
		return
	}
	ConnPgx.Close()
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
