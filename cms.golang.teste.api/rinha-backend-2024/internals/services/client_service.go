package services

import (
	"context"
	"errors"
	"time"

	"github.com/chrismarsilva/rinha-backend-2024/internals/dtos"
	"github.com/chrismarsilva/rinha-backend-2024/internals/models"
	"github.com/chrismarsilva/rinha-backend-2024/internals/repositories"

	//"github.com/jmoiron/sqlx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IClientService interface {
	CreateTransaction(id int, request dtos.TransacaoRequestDto) (dtos.TransacaoResponseDto, error)
	GetExtract(id int) (dtos.ExtratoResponseDto, error)
}

type ClientService struct {
	db                    *pgxpool.Pool
	clientRepo            repositories.IClientRepository
	clientTransactionRepo repositories.IClientTransactionRepository
}

func NewClientService(db *pgxpool.Pool, clientRepo repositories.IClientRepository, clientTransactionRepo repositories.IClientTransactionRepository) *ClientService {
	return &ClientService{
		db:                    db,
		clientRepo:            clientRepo,
		clientTransactionRepo: clientTransactionRepo,
	}
}

func (s *ClientService) CreateTransaction(id int, request dtos.TransacaoRequestDto) (dtos.TransacaoResponseDto, error) {

	// conn := database.GetConnection()
	// defer conn.Close()

	var transacao dtos.TransacaoResponseDto

	// s.db.Exec("SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;")

	var cliente models.Cliente
	err := s.clientRepo.Get(&cliente, id)
	if err != nil {
		return transacao, err
	}

	if request.Tipo == "d" {
		cliente.Saldo += cliente.Limite - request.Valor
		if cliente.Saldo < 0 {
			return transacao, errors.New("Novo saldo do cliente menor que seu limite disponível.")
		}
		cliente.Saldo -= cliente.Limite
	} else {
		cliente.Saldo += request.Valor
	}

	// if request.Tipo == "d" && cliente.Saldo < -cliente.Limite {
	// 	return nil, ErroTransacaoDebito
	// }

	ctx := context.Background()

	//go func() {
	//tx := s.db.MustBegin()
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return transacao, err
	}

	// defer func() {
	// 	err = txn.Rollback()
	// 	if err != nil {
	// 		if !errors.Is(err, sql.ErrTxDone) {
	// 			log.Println("Error txn.Rollback(): "  +err.Error())
	// 		}
	// 	}
	// }()

	err = s.clientRepo.UpdSaldo(tx, id, request.Valor, request.Tipo)
	if err != nil {
		tx.Rollback(ctx)
		return transacao, err
	}

	err = s.clientTransactionRepo.Add(tx, id, request.Valor, request.Tipo, request.Descricao)
	if err != nil {
		tx.Rollback(ctx)
		return transacao, err
	}

	tx.Commit(ctx)
	//}()

	transacao = dtos.TransacaoResponseDto{
		Limite: cliente.Limite,
		Saldo:  cliente.Saldo,
	}

	return transacao, nil
}

func (s *ClientService) GetExtract(id int) (dtos.ExtratoResponseDto, error) {
	var extrato dtos.ExtratoResponseDto

	// s.db.Exec("SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED;")

	var cliente models.Cliente
	err := s.clientRepo.Get(&cliente, id)
	if err != nil {
		return extrato, err
	}

	clienteTransacoes := map[int]models.ClienteTransacao{}
	err = s.clientTransactionRepo.GetAll(&clienteTransacoes, id)
	if err != nil {
		return extrato, err
	}

	transacoes := make([]dtos.ExtratoTransacoesResponseDto, 0, 100) // transacoes := make([]dtos.ExtratoTransacoesResponseDto, len(clienteTransacoes))

	for _, value := range clienteTransacoes {
		transacao := dtos.ExtratoTransacoesResponseDto{
			Valor:       value.Valor,
			Tipo:        value.Tipo,
			Descricao:   value.Descricao,
			RealizadaEm: value.DtHrRegistro.Format("2006-01-02T15:04:05.000000Z"),
		}

		transacoes = append(transacoes, transacao) //transacoes[key] = transacao
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
