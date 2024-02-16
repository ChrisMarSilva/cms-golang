package services

import (
	"errors"
	"time"

	"github.com/chrismarsilva/rinha-backend-2024/internals/dtos"
	"github.com/chrismarsilva/rinha-backend-2024/internals/models"
	"github.com/chrismarsilva/rinha-backend-2024/internals/repositories"
)

type IClientService interface {
	CreateTransaction(id int, request dtos.TransacaoRequestDto) (dtos.TransacaoResponseDto, error)
	GetExtract(id int) (dtos.ExtratoResponseDto, error)
}

type ClientService struct {
	clientRepo            repositories.IClientRepository
	clientTransactionRepo repositories.IClientTransactionRepository
}

func NewClientService(clientRepo repositories.IClientRepository, clientTransactionRepo repositories.IClientTransactionRepository) *ClientService {
	return &ClientService{clientRepo: clientRepo, clientTransactionRepo: clientTransactionRepo}
}

func (s *ClientService) CreateTransaction(id int, request dtos.TransacaoRequestDto) (dtos.TransacaoResponseDto, error) {
	var transacao dtos.TransacaoResponseDto

	var cliente models.Cliente
	err := s.clientRepo.Get(&cliente, id)
	if err != nil {
		return transacao, err
	}

	var novoSado int64 = 0

	if request.Tipo == "d" {
		novoSado = cliente.Saldo + cliente.Limite - request.Valor
	} else {
		novoSado = cliente.Saldo + cliente.Limite + request.Valor
	}

	if novoSado < 0 {
		return transacao, errors.New("Novo saldo do cliente menor que seu limite disponível.")
	}

	novoSado -= cliente.Limite

	// tx, err := db.Begin() //tx := db.MustBegin()
	// if err != nil {
	// 	return transacao, err
	// }

	err = s.clientRepo.UpdSaldo(id, request.Valor, request.Tipo)
	if err != nil {
		// tx.Rollback()
		return transacao, err
	}

	err = s.clientTransactionRepo.Add(id, request.Valor, request.Tipo, request.Descricao)
	if err != nil {
		// tx.Rollback()
		return transacao, err
	}

	//err = tx.Commit() 
	// if err != nil {
	// 	// tx.Rollback()
	// 	return transacao, err
	// }

	transacao = dtos.TransacaoResponseDto{
		Limite: cliente.Limite,
		Saldo:  novoSado,
	}

	return transacao, nil
}

func (s *ClientService) GetExtract(id int) (dtos.ExtratoResponseDto, error) {
	var extrato dtos.ExtratoResponseDto

	var cliente models.Cliente
	err := s.clientRepo.Get(&cliente, id)
	if err != nil {
		return extrato, err
	}

	clienteTransacoes := []models.ClienteTransacao{}
	err = s.clientTransactionRepo.GetAll(&clienteTransacoes, id)
	if err != nil {
		return extrato, err
	}

	transacoes := []dtos.ExtratoTransacoesResponseDto{}

	for _, value := range clienteTransacoes {
		transacao := dtos.ExtratoTransacoesResponseDto{
			Valor:       value.Valor,
			Tipo:        value.Tipo,
			Descricao:   value.Descricao,
			RealizadaEm: value.DtHrRegistro.Format("2006-01-02T15:04:05.000000Z"),
		}

		transacoes = append(transacoes, transacao)
	}

	extrato = dtos.ExtratoResponseDto{
		Saldo: dtos.ExtratoSaldoResponseDto{
			Total:       cliente.Saldo,
			DataExtrato: time.Now().Format("2006-01-02T15:04:05.000000Z"),
			Limite:      cliente.Limite,
		},
		Transacoes: transacoes,
	}

	return extrato, nil
}
