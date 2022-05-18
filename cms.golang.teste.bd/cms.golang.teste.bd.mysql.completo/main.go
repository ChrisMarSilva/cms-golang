package main

import (
	"context"
	"os"

	//"database/sql"
	"fmt"
	"log"
	"time"

	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rocketlaunchr/dbq/v2"
	sql "github.com/rocketlaunchr/mysql-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	//"context"
	// "database/sql"
	// "fmt"
	// "log"
	// "strconv"
	// "strings"
	// "time"
	// _ "github.com/go-sql-driver/mysql"
	//"github.com/rocketlaunchr/dbq/v2"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
	// "gorm.io/hints"
	// "github.com/jmoiron/sqlx"
)

type Pessoa struct {
	Id     int32  `dbq:"ID"       db:"ID"        gorm:"primaryKey; column:ID"`
	Name   string `dbq:"NOME"     db:"NOME"      gorm:"column:NOME"`
	Status string `dbq:"SITUACAO" db:"SITUACAO"  gorm:"column:SITUACAO"`
}

type Person struct {
	Id     int32  `db:"ID"`
	Name   string `db:"NOME"`
	Status string `db:"SITUACAO"`
}

// Recommended by dbq
func (m *Pessoa) ScanFast() []interface{} {
	return []interface{}{&m.Id, &m.Name, &m.Status}
}

func (Pessoa) TableName() string {
	return "TBTESTE"
}

func (row Pessoa) ToString() string {
	return fmt.Sprintf("Id: %s; Name: %s; Status: %s", row.Id, row.Name, row.Status)
}

// DROP TABLE IF EXISTS TBTESTE;
// CREATE TABLE IF NOT EXISTS TBTESTE (
//   ID       int(11)      NOT NULL AUTO_INCREMENT,
//   NOME     varchar(150) NOT NULL,
//   SITUACAO varchar(1)   NOT NULL,
//   PRIMARY KEY (ID),
//   KEY IDX_TESTE_01 (NOME)
// ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

func processar_insert(database *gorm.DB, done chan int) {

	var pessoas2 []Pessoa
	for num := 1; num <= 100000; num++ {
		pessoa := Pessoa{Name: "Nome tete1 " + strconv.Itoa(num), Status: "A"}
		pessoas2 = append(pessoas2, pessoa)
	}
	start := time.Now()
	database.CreateInBatches(&pessoas2, 1000)               // len(pessoas2)
	log.Println("gorm.CreateInBatches=", time.Since(start)) // 1000 = 177.0071ms // 63.0944ms

}

func main() {

	var start time.Time
	var total int32
	//	max := 1000
	//	log.Println(max)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------

	log.Println("")

	db, err := sql.Open("mysql", "root:senha@tcp(localhost:3306)/database")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close() // 0ms

	//tx, _ := db.Begin()
	//tx.Exec("TRUNCATE TABLE TBTESTE")
	//tx.Commit()

	selDB, _ := db.Query("SELECT COUNT(1) FROM TBTESTE")
	selDB.Next()
	selDB.Scan(&total)
	//log.Println("Total=", total)

	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------

	// start = time.Now()
	// for num := 1; num <= max; num++ {
	// 	db.Exec("INSERT INTO TBTESTE (NOME, SITUACAO) values (?, ?) ", "Nome "+strconv.Itoa(num), "A")
	// }
	// log.Println("db.Exec=", time.Since(start)) // 1000 = 51.1656087s

	// tx, _ = db.Begin()
	// tx.Exec("TRUNCATE TABLE TBTESTE")
	// tx.Commit()

	// // ------------------------------------------------------------------------------------------------------
	// // ------------------------------------------------------------------------------------------------------

	// start = time.Now()
	// tx, _ = db.Begin()
	// for num := 1; num <= max; num++ {
	// 	tx.Exec("INSERT INTO TBTESTE (NOME, SITUACAO) values (?, ?) ", "Nome "+strconv.Itoa(num), "A")
	// }
	// tx.Commit()
	// log.Println("tx.Exec=", time.Since(start)) // 1000 = 50.2818532s

	// tx, _ = db.Begin()
	// tx.Exec("TRUNCATE TABLE TBTESTE")
	// tx.Commit()

	// // ------------------------------------------------------------------------------------------------------
	// // ------------------------------------------------------------------------------------------------------

	// start = time.Now()
	// tx, _ = db.Begin()
	// stmt, _ := tx.Prepare("INSERT INTO TBTESTE (NOME, SITUACAO) values (?, ?)")
	// for num := 1; num <= max; num++ {
	// 	stmt.Exec("Nome "+strconv.Itoa(num), "A")
	// }
	// tx.Commit()
	// log.Println("db.Prepare=", time.Since(start)) // 1000 = 1.6254313s

	// tx, _ = db.Begin()
	// tx.Exec("TRUNCATE TABLE TBTESTE")
	// tx.Commit()

	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------

	start = time.Now()
	selDB, _ = db.Query("SELECT ID, NOME, SITUACAO FROM TBTESTE WHERE ID = ?", 99999)
	selDB.Next()
	var id int32
	var nome, situacao string
	selDB.Scan(&id, &nome, &situacao)
	log.Println("db.Query=", time.Since(start)) // db.Query= 9.0129ms
	log.Println(id, nome, situacao)

	log.Println("")

	// start = time.Now()
	// selDB, _ = db.Query("SELECT ID, NOME, SITUACAO FROM TBTESTE")
	// log.Println("db.Query=", time.Since(start)) // db.Query= 7.93ms
	// defer selDB.Close()

	// start = time.Now()
	// for selDB.Next() {
	// 	var id int32
	// 	var nome, situacao string
	// 	selDB.Scan(&id, &nome, &situacao)
	// }
	// log.Println("sql.Next=", time.Since(start)) // 0s

	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------

	start = time.Now()
	database, err := gorm.Open(mysql.Open("root:senha@tcp(localhost:3306)/database?parseTime=true&loc=Local"), &gorm.Config{Logger: newLogger, SkipDefaultTransaction: true, QueryFields: true, PrepareStmt: true})
	if err != nil {
		log.Fatalln(err.Error())
	}
	//log.Println("gorm.Open=", time.Since(start)) // 8.1117ms

	// database.Exec("TRUNCATE TABLE TBTESTE")

	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------

	// start = time.Now()
	// for num := 1; num <= max; num++ {
	// 	database.Exec("INSERT INTO TBTESTE (NOME, SITUACAO) values (?, ?)", "Nome "+strconv.Itoa(num), "A")
	// }
	// log.Println("gorm.Exec=", time.Since(start)) // 1000 = 51.4150896s
	// database.Exec("TRUNCATE TABLE TBTESTE")

	// // ------------------------------------------------------------------------------------------------------
	// // ------------------------------------------------------------------------------------------------------

	// start = time.Now()
	// tx2 := database.Begin()
	// for num := 1; num <= max; num++ {
	// 	tx2.Exec("INSERT INTO TBTESTE (NOME, SITUACAO) values (?, ?)", "Nome "+strconv.Itoa(num), "A")
	// }
	// tx2.Commit()
	// log.Println("gorm.tx2.Exec=", time.Since(start)) // 1000 = 50.4685681s
	// database.Exec("TRUNCATE TABLE TBTESTE")

	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------

	// start = time.Now()
	// for num := 1; num <= max; num++ {
	// 	pessoa := Pessoa{Name: "Nome " + strconv.Itoa(num), Status: "A"}
	// 	database.Create(&pessoa)
	// }
	// log.Println("gorm.Create=", time.Since(start)) // 1000 = 51.4910592s //
	// database.Exec("TRUNCATE TABLE TBTESTE")

	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------

	// var pessoas []Pessoa
	// for num := 1; num <= max; num++ {
	// 	pessoa := Pessoa{Name: "Nome " + strconv.Itoa(num), Status: "A"}
	// 	pessoas = append(pessoas, pessoa)
	// }
	// start = time.Now()
	// database.Create(&pessoas)
	// log.Println("gorm.Create.Batch=", time.Since(start)) // 1000 = 155.6042ms // 137.0815ms
	// database.Exec("TRUNCATE TABLE TBTESTE")

	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------

	// var pessoas2 []Pessoa
	// for num := 1; num <= 1000000; num++ {
	// 	pessoa := Pessoa{Name: "Nome tete1 " + strconv.Itoa(num), Status: "A"}
	// 	pessoas2 = append(pessoas2, pessoa)
	// }
	// start = time.Now()
	// database.CreateInBatches(&pessoas2, 1000)               // len(pessoas2)
	// log.Println("gorm.CreateInBatches=", time.Since(start)) // 1000 = 177.0071ms // 63.0944ms
	// //database.Exec("TRUNCATE TABLE TBTESTE")

	// done := make(chan int)
	// start = time.Now()
	// for num := 1; num <= 10; num++ {
	// 	go processar_insert(database, done)
	// }
	// for num := 1; num <= 10; num++ {
	// 	<-done
	// }
	// log.Println("gorm.goroutine=", time.Since(start)) // c

	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------

	// valueStrings := []string{}
	// valueArgs := []interface{}{}
	// for num := 1; num <= max; num++ {
	// 	valueStrings = append(valueStrings, "(?, ?)")
	// 	valueArgs = append(valueArgs, "Nome "+strconv.Itoa(num))
	// 	valueArgs = append(valueArgs, "A")
	// }
	// smt := fmt.Sprintf("INSERT INTO TBTESTE (NOME, SITUACAO) values %s", strings.Join(valueStrings, ","))

	// start = time.Now()
	// tx3 := database.Begin()
	// tx3.Exec(smt, valueArgs...)
	// tx3.Commit()
	// log.Println("gorm.Exec.smt=", time.Since(start)) // 1000 = 162.6233ms

	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------

	var pessoa Pessoa

	// var pessoas []Pessoa
	// start = time.Now()
	// database.Find(&pessoas)
	// log.Println("gorm.Find=", time.Since(start)) //  11.1035ms

	// start = time.Now()
	// for _, pessoa := range pessoas {
	// 	pessoa.Name = pessoa.Name + "11"
	// }
	// log.Println("gorm.Find.Next=", time.Since(start)) // 0s

	// start = time.Now()
	// database.Raw("SELECT ID, NOME, SITUACAO FROM TBTESTE WHERE ID = ?", 999).Scan(&pessoa)
	// log.Println("gorm.Raw=", time.Since(start)) //3.9867ms

	// start = time.Now()
	// database.Where("ID = ?", 999).First(&pessoa)
	// log.Println("gorm.Where=", time.Since(start)) // 3.6303ms

	start = time.Now()
	database.First(&pessoa, "ID = ?", 99998)
	log.Println("gorm.First=", time.Since(start)) // 1.543ms
	log.Println(pessoa)

	// start = time.Now()
	// database.Clauses(hints.UseIndex("IDX_TESTE_01")).Find(&Pessoa{})
	// log.Println("gorm.Clauses=", time.Since(start)) //4.9509ms

	// start = time.Now()
	// database.Order("id").Limit(1).Find(&pessoa)
	// log.Println("gorm.Order.Limit.Find=", time.Since(start)) //3.6444ms

	log.Println("")

	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------

	ctx := context.Background()
	// pool, _ := sql.Open("mysql", "root:senha@tcp(localhost:3306)/database")

	conn, err := db.Conn(ctx)
	defer conn.Close() // Return the connection back to the pool

	// start = time.Now()
	// conn.QueryContext(ctx, "SELECT ID, NOME, SITUACAO FROM TBTESTE WHERE ID = 999 LIMIT 1")
	// log.Println("pool.conn.QueryContext=", time.Since(start)) // pool.conn.QueryContext= 2.7171ms

	// start = time.Now()
	// dbq.Qs(ctx, db, "SELECT ID, NOME, SITUACAO FROM TBTESTE WHERE ID = 999 LIMIT 1", Pessoa{}, nil)
	// log.Println("dbq.Qs=", time.Since(start)) //  dbq.Qs= 7.4032ms

	start = time.Now()
	result := dbq.MustQ(ctx, db, "SELECT ID, NOME, SITUACAO FROM TBTESTE WHERE ID = 99997 LIMIT 1", dbq.SingleResult)
	log.Println("dbq.MustQ=", time.Since(start)) //  dbq.MustQ= 2.8238ms
	pes := result.(map[string]interface{})
	log.Println(pes["ID"], pes["NOME"], pes["SITUACAO"])

	log.Println("")

	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------

	dbx, _ := sqlx.Connect("mysql", "root:senha@tcp(localhost:3306)/database")
	defer dbx.Close()

	// https://github.com/kgoralski/go-crud-template/blob/master/internal/banks/domain/store.go
	var jason Person
	start = time.Now()
	dbx.Get(&jason, "SELECT ID, NOME, SITUACAO FROM TBTESTE WHERE ID = ? LIMIT 1", 99996)
	log.Println("dbx.Get=", time.Since(start)) // dbx.Get= 1.744ms
	log.Println(jason)

	// places := []Person{}
	// start = time.Now()
	// dbx.Select(&places, "SELECT ID, NOME, SITUACAO FROM TBTESTE WHERE ID = 999 LIMIT 1")
	// log.Println("dbx.Select=", time.Since(start)) // dbx.Select= 1.5303ms
	// //log.Println(places[0])

	// start = time.Now()
	// dbx.NamedQuery(`SELECT ID, NOME, SITUACAO FROM TBTESTE WHERE ID=:ID`, map[string]interface{}{"ID": "999"})
	// log.Println("dbx.NamedQuery=", time.Since(start)) // dbx.NamedQuery= 5.5337ms
	//place := Person{}
	// for rows.Next() {
	// 	rows.StructScan(&place)
	// 	fmt.Printf("%#v\n", place)
	// }

	log.Println("")

	// ------------------------------------------------------------------------------------------------------
	// ------------------------------------------------------------------------------------------------------

}
