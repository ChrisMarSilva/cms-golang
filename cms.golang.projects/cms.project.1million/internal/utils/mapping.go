package utils

import (
	"github.com/chrismarsilva/cms.project.1million/internal/dtos"
	"github.com/chrismarsilva/cms.project.1million/internal/models"
)

func MapToResponseDto(model *models.PersonModel) *dtos.PersonResponseDto {
	return dtos.NewPersonResponseDto(model.ID, model.Name, model.CreatedAt)
}

func MapToModel(dto *dtos.PersonRequestDto) *models.PersonModel {
	return models.NewPersonModel(dto.Name)
}
