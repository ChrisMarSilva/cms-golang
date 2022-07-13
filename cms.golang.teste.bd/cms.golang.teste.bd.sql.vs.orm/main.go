package main

import (
	"context"
	"log"
	"time"

	"github.com/ChrisMarSilva/cms.golang.teste.bd.sql.vs.orm/configs"
	"github.com/ChrisMarSilva/cms.golang.teste.bd.sql.vs.orm/entities"
	"github.com/ChrisMarSilva/cms.golang.teste.bd.sql.vs.orm/repositories"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.bd.sql.vs.orm
// go get -u github.com/jmoiron/sqlx
// go get -u gorm.io/gorm
// go get -u github.com/jinzhu/gorm
// go get -u github.com/go-sql-driver/mysql
// go mod tidy

// go run main.go

func main() {

	var start time.Time

	repos := repositories.New(repositories.Options{
		ReaderSqlx: configs.GetReaderSqlx(),
		WriterSqlx: configs.GetWriterSqlx(),
		ReaderGorm: configs.GetReaderGorm(),
		WriterGorm: configs.GetWriterGorm(),
	})

	start = time.Now()

	for i := 0; i < 10000; i++ {
		repos.User.GetAll(context.Background())
	}

	log.Println("Success = ", time.Since(start))
}

func main_old() {

	var start time.Time

	repos := repositories.New(repositories.Options{
		ReaderSqlx: configs.GetReaderSqlx(),
		WriterSqlx: configs.GetWriterSqlx(),
		ReaderGorm: configs.GetReaderGorm(),
		WriterGorm: configs.GetWriterGorm(),
	})

	start = time.Now()

	for i := 0; i < 10000; i++ {
		err := repos.User.Create(context.Background(), entities.User{
			Name:     "Leonardo Miranda",
			Email:    "email@email.com",
			NickName: "leomirandadev",
			Password: "e8d95a51f3af4a3b134bf6bb680a213a",
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Success = ", time.Since(start))
}
