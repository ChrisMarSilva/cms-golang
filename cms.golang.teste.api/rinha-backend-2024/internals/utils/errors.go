package utils

import (
	"errors"
)

var (
	// ErrInvalidEmail       = errors.New("invalid email")
	// ErrEmailAlreadyExists = errors.New("email already exists")
	// ErrEmptyPassword      = errors.New("password can't be empty")
	// ErrInvalidAuthToken   = errors.New("invalid auth-token")
	// ErrInvalidCredentials = errors.New("invalid credentials")
	// ErrUnauthorized       = errors.New("Unauthorized")

	// ErrApelidoJaExiste = errors.New("apelido já existe")

	ErrClienteNaoExiste = errors.New("cliente não existe")
	ErrValorDaTransacao = errors.New("o valor da transação deve ser maior que zero")
	ErrTipoDaTransacao  = errors.New("a transação deve ser do tipo 'c' (crédito) ou 'd' (débito)")
	ErrDescricao        = errors.New("a descrição deve ter de 1 a 10 caractéres")
	ErrTransacaoDebito  = errors.New("a transação do tipo 'd' (débito) nunca pode deixar o saldo do cliente menor que seu limite disponível")
)
