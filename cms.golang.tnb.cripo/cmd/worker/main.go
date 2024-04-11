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
	//filename = "../../docs/99. cms - export foxbit - Histórico Chris.xlsx"
	filename = "./docs/99. cms - export foxbit - Histórico Chris.xlsx"
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
