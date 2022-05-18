package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

// go mod init github.com/ChrisMarSilva/cms.golang.tnb.bd.backup
// go get -u github.com/go-sql-driver/mysql
// go get -u github.com/denisenkom/go-mssqldb
// go get -u -v github.com/simukti/sqldb-logger
// go get -u github.com/simukti/sqldb-logger/logadapter/zerologadapter
// go mod tidy

// go run main.go

func main() {

	log.Println("INI")
	var start time.Time = time.Now()

	base := NewBase()
	base.CarregarTabelas()
	base.CarregarDeletes() // Deletes :   17.3904746s // antes era 4m
	base.CarregarCampos()  // Campos  :   43.6787677s // antes era 1m
	base.CarregarInserts() // Inserts : 6m43.7069266s // antes era 14m

	log.Println("FIM:", time.Since(start)) // FIM: 6m43.7069266s // antes era 19m
}

type Base struct {
	Tabelas []Tabela
}

func NewBase() *Base {
	return &Base{}
}

func (b *Base) CarregarTabelas() {
	b.Tabelas = make([]Tabela, 0)
	b.Tabelas = append(b.Tabelas, *NewTabela("TBSITUACAO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBBDR_EMPRESA_SEGMENTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBBDR_EMPRESA_SETOR"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBBDR_EMPRESA_SUBSETOR"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_SEGMENTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_SETOR"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_SUBSETOR"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBFII_FUNDOIMOB_TIPO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBUSUARIO_CONFIG_TIPO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBUSUARIO_EMAIL_INFO_PERIODO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCONFIG"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_ATIVO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_ATIVOCOTACAO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_ATIVOINDICADOR"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_FATORELEVANTE"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_PROVENTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_FINAN"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_FINAN_AGENDA"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_FINAN_BPA_ANUAL"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_FINAN_BPA_TRIMESTRAL"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_FINAN_BPP_ANUAL"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_FINAN_BPP_TRIMESTRAL"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_FINAN_DFC_ANUAL"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_FINAN_DFC_TRIMESTRAL"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_FINAN_DRE_ANUAL"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_FINAN_DRE_TRIMESTRAL"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBFII_FUNDOIMOB_ADMIN"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBFII_FUNDOIMOB"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBFII_FUNDOIMOB_COTACAO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBFII_FUNDOIMOB_FATORELEVANTE"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBFII_FUNDOIMOB_PROVENTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBETF_INDICE"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBETF_INDICE_COTACAO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBETF_FATORELEVANTE"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBBDR_EMPRESA"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBBDR_EMPRESA_COTACAO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBBDR_EMPRESA_FATORELEVANTE"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBBDR_EMPRESA_PROVENTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCRIPTO_EMPRESA"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBUSUARIO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBUSUARIO_LOG"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBUSUARIO_HASH"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBUSUARIO_CONFIG"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBLOGERRO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_10"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_188"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_2"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_212"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_227"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_231"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_238"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_26"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_261"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_321"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_357"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_358"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_382"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_384"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_391"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_396"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_404"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_410"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_411"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_60"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_64"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_67"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_68"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_OPER_USER_7"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_188"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_2"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_212"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_227"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_231"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_238"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_26"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_261"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_321"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_357"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_358"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_382"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_384"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_391"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_404"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_410"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_411"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_60"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_64"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_67"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_68"))
	// b.Tabelas = append(b.Tabelas, *NewTabela("TBCEI_PROV_USER_7"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCORRETORA_LISTA"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCORRETORA"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCOMENTARIO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCOMENTARIO_DENUNCIA"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCOMENTARIO_REACAO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCOMENTARIO_ALERTA"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBALERTA"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBALERTA_ASSINATURA"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBALERTA_NOTICIA"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBAPURACAO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBAPURACAO_CALCULADA"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBALUGUEL_ATIVO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBLANCAMENTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBOPERACAO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBPROVENTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBFII_LANCAMENTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBFII_PROVENTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBETF_LANCAMENTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBETF_OPERACAO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBBDR_LANCAMENTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBBDR_OPERACAO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBBDR_PROVENTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCRIPTO_LANCAMENTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBUSUARIO_ACOMP_GRUPO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBUSUARIO_ACOMP_ATIVO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBUSUARIO_ACOMP_BDR"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBUSUARIO_ACOMP_CRIPTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBUSUARIO_ACOMP_FUNDO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBUSUARIO_ACOMP_INDICE"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCARTEIRA"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCARTEIRA_ATIVO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCARTEIRA_BDR"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCARTEIRA_COTAS"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCARTEIRA_CRIPTO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCARTEIRA_FUNDO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCARTEIRA_INDICE"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCARTEIRA_PROJECAO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBCARTEIRA_PROJECAO_ITEM"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBUSUARIO_EMAIL_INFO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBUSUARIO_EMAIL_INFO_CONTEUDO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBFII_FUNDOIMOB_FATORELEVANTE_ATIVO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBFII_FUNDOIMOB_PROVENTO_ATIVO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_FATORELEVANTE_ATIVO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBEMPRESA_PROVENTO_ATIVO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBETF_FATORELEVANTE_ATIVO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBBDR_EMPRESA_FATORELEVANTE_ATIVO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBBDR_EMPRESA_PROVENTO_ATIVO"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBNOTA_CORRETAGEM"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBNOTA_CORRETAGEM_DATA"))
	b.Tabelas = append(b.Tabelas, *NewTabela("TBNOTA_CORRETAGEM_OPER"))
}

func (b *Base) CarregarDeletes() {
	var start time.Time = time.Now()

	dbHist := b.ConectarBDHist()
	defer dbHist.Close()

	txHist, err := dbHist.Begin()
	if err != nil {
		log.Fatalln("Erro.Delete.Begin: ", err)
	}

	var wg sync.WaitGroup
	iTotLen := len(b.Tabelas)
	iIdx := iTotLen - 1

	for {

		wg.Add(1)
		go func(ww *sync.WaitGroup, iiIdx int, iiTotLen int) {

			defer ww.Done()
			// log.Println("#"+strconv.Itoa(iiIdx+1)+"/"+strconv.Itoa(iiTotLen)+" - Tab:", b.Tabelas[iiIdx].Nome)

			tabela := b.Tabelas[iiIdx]

			if tabela.Nome == "TBCOMENTARIO" {
				_, err = txHist.Exec("DELETE FROM TBCOMENTARIO WHERE IDPAI IS NOT NULL")
				if err != nil {
					// txHist.Rollback()
					log.Println("Erro.Delete.Tabela:", tabela.Nome)
					log.Fatalln("Erro.Delete:", err)
				}
			}

			for {
				_, err = txHist.Exec("DELETE FROM " + tabela.Nome)
				if err != nil {
					// txHist.Rollback()
					// log.Println("Erro.Delete.Tabela:", tabela.Nome)
					// log.Fatalln("Erro.Delete:", tabela.Nome, err)
					continue
				}
				break
			} // for {

		}(&wg, iIdx, iTotLen)

		iIdx--
		if iIdx < 0 {
			break
		}
	} // for _, tabela := range b.Tabelas {

	wg.Wait()

	err = txHist.Commit()
	if err != nil {
		log.Fatalln("Erro.Delete.Commit: ", err)
	}

	log.Println("FIM Deletes:", time.Since(start))
}

func (b *Base) CarregarCampos() {

	var start time.Time = time.Now()

	dbAtivo := b.ConectarBDAtivo() // ConectarBDAtivo // ConectarBDHist
	defer dbAtivo.Close()

	sQuery := "SELECT COLUMN_NAME, DATA_TYPE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = ? ORDER BY ORDINAL_POSITION"
	stmt, err := dbAtivo.Prepare(sQuery)
	if err != nil {
		log.Fatalln("Erro.Campos.Prepare: ", err)
	}
	defer stmt.Close()

	// Campos: 1m1.6216909s // antes
	// Campos: 47.252127s   // depois - ConectarBDAtivo
	// Campos: 675.1423ms   // depois - ConectarBDHist

	// dbAtivo.Exec("SET @global.max_allowed_packet  = 10737418240 ")
	// dbAtivo.Exec("SET @@global.max_allowed_packet = 10737418240 ")
	// dbAtivo.Exec("SET max_allowed_packet          = 10737418240 ")
	// dbAtivo.Exec("SET GLOBAL max_allowed_packet   = 10737418240 ")
	// dbAtivo.Exec("SET SESSION max_allowed_packet  = 10737418240 ")
	// dbAtivo.Exec("SET GLOBAL net_buffer_length    = 10737418240 ")

	// dbAtivo.Exec("SET GLOBAL max_user_connections = 300 ")
	// dbAtivo.Exec("SET GLOBAL max_connections = 300")

	var wg sync.WaitGroup
	iTotLen := len(b.Tabelas)

	for iIdx := 0; iIdx < iTotLen; iIdx++ {

		if iIdx > 0 && iIdx%20 == 0 {
			// log.Println("#"+strconv.Itoa(iIdx+1)+"/"+strconv.Itoa(iTotLen)+" - Sleep")
			// time.Sleep(1 * time.Second)
			wg.Wait()
		}

		wg.Add(1)
		go func(ww *sync.WaitGroup, iiIdx int, iiTotLen int) {

			defer ww.Done()
			// log.Println("#"+strconv.Itoa(iiIdx+1)+"/"+strconv.Itoa(iiTotLen)+" - Tab:", b.Tabelas[iiIdx].Nome)

			selDBAtivo, err := stmt.Query(b.Tabelas[iiIdx].Nome)
			if err != nil {
				log.Fatalln(err)
			}
			defer selDBAtivo.Close()

			var campos []TabelaCampo

			for selDBAtivo.Next() {
				var sColuna, sTipo string
				err = selDBAtivo.Scan(&sColuna, &sTipo)
				if err != nil {
					log.Fatalln(err)
				}
				sColuna = strings.ToUpper(sColuna)
				sTipo = strings.ToUpper(sTipo)
				campos = append(campos, *NewTabelaCampo(sColuna, sTipo))
			} // for selDBAtivo.Next() {

			b.Tabelas[iiIdx].Campos = campos

		}(&wg, iIdx, iTotLen)

	} // for iIdx := 0; iIdx < iTotLen; iIdx++ {

	wg.Wait()

	for iIdx := 0; iIdx < iTotLen; iIdx++ {
		wg.Add(1)
		go func(ww *sync.WaitGroup, iiIdx int) {
			defer ww.Done()
			b.Tabelas[iiIdx].MontarInsert()
		}(&wg, iIdx)
	} // for iIdx := 0; iIdx < iTotLen; iIdx++ {

	wg.Wait()

	log.Println("FIM Campos:", time.Since(start))
}

func (b *Base) CarregarInserts() {

	var start time.Time = time.Now()

	dbAtivo := b.ConectarBDAtivo()
	defer dbAtivo.Close()

	dbHist := b.ConectarBDHist()
	defer dbHist.Close()

	iTot := len(b.Tabelas)
	for iIndx, tabela := range b.Tabelas {

		if tabela.Nome == "TBCEI_OPER_USER_10" ||
			tabela.Nome == "TBCEI_OPER_USER_188" ||
			tabela.Nome == "TBCEI_OPER_USER_2" ||
			tabela.Nome == "TBCEI_OPER_USER_212" ||
			tabela.Nome == "TBCEI_OPER_USER_227" ||
			tabela.Nome == "TBCEI_OPER_USER_231" ||
			tabela.Nome == "TBCEI_OPER_USER_238" ||
			tabela.Nome == "TBCEI_OPER_USER_26" ||
			tabela.Nome == "TBCEI_OPER_USER_261" ||
			tabela.Nome == "TBCEI_OPER_USER_321" ||
			tabela.Nome == "TBCEI_OPER_USER_357" ||
			tabela.Nome == "TBCEI_OPER_USER_358" ||
			tabela.Nome == "TBCEI_OPER_USER_382" ||
			tabela.Nome == "TBCEI_OPER_USER_384" ||
			tabela.Nome == "TBCEI_OPER_USER_391" ||
			tabela.Nome == "TBCEI_OPER_USER_396" ||
			tabela.Nome == "TBCEI_OPER_USER_404" ||
			tabela.Nome == "TBCEI_OPER_USER_410" ||
			tabela.Nome == "TBCEI_OPER_USER_411" ||
			tabela.Nome == "TBCEI_OPER_USER_60" ||
			tabela.Nome == "TBCEI_OPER_USER_64" ||
			tabela.Nome == "TBCEI_OPER_USER_67" ||
			tabela.Nome == "TBCEI_OPER_USER_68" ||
			tabela.Nome == "TBCEI_OPER_USER_7" ||
			tabela.Nome == "TBCEI_PROV_USER_188" ||
			tabela.Nome == "TBCEI_PROV_USER_2" ||
			tabela.Nome == "TBCEI_PROV_USER_212" ||
			tabela.Nome == "TBCEI_PROV_USER_227" ||
			tabela.Nome == "TBCEI_PROV_USER_231" ||
			tabela.Nome == "TBCEI_PROV_USER_238" ||
			tabela.Nome == "TBCEI_PROV_USER_26" ||
			tabela.Nome == "TBCEI_PROV_USER_261" ||
			tabela.Nome == "TBCEI_PROV_USER_321" ||
			tabela.Nome == "TBCEI_PROV_USER_357" ||
			tabela.Nome == "TBCEI_PROV_USER_358" ||
			tabela.Nome == "TBCEI_PROV_USER_382" ||
			tabela.Nome == "TBCEI_PROV_USER_384" ||
			tabela.Nome == "TBCEI_PROV_USER_391" ||
			tabela.Nome == "TBCEI_PROV_USER_404" ||
			tabela.Nome == "TBCEI_PROV_USER_410" ||
			tabela.Nome == "TBCEI_PROV_USER_411" ||
			tabela.Nome == "TBCEI_PROV_USER_60" ||
			tabela.Nome == "TBCEI_PROV_USER_64" ||
			tabela.Nome == "TBCEI_PROV_USER_67" ||
			tabela.Nome == "TBCEI_PROV_USER_68" ||
			tabela.Nome == "TBCEI_PROV_USER_7" {
			continue
		}

		// var startLocal time.Time = time.Now()

		// log.Println(" --> "+tabela.Nome, "=", tabela.Insert)

		selDBAtivo, err := dbAtivo.Query("SELECT * FROM " + tabela.Nome)
		if err != nil {
			log.Fatalln(err)
		}
		defer selDBAtivo.Close()

		// dbAtivo.Ping()
		// dbHist.Ping()

		colNames, err := selDBAtivo.Columns()
		lenCN := len(colNames)

		vals := make([]interface{}, lenCN)
		values := make([][]interface{}, 0, lenCN)
		row := make(map[string]string, lenCN)

		for i := 0; i < lenCN; i++ {
			vals[i] = new(sql.RawBytes)
		}

		for selDBAtivo.Next() {

			err := selDBAtivo.Scan(vals...)
			if err != nil {
				log.Fatalln(err)
			}

			for i := 0; i < lenCN; i++ {
				if rb, ok := vals[i].(*sql.RawBytes); ok {
					row[colNames[i]] = string(*rb)
					*rb = nil // reset pointer to discard current value to avoid a bug
				}
			}

			vls := make([]interface{}, 0)
			for _, col := range colNames {
				if (tabela.Nome == "TBEMPRESA" ||
					tabela.Nome == "TBFII_FUNDOIMOB" ||
					tabela.Nome == "TBCORRETORA" ||
					tabela.Nome == "TBCOMENTARIO" ||
					tabela.Nome == "TBLANCAMENTO" ||
					tabela.Nome == "TBOPERACAO" ||
					tabela.Nome == "TBPROVENTO" ||
					tabela.Nome == "TBFII_LANCAMENTO" ||
					tabela.Nome == "TBFII_PROVENTO" ||
					tabela.Nome == "TBETF_LANCAMENTO" ||
					tabela.Nome == "TBETF_OPERACAO" ||
					tabela.Nome == "TBBDR_LANCAMENTO" ||
					tabela.Nome == "TBBDR_OPERACAO" ||
					tabela.Nome == "TBBDR_PROVENTO" ||
					tabela.Nome == "TBCRIPTO_LANCAMENTO" ||
					tabela.Nome == "TBCARTEIRA_ATIVO" ||
					tabela.Nome == "TBCARTEIRA_FUNDO" ||
					tabela.Nome == "TBCARTEIRA_INDICE" ||
					tabela.Nome == "TBCARTEIRA_BDR" ||
					tabela.Nome == "TBCARTEIRA_COTAS" ||
					tabela.Nome == "TBCARTEIRA_CRIPTO") &&
					row[col] == "" {
					vls = append(vls, nil)
				} else {
					vls = append(vls, row[col])
				}
			}
			values = append(values, vls)

		} // for selDBAtivo.Next() {

		// dbAtivo.Ping()
		// dbHist.Ping()

		txHist, err := dbHist.Begin()
		if err != nil {
			log.Fatalln("Erro.Begin: ", err)
		}

		_, err = txHist.Exec("SET IDENTITY_INSERT " + tabela.Nome + " ON")
		if err != nil {
			txHist.Rollback()
			txHist, _ = dbHist.Begin()
		}

		stmt, err := txHist.Prepare(tabela.Insert)
		if err != nil {
			log.Fatalln("Erro.Prepare: ", err)
		}
		defer stmt.Close()

		iQtdTot := 0
		for _, v := range values {
			iQtdTot++

			_, err = stmt.Exec(v...)
			if err != nil {
				log.Println("Erro.Exec.tabela: ", tabela.Nome)
				log.Println("Erro.Exec.Tentando.Novamente")
				_, err = stmt.Exec(v...)
			}

			if err != nil {
				txHist.Rollback()
				log.Println("Erro.Exec.tabela: ", tabela.Nome)
				log.Println("Erro.Exec.Insert: ", tabela.Insert)
				log.Println("Erro.Exec.tamanho: ", len(v), "valores:", v)
				//log.Println("Erro.Exec.values: ", values)
				log.Fatalln("Erro.Exec: ", err)
			}

			// dbAtivo.Ping()
			// dbHist.Ping()

		}

		err = txHist.Commit()
		if err != nil {
			log.Fatalln("Erro.Commit: ", err)
		}

		// dbAtivo.Ping()
		// dbHist.Ping()

		// log.Println("#"+strconv.Itoa(iIndx+1)+"/"+strconv.Itoa(iTot)+" - Inserts "+tabela.Nome+"("+strconv.Itoa(iQtdTot)+"):", time.Since(startLocal))
		if iIndx == 0 || iIndx%10 == 0 {
			log.Println("#" + strconv.Itoa(iIndx+1) + "/" + strconv.Itoa(iTot) + " - ok")
		}

	} // for _, tabela := range b.Tabelas {

	// dbHist.Exec(" UPDATE TBUSUARIO SET SENHA = 'a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3', TENTATIVA=0, SITUACAO='A' ")

	log.Println("FIM Inserts:", time.Since(start))
}

func (b *Base) ConsultaTeste() {

	var start time.Time = time.Now()

	db := b.ConectarBDHist() // ConectarBDAtivo // ConectarBDHist
	defer db.Close()

	sel, err := db.Query(" SELECT NOME FROM TBUSUARIO WHERE ID = 2 ")
	if err != nil {
		log.Fatalln("Erro.Query: ", err)
	}
	defer sel.Close()

	for sel.Next() {
		var sNome string
		err = sel.Scan(&sNome)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("  --> Nome:", sNome)
	} // for sel.Next() {

	log.Println("FIM Teste Consulta:", time.Since(start))
}

func (b *Base) ConectarBDAtivo() *sql.DB {

	driverName := "mysql"
	// dataSourceName := "root:senha@tcp(localhost:3306)/database" // Notebook
	// dataSourceName := "root:senha@tcp(localhost:3306)/database?parseTime=true&timeout=5m&readTimeout=5m&net_write_timeout=6000&tls=skip-verify&charset=utf8" // Hostgator
	dataSourceName := "root:senha@tcp(localhost:3306)/database?parseTime=true&timeout=5m&readTimeout=5m&net_write_timeout=6000&tls=skip-verify&charset=utf8" // DigitalOcean

	log.Println("ConectarBDAtivo:", dataSourceName)

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}
	// defer db.Close()

	db.SetMaxIdleConns(1000) // db.SetMaxIdleConns(0) // SetMaxIdleConns define o número máximo de conexões no pool de conexão ociosa.
	db.SetMaxOpenConns(2000) // SetMaxOpenConns define o número máximo de conexões abertas com o banco de dados.
	db.SetConnMaxIdleTime(time.Hour)
	db.SetConnMaxLifetime(time.Minute * 60) // 24 *time.Hour // SetConnMaxLifetime define a quantidade máxima de tempo que uma conexão pode ser reutilizada.

	loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))

	// db = sqldblogger.OpenDriver(dataSourceName, db.Driver(), loggerAdapter)

	db = sqldblogger.OpenDriver(
		dataSourceName,
		db.Driver(),
		loggerAdapter,
		sqldblogger.WithErrorFieldname("sql_error"),                   // default: error
		sqldblogger.WithDurationFieldname("query_duration"),           // default: duration
		sqldblogger.WithTimeFieldname("log_time"),                     // default: time
		sqldblogger.WithSQLQueryFieldname("sql_query"),                // default: query
		sqldblogger.WithSQLArgsFieldname("sql_args"),                  // default: args
		sqldblogger.WithMinimumLevel(sqldblogger.LevelError),          // default: LevelDebug
		sqldblogger.WithLogArguments(true),                            // default: true
		sqldblogger.WithDurationUnit(sqldblogger.DurationMillisecond), // default: DurationMillisecond
		sqldblogger.WithTimeFormat(sqldblogger.TimeFormatRFC3339),     // default: TimeFormatUnix
		sqldblogger.WithLogDriverErrorSkip(false),                     // default: false
		sqldblogger.WithSQLQueryAsMessage(false),                      // default: false
		//sqldblogger.WithUIDGenerator(sqldblogger.UIDGenerator),       // default: *defaultUID
		sqldblogger.WithConnectionIDFieldname("con_id"),       // default: conn_id
		sqldblogger.WithStatementIDFieldname("stm_id"),        // default: stmt_id
		sqldblogger.WithTransactionIDFieldname("trx_id"),      // default: tx_id
		sqldblogger.WithWrapResult(false),                     // default: true
		sqldblogger.WithIncludeStartTime(false),               // default: false
		sqldblogger.WithStartTimeFieldname("start_time"),      // default: start
		sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug), // default: LevelInfo
		sqldblogger.WithQueryerLevel(sqldblogger.LevelDebug),  // default: LevelInfo
		sqldblogger.WithExecerLevel(sqldblogger.LevelDebug),   // default: LevelInfo
	)

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func (b *Base) ConectarBDHist() *sql.DB {

	// driverName := "mssql"
	driverName := "mysql"
	// dataSourceName := "sqlserver://sa:sa@127.0.0.1:5401?database=CMS_TESTE_TNB" // Notebook
	// dataSourceName := "root:senha@tcp(localhost:3306)/database?parseTime=true&timeout=5m&readTimeout=5m&net_write_timeout=6000&tls=skip-verify&charset=utf8" // DigitalOcean
	dataSourceName := "root:senha@tcp(localhost:3306)/database?parseTime=true&timeout=5m&readTimeout=5m&net_write_timeout=6000&tls=skip-verify&charset=utf8" // Docker
	log.Println("ConectarBDHist:", dataSourceName)

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}
	// defer db.Close()

	db.SetMaxIdleConns(1000) // db.SetMaxIdleConns(0) // SetMaxIdleConns define o número máximo de conexões no pool de conexão ociosa.
	db.SetMaxOpenConns(2000) // SetMaxOpenConns define o número máximo de conexões abertas com o banco de dados.
	db.SetConnMaxIdleTime(time.Hour)
	db.SetConnMaxLifetime(time.Minute * 60) // 24 *time.Hour // SetConnMaxLifetime define a quantidade máxima de tempo que uma conexão pode ser reutilizada.

	// loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))

	// db = sqldblogger.OpenDriver(dataSourceName, db.Driver(), loggerAdapter)

	// db = sqldblogger.OpenDriver(
	// 	dataSourceName,
	// 	db.Driver(),
	// 	loggerAdapter,
	// 	sqldblogger.WithErrorFieldname("sql_error"),                   // default: error
	// 	sqldblogger.WithDurationFieldname("query_duration"),           // default: duration
	// 	sqldblogger.WithTimeFieldname("log_time"),                     // default: time
	// 	sqldblogger.WithSQLQueryFieldname("sql_query"),                // default: query
	// 	sqldblogger.WithSQLArgsFieldname("sql_args"),                  // default: args
	// 	sqldblogger.WithMinimumLevel(sqldblogger.LevelInfo),           // default: LevelDebug
	// 	sqldblogger.WithLogArguments(true),                            // default: true
	// 	sqldblogger.WithDurationUnit(sqldblogger.DurationMillisecond), // default: DurationMillisecond
	// 	sqldblogger.WithTimeFormat(sqldblogger.TimeFormatRFC3339),     // default: TimeFormatUnix
	// 	sqldblogger.WithLogDriverErrorSkip(false),                     // default: false
	// 	sqldblogger.WithSQLQueryAsMessage(false),                      // default: false
	// 	//sqldblogger.WithUIDGenerator(sqldblogger.UIDGenerator),       // default: *defaultUID
	// 	sqldblogger.WithConnectionIDFieldname("con_id"),       // default: conn_id
	// 	sqldblogger.WithStatementIDFieldname("stm_id"),        // default: stmt_id
	// 	sqldblogger.WithTransactionIDFieldname("trx_id"),      // default: tx_id
	// 	sqldblogger.WithWrapResult(false),                     // default: true
	// 	sqldblogger.WithIncludeStartTime(false),               // default: false
	// 	sqldblogger.WithStartTimeFieldname("start_time"),      // default: start
	// 	sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug), // default: LevelInfo
	// 	sqldblogger.WithQueryerLevel(sqldblogger.LevelDebug),  // default: LevelInfo
	// 	sqldblogger.WithExecerLevel(sqldblogger.LevelDebug),   // default: LevelInfo
	// )

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

type Tabela struct {
	Nome   string
	Campos []TabelaCampo
	Insert string
}

func NewTabela(nome string) *Tabela {
	return &Tabela{
		Nome:   nome,
		Campos: []TabelaCampo{},
		Insert: "",
	}
}

func (t *Tabela) MontarInsert() {
	t.Insert = ""
	sqlCampo := ""
	sqlValor := ""
	for _, campo := range t.Campos {
		sqlCampo += ", " + campo.Nome
		sqlValor += ", ?"
	}
	t.Insert = "INSERT INTO " + t.Nome + " (" + sqlCampo[2:] + ") VALUES (" + sqlValor[2:] + ")"
}

type TabelaCampo struct {
	Nome string
	Tipo string
}

func NewTabelaCampo(nome string, tipo string) *TabelaCampo {
	return &TabelaCampo{
		Nome: nome,
		Tipo: tipo,
	}
}
