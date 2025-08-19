package stores

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/utils"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabaseGorm(ctx context.Context) *gorm.DB {
	opts := &gorm.Config{
		// QueryFields:     true, // Selecionar por campos // SELECT * FROM `users` // without this option
		// PrepareStmt:     true, // ria instruções preparadas ao executar qualquer SQL e as armazena em cache para acelerar chamadas futuras
		// CreateBatchSize: 1000,
		Logger: logger.Default.LogMode(logger.Silent), // Silent // Info
		//SkipDefaultTransaction: true, // Ignorar transações padrão // executa operações de gravação dentro de uma transação para garantir a consistência dos dados.
		//DisableAutomaticPing:                     true,
		//DisableForeignKeyConstraintWhenMigrating: true,
		//DryRun:                                   false, // Gera SQLsem executar.
	}
	db, err := gorm.Open(postgres.Open(utils.DBUri), opts)
	if err != nil {
		slog.Error("Failed to connect to database", slog.Any("error", err))
		return nil
	}

	// db.AutoMigrate(&models.PersonModel{})
	// db.AutoMigrate(&models.UserModel{}, &models.ProductModel{}, &models.OrderModel{})

	// gorm.conn.DB()
	return db
}

func CloseDatabaseGorm(ctx context.Context, db *gorm.DB) {
	if db != nil {
		// db.Close()
	}
}

func GetAllGorm(ctx context.Context, db *gorm.DB, models *[]models.PersonModel) *gorm.DB {
	return db.Find(&models)
}

func FetchAllGorm(ctx context.Context) {
	slog.Info("Connecting to database using GORM...")

	db := NewDatabaseGorm(ctx)
	defer CloseDatabaseGorm(ctx, db)

	var persons []models.PersonModel
	// err := db.Find(&persons).Error
	err := GetAllGorm(ctx, db, &persons).Error
	if err != nil {
		slog.Error("Failed to query persons from database", slog.Any("error", err))
		return
	}

	for _, p := range persons {
		slog.Info("PersonModel", "person", p.Name)
	}

	slog.Info("Persons retrieved from database", slog.Int("count", len(persons)))
}

func ClearGorm(ctx context.Context) {
	db := NewDatabaseGorm(ctx)
	defer CloseDatabaseGorm(ctx, db)

	db.Exec(utils.DBDeleteGorm)
}

func AddOneGorm(ctx context.Context, count int) time.Duration {
	db := NewDatabaseGorm(ctx)
	defer CloseDatabaseGorm(ctx, db)

	start := time.Now()
	for i := 0; i < count; i++ {
		row := models.NewPerson("Gorm AddOne #" + strconv.Itoa(i))
		db.Create(&row)
	}

	return time.Since(start)
}

func AddManyGorm(ctx context.Context, count int) time.Duration {
	db := NewDatabaseGorm(ctx)
	defer CloseDatabaseGorm(ctx, db)

	start := time.Now()

	rows := make([]*models.PersonModel, 0, count)
	for i := 0; i < count; i++ {
		row := models.NewPerson("Gorm1 AddMany #" + strconv.Itoa(i))
		rows = append(rows, row)
	}

	tx := db.Begin()
	//tx := db.Session(&gorm.Session{PrepareStmt: true, CreateBatchSize: count})
	defer tx.Rollback()

	for i := 0; i < count; i++ {
		tx.Create(&rows[i])
	}

	tx.Commit()
	return time.Since(start)
}

func AddManyGorm2(ctx context.Context, count int) time.Duration {
	db := NewDatabaseGorm(ctx)
	defer CloseDatabaseGorm(ctx, db)

	start := time.Now()

	rows := make([]*models.PersonModel, 0, count)
	for i := 0; i < count; i++ {
		row := models.NewPerson("Gorm2 AddMany #" + strconv.Itoa(i))
		rows = append(rows, row)
	}

	//tx := db.Begin()
	tx := db.Session(&gorm.Session{PrepareStmt: true, CreateBatchSize: count})
	defer tx.Rollback()

	tx.Create(&rows)

	tx.Commit()

	return time.Since(start)
}

func AddManyGorm3(ctx context.Context, count int) time.Duration {
	db := NewDatabaseGorm(ctx)
	defer CloseDatabaseGorm(ctx, db)

	start := time.Now()

	query := ""
	args := []interface{}{}
	num := 1

	for i := 0; i < count; i++ {
		row := models.NewPerson("Gorm3 AddMany #" + strconv.Itoa(i))

		query = fmt.Sprintf("%s($%d, $%d, $%d),", query, num, num+1, num+2)
		args = append(args, row.ID, row.Name, row.CreatedAt)
		num += 3
	}

	query = `INSERT INTO "TbPerson" ("id", "name", "created_at") VALUES ` + query
	query = query[0 : len(query)-1]

	//tx := db.Begin()
	tx := db.Session(&gorm.Session{PrepareStmt: true, CreateBatchSize: count})
	defer tx.Rollback()

	tx.Exec(query, args...)

	tx.Commit()

	return time.Since(start)
}
