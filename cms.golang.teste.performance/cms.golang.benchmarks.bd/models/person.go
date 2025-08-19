package models

import (
	"errors"
	"log/slog"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
)

var (
	faker = gofakeit.New(0)
)

// const (
// 	PersonActive   string = "A"
// 	PersonInactive        = "I"
// )

type PersonModel struct {
	ID   uuid.UUID `db:"id" gorm:"column:id; type:uuid; primaryKey"`
	Name string    `db:"name" gorm:"column:name"`
	//Status    string    `db:"status" gorm:"column:status"`
	CreatedAt time.Time `db:"created_at" gorm:"column:created_at; index"`
}

func NewPerson(name string) *PersonModel {
	return &PersonModel{
		ID:   uuid.New(),
		Name: name,
		//Status : PersonActive,
		CreatedAt: time.Now().UTC(),
	}
}

func (PersonModel) TableName() string {
	return "TbPerson"
}

func (p PersonModel) ToString() string {
	return "PersonModel: id=" + p.ID.String() +
		", name=" + p.Name +
		// ", status=" + p.Status +
		// ", status descr.=" + p.StatusString() +
		", created_at=" + p.CreatedAt.Format(time.RFC3339)
}

// func (p PersonModel) StatusString() string {
// 	switch p.Status {
// 	case PersonActive:
// 		return "Active"
// 	case PersonInactive:
// 		return "Inactive"
// 	}
// 	return "unknown"
// }

func (row PersonModel) Validate() error {
	if row.ID == uuid.Nil {
		return errors.New("id is NIL")
	}

	if row.ID.String() == "" {
		return errors.New("id is required")
	}

	if row.Name == "" {
		return errors.New("name is required")
	}

	// if row.Status == "" {
	// 	return errors.New("status is required")
	// }

	return nil
}

func GenerateFakePerson() *PersonModel {
	return NewPerson(faker.Name())
}

func GenerateFakePersons(prefix string, count int) []*PersonModel {
	rows := make([]*PersonModel, count)

	for i := 0; i < count; i++ {
		//rows = append(rows, GenerateFakePerson())
		rows[i] = NewPerson(prefix + " #" + strconv.Itoa(i))
	}

	return rows
}

func FetchAll() {
	person := GenerateFakePerson()
	slog.Info("PersonModel", "person", person.Name)

	persons := GenerateFakePersons("FetchAll", 5)
	for i, p := range persons {
		slog.Info("PersonModel", "index", i, "person", p.Name)
	}
}
