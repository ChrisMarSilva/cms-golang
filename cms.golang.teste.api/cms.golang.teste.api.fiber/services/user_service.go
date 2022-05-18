package service

import (
	"context"
	"errors"

	entity "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/entities"
	repository "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/repositories"
	"github.com/google/uuid"
	opentracing "github.com/opentracing/opentracing-go"
	spanlog "github.com/opentracing/opentracing-go/log"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo repository.IUserRepository
	logRepo  repository.LogMonitorRepositoryMSSQL
}

func NewUserService(userRepo repository.IUserRepository, logRepo repository.LogMonitorRepositoryMSSQL) *UserService {
	return &UserService{
		userRepo: userRepo,
		logRepo:  logRepo,
	}
}

func (service *UserService) GetAll(ctx context.Context) (users []entity.User, err error) {

	metodo := "UserService.GetAll"

	sp, ctx := opentracing.StartSpanFromContext(ctx, metodo)
	defer sp.Finish()

	err = service.userRepo.GetAll(ctx, &users)
	if err != nil {
		service.logRepo.Inserir(metodo, "1", "Erro: "+err.Error())
		sp.SetTag("error", true)
		sp.LogFields(spanlog.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.User{}, err
		}
		return []entity.User{}, err
	}

	return users, nil
}

func (service *UserService) Get(ctx context.Context, id uuid.UUID) (user entity.User, err error) {

	metodo := "UserService.Get"

	sp, ctx := opentracing.StartSpanFromContext(ctx, metodo)
	defer sp.Finish()

	err = service.userRepo.Get(ctx, &user, id)
	if err != nil {
		service.logRepo.Inserir(metodo, "1", "Erro: "+err.Error())
		sp.SetTag("error", true)
		sp.LogFields(spanlog.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, err
		}
		return entity.User{}, err
	}

	return user, nil
}

func (service *UserService) Create(ctx context.Context, user entity.User) (err error) {

	metodo := "UserService.Create"

	sp, ctx := opentracing.StartSpanFromContext(ctx, metodo)
	defer sp.Finish()

	err = user.Validate()
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(spanlog.Error(err))
		return err
	}

	err = service.userRepo.Create(ctx, &user)
	if err != nil {
		service.logRepo.Inserir(metodo, "1", "Erro: "+err.Error())
		sp.SetTag("error", true)
		sp.LogFields(spanlog.Error(err))
		return err
	}

	return nil
}

func (service *UserService) Update(ctx context.Context, id uuid.UUID, body entity.User) (err error) {

	metodo := "UserService.Update"

	sp, ctx := opentracing.StartSpanFromContext(ctx, metodo)
	defer sp.Finish()

	sp.SetTag("body.New", body)

	user, err := service.Get(ctx, id)
	if err != nil {
		return err
	}

	sp.SetTag("user.Old", user)

	// user.ID = body.ID
	user.Nome = body.Nome
	//user.Status = body.Status

	err = user.Validate()
	if err != nil {
		sp.SetTag("error", true)
		sp.LogFields(spanlog.Error(err))
		return err
	}

	err = service.userRepo.Update(ctx, &user)
	if err != nil {
		service.logRepo.Inserir(metodo, "1", "Erro: "+err.Error())
		sp.SetTag("error", true)
		sp.LogFields(spanlog.Error(err))
		return err
	}

	return nil
}

func (service *UserService) Delete(ctx context.Context, id uuid.UUID) (err error) {

	metodo := "UserService.Delete"

	sp, ctx := opentracing.StartSpanFromContext(ctx, metodo)
	defer sp.Finish()

	user, err := service.Get(ctx, id)
	if err != nil {
		return err
	}

	sp.SetTag("user.Old", user)

	err = service.userRepo.Delete(ctx, id)
	if err != nil {
		service.logRepo.Inserir(metodo, "1", "Erro: "+err.Error())
		sp.SetTag("error", true)
		sp.LogFields(spanlog.Error(err))
		return err
	}

	return nil
}
