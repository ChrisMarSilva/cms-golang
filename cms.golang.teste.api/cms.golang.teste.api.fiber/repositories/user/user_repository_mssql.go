package user

import (
	"context"

	entity "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/entities"
	repository "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/repositories"
	"github.com/google/uuid"
	opentracing "github.com/opentracing/opentracing-go"
	spanlog "github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

type UserRepositoryMSSQL struct {
	ctx     context.Context
	db      *gorm.DB
	logRepo repository.LogMonitorRepositoryMSSQL
}

func NewUserRepositoryMSSQL(ctx context.Context, db *gorm.DB, logRepo repository.LogMonitorRepositoryMSSQL) *UserRepositoryMSSQL {
	return &UserRepositoryMSSQL{
		ctx:     ctx,
		db:      db,
		logRepo: logRepo,
	}
}

func (repo UserRepositoryMSSQL) GetAll(ctx context.Context, users *[]entity.User) (err error) {

	metodo := "UserRepositoryMSSQL.GetAll"

	sp, ctx := opentracing.StartSpanFromContext(ctx, metodo)
	defer sp.Finish()

	err = repo.db.Find(users).Error
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(spanlog.Error(err))
		query := repo.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Find(users)
		})
		sp.SetTag("query", query)
		repo.logRepo.Inserir(metodo, "1", "query: "+query+"; Erro: "+err.Error())
		return err
	}

	return nil
}

func (repo UserRepositoryMSSQL) Get(ctx context.Context, user *entity.User, id uuid.UUID) (err error) {

	metodo := "UserRepositoryMSSQL.Get"

	sp, ctx := opentracing.StartSpanFromContext(ctx, metodo)
	defer sp.Finish()

	err = repo.db.First(user, "ID = ?", id).Error

	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(spanlog.Error(err))
		query := repo.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.First(user, "ID = ?", id)
		})
		sp.SetTag("query", query)
		repo.logRepo.Inserir(metodo, "1", "query: "+query+"; Erro: "+err.Error())
		return err
	}

	return nil
}

func (repo UserRepositoryMSSQL) Create(ctx context.Context, user *entity.User) (err error) {

	metodo := "UserRepositoryMSSQL.Create"

	sp, ctx := opentracing.StartSpanFromContext(ctx, metodo)
	defer sp.Finish()

	err = repo.db.Create(user).Error

	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(spanlog.Error(err))
		query := repo.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Create(user)
		})
		sp.SetTag("query", query)
		repo.logRepo.Inserir(metodo, "1", "query: "+query+"; Erro: "+err.Error())
		return err
	}

	return nil
}

func (repo UserRepositoryMSSQL) Update(ctx context.Context, user *entity.User) (err error) {

	metodo := "UserRepositoryMSSQL.Update"

	sp, ctx := opentracing.StartSpanFromContext(ctx, metodo)
	defer sp.Finish()

	err = repo.db.Save(user).Error

	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(spanlog.Error(err))
		query := repo.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Save(user)
		})
		sp.SetTag("query", query)
		repo.logRepo.Inserir(metodo, "1", "query: "+query+"; Erro: "+err.Error())
		return err
	}

	return nil
}

func (repo UserRepositoryMSSQL) Delete(ctx context.Context, id uuid.UUID) (err error) {

	metodo := "UserRepositoryMSSQL.Delete"

	sp, ctx := opentracing.StartSpanFromContext(ctx, metodo)
	defer sp.Finish()

	//err = u.db.Delete(&user).Error
	//err = u.db.Delete(&user, "ID = ?", id).Error
	//err = u.db.Where("ID = ?", id).Delete(user).Error
	//err = u.db.First(user, "ID = ?", id).Error
	err = repo.db.Where("ID = ?", id).Delete(&entity.User{}).Error
	// err = u.db.Delete(entity.User{}, "ID = ?", id).Error

	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(spanlog.Error(err))
		query := repo.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
			return tx.Where("ID = ?", id).Delete(&entity.User{})
		})
		sp.SetTag("query", query)
		repo.logRepo.Inserir(metodo, "1", "query: "+query+"; Erro: "+err.Error())
		return err
	}

	return nil
}
