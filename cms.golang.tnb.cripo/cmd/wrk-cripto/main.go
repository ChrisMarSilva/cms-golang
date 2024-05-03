package main

import (
	"log"

	"github.com/chrismarsilva/cms.golang.tnb.cripo.domain/services"
)

var (
	filename string
	sheet    string
)

func init() {
	filename = "./../../docs/foxbit.xlsx"
	sheet = "Planilha1" // "Planilha1" // "Query result"
}

func main() {
	operacionService := services.NewOperacionService(filename, sheet)

	err := operacionService.ProcessFile()
	if err != nil {
		log.Println("operacionService.Process():", err)
		return
	}
}
