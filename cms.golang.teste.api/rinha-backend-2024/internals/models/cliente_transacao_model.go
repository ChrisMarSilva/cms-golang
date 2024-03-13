package models

import (
	"fmt"
	"time"
)

type ClienteTransacao struct {
	IdTransacao  int32     `db:"id"`
	IdCliente    int       `db:"cliente_id"`
	Valor        int64     `db:"valor"`
	Tipo         string    `db:"tipo"`
	Descricao    string    `db:"descricao"`
	DtHrRegistro time.Time `db:"dthrregistro"`
}

func (row ClienteTransacao) ToString() string {
	return fmt.Sprintf("valor: %s; tipo: %s; descricao: %s; data: %s", row.Valor, row.Tipo, row.Descricao, row.DtHrRegistro.Format("2006-01-02T15:04:05.000000Z"))
}

type Job struct {
	Name          string
	Transacao     *ClienteTransacao
	TransacaoJson []byte
}

type ParseTransactionResult struct {
	Transacao *ClienteTransacao
	Error     error
}
