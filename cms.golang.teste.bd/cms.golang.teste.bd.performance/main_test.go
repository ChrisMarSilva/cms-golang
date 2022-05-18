package main_test

import (
	"context"
	"database/sql"
	"strconv"
	"testing"

	database "github.com/ChrisMarSilva/cms.golang.teste.bd.performance/databases"
	entity "github.com/ChrisMarSilva/cms.golang.teste.bd.performance/entities"
	repository "github.com/ChrisMarSilva/cms.golang.teste.bd.performance/repositories"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

// go test
// go test -v
// go test -bench=.
// go test -bench=Add
// go test -run BenchmarkMSSQLDriverSqlCreate -v
// go test -run=XXX -bench . -benchmem

// --------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------

func GetConMSSQLDriverSql() *sql.DB {
	driverDB := database.DatabaseSQLServerSQL{}
	driverDB.StartDB()
	return driverDB.GetDatabase()
}

func GetRepoMSSQLDriverSql(db *sql.DB) repository.IUserRepository {
	return *repository.NewUserRepositorySQLServerSQL(db)
}

func TestMSSQLDriverSqlCreate(t *testing.T) {

	db := GetConMSSQLDriverSql()
	defer db.Close()
	userRepo := GetRepoMSSQLDriverSql(db)
	ctx := context.Background()

	user := entity.User{ID: uuid.New(), Nome: "PESSOA SQL 1", Status: entity.UserActive}
	err := userRepo.Create(ctx, &user)
	if err != nil {
		t.Error("Erro on create", err)
	}
}

func BenchmarkMSSQLDriverSqlCreate(b *testing.B) {

	//var start time.Time = time.Now()

	//b.StopTimer()
	db := GetConMSSQLDriverSql()
	defer db.Close()
	userRepo := GetRepoMSSQLDriverSql(db)
	ctx := context.Background()
	//b.StartTimer()

	b.ResetTimer()

	//count := 0
	for i := 0; i < b.N; i++ {
		user := entity.User{ID: uuid.New(), Nome: "PESSOA SQL " + strconv.Itoa(i), Status: entity.UserActive}
		err := userRepo.Create(ctx, &user)
		if err != nil {
			b.Error("Erro on create", err)
		}
		//count++
	}

	// vai rodar 4 threads, vai pegar o tempo da thread de rodou o maior numero de vezes com o menor tempo
	//println("BenchmarkMSSQLDriverSqlCreate = ", count, "Tempo=", time.Since(start))
}

// --------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------

func GetConMSSQLDriverGorm() *gorm.DB {
	driverDB := database.DatabaseSQLServerGorm{}
	driverDB.StartDB()
	return driverDB.GetDatabase()
}

func GetRepoMSSQLDriverGorm(db *gorm.DB) repository.IUserRepository {
	return *repository.NewUserRepositorySQLServerGorm(db)
}

func TestMSSQLDriverGormCreate(t *testing.T) {

	db := GetConMSSQLDriverGorm()
	//defer db.Close()
	userRepo := GetRepoMSSQLDriverGorm(db)
	ctx := context.Background()

	user := entity.User{ID: uuid.New(), Nome: "PESSOA GORM 1", Status: entity.UserActive}
	err := userRepo.Create(ctx, &user)
	if err != nil {
		t.Error("Erro on create", err)
	}
}

func BenchmarkMSSQLDriverGormCreate(b *testing.B) {

	//var start time.Time = time.Now()

	//b.StopTimer()
	db := GetConMSSQLDriverGorm()
	//defer db.Close()
	userRepo := GetRepoMSSQLDriverGorm(db)
	ctx := context.Background()
	//b.StartTimer()

	b.ResetTimer()

	//count := 0
	for i := 0; i < b.N; i++ {
		user := entity.User{ID: uuid.New(), Nome: "PESSOA GORM " + strconv.Itoa(i), Status: entity.UserActive}
		err := userRepo.Create(ctx, &user)
		if err != nil {
			b.Error("Erro on create", err)
		}
		//count++
	}

	// vai rodar 4 threads, vai pegar o tempo da thread de rodou o maior numero de vezes com o menor tempo
	//println("BenchmarkMSSQLDriverSqlCreate = ", count, "Tempo=", time.Since(start))
}

// --------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------

func GetConMSSQLDriverSqlX() *sqlx.DB {
	driverDB := database.DatabaseSQLServerSQLX{}
	driverDB.StartDB()
	return driverDB.GetDatabase()
}

func GetRepoMSSQLDriverSqlX(db *sqlx.DB) repository.IUserRepository {
	return *repository.NewUserRepositorySQLServerSQLX(db)
}

func TestMSSQLDriverSqlXCreate(t *testing.T) {

	db := GetConMSSQLDriverSqlX()
	defer db.Close()
	userRepo := GetRepoMSSQLDriverSqlX(db)
	ctx := context.Background()

	user := entity.User{ID: uuid.New(), Nome: "PESSOA SQL 1", Status: entity.UserActive}
	err := userRepo.Create(ctx, &user)
	if err != nil {
		t.Error("Erro on create", err)
	}
}

func BenchmarkMSSQLDriverSqlXCreate(b *testing.B) {

	//var start time.Time = time.Now()

	//b.StopTimer()
	db := GetConMSSQLDriverSqlX()
	defer db.Close()
	userRepo := GetRepoMSSQLDriverSqlX(db)
	ctx := context.Background()
	//b.StartTimer()

	b.ResetTimer()

	//count := 0
	for i := 0; i < b.N; i++ {
		user := entity.User{ID: uuid.New(), Nome: "PESSOA SQLX " + strconv.Itoa(i), Status: entity.UserActive}
		err := userRepo.Create(ctx, &user)
		if err != nil {
			b.Error("Erro on create", err)
		}
		//count++
	}

	// vai rodar 4 threads, vai pegar o tempo da thread de rodou o maior numero de vezes com o menor tempo
	//println("BenchmarkMSSQLDriverSqlCreate = ", count, "Tempo=", time.Since(start))
}

// --------------------------------------------------------------------------------------
// --------------------------------------------------------------------------------------
