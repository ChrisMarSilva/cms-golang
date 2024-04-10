package main

// go mod init github.com/chrismarsilva/cms.golang.tnb.cripo
// go get -u github.com/xuri/excelize/v2
// go get -u github.com/mattn/go-sqlite3
// go get -u github.com/shopspring/decimal
// go mod tidy
// go run main.go
// go run .

// go install github.com/cosmtrek/air@latest
// air init
// air

import (
	"log"

	"github.com/chrismarsilva/cms.golang.tnb.cripo/internals/services"
)

var (
	filename string
	sheet    string
)

func init() {
	filename = "./../../docs/99. cms - export foxbit - Histórico Chris.xlsx"
	sheet = "Planilha1" // Planilha1 // Query result
}

func main() {
	op := services.NewOperacionService(filename, sheet)
	err := op.ProcessFile()
	if err != nil {
		log.Println("op.Process():", err)
		return
	}
}

// func mainOld() {
// 	f, err := excelize.OpenFile(filename)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer func() {
// 		if err := f.Close(); err != nil {
// 			log.Println(err)
// 		}
// 	}()
//
// 	rows, err := f.GetRows(sheet)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
//
// 	var totDeposito float64 = 0
// 	var totRetirada float64 = 0
// 	p := message.NewPrinter(language.BrazilianPortuguese)
//
// 	for idx, row := range rows {
// 		//sn := row[0]
// 		datahora := strings.TrimSpace(row[1]) // DateTime.ParseExact(row[1]?.ToString()?.Trim()!, "dd/MM/yyyy HH:mm:ss", CultureInfo.InvariantCulture);
// 		descricao := strings.TrimSpace(row[2])
// 		moeda := strings.TrimSpace(row[3])
// 		valor := strings.Replace(strings.TrimSpace(row[4]), "R$", "", -1)
// 		taxa := strings.Replace(strings.TrimSpace(row[5]), "R$", "", -1)
// 		saldo := strings.Replace(strings.TrimSpace(row[6]), "R$", "", -1)
// 		dados := ""
// 		if len(row) > 7 {
// 			dados = strings.ReplaceAll(strings.TrimSpace(row[7]), "\n", "")
// 		}
//
// 		if moeda == "moeda" {
// 			continue
// 		}
//
// 		if descricao == "EarnApplication" {
// 			continue
// 		}
//
// 		dthr, err := time.Parse("02/01/2006 15:04:05", datahora)
// 		if err != nil {
// 			log.Println("datahora", err)
// 			return
// 		}
// 		datahora = dthr.Format("02/01/2006 15:04:05")
//
// 		valor = strings.Replace(valor, ".", "", -1)
// 		valor = strings.Replace(valor, ",", ".", -1)
// 		fValor, err := strconv.ParseFloat(valor, 64)
// 		if err != nil {
// 			log.Println("valor", err)
// 			return
// 		}
//
// 		if descricao == "Depósito" && moeda == "BRL" {
// 			totDeposito += fValor
// 			log.Println(idx, "\t", datahora, "\t", descricao, "\t", moeda, "\t", "\tDepósito\t", valor, "\t", taxa, "\t", saldo, "\t", dados)
// 			//continue
// 		} else if descricao == "Retirada" && moeda == "BRL" {
// 			totRetirada += math.Abs(fValor)
// 			log.Println(idx, "\t", datahora, "\t", descricao, "\t", moeda, "\t", "\tRetirada\t", valor, "\t", taxa, "\t", saldo, "\t", dados)
// 			//continue
// 		} else if descricao == "Trade" {
// 			taxa = strings.Replace(taxa, ".", "", -1)
// 			taxa = strings.Replace(taxa, ",", ".", -1)
// 			fTaxa, err := strconv.ParseFloat(taxa, 64)
// 			if err != nil {
// 				log.Println("taxa", err)
// 				return
// 			}
//
// 			saldo = strings.Replace(saldo, ".", "", -1)
// 			saldo = strings.Replace(saldo, ",", ".", -1)
// 			fSaldo, err := strconv.ParseFloat(saldo, 64)
// 			if err != nil {
// 				log.Println("saldo", err)
// 				return
// 			}
//
// 			if strings.Contains(dados, "Preço") {
// 				re := regexp.MustCompile("[^0-9.,]")
// 				dados = re.ReplaceAllString(dados, "")
// 				dados = strings.TrimSpace(dados)
//
// 			}
//
// 			dados = strings.Replace(dados, ".", "", -1)
// 			dados = strings.Replace(dados, ",", ".", -1)
// 			fDados, err := strconv.ParseFloat(dados, 64)
// 			if err != nil {
// 				log.Println("dados", err)
// 				return
// 			}
//
// 			log.Println(idx, "\t", datahora, "\t", descricao, "\t", moeda, "\tTrade\t", fValor, "\t", fTaxa, "\t", fSaldo, "\t", fDados)
// 			//continue
// 		} else {
// 			log.Println(idx, "\t", "TIPO NAO DEFINIDOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO", "\t", datahora, "\t", descricao, "\t", moeda, "\t", "\tRetirada\t", valor, "\t", taxa, "\t", saldo, "\t", dados)
// 			//continue
// 		}
//
// 		if idx >= 3 {
// 			break
// 		}
// 	}
//
// 	log.Println("")
// 	log.Println("Total Depósito\t: R$", p.Sprintf("%.2f", totDeposito))
// 	log.Println("Retirada Depósito\t: R$", p.Sprintf("%.2f", totRetirada))
// 	log.Println("Saldo Atual\t\t: R$", p.Sprintf("%.2f", totDeposito-totRetirada))
// }
