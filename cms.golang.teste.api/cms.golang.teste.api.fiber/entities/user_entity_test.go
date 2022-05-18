package entity_test

import (
	"testing"

	entity "github.com/ChrisMarSilva/cms-golang-teste-api-fiber/entities"
	"github.com/google/uuid"
)

// // go test
// // go test -v
// // go test -bench=.
// // go test -bench=Add
// go test -run TestUserrepoGetAll -v
// go test -run=XXX -bench . -benchmem

// log.Println(uuid.New().String())
// log.Println(uuid.NewString())
// id := uuid.New()
// log.Println(id.String())

func TestUser(t *testing.T) {

	user := entity.User{ID: uuid.New(), Nome: "PESSOA 1", Status: entity.UserActive}
	userEmpty := entity.User{}

	if user == userEmpty {
		t.Fatalf("user null ")
	}

	err := user.Validate()
	if err != nil {
		t.Fatalf("error %t", err)
	}

}

// func TestUserValidate(t *testing.T) {
// 	t.Parallel()

// 	tests := []struct {
// 		name    string
// 		input   entity.User
// 		withErr bool
// 	}{
// 		{"OK", entity.User{ID: uuid.New(), Nome: "PESSOA 1", Status: entity.UserActive}, false},
// 		{"ERR: Nome", entity.User{ID: uuid.New(), Nome: "", Status: entity.UserActive}, true},
// 		{"ERR: Status", entity.User{ID: uuid.New(), Nome: "PESSOA 3", Status: ""}, true},
// 		{"ERR: ID", entity.User{ID: uuid.New(), Nome: "PESSOA ", Status: entity.UserActive}, true},
// 	}

// 	for _, tt := range tests {
// 		tt := tt
// 		t.Run(tt.name, func(t *testing.T) {
// 			t.Parallel()
// 			if actualErr := tt.input.Validate(); (actualErr != nil) != tt.withErr {
// 				t.Fatalf("expected error %t, got %s", tt.withErr, actualErr)
// 			}
// 		})
// 	}
// }

func BenchmarkUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := entity.User{ID: uuid.New(), Nome: "PESSOA 1", Status: entity.UserActive}
		if err := user.Validate(); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkUuid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid := uuid.New()
		if uuid.String() == "" {
			b.Fatal("invalid uuid")
		}
	}
}

func BenchmarkUuidString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid := uuid.New().String()
		if uuid == "" {
			b.Fatal("invalid uuid")
		}
	}
}

func BenchmarkUuidNewString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid := uuid.NewString()
		if uuid == "" {
			b.Fatal("invalid uuid")
		}
	}
}
