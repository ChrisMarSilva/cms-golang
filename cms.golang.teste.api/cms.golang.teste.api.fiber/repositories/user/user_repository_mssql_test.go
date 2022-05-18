package user_test

import (
	"context"
	"math/rand"
	"strconv"
	"testing"
	"time"

	database "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/databases"
	entity "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/entities"
	repository "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/repositories"
	"github.com/ChrisMarSilva/cms-golang-teste-api-fiber/repositories/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// go test
// go test -v
// go test -bench=.
// go test -bench=Add
// go test -run TestUserrepoGetAll -v
// go test -run=XXX -bench . -benchmem

// log.Println(uuid.New().String())
// log.Println(uuid.NewString())
// id := uuid.New()
// log.Println(id.String())

func GetDbPadrao() *gorm.DB {
	driverDB := database.DatabaseSQLServer{} // DatabaseMySQL // DatabaseSQLServer
	driverDB.StartDB()
	return driverDB.GetDatabase()
}

func GetUserRepositoryPadrao(db *gorm.DB) repository.IUserRepository {
	return user.NewUserRepositoryMSSQL(nil, db)
}

func GetNumeroPadrao() int {
	min := 10
	max := 30
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func GetUuidPadrao() string {
	return "A4147F69-AD68-49AC-B7CF-0397AE22C574" // pegar do direto do banco de dados
}

func TestUserRepoGetAll(t *testing.T) {

	ctx := context.Background()
	db := GetDbPadrao()
	repo := GetUserRepositoryPadrao(db)
	var users []entity.User

	err := repo.GetAll(ctx, &users)
	if err != nil {
		t.Fatalf("error %t", err)
	}

	//fmt.Println(users)
	// for _, user := range users {
	// 	fmt.Println(user)
	// 	fmt.Printf("id: %s; nome: %s; status: %s; statusS: %s\n", user.ID.String(), user.Nome, user.Status, user.StatusString())
	// }

}

func TestUserRepoGet(t *testing.T) {

	ctx := context.Background()
	db := GetDbPadrao()
	repo := GetUserRepositoryPadrao(db)
	var user entity.User

	uuid, err := uuid.Parse(GetUuidPadrao())
	if err != nil {
		t.Fatalf("error uuid %t", err)
	}

	err = repo.Get(ctx, &user, uuid)
	if err != nil {
		t.Fatalf("error %t", err)
	}

	// fmt.Println(user)
	// fmt.Printf("id: %s; nome: %s; status: %s; statusS: %s\n", user.ID.String(), user.Nome, user.Status, user.StatusString())

}

func TestUserRepoCreate(t *testing.T) {

	ctx := context.Background()
	db := GetDbPadrao()
	repo := GetUserRepositoryPadrao(db)

	num := GetNumeroPadrao()
	user := entity.User{ID: uuid.New(), Nome: "PESSOA " + strconv.Itoa(num), Status: entity.UserActive}

	err := repo.Create(ctx, &user)
	if err != nil {
		t.Fatalf("error %t", err)
	}

	//fmt.Println(user)
	//fmt.Printf("id: %s; nome: %s; status: %s; statusS: %s\n", user.ID.String(), user.Nome, user.Status, user.StatusString())

}

func TestUserRepoUpdate(t *testing.T) {

	ctx := context.Background()
	db := GetDbPadrao()
	repo := GetUserRepositoryPadrao(db)
	var user entity.User

	uuid, err := uuid.Parse(GetUuidPadrao())
	if err != nil {
		t.Fatalf("error uuid %t", err)
	}

	err = repo.Get(ctx, &user, uuid)
	if err != nil {
		t.Fatalf("error Get %t", err)
	}

	//fmt.Println(user)
	//fmt.Printf("ANTES --> id: %s; nome: %s; status: %s; statusS: %s\n", user.ID.String(), user.Nome, user.Status, user.StatusString())

	num := GetNumeroPadrao()
	user.Nome = "PESSOA 2 => PESSOA " + strconv.Itoa(num)
	err = repo.Update(ctx, &user)
	if err != nil {
		t.Fatalf("error Update %t", err)
	}

	//fmt.Printf("DEPOIS --> id: %s; nome: %s; status: %s; statusS: %s\n", user.ID.String(), user.Nome, user.Status, user.StatusString())

}

func TestUserRepoDelete(t *testing.T) {

	ctx := context.Background()
	db := GetDbPadrao()
	repo := GetUserRepositoryPadrao(db)

	uuid, err := uuid.Parse(GetUuidPadrao())
	if err != nil {
		t.Fatalf("error uuid %t", err)
	}

	//var user entity.User
	// err = repo.Get(&user, uuid)
	// if err != nil {
	// 	t.Fatalf("error Get %t", err)
	// }

	//fmt.Println(user)
	//fmt.Printf("id: %s; nome: %s; status: %s; statusS: %s\n", user.ID.String(), user.Nome, user.Status, user.StatusString())

	err = repo.Delete(ctx, uuid)
	if err != nil {
		t.Fatalf("error Delete %t", err)
	}

}

func BenchmarkRepoGetAll(b *testing.B) {
	ctx := context.Background()
	db := GetDbPadrao()
	repo := GetUserRepositoryPadrao(db)
	for i := 0; i < b.N; i++ {
		var users []entity.User
		if err := repo.GetAll(ctx, &users); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkRepoGet(b *testing.B) {
	ctx := context.Background()
	db := GetDbPadrao()
	repo := GetUserRepositoryPadrao(db)
	uuid, _ := uuid.Parse(GetUuidPadrao())
	for i := 0; i < b.N; i++ {
		var user entity.User
		if err := repo.Get(ctx, &user, uuid); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkRepoCreate(b *testing.B) {
	ctx := context.Background()
	db := GetDbPadrao()
	repo := GetUserRepositoryPadrao(db)
	for i := 0; i < b.N; i++ {
		user := entity.User{ID: uuid.New(), Nome: "PESSOA " + strconv.Itoa(GetNumeroPadrao()), Status: entity.UserActive}
		if err := repo.Create(ctx, &user); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

// func BenchmarkRepoCreate10Mil(b *testing.B) {
// ctx := context.Background()
// 	db := GetDbPadrao()
// 	repo := GetUserRepositoryPadrao(db)
// 	// tx := db.Session(&gorm.Session{PrepareStmt: true})
// 	// tx := repository.Db.Begin()
// 	// tx.Commit()
// 	for i := 0; i < 10001; i++ {
// 		user := entity.User{ID: uuid.New(), Nome: "PESSOA " + strconv.Itoa(i), Status: entity.UserActive}
// 		if err := repo.Create(ctx, &user); err != nil {
// 			b.Fatalf("error %t", err)
// 		}
// 	}
// }

func BenchmarkRepoUpdate(b *testing.B) {
	ctx := context.Background()
	db := GetDbPadrao()
	repo := GetUserRepositoryPadrao(db)
	uuid, _ := uuid.Parse(GetUuidPadrao())
	for i := 0; i < b.N; i++ {
		var user entity.User
		if err := repo.Get(ctx, &user, uuid); err != nil {
			b.Fatalf("error Get %t", err)
		}
		user.Nome = "NOVA PESSOA " + strconv.Itoa(GetNumeroPadrao())
		if err := repo.Update(ctx, &user); err != nil {
			b.Fatalf("error Update %t", err)
		}
	}
}
