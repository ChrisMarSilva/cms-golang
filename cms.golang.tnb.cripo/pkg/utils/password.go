package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // bcrypt.DefaultCost
	return string(bytes), err
}

func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type IHashAdapter interface {
	GenerateHash(str string) (string, error)
	CheckHash(hash, str string) bool
}

type HashAdapter struct {
}

func NewHashAdapter() IHashAdapter {
	return &HashAdapter{}
}

func (h *HashAdapter) GenerateHash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 5)
	return string(bytes), err
}

func (h *HashAdapter) CheckHash(hash, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
