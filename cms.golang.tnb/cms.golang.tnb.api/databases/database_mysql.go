package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseMySQL struct {
	dbMySQL *gorm.DB
}

func (data *DatabaseMySQL) StartDB() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Silent // Error // Warn // Info
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	newConfig := &gorm.Config{
		Logger:                                   newLogger,
		SkipDefaultTransaction:                   true,
		QueryFields:                              true,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     true,
		DisableForeignKeyConstraintWhenMigrating: true,
		DryRun:                                   false,
	}

	dsn := ""
	dsn = "root:senha@tcp(localhost:3306)/database?parseTime=true&loc=Local"
	dsn = "root:senha@tcp(localhost:3306)/database?parseTime=true"

	database, err := gorm.Open(mysql.Open(dsn), newConfig)
	if err != nil {
		log.Fatalf("Error connecting to database : error=%v", err)
	}

	config, err := database.DB()
	if err != nil {
		log.Fatalf("Error connecting to database : error=%v", err)
	}

	config.SetMaxIdleConns(100) // SetMaxIdleConns define o número máximo de conexões no pool de conexão ociosa.
	config.SetMaxOpenConns(200) // SetMaxOpenConns define o número máximo de conexões abertas com o banco de dados.
	config.SetConnMaxIdleTime(time.Hour)
	config.SetConnMaxLifetime(time.Minute * 5) // 24 *time.Hour // SetConnMaxLifetime define a quantidade máxima de tempo que uma conexão pode ser reutilizada.

	data.dbMySQL = database
}

func (data *DatabaseMySQL) GetDatabase() *gorm.DB {
	return data.dbMySQL
}

func (DatabaseMySQL) Close() {
}
