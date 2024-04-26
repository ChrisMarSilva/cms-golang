package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	//_ "github.com/jackc/pgx/v5"
	//_ "github.com/jackc/pgx/v5/pgxpool"
	//_ "github.com/jackc/pgx/v5/stdlib"
)

type Database struct {
	*sql.DB
}

// var Db *sql.DB // *gorm.DB
// var dbConn = &Database{}
// var Connection *pgxpool.Pool

// func InitDb() *gorm.DB {
// 	Db = connectDB()
// 	return Db
// }

func Connect() (*Database, error) {
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

// func Close() {
// 	if Connection == nil {
// 		return
// 	}
// 	Connection.Close()
// }
