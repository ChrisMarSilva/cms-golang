package database

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {

	const DB_USERNAME = "root"
	const DB_PASSWORD = "senha"
	const DB_NAME = "database"
	const DB_HOST = "localhost"
	const DB_PORT = "3306"
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?parseTime=true&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database : error=%v", err)
	}

	db = database

	config, _ := db.DB()
	config.SetMaxIdleConns(10)                 // SetMaxIdleConns define o número máximo de conexões no pool de conexão ociosa.
	config.SetMaxOpenConns(100)                // SetMaxOpenConns define o número máximo de conexões abertas com o banco de dados.
	config.SetConnMaxLifetime(time.Minute * 5) // SetConnMaxLifetime define a quantidade máxima de tempo que uma conexão pode ser reutilizada.

	log.Println("CMS - StartDB")
}

func GetDatabase() *gorm.DB {
	return db
}
