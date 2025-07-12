package usecases

import (
	"context"
	"errors"
	"log/slog"

	"github.com/chrismarsilva/rinha-backend-2025/internals/dtos"
	"github.com/chrismarsilva/rinha-backend-2025/internals/models"
	"github.com/chrismarsilva/rinha-backend-2025/internals/repositories"
	"github.com/google/uuid"
)

type UseCases struct {
	repos *repositories.Repositories
}

func New(repos *repositories.Repositories) *UseCases {
	return &UseCases{
		repos: repos,
	}
}

func (u UseCases) GetAll(ctx context.Context) ([]dtos.UserResponseDto, error) {
	users, err := u.repos.User.GetAll(ctx)
	if err != nil {
		slog.Error("Error fetching users: ", err)
		return nil, err
	}

	response := make([]dtos.UserResponseDto, 0)

	if len(users) > 0 {
		for _, user := range users {
			userDto := dtos.UserResponseDto{ID: user.ID, Name: user.Name, Email: user.Email}
			response = append(response, userDto)
		}
	}

	return response, nil
}

func getNewUuid(resultChan chan uuid.UUID) {
	resultChan <- uuid.New()
}

func (u UseCases) Create(ctx context.Context, request dtos.UserRequestDto) (dtos.UserResponseDto, error) {
	uuidChannel := make(chan uuid.UUID)
	go getNewUuid(uuidChannel)

	user, err := u.repos.User.GetByEmail(ctx, request.Email)
	if err != nil {
		slog.Error("User with error this email: ", "email", request.Email)
		return dtos.UserResponseDto{}, err
	}

	// if user.ID == uuid.Nil {http.StatusNotFound}
	if user != nil {
		slog.Error("User already exists with this email", "email", request.Email)
		return dtos.UserResponseDto{}, errors.New("User already exists with this email")
	}

	model := models.User{
		ID:    <-uuidChannel, /// uuid.New(),
		Name:  request.Name,
		Email: request.Email}

	err = u.repos.User.Create(ctx, model)
	if err != nil {
		return dtos.UserResponseDto{}, err
	}

	response := dtos.UserResponseDto{ID: model.ID, Name: model.Name, Email: model.Email}
	return response, nil
}
