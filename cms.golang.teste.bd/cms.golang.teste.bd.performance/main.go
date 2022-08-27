package main

import (
	"context"
	"log"
	"strconv"
	"time"

	database "github.com/ChrisMarSilva/cms.golang.teste.bd.performance/databases"
	entity "github.com/ChrisMarSilva/cms.golang.teste.bd.performance/entities"
	repository "github.com/ChrisMarSilva/cms.golang.teste.bd.performance/repositories"
	"github.com/google/uuid"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.bd.performance
// go get -u github.com/google/uuid
// go get -u github.com/denisenkom/go-mssqldb  // SQL
// go get -u gorm.io/gorm
// go get -u gorm.io/hints
// go get -u gorm.io/driver/sqlserver
// go get -u github.com/jmoiron/sqlx
// go get -u github.com/kyleconroy/sqlc/cmd/sqlc
// go mod tidy

// go run main.go

func main() {

	// var startGeral time.Time
	// startGeral = time.Now()
	// usersGeral1 := GetListUsers()
	// log.Println("GetListUsers     =", time.Since(startGeral), "Len=", len(usersGeral1))
	// startGeral = time.Now()
	// usersGeral2 := GetListUsersFast()
	// log.Println("GetListUsersFast =", time.Since(startGeral), "Len=", len(usersGeral2))

	// TesteDbSQLServerDriverSQL()
	// TesteDbSQLServerDriverSQLX()
	// TesteDbSQLServerDriverGorm()
	// TesteDbSQLServerDriverSQLC()

}

const QTDE = 1000

func GetListUsers() []entity.User {
	var users []entity.User
	for i := 0; i < QTDE; i++ {
		user := entity.User{ID: uuid.New(), Nome: "PESSOA GERAL " + strconv.Itoa(i), Status: entity.UserActive}
		users = append(users, user)
	}
	return users
}

func GetListUsersFast() map[int]entity.User {
	//users := make([]entity.User, iQtde)
	users := make(map[int]entity.User, QTDE)
	for i := 0; i < QTDE; i++ {
		user := entity.User{ID: uuid.New(), Nome: "PESSOA GERAL " + strconv.Itoa(i), Status: entity.UserActive}
		users[i] = user
	}
	return users
}

func TesteDbSQLServerDriverSQL() {

	var startGeral time.Time
	var startUnico time.Time

	startGeral = time.Now()

	startUnico = time.Now()
	driverDB := database.DatabaseSQLServerSQL{}
	driverDB.StartDB()
	db := driverDB.GetDatabase()
	defer driverDB.Close()
	userRepo := repository.NewUserRepositorySQLServerSQL(db)
	ctx := context.Background()
	log.Println("Teste Db SQLServer Driver SQL  StartDB=", time.Since(startUnico))

	startUnico = time.Now()
	user := entity.User{ID: uuid.New(), Nome: "PESSOA SQL 1", Status: entity.UserActive}
	err := userRepo.Create(ctx, &user)
	if err != nil {
		log.Println("Erro on create", err)
	}
	log.Println("Teste Db SQLServer Driver SQL  Create=", time.Since(startUnico))

	usersGeral := GetListUsers() // GetListUsers()// GetListUsersFast()
	startUnico = time.Now()
	err = userRepo.CreateInBatch(ctx, usersGeral)
	if err != nil {
		log.Println("Erro on CreateInBatch", err)
	}
	log.Println("Teste Db SQLServer Driver SQL  CreateInBatch=", time.Since(startUnico))

	startUnico = time.Now()
	user.Nome = user.Nome + " -> " + user.Nome
	user.Status = entity.UserInactive
	err = userRepo.Update(ctx, &user)
	if err != nil {
		log.Println("Erro on update", err)
	}
	log.Println("Teste Db SQLServer Driver SQL  Update=", time.Since(startUnico))

	startUnico = time.Now()
	err = userRepo.Delete(ctx, user.ID)
	if err != nil {
		log.Println("Erro on delete", err)
	}
	// uuid1, err := uuid.Parse("7AA3CB84-FD5E-4631-8928-D2EEB5D53E17")
	// if err != nil {
	// 	log.Println("Erro on Parse uuid", err)
	// }
	// err = userRepo.Delete(ctx, uuid1)
	// if err != nil {
	// 	log.Println("Erro on delete", err)
	// }
	log.Println("Teste Db SQLServer Driver SQL  Delete=", time.Since(startUnico))

	startUnico = time.Now()
	var users []entity.User
	err = userRepo.GetAll(ctx, &users)
	if err != nil {
		log.Println("Erro on get all", err)
	}
	// for _, user := range users {
	// 	log.Println("GetAll => ID:", user.ID, "Nome:", user.Nome, "Status:", user.Status)
	// }
	log.Println("Teste Db SQLServer Driver SQL  GetAll=", time.Since(startUnico))

	startUnico = time.Now()
	uuid2, err := uuid.Parse("0A170649-F3D4-4D66-8F39-6BB89DD75854")
	if err != nil {
		log.Println("Erro on Parse uuid", err)
	}
	var user1 entity.User
	err = userRepo.Get(ctx, &user1, uuid2)
	// if err != nil {
	// 	log.Println("Erro on get", err)
	// } else {
	// 	log.Println("Get => ID:", user1.ID, "Nome:", user1.Nome, "Status:", user1.Status)
	// }
	log.Println("Teste Db SQLServer Driver SQL  Get=", time.Since(startUnico))

	log.Println("Teste Db SQLServer Driver SQL  Fim=", time.Since(startGeral))
}

func TesteDbSQLServerDriverSQLX() {

	var startGeral time.Time
	var startUnico time.Time

	startGeral = time.Now()

	startUnico = time.Now()
	driverDB := database.DatabaseSQLServerSQLX{}
	driverDB.StartDB()
	db := driverDB.GetDatabase()
	defer db.Close()
	userRepo := repository.NewUserRepositorySQLServerSQLX(db)
	log.Println("Teste Db SQLServer Driver SQLX StartDB=", time.Since(startUnico))

	ctx := context.Background()

	startUnico = time.Now()
	user := entity.User{ID: uuid.New(), Nome: "PESSOA SQLX 1", Status: entity.UserActive}
	err := userRepo.Create(ctx, &user)
	if err != nil {
		log.Println("Erro on create", err)
	}
	log.Println("Teste Db SQLServer Driver SQLX Create=", time.Since(startUnico))

	usersGeral := GetListUsers() // GetListUsers()// GetListUsersFast()
	startUnico = time.Now()
	err = userRepo.CreateInBatch(ctx, usersGeral)
	if err != nil {
		log.Println("Erro on CreateInBatch", err)
	}
	log.Println("Teste Db SQLServer Driver SQLX CreateInBatch=", time.Since(startUnico))

	startUnico = time.Now()
	user.Nome = user.Nome + " -> " + user.Nome
	user.Status = entity.UserInactive
	err = userRepo.Update(ctx, &user)
	if err != nil {
		log.Println("Erro on update", err)
	}
	log.Println("Teste Db SQLServer Driver SQLX Update=", time.Since(startUnico))

	startUnico = time.Now()
	err = userRepo.Delete(ctx, user.ID)
	if err != nil {
		log.Println("Erro on delete", err)
	}
	// uuid1, err := uuid.Parse("7AA3CB84-FD5E-4631-8928-D2EEB5D53E17")
	// if err != nil {
	// 	log.Println("Erro on Parse uuid", err)
	// }
	// err = userRepo.Delete(ctx, uuid1)
	// if err != nil {
	// 	log.Println("Erro on delete", err)
	// }
	log.Println("Teste Db SQLServer Driver SQLX Delete=", time.Since(startUnico))

	startUnico = time.Now()
	var users []entity.User
	err = userRepo.GetAll(ctx, &users)
	if err != nil {
		log.Println("Erro on get all", err)
	}
	// for _, user := range users {
	// 	log.Println("GetAll => ID:", user.ID, "Nome:", user.Nome, "Status:", user.Status)
	// }
	log.Println("Teste Db SQLServer Driver SQLX GetAll=", time.Since(startUnico))

	startUnico = time.Now()
	uuid2, err := uuid.Parse("1A46F0E7-3CB4-47B8-85A0-0E05541D52E1")
	if err != nil {
		log.Println("Erro on Parse uuid", err)
	}
	var user1 entity.User
	err = userRepo.Get(ctx, &user1, uuid2)
	// if err != nil {
	// 	log.Println("Erro on get", err)
	// } else {
	// 	log.Println("Get => ID:", user1.ID, "Nome:", user1.Nome, "Status:", user1.Status)
	// }
	log.Println("Teste Db SQLServer Driver SQLX Get=", time.Since(startUnico))

	log.Println("Teste Db SQLServer Driver SQLX Fim=", time.Since(startGeral))
}

func TesteDbSQLServerDriverGorm() {

	var startGeral time.Time
	var startUnico time.Time

	startGeral = time.Now()

	startUnico = time.Now()
	driverDB := database.DatabaseSQLServerGorm{}
	driverDB.StartDB()
	db := driverDB.GetDatabase()
	defer driverDB.Close()
	userRepo := repository.NewUserRepositorySQLServerGorm(db)
	ctx := context.Background()
	log.Println("Teste Db SQLServer Driver GORM StartDB=", time.Since(startUnico))

	startUnico = time.Now()
	user := entity.User{ID: uuid.New(), Nome: "PESSOA GORM 1", Status: entity.UserActive}
	err := userRepo.Create(ctx, &user)
	if err != nil {
		log.Println("Erro on create", err)
	}
	log.Println("Teste Db SQLServer Driver GORM Create=", time.Since(startUnico))

	usersGeral := GetListUsers() // GetListUsers()// GetListUsersFast()
	startUnico = time.Now()
	err = userRepo.CreateInBatch(ctx, usersGeral)
	if err != nil {
		log.Println("Erro on CreateInBatch", err)
	}
	log.Println("Teste Db SQLServer Driver GORM CreateInBatch=", time.Since(startUnico))

	startUnico = time.Now()
	user.Nome = user.Nome + " -> " + user.Nome
	user.Status = entity.UserInactive
	err = userRepo.Update(ctx, &user)
	if err != nil {
		log.Println("Erro on update", err)
	}
	log.Println("Teste Db SQLServer Driver GORM Update=", time.Since(startUnico))

	startUnico = time.Now()
	err = userRepo.Delete(ctx, user.ID)
	if err != nil {
		log.Println("Erro on delete", err)
	}
	// uuid1, err := uuid.Parse("5B908935-77F6-4644-94C7-C3DD307D9491")
	// if err != nil {
	// 	log.Println("Erro on Parse uuid", err)
	// }
	// err = userRepo.Delete(ctx, uuid1)
	// if err != nil {
	// 	log.Println("Erro on delete", err)
	// }
	log.Println("Teste Db SQLServer Driver GORM Delete=", time.Since(startUnico))

	startUnico = time.Now()
	var users []entity.User
	err = userRepo.GetAll(ctx, &users)
	if err != nil {
		log.Println("Erro on get all", err)
	}
	// for _, user := range users {
	// 	log.Println("GetAll => ID:", user.ID, "Nome:", user.Nome, "Status:", user.Status)
	// }
	log.Println("Teste Db SQLServer Driver GORM GetAll=", time.Since(startUnico))

	startUnico = time.Now()
	uuid2, err := uuid.Parse("F548743D-B71B-4411-AB8D-D90E24569F39")
	if err != nil {
		log.Println("Erro on Parse uuid", err)
	}
	var user1 entity.User
	err = userRepo.Get(ctx, &user1, uuid2)
	// if err != nil {
	// 	log.Println("Erro on get", err)
	// } else {
	// 	log.Println("Get => ID:", user1.ID, "Nome:", user1.Nome, "Status:", user1.Status)
	// }
	log.Println("Teste Db SQLServer Driver GORM Get=", time.Since(startUnico))

	log.Println("Teste Db SQLServer Driver GORM Fim=", time.Since(startGeral))
}

func TesteDbSQLServerDriverSQLC() {

	var startGeral time.Time
	//var startUnico time.Time

	startGeral = time.Now()

	// driverDB := database.DatabaseSQLServerSQLC{}
	// driverDB.StartDB()
	// db := driverDB.GetDatabase()
	// defer driverDB.Close()

	// userRepo := repository.NewUserRepositorySQLServerSQLC(db)

	// user := entity.User{ID: uuid.New(), Nome: "PESSOA SQL 1", Status: entity.UserActive}
	// err := userRepo.Create(&user)
	// if err != nil {
	// 	log.Println("Erro on create", err)
	// }

	// user.Nome = user.Nome + " -> " + user.Nome
	// user.Status = entity.UserInactive
	// err = userRepo.Update(&user)
	// if err != nil {
	// 	log.Println("Erro on update", err)
	// }

	// err = userRepo.Delete(user.ID)
	// if err != nil {
	// 	log.Println("Erro on delete", err)
	// }
	// // uuid1, err := uuid.Parse("7AA3CB84-FD5E-4631-8928-D2EEB5D53E17")
	// // if err != nil {
	// // 	log.Println("Erro on Parse uuid", err)
	// // }
	// // err = userRepo.Delete(uuid1)
	// // if err != nil {
	// // 	log.Println("Erro on delete", err)
	// // }

	// var users []entity.User
	// err = userRepo.GetAll(&users)
	// if err != nil {
	// 	log.Println("Erro on get all", err)
	// }
	// // for _, user := range users {
	// // 	log.Println("GetAll => ID:", user.ID, "Nome:", user.Nome, "Status:", user.Status)
	// // }

	// uuid2, err := uuid.Parse("0A170649-F3D4-4D66-8F39-6BB89DD75854")
	// if err != nil {
	// 	log.Println("Erro on Parse uuid", err)
	// }
	// var user1 entity.User
	// err = userRepo.Get(&user1, uuid2)
	// // if err != nil {
	// // 	log.Println("Erro on get", err)
	// // } else {
	// // 	log.Println("Get => ID:", user1.ID, "Nome:", user1.Nome, "Status:", user1.Status)
	// // }

	log.Println("Teste Db SQLServer Driver SQLC Fim=", time.Since(startGeral))
}
