package main

import (
	"log"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"

    "github.com/xuri/excelize/v2"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.arquivos.excel2
// go get github.com/xuri/excelize/v2
// go mod tidy

// go run main.go

func main() {

	f := excelize.NewFile()

    // Create a new sheet.
	index := f.NewSheet("Sheet2")  

    // Set value of a cell.
    f.SetCellValue("Sheet1", "B2", 100)
    f.SetCellValue("Sheet2", "A2", "Hello world.")

	// Set active sheet of the workbook.
    f.SetActiveSheet(index)

	// Save spreadsheet by the given path.
    if err := f.SaveAs("Book1.xlsx"); err != nil {
        log.Println(err)
    }

	log.Println("Fim - SaveAs")

	f, err := excelize.OpenFile("Book1.xlsx")
    if err != nil {
        log.Println(err)
        return
    }

    defer func() {
        // Close the spreadsheet.
        if err := f.Close(); err != nil {
            log.Println(err)
        }
    }()

    // Get value from cell by given worksheet name and axis.
    cell, err := f.GetCellValue("Sheet1", "B2")
    if err != nil {
        log.Println(err)
        return
    }
    log.Println("Sheet1", "B2", cell)

    // Get all the rows in the Sheet1.
    rows, err := f.GetRows("Sheet1")
    if err != nil {
        log.Println(err)
        return
    }

    log.Println("Sheet1")
	for idxRow, row := range rows {
        for idxCell, colCell := range row {
            log.Print("idxRow: ", idxRow, " - idxCell: ", idxCell, " - colCell: ", colCell, "\t")
        }
        // log.Println()
    }

	log.Println("Fim - OpenFile")
    
    // f := excelize.NewFile()

    f.NewSheet("Sheet3")

    categories := map[string]string{ "A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
    for k, v := range categories {
        f.SetCellValue("Sheet3", k, v)
    }

    values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
    for k, v := range values {
        f.SetCellValue("Sheet3", k, v)
    }

    if err := f.AddChart("Sheet3", "E1", `{
        "type": "col3DClustered",
        "series": [
        {
            "name": "Sheet3!$A$2",
            "categories": "Sheet3!$B$1:$D$1",
            "values": "Sheet3!$B$2:$D$2"
        },
        {
            "name": "Sheet3!$A$3",
            "categories": "Sheet3!$B$1:$D$1",
            "values": "Sheet3!$B$3:$D$3"
        },
        {
            "name": "Sheet3!$A$4",
            "categories": "Sheet3!$B$1:$D$1",
            "values": "Sheet3!$B$4:$D$4"
        }],
        "title":
        {
            "name": "Fruit 3D Clustered Column Chart"
        }
    }`); err != nil {
        log.Println(err)
        return
    }

    // Save spreadsheet by the given path.
    if err := f.SaveAs("Book1.xlsx"); err != nil {
        log.Println(err)
    }
    
	log.Println("Fim - SaveAs")


    f, err = excelize.OpenFile("Book1.xlsx")
    if err != nil {
        log.Println(err)
        return
    }

    defer func() {
        // Close the spreadsheet.
        if err := f.Close(); err != nil {
            log.Println(err)
        }
    }()

    // Insert a picture.
    if err := f.AddPicture("Sheet1", "A2", "chart.png", ""); err != nil {
        log.Println("#1", err)
    }

    // Insert a picture to worksheet with scaling.
    if err := f.AddPicture("Sheet1", "D2", "excel.jpg",`{"x_scale": 0.5, "y_scale": 0.5}`); err != nil {
        log.Println("#2", err)
    }

    // Insert a picture offset in the cell with printing support.
    if err := f.AddPicture("Sheet1", "H2", "excel.gif", `{
        "x_offset": 15,
        "y_offset": 10,
        "print_obj": true,
        "lock_aspect_ratio": false,
        "locked": false
    }`); err != nil {
        log.Println("#3", err)
    }

    // Save the spreadsheet with the origin path.
    if err = f.Save(); err != nil {
        log.Println(err)
    }

}
