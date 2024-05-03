package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type IDatabase interface {
	CloseDb()
}

type Database struct {
	*sql.DB
}

func NewDatabase() (*Database, error) {
	// config := configs.New()

	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
	//dsn := fmt.Sprintf( "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	//db, err := sqlx.Connect("postgres", dsn)
	//db, err := sql.Open("pgx", dsn)

	// Connection, err = pgxpool.Connect(context.Background(), config.DatabaseURL)
	//dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	//defer dbpool.Close()

	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	// defer conn.Close(context.Background())

	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 	ctxTimeout, ctxCancel := context.WithTimeout(context.Background(), time.Second*3)
	// defer ctxCancel()
	// db, err := sqlx.ConnectContext(ctxTimeout, "postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",  postgresConfig.Host, postgresConfig.Port, postgresConfig.Username, postgresConfig.Password, postgresConfig.DBName, postgresConfig.SSLMode))

	db, err := sql.Open("sqlite3", "./banco.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	// db.Logger = logger.Default.LogMode(logger.Info)

	// err = DB.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.User{})

	// dbConn.SQL = db
	// return dbConn, err

	return &Database{db}, nil
}

func (this Database) CloseDb() {
	if this.DB == nil {
		return
	}

	err := this.DB.Close()
	if err != nil {
		log.Println("database close failure: %v", err)
	}
}

/*
// ./platform/database/postgres.go

package database

import (
    "fmt"
    "os"
    "strconv"
    "time"

    "github.com/jmoiron/sqlx"

    _ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
)

// PostgreSQLConnection func for connection to PostgreSQL database.
func PostgreSQLConnection() (*sqlx.DB, error) {
    // Define database connection settings.
    maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
    maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
    maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

    // Define database connection for PostgreSQL.
    db, err := sqlx.Connect("pgx", os.Getenv("DB_SERVER_URL"))
    if err != nil {
        return nil, fmt.Errorf("error, not connected to database, %w", err)
    }

    // Set database connection settings.
    db.SetMaxOpenConns(maxConn)                           // the default is 0 (unlimited)
    db.SetMaxIdleConns(maxIdleConn)                       // defaultMaxIdleConns = 2
    db.SetConnMaxLifetime(time.Duration(maxLifetimeConn)) // 0, connections are reused forever

    // Try to ping database.
    if err := db.Ping(); err != nil {
        defer db.Close() // close database connection
        return nil, fmt.Errorf("error, not sent ping to database, %w", err)
    }

    return db, nil
}

*/
