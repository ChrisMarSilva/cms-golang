package dtos

type ExtratoResponseDto struct {
	Saldo      ExtratoSaldoResponseDto        `json:"saldo"`
	Transacoes []ExtratoTransacoesResponseDto `json:"transacoes"`
}

// func NewExtratoResponseDto(saldo ExtratoSaldoResponseDto, transacoes *[]ExtratoTransacoesResponseDto) *ExtratoResponseDto {
// 	return &ExtratoResponseDto{
// 		Saldo:      saldo,
// 		Transacoes: transacoes,
// 	}
// }

type ExtratoSaldoResponseDto struct {
	Total       int64  `json:"total"`
	DataExtrato string `json:"data_extrato"`
	Limite      int64  `json:"limite"`
}

// func NewExtratoSaldoResponseDto(total int64, limite int64) *ExtratoSaldoResponseDto {
// 	return &ExtratoSaldoResponseDto{
// 		Total:      total,
// 		DataExtrato: time.Now(),
// 		Limite: limite,
// 	}
// }

type ExtratoTransacoesResponseDto struct {
	Valor       int64  `json:"valor,omitempty"`
	Tipo        string `json:"tipo,omitempty"`
	Descricao   string `json:"descricao,omitempty"`
	RealizadaEm string `json:"realizada_em,omitempty"`
}

// func NewExtratoTransacoesResponseDto(valor int64, tipo string, descricao string, realizadaEm time.Time) *ExtratoTransacoesResponseDto {
// 	return &ExtratoTransacoesResponseDto{
// 		Valor:      valor,
// 		Tipo:     tipo,
// 		Descricao: descricao,
// 		RealizadaEm: realizadaEm,
// 	}
// }

// type Date struct {
// 	time.Time
// }

// func (t Date) MarshalJSON() (string, error) {
// 	return t.Time.Format(`"2006-01-02T15:04:05.000000Z"`), nil
// }

// func (t *Date) UnmarshalJSON(b []byte) (err error) {
// 	date, err := time.Parse(`"2006-01-02"`, string(b))
// 	if err != nil {
// 		return fmt.Errorf("error on converting date: %w", err)
// 	}
// 	t.Time = date
// 	return
// }
