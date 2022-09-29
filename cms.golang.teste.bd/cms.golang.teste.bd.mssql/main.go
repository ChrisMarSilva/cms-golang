package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
	"errors"
	//"runtime"

	"github.com/denisenkom/go-mssqldb"
	// _ "github.com/denisenkom/go-mssqldb"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.bd.mssql
// go get -u github.com/denisenkom/go-mssqldb
// go mod tidy

// go run main.go

func main() {
	log.Println("INI")
	var startGeral time.Time = time.Now()

	var wg sync.WaitGroup
	var iTotal int = 1_000_000 // 100 // 1_000 // 10_000 // 100_000 // 1_000_000 // 10_000_000

	// wg.Add(1)
	// go DoInsertWithFor(iTotal, &wg, "TBTESTE01") 

	wg.Add(1)
	go DoInsertWithBulkOptions(iTotal, &wg, "TBTESTE02") 
	// DoInsertWithBulkOptions( 1_000_000 ):              41.3235301s

	// wg.Add(1)
	// go DoInsertWithGoRoutines(iTotal, &wg, "TBTESTE03")

	// wg.Add(1)
	// go DoInsertWithBulkOptionsAndGoRoutines(iTotal, &wg, "TBTESTE04")

	wg.Wait()

	log.Println("FIM   ===> Geral:", time.Since(startGeral))
}

/*

CREATE TABLE [dbo].[TBTESTE04](
	[ISPBPrincipal] [numeric](8, 0) NOT NULL,
	[ISPBAdministrado] [numeric](8, 0) NOT NULL,
	[TpPessoaPagdr] [varchar](1) NOT NULL,
	[CPFCNPJPagdr] [numeric](14, 0) NOT NULL,
	[NumCtrlPart] [varchar](20) NULL,
	[NumIdentcPagdr] [numeric](19, 0) NULL,
	[NumRefAtlCadCliPagdr] [numeric](19, 0) NULL,
	[NumSeqAtlzCadCliPagdr] [numeric](19, 0) NULL,
	[IndrAdesCliPagdrDDA] [varchar](1) NULL,
	[SitCliPagdrPart] [varchar](1) NULL,
	[TP_MSG_ARQUIVO] [varchar](1) NOT NULL,
	[ID_MSG_ARQUIVO] [numeric](9, 0) NULL,
	[DH_Registro] [numeric](17, 0) NOT NULL,
	[ST_Pagdr] [varchar](3) NOT NULL,
CONSTRAINT [PKTESTE04] PRIMARY KEY CLUSTERED 
(
	[ISPBPrincipal] ASC,
	[ISPBAdministrado] ASC,
	[TpPessoaPagdr] ASC,
	[CPFCNPJPagdr] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO

*/


func DoInsertWithFor(iTotal int, wgGeral *sync.WaitGroup, nomeTabela string) error {	
	// log.Println("INI   ===> DoInsertWithFor")
	defer wgGeral.Done()

	conn, err := GetConnection()
	if err != nil {
		// log.Println("Error GetConnection: "  +err.Error())
		return err
	}
	defer conn.Close()

	_, err = conn.Exec("TRUNCATE TABLE " + nomeTabela)
	if err != nil {
		// log.Println("Error conn.Exec(TruncateTable): "  +err.Error())
		return err
	}

	var startGeral time.Time = time.Now()
	// var startEtapas time.Time

	txn, err := conn.Begin()
	if err != nil {
		// log.Println("Error conn.Begin(): " + err.Error())
		return err
	}

	defer func() {
		err = txn.Rollback()
		if err != nil {
			if !errors.Is(err, sql.ErrTxDone) {
				log.Println("Error txn.Rollback(): "  +err.Error())
			}
		}
	}()

	sql := ""
	sql += "INSERT INTO " + nomeTabela + "(ISPBPrincipal, ISPBAdministrado, TpPessoaPagdr, CPFCNPJPagdr, NumCtrlPart, NumIdentcPagdr, NumRefAtlCadCliPagdr, NumSeqAtlzCadCliPagdr, IndrAdesCliPagdrDDA, SitCliPagdrPart, TP_MSG_ARQUIVO, ID_MSG_ARQUIVO, DH_Registro, ST_Pagdr) "
	sql += "VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "

	// startEtapas = time.Now()
	for iIndex := 0; iIndex < iTotal; iIndex++ {
		_, err = txn.Exec(sql, 4358798, 4358798, "F", iIndex+1, "22092720370000001000", iIndex+1, 1, 1, "S", "", "A", 101, 20220901102030, "PI9")
		if err != nil {
			// log.Println("Error stmt.Exec(1): " + err.Error())
			return err
		}
	}
	// log.Println("   DoInsertWithFor.TEMPO ===> stmt.Exec(for):", time.Since(startEtapas))

	// startEtapas = time.Now()
	err = txn.Commit()
	if err != nil {
		// log.Println("Error txn.Commit(): " + err.Error())
		return err
	}
	// log.Println("   DoInsertWithFor.TEMPO ===> txn.Commit():", time.Since(startEtapas))

	log.Println("TEMPO ===> DoInsertWithFor(", iTotal, "):                     ", time.Since(startGeral))
	return nil
}

func DoInsertWithBulkOptions(iTotal int, wgGeral *sync.WaitGroup, nomeTabela string) error {	
	// log.Println("INI   ===> DoInsertWithBulkOptions")
	defer wgGeral.Done()

	conn, err := GetConnection()
	if err != nil {
		// log.Println("Error GetConnection: "  +err.Error())
		return err
	}
	defer conn.Close()

	// https://github.com/denisenkom/go-mssqldb/blob/master/examples/bulk/bulk.go

	_, err = conn.Exec("TRUNCATE TABLE " + nomeTabela)
	if err != nil {
		// log.Println("Error conn.Exec(TruncateTable): "  +err.Error())
		return err
	}

	var startGeral time.Time = time.Now()
	// var startEtapas time.Time

	txn, err := conn.Begin()
	if err != nil {
		// log.Println("Error conn.Begin(): " + err.Error())
		return err
	}

	defer func() {
		err = txn.Rollback()
		if err != nil {
			if !errors.Is(err, sql.ErrTxDone) {
				log.Println("Error txn.Rollback(): "  +err.Error())
			}
		}
	}()

	bulkImport := mssql.CopyIn(nomeTabela, mssql.BulkOptions{}, "ISPBPrincipal", "ISPBAdministrado", "TpPessoaPagdr", "CPFCNPJPagdr", "NumCtrlPart", "NumIdentcPagdr", "NumRefAtlCadCliPagdr", "NumSeqAtlzCadCliPagdr", "IndrAdesCliPagdrDDA", "SitCliPagdrPart", "TP_MSG_ARQUIVO", "ID_MSG_ARQUIVO", "DH_Registro", "ST_Pagdr")
	// INSERTBULK {"TableName":"TBTESTE","ColumnsName":["ISPBPrincipal", "ISPBAdministrado", "TpPessoaPagdr", "CPFCNPJPagdr", "NumCtrlPart", "NumIdentcPagdr", "NumRefAtlCadCliPagdr", "NumSeqAtlzCadCliPagdr", "IndrAdesCliPagdrDDA", "SitCliPagdrPart", "TP_MSG_ARQUIVO", "ID_MSG_ARQUIVO", "DH_Registro", "ST_Pagdr"],"Options":{"CheckConstraints":false,"FireTriggers":false,"KeepNulls":false,"KilobytesPerBatch":0,"RowsPerBatch":0,"Order":null,"Tablock":false}}

	stmt, err := txn.Prepare(bulkImport)
	if err != nil {
		// log.Println("Error txn.Prepare(): " + err.Error())
		return err
	}
	defer stmt.Close()

	// startEtapas = time.Now()
	for iIndex := 0; iIndex < iTotal; iIndex++ {
		_, err = stmt.Exec(4358798, 4358798, "F", iIndex+1, "22092720370000001000", iIndex+1, 1, 1, "S", "", "A", 101, 20220901102030, "PI9")
		if err != nil {
			// log.Println("Error stmt.Exec(1): " + err.Error())
			return err
		}
	}
	// log.Println("   DoInsertWithBulkOptions.TEMPO ===> stmt.Exec(for):", time.Since(startEtapas))

	// startEtapas = time.Now()
	_, err = stmt.Exec()
	if err != nil {
		// log.Println("Error stmt.Exec(2): " + err.Error())
		return err
	}
	// rowCount, _ := result.RowsAffected()
	// log.Println("   DoInsertWithBulkOptions.TEMPO ===> stmt.Exec(", rowCount, "RowsAffected):", time.Since(startEtapas))

	// startEtapas = time.Now()
	err = txn.Commit()
	if err != nil {
		// log.Println("Error txn.Commit(): " + err.Error())
		return err
	}
	// log.Println("   DoInsertWithBulkOptions.TEMPO ===> txn.Commit():", time.Since(startEtapas))

	log.Println("TEMPO ===> DoInsertWithBulkOptions(", iTotal, "):             ", time.Since(startGeral))
	return nil
}

func DoInsertWithGoRoutines(iTotal int, wgGeral *sync.WaitGroup, nomeTabela string) error {	
	// log.Println("INI   ===> DoInsertWithGoRoutines")
	defer wgGeral.Done()

	conn, err := GetConnection()
	if err != nil {
		// log.Println("Error GetConnection: "  +err.Error())
		return err
	}
	defer conn.Close()

	// https://github.com/denisenkom/go-mssqldb/blob/master/examples/routine/routine.go

	_, err = conn.Exec("TRUNCATE TABLE " + nomeTabela)
	if err != nil {
		// log.Println("Error conn.Exec(TruncateTable): "  +err.Error())
		return err
	}

	var startGeral time.Time = time.Now()
	// var startEtapas time.Time

	txn, err := conn.Begin()
	if err != nil {
		// log.Println("Error conn.Begin(): " + err.Error())
		return err
	}

	defer func() {
		err = txn.Rollback()
		if err != nil {
			if !errors.Is(err, sql.ErrTxDone) {
				log.Println("Error txn.Rollback(): "  +err.Error())
			}
		}
	}()

	insertSql := ""
	insertSql += "INSERT INTO " + nomeTabela + "(ISPBPrincipal, ISPBAdministrado, TpPessoaPagdr, CPFCNPJPagdr, NumCtrlPart, NumIdentcPagdr, NumRefAtlCadCliPagdr, NumSeqAtlzCadCliPagdr, IndrAdesCliPagdrDDA, SitCliPagdrPart, TP_MSG_ARQUIVO, ID_MSG_ARQUIVO, DH_Registro, ST_Pagdr) "
	insertSql += "VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?) "

	stmt, err := txn.Prepare(insertSql)
	if err != nil {
		// log.Println("Error txn.Prepare(): " + err.Error())
		return err
	}
	defer stmt.Close()

	var wg sync.WaitGroup
	// startEtapas = time.Now()

	wg.Add(iTotal)

	for iIndex := 0; iIndex < iTotal; iIndex++ {
		// wg.Add(1)
		go func(iIndexLocal int) { // , wgLocal *sync.WaitGroup, stmtLocal *sql.Stmt
			defer wg.Done()
			_, err = stmt.Exec(4358798, 4358798, "F", iIndexLocal+1, "22092720370000001000", iIndexLocal+1, 1, 1, "S", "", "A", 101, 20220901102030, "PI9")
			if err != nil {
				log.Fatal(err)
				// log.Println("Error stmt.Exec(iIndexLocal): " + err.Error())
			}
		}(iIndex) // , &wg, stmt
	}

	wg.Wait()
	// log.Println("   DoInsertWithGoRoutines.TEMPO ===> stmt.Exec(for):", time.Since(startEtapas))

	// done := make(chan bool)
	// selectSql := "select idstr from " + nomeTabela + " where id = "
	// for iIndex := 0; iIndex < iTotal; iIndex++ {
	// 	go func(iIndexLocal int) {
	// 		rows, err := db.Query(selectSql + strconv.Itoa(iIndexLocal))
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		defer rows.Close()
	// 		for rows.Next() {
	// 			var id int64
	// 			err := rows.Scan(&id)
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			} else {
	// 				log.Printf("Found %d\n", iIndexLocal)
	// 			}
	// 		}
	// 		done <- true
	// 	}(iIndex)
	// }

	// for i := 0; i < cExec; i++ {
	// 	<-done
	// }

	// startEtapas = time.Now()
	err = txn.Commit()
	if err != nil {
		// log.Println("Error txn.Commit(): " + err.Error())
		return err
	}
	// log.Println("   DoInsertWithGoRoutines.TEMPO ===> txn.Commit():", time.Since(startEtapas))

	log.Println("TEMPO ===> DoInsertWithGoRoutines(", iTotal, "):              ", time.Since(startGeral))
	return nil
}

func DoInsertWithBulkOptionsAndGoRoutines(iTotal int, wgGeral *sync.WaitGroup, nomeTabela string) error {	
	// log.Println("INI   ===> DoInsertWithBulkOptionsAndGoRoutines")
	defer wgGeral.Done()

	conn, err := GetConnection()
	if err != nil {
		// log.Println("Error GetConnection: "  +err.Error())
		return err
	}
	defer conn.Close()

	// https://github.com/denisenkom/go-mssqldb/blob/master/examples/bulk/bulk.go
	// https://github.com/denisenkom/go-mssqldb/blob/master/examples/routine/routine.go

	_, err = conn.Exec("TRUNCATE TABLE " + nomeTabela) 
	if err != nil {
		// log.Println("Error conn.Exec(TruncateTable): "  +err.Error())
		return err
	}

	var startGeral time.Time = time.Now()
	// var startEtapas time.Time

	txn, err := conn.Begin()
	if err != nil {
		// log.Println("Error conn.Begin(): " + err.Error())
		return err
	}

	defer func() {
		err = txn.Rollback()
		if err != nil {
			if !errors.Is(err, sql.ErrTxDone) {
				log.Println("Error txn.Rollback(): "  +err.Error())
			}
		}
	}()

	bulkImport := mssql.CopyIn(nomeTabela, mssql.BulkOptions{}, "ISPBPrincipal", "ISPBAdministrado", "TpPessoaPagdr", "CPFCNPJPagdr", "NumCtrlPart", "NumIdentcPagdr", "NumRefAtlCadCliPagdr", "NumSeqAtlzCadCliPagdr", "IndrAdesCliPagdrDDA", "SitCliPagdrPart", "TP_MSG_ARQUIVO", "ID_MSG_ARQUIVO", "DH_Registro", "ST_Pagdr")
	// INSERTBULK {"TableName":"TBTESTE","ColumnsName":["ISPBPrincipal", "ISPBAdministrado", "TpPessoaPagdr", "CPFCNPJPagdr", "NumCtrlPart", "NumIdentcPagdr", "NumRefAtlCadCliPagdr", "NumSeqAtlzCadCliPagdr", "IndrAdesCliPagdrDDA", "SitCliPagdrPart", "TP_MSG_ARQUIVO", "ID_MSG_ARQUIVO", "DH_Registro", "ST_Pagdr"],"Options":{"CheckConstraints":false,"FireTriggers":false,"KeepNulls":false,"KilobytesPerBatch":0,"RowsPerBatch":0,"Order":null,"Tablock":false}}

	stmt, err := txn.Prepare(bulkImport)
	if err != nil {
		// log.Println("Error txn.Prepare(): " + err.Error())
		return err
	}
	defer stmt.Close()

	var wg sync.WaitGroup
	// startEtapas = time.Now()

	wg.Add(iTotal)

	for iIndex := 0; iIndex < iTotal; iIndex++ {
		// wg.Add(1)
		go func(iIndexLocal int) { // , wgLocal *sync.WaitGroup
			defer wg.Done()
			_, err = stmt.Exec(4358798, 4358798, "F", iIndexLocal+1, "22092720370000001000", iIndexLocal+1, 1, 1, "S", "", "A", 101, 20220901102030, "PI9")
			if err != nil {
				log.Fatal(err)
				// log.Println("Error stmt.Exec(iIndexLocal): " + err.Error())
				// return err
			}
		}(iIndex) // , &wg
	}

	wg.Wait()
	// log.Println("   DoInsertWithBulkOptionsAndGoRoutines.TEMPO ===> stmt.Exec(for):", time.Since(startEtapas))

	// startEtapas = time.Now()
	_, err = stmt.Exec()
	if err != nil {
		// log.Println("Error stmt.Exec(2): " + err.Error())
		return err
	}
	// rowCount, _ := result.RowsAffected()
	// log.Println("   DoInsertWithBulkOptionsAndGoRoutines.TEMPO ===> stmt.Exec(", rowCount, "RowsAffected):", time.Since(startEtapas))

	// startEtapas = time.Now()
	err = txn.Commit()
	if err != nil {
		// log.Println("Error txn.Commit(): " + err.Error())
		return err
	}
	// log.Println("   DoInsertWithBulkOptionsAndGoRoutines.TEMPO ===> txn.Commit():", time.Since(startEtapas))

	log.Println("TEMPO ===> DoInsertWithBulkOptionsAndGoRoutines(", iTotal, "):", time.Since(startGeral))
	return nil
}

func GetConnection() (*sql.DB, error) {

	// err := godotenv.Load()
	// if err != nil {
	//  log.Println("Error godotenv.Load(): "  +err.Error())
	// }
	// server := os.Getenv("MSSQL_DB_SERVER")
	// port = os.Getenv("MSSQL_DB_PORT")
	// user = os.Getenv("MSSQL_DB_USER")
	// password = os.Getenv("MSSQL_DB_PASSWORD")
	// database = os.Getenv("MSSQL_DB_DATABASE")

	server := "localhost" // "127.0.0.1" // "CMS-NOTE-2020\\SQLEXPRESS2019"
	port := 1433
	user := "sa"
	password := "sa123"
	database := "JDNPC"

	// connString := fmt.Sprintf("Provider=SQLOLEDB.1;Password=%s;Persist Security Info=True;User ID=%s;Initial Catalog=%s;Data Source=%s",  password, user, database, server)
	// connString := fmt.Sprintf("Provider=SQLNCLI11.1;Password=%s;Persist Security Info=False;User ID=%s;Initial Catalog=%s;Data Source=%s", password, user, database, server)
	// connString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=300&timeout=5m30s&charset=utf8", user, password, server, database)
	connString := fmt.Sprintf("server=%s;port=%d;user id=%s;password=%s;database=%s;connection timeout=300;timeout=300;charset=utf8;Trusted_Connection=True;", server, port, user, password, database)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		// loglog.Println("Error sql.Open(): " + err.Error())
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		// loglog.Println("Error conn.Ping(): " + err.Error())
		return nil, err
	}

	return conn, nil
}

func main_old() {

	var wg sync.WaitGroup
	connString := ""
	
	server :="127.0.0.1" //  "192.168.1.106" // "localhost" // "CMS-NOTE-2020//SQLEXPRESS2019" // "CMS-NOTE-2020//SQLEXPRESS2019" // "CMS-NOTE-2020\SQLEXPRESS2019"
	port := 1433
	password := "sa123"
	user := "sa"

	fmt.Printf("BD\n")
	fmt.Printf(" server:%s\n", server)
	fmt.Printf(" port:%d\n", port)
	fmt.Printf(" user:%s\n", user)
	fmt.Printf(" password:%s\n", password)

	wg.Add(1)
	connString = "Provider=SQLOLEDB.1;Password=sa123;Persist Security Info=True;User ID=sa;Initial Catalog=CMS_TESTE_CNAB240;Data Source=CMS-NOTE-2020"
	go TesteConnect("Teste #1.0", connString, &wg)
	wg.Add(1)
	connString = "Provider=SQLOLEDB.1;Password=sa123;Persist Security Info=True;User ID=sa;Initial Catalog=CMS_TESTE_CNAB240;Data Source=127.0.0.1"
	go TesteConnect("Teste #1.1", connString, &wg)
	wg.Add(1)
	connString = "Provider=SQLOLEDB.1;Password=sa123;Persist Security Info=True;User ID=sa;Initial Catalog=CMS_TESTE_CNAB240;Data Source=localhost"
	go TesteConnect("Teste #1.2", connString, &wg)

	wg.Add(1)
	connString = "Provider=SQLNCLI11.1;Password=sa123;Persist Security Info=False;User ID=sa;Initial Catalog=CMS_TESTE_CNAB240;Data Source=CMS-NOTE-2020"
	go TesteConnect("Teste #2.0", connString, &wg)
	wg.Add(1)
	connString = "Provider=SQLNCLI11.1;Password=sa123;Persist Security Info=False;User ID=sa;Initial Catalog=CMS_TESTE_CNAB240;Data Source=127.0.0.1"
	go TesteConnect("Teste #2.1", connString, &wg)
	wg.Add(1)
	connString = "Provider=SQLNCLI11.1;Password=sa123;Persist Security Info=False;User ID=sa;Initial Catalog=CMS_TESTE_CNAB240;Data Source=localhost"
	go TesteConnect("Teste #2.2", connString, &wg)

	wg.Add(1)
	connString = "sqlserver://sa:sa123@127.0.0.1?database=CMS_TESTE_CNAB240&connection+timeout=10&timeout=10s"
	go TesteConnect("Teste #3.0", connString, &wg)
	wg.Add(1)
	connString = "sqlserver://sa:sa123@localhost?database=CMS_TESTE_CNAB240&connection+timeout=10&timeout=10s"
	go TesteConnect("Teste #3.1", connString, &wg)
	wg.Add(1)
	connString = "sqlserver://sa:sa123@CMS-NOTE-2020?database=CMS_TESTE_CNAB240&connection+timeout=10&timeout=10s"
	go TesteConnect("Teste #3.2", connString, &wg)
	wg.Add(1)
	connString = "sqlserver://sa:sa123@192.168.1.106?database=CMS_TESTE_CNAB240"
	go TesteConnect("Teste #3.5", connString, &wg)

	// wg.Add(1)
	// connString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, password, port)
	// go TesteConnect("Teste #4", connString, &wg)
	wg.Add(1)
	connString = "server=localhost;user id=sa;password=sa123;database=CMS_TESTE_CNAB240;port=1433;connection timeout=30;Trusted_Connection=True;"
	go TesteConnect("Teste #4.1", connString, &wg)
	wg.Add(1)
	connString = "server=192.168.1.106;user id=sa;password=sa123;database=CMS_TESTE_CNAB240;port=1433;TrustServerCertificate=true;"
	go TesteConnect("Teste #4.2", connString, &wg)
	wg.Add(1)
	connString = "server=192.168.1.106;user id=sa;password=sa123;database=CMS_TESTE_CNAB240"
	go TesteConnect("Teste #4.3", connString, &wg)

	// wg.Add(1)
	// connString = "sqlserver://sa:sa123@127.0.0.1:1433/SQLEXPRESS2019?database=CMS_TESTE_CNAB240&encrypt=true"
	// go TesteConnect("Teste #22.6", connString, &wg)

	wg.Wait()

	// for i := 33; i <= 126; i++ { // for i := 33; i <= 126; i++ {
	// 	// fmt.Printf("%d: %c  ", i, i) //Prints ONLY the unicode chars
	// 	fmt.Printf("%d: %#U  ", i, i) //Prints the unicode chars and values as well
	// }

	// conn, err := sql.Open("mssql", connString)
	// if err != nil {
	// 	fmt.Println("Error Open: " + err.Error())
	// 	return
	// }
	// defer conn.Close()

	
	// Use SessionInitSql to set any options that cannot be set with the dsn string
	// With ANSI_NULLS set to ON, compare NULL data with = NULL or <> NULL will return 0 rows
	// conn.SessionInitSQL = "SET ANSI_NULLS ON"

	// rows, err := conn.Query("SELECT TPARQV, DESCRICAO FROM TBJDSPBCAB_CNAB240_ARQUIVO_TP WITH(NOLOCK)")
	// if err != nil {
	// 	log.Println("Error Query: " + err.Error())
	// 	return
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var sit, desc string
	// 	err := rows.Scan(&sit, &desc)
	// 	if err != nil {
	// 		fmt.Println("Error Scan: " + err.Error())
	// 		return
	// 	}
	// 	fmt.Printf("TPARQV: %s, DESCRICAO: %s\n", sit, desc)
	// }

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

	// runtime.Goexit()

	log.Print("FIM")
}

func TesteConnect(versao string, connString string, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		fmt.Println("Error Conn (" + versao + "): " + err.Error())
		// fmt.Println("Error Conn: " + err.Error() + " - connString: " + connString)
		return
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		fmt.Println("Error Ping (" + versao + "): " + err.Error())
		// fmt.Println("Error Ping: " + err.Error() + " - connString: " + connString)
		return
	}

	rows, err := conn.Query("SELECT TPARQV, DESCRICAO FROM TBJDSPBCAB_CNAB240_ARQUIVO_TP WITH(NOLOCK)")
	if err != nil {
		fmt.Println("Error Query (" + versao + "): " + err.Error())
		// log.Println("Error Query: " + err.Error() + " - connString: " + connString)
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
		fmt.Printf("OKKKKKKKK: %s - TPARQV: %s, DESCRICAO: %s\n", versao, sit, desc)
		break
	}

	// fmt.Println("OKKKKKKKK: " + versao)
	// fmt.Println("OKKKKKKKK - connString: " + connString)
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
