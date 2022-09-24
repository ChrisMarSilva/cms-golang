package main

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"

	_ "github.com/denisenkom/go-mssqldb"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.bd.mssql
// go get -u github.com/denisenkom/go-mssqldb
// go mod tidy

// go run main.go

func main() {

	connString := ""
	connString = "Provider=SQLOLEDB.1;Password=jddesenv;Persist Security Info=True;User ID=jddesenv;Initial Catalog=JDSPB;Data Source=JDSP108"
	connString = "Provider=SQLOLEDB.1;Password=sa;Persist Security Info=True;User ID=sa;Initial Catalog=CMS_TESTE_CNAB240;Data Source=CMS-NOTE-2020\\SQLEXPRESS"
	connString = "sqlserver://sa:sa@127.0.0.1:5401?database=CMS_TESTE_CNAB240"
	connString = "sqlserver://sa:sa@127.0.0.1:1433?database=CMS_TESTE_CNAB240"
	connString = "sqlserver://sa:sa@localhost:5401?database=CMS_TESTE_CNAB240"
	connString = "sqlserver://sa:sa@localhost:1433?database=CMS_TESTE_CNAB240"
	connString = "sqlserver://sa:sa@CMS-NOTE-2020\\SQLEXPRESS:5401?database=CMS_TESTE_CNAB240"
	connString = "sqlserver://sa:sa@localhost?database=CMS_TESTE_CNAB240"
	connString = "Provider=SQLNCLI11.1;Persist Security Info=False;User ID=sa;Initial Catalog=JDNPC;Data Source=CMS-NOTE-2020\\SQLEXPRESS"

	connString = "sqlserver://sa:sa@CMS-NOTE-2020" + fmt.Sprintf("%c", 92) + "SQLEXPRESS?database=CMS_TESTE_CNAB240"
	connString = "Provider=SQLNCLI11.1;Persist Security Info=False;User ID=sa;Initial Catalog=JDNPC;Data Source=CMS-NOTE-2020" + fmt.Sprintf("%c", 92) + "SQLEXPRESS"
	fmt.Println(connString)
	// return

	// for i := 33; i <= 126; i++ { // for i := 33; i <= 126; i++ {
	// 	// fmt.Printf("%d: %c  ", i, i) //Prints ONLY the unicode chars
	// 	fmt.Printf("%d: %#U  ", i, i) //Prints the unicode chars and values as well
	// }
	// return

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		fmt.Println("Error Open: " + err.Error())
		return
	}

	defer conn.Close()

	rows, err := conn.Query("SELECT TPARQV, DESCRICAO FROM TBJDSPBCAB_CNAB240_ARQUIVO_TP WITH(NOLOCK)")
	if err != nil {
		log.Println("Error Query: " + err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var sit, desc string
		err := rows.Scan(&sit, &desc)
		if err != nil {
			fmt.Println("Error Scan: " + err.Error())
			return
		}
		fmt.Printf("TPARQV: %s, DESCRICAO: %s\n", sit, desc)
	}

	// createID, err := CreateEmployee(conn, "Jake", "United States")
	// if err != nil {
	// 	log.Fatal("CreateEmployee failed:", err.Error())
	// }
	// fmt.Printf("Inserted ID: %d successfully.\n", createID)

	// count, err := ReadEmployees(conn)
	// if err != nil {
	// 	log.Fatal("ReadEmployees failed:", err.Error())
	// }
	// fmt.Printf("Read %d rows successfully.\n", count)

	// updateID, err := UpdateEmployee(conn, "Jake", "Poland")
	// if err != nil {
	// 	log.Fatal("UpdateEmployee failed:", err.Error())
	// }
	// fmt.Printf("Updated row with ID: %d successfully.\n", updateID)

	// rows, err := DeleteEmployee(conn, "Jake")
	// if err != nil {
	// 	log.Fatal("DeleteEmployee failed:", err.Error())
	// }
	// fmt.Printf("Deleted %d rows successfully.\n", rows)

	runtime.Goexit()

	fmt.Print("ok")
}

func CreateEmployee(db *sql.DB, name string, location string) (int64, error) {
	tsql := fmt.Sprintf("INSERT INTO TestSchema.Employees (Name, Location) VALUES ('%s','%s');",
		name, location)
	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error inserting new row: " + err.Error())
		return -1, err
	}
	return result.LastInsertId()
}

func ReadEmployees(db *sql.DB) (int, error) {
	tsql := fmt.Sprintf("SELECT Id, Name, Location FROM TestSchema.Employees;")
	tsql = fmt.Sprintf("SELECT ST_EXPURGO, DSC_SITUACAO FROM TBJDLGPD_EXPURGO_SITUACAO WITH(NOLOCK);")
	rows, err := db.Query(tsql)
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
		return -1, err
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		// var name, location string
		// var id int
		// err := rows.Scan(&id, &name, &location)
		// if err != nil {
		// 	fmt.Println("Error reading rows: " + err.Error())
		// 	return -1, err
		// }
		// fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, name, location)
		// count++
		var sit, desc string
		err := rows.Scan(&sit, &desc)
		if err != nil {
			fmt.Println("Error reading rows: " + err.Error())
			return -1, err
		}
		fmt.Printf("ST_EXPURGO: %s, DSC_SITUACAO: %s\n", sit, desc)
		count++
	}

	return count, nil
}

func UpdateEmployee(db *sql.DB, name string, location string) (int64, error) {
	tsql := fmt.Sprintf("UPDATE TestSchema.Employees SET Location = '%s' WHERE Name= '%s'",
		location, name)
	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error updating row: " + err.Error())
		return -1, err
	}
	return result.LastInsertId()
}

func DeleteEmployee(db *sql.DB, name string) (int64, error) {
	tsql := fmt.Sprintf("DELETE FROM TestSchema.Employees WHERE Name='%s';", name)
	result, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("Error deleting row: " + err.Error())
		return -1, err
	}
	return result.RowsAffected()
}
