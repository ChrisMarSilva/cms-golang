package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbSqlServer *gorm.DB

type DatabaseSQLServer struct {
}

func (DatabaseSQLServer) StartDB() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Silent // Error // Warn // Info
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	// &gorm.Config{Logger: newLogger, SkipDefaultTransaction: true, QueryFields: true, PrepareStmt: true}
	newConfig := &gorm.Config{
		Logger:                                   newLogger,
		SkipDefaultTransaction:                   true,
		QueryFields:                              true,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     true,
		DisableForeignKeyConstraintWhenMigrating: true,
		DryRun:                                   false,
	}

	// dsn := "sqlserver://sa:sa@10.0.0.82:1433?database=CMS_TESTE_CNAB240" // localhost
	//dsn := "sqlserver://sa:sa@127.0.0.1:5401?database=CMS_TESTE_CNAB240" // docker
	dsn := "sqlserver://sa:sa@127.0.0.1:5401?database=CMS_TESTE_JDSPB" // docker

	database, err := gorm.Open(sqlserver.Open(dsn), newConfig)
	if err != nil {
		log.Fatalf("Error connecting to database : error=%v", err)
	}

	// for {
	// 	pingErr := db.DB().Ping()
	// 	if pingErr != nil {
	// 		fmt.Println(pingErr)
	// 	} else {
	// 		fmt.Println("Connected")
	// 	}
	// 	time.Sleep(time.Duration(3) * time.Second)
	// }

	config, err := database.DB()
	if err != nil {
		log.Fatalf("Error connecting to database : error=%v", err)
	}

	config.SetMaxIdleConns(100) // SetMaxIdleConns define o número máximo de conexões no pool de conexão ociosa.
	config.SetMaxOpenConns(200) // SetMaxOpenConns define o número máximo de conexões abertas com o banco de dados.
	config.SetConnMaxIdleTime(time.Hour)
	config.SetConnMaxLifetime(time.Minute * 5) // 24 *time.Hour // SetConnMaxLifetime define a quantidade máxima de tempo que uma conexão pode ser reutilizada.

	dbSqlServer = database
}

func (DatabaseSQLServer) GetDatabase() *gorm.DB {
	return dbSqlServer
}

func (DatabaseSQLServer) Close() {
}
