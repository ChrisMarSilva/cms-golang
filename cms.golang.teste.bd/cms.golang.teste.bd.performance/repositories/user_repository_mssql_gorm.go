package repository

import (
	"context"

	entity "github.com/ChrisMarSilva/cms.golang.teste.bd.performance/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositorySQLServerGorm struct {
	db *gorm.DB
	// WithTrx() UserRepositorySQLServerGorm
	// CommitTrx() error
	// RollbackTrx()
}

func NewUserRepositorySQLServerGorm(db *gorm.DB) *UserRepositorySQLServerGorm {
	return &UserRepositorySQLServerGorm{db: db}
}

// func (repo UserRepositorySQLServerGorm) WithTrx() UserRepositorySQLServerGorm {
// 	return &UserRepositorySQLServerGorm{db: repo.db.Begin()}
// }

// trx := repoUser.WithTrx()
// err := trx.Create(ctx, name)
// err := trx.CommitTrx()

func (repo UserRepositorySQLServerGorm) GetAll(ctx context.Context, users *[]entity.User) (err error) {
	err = repo.db.Find(users).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo UserRepositorySQLServerGorm) Get(ctx context.Context, user *entity.User, id uuid.UUID) (err error) {
	err = repo.db.First(user, "ID = ?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo UserRepositorySQLServerGorm) Create(ctx context.Context, user *entity.User) (err error) {
	err = repo.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo UserRepositorySQLServerGorm) CreateInBatch(ctx context.Context, users []entity.User) (err error) {

	tx := repo.db.Begin()

	for _, user := range users {
		err = tx.Create(user).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func (repo UserRepositorySQLServerGorm) Update(ctx context.Context, user *entity.User) (err error) {
	repo.db.Save(user)
	return nil
}

func (repo UserRepositorySQLServerGorm) Delete(ctx context.Context, id uuid.UUID) (err error) {
	err = repo.db.Where("ID = ?", id).Delete(&entity.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
