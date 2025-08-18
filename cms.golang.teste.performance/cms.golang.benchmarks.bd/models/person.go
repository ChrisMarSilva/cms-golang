package models

import (
	"log/slog"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

var (
	faker = gofakeit.New(0)
)

type PersonModel struct {
	ID        uuid.UUID `db:"id" gorm:"primaryKey"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at" gorm:"index"`
}

func NewPerson(name string) *PersonModel {
	return &PersonModel{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}
}

func (PersonModel) TableName() string {
	return "TbPerson"
}

func GenerateFakePerson() *PersonModel {
	return NewPerson(faker.Name())
}

func GenerateFakePersons(count int) []*PersonModel {
	persons := make([]*PersonModel, 0, count)
	for i := 0; i < count; i++ {
		persons = append(persons, GenerateFakePerson())
	}

	return persons
}

func FetchAll() {
	person := GenerateFakePerson()
	slog.Info("PersonModel", "person", person.Name)

	persons := GenerateFakePersons(5)
	for i, p := range persons {
		slog.Info("PersonModel", "index", i, "person", p.Name)
	}
}
