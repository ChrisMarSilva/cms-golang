package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// go get github.com/360EntSecGroup-Skylar/excelize/v2

func main() {

	f, err := excelize.OpenFile("Pasta1.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	c1, err := f.GetCellValue("Planilha1", "A1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c1)

	c2, err := f.GetCellValue("Planilha1", "A3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c2)

	c3, err := f.GetCellValue("Planilha1", "B2")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c3)

	rows, err := f.GetRows("Planilha1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Print("\n")
	}

	fmt.Println("FIM")
}
