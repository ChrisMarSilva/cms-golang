package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	filename := "conf.csv"

	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1 // see the Reader struct information below
	csvLines, err := reader.ReadAll()
	// csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		fmt.Println(line[0] + " " + line[1] + " " + line[2])
	}

	// for {
	// 	row, err := csvr.Read()
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			err = nil
	// 		}
	// 	}
	// 	fmt.Println(row)
	// }

	// 	for _, row := range rawCSVdata {
	// 		for _, col := range row {
	// 				_,err := fmt.Print(col)
	// 				if err != nil {
	// 						fmt.Println(err)
	// 				}
	// 		}
	// 		fmt.Println("")
	//  }

	fmt.Println("FIM")
}
