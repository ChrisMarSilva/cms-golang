package main

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/rs/zerolog"
	// sqldblogger "github.com/simukti/sqldb-logger"
	// "github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.bd.mysql.basico
// go get -u github.com/go-sql-driver/mysql
// go get -u github.com/denisenkom/go-mssqldb
// go get -u -v github.com/simukti/sqldb-logger
// go get -u github.com/simukti/sqldb-logger/logadapter/zerologadapter
// go mod tidy

// go run main.go

func main() {

	var start time.Time

	//driverName := "mysql"
	//dataSourceName := "root:senha@tcp(localhost:3306)/database"

	driverName := "mssql"
	dataSourceName := "sqlserver://sa:sa@127.0.0.1:5401?database=CMS_TESTE_TNB"

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalln("sql.Open", err)
	}
	defer db.Close() // adiar o fechamento até depois que a função principal terminar

	db.SetMaxIdleConns(100) // SetMaxIdleConns define o número máximo de conexões no pool de conexão ociosa.
	db.SetMaxOpenConns(200) // SetMaxOpenConns define o número máximo de conexões abertas com o banco de dados.
	db.SetConnMaxIdleTime(time.Hour)
	db.SetConnMaxLifetime(time.Minute * 30) // 24 *time.Hour // SetConnMaxLifetime define a quantidade máxima de tempo que uma conexão pode ser reutilizada.

	//loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	// db = sqldblogger.OpenDriver(dsn, db.Driver(), loggerAdapter /*, using_default_options*/) // db is STILL *sql.DB

	// db = sqldblogger.OpenDriver(
	// 	dataSourceName,
	// 	db.Driver(),
	// 	loggerAdapter,
	// 	// AVAILABLE OPTIONS
	// 	sqldblogger.WithErrorFieldname("sql_error"),                  // default: error
	// 	sqldblogger.WithDurationFieldname("query_duration"),          // default: duration
	// 	sqldblogger.WithTimeFieldname("log_time"),                    // default: time
	// 	sqldblogger.WithSQLQueryFieldname("sql_query"),               // default: query
	// 	sqldblogger.WithSQLArgsFieldname("sql_args"),                 // default: args
	// 	sqldblogger.WithMinimumLevel(sqldblogger.LevelError),         // default: LevelDebug
	// 	sqldblogger.WithLogArguments(false),                          // default: true
	// 	sqldblogger.WithDurationUnit(sqldblogger.DurationNanosecond), // default: DurationMillisecond
	// 	sqldblogger.WithTimeFormat(sqldblogger.TimeFormatRFC3339),    // default: TimeFormatUnix
	// 	sqldblogger.WithLogDriverErrorSkip(true),                     // default: false
	// 	sqldblogger.WithSQLQueryAsMessage(true),                      // default: false
	// 	//sqldblogger.WithUIDGenerator(sqldblogger.UIDGenerator),       // default: *defaultUID
	// 	sqldblogger.WithConnectionIDFieldname("con_id"),       // default: conn_id
	// 	sqldblogger.WithStatementIDFieldname("stm_id"),        // default: stmt_id
	// 	sqldblogger.WithTransactionIDFieldname("trx_id"),      // default: tx_id
	// 	sqldblogger.WithWrapResult(false),                     // default: true
	// 	sqldblogger.WithIncludeStartTime(true),                // default: false
	// 	sqldblogger.WithStartTimeFieldname("start_time"),      // default: start
	// 	sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug), // default: LevelInfo
	// 	sqldblogger.WithQueryerLevel(sqldblogger.LevelDebug),  // default: LevelInfo
	// 	sqldblogger.WithExecerLevel(sqldblogger.LevelDebug),   // default: LevelInfo
	// )

	err = db.Ping()
	if err != nil {
		log.Fatalln("sql.Ping", err)
	}

	selDB, err := db.Query("SELECT CODIGO, DESCRICAO FROM TBCRIPTO_EMPRESA_ST")
	if err != nil {
		log.Fatalln("sql.Query", err)
	}
	defer selDB.Close()

	res := []Situacao{}
	for selDB.Next() {
		var codigo, descricao string
		err = selDB.Scan(&codigo, &descricao)
		if err != nil {
			log.Fatalln("sql.Scan", err)
		}
		emp := Situacao{}
		emp.Codigo = codigo
		emp.Descricao = descricao
		res = append(res, emp)
	}

	db.Exec("TRUNCATE TABLE TBTESTE_WRK")

	start = time.Now()
	for i := 1; i <= 10000; i++ { // 10000=10.000= // 100000=100.000=
		t := time.Now()
		result, err := db.Exec("INSERT INTO TBTESTE_WRK (NOME, DTHR, SITUACAO) VALUES (?,?,?)", "Teste "+strconv.Itoa(i), t.Format("20060102150405"), "A")
		if err != nil {
			log.Fatalln("sql.Exec", err)
		}
		if rowsAffected, _ := result.RowsAffected(); rowsAffected != 1 {
			log.Fatalln("sql.RowsAffected", err)
		}
	}
	log.Println("sql.Insert=", time.Since(start)) // 9m23.7323813s  // 1m45.8338232s

	/*


		USE CMS_TESTE_TNB;

		CREATE TABLE TBTESTE_WRK (
		  ID       numeric(11,0)      NOT NULL IDENTITY(1,1),
		  NOME     varchar(150) NOT NULL,
		  DTHR     varchar(14)  NOT NULL,
		  SITUACAO varchar(1)   NOT NULL,
		  PRIMARY KEY (ID)
		);


		CREATE TABLE TBTESTE_API (
		  ID       numeric(11,0)      NOT NULL IDENTITY(1,1),
		  NOME     varchar(150) NOT NULL,
		  DTHR     varchar(14)  NOT NULL,
		  SITUACAO varchar(1)   NOT NULL,
		  PRIMARY KEY (ID)
		) ;

		-- TRUNCATE TABLE TBTESTE_WRK
		-- TRUNCATE TABLE TBTESTE_API


		SELECT COUNT(1), MIN(DTHR) MIN, MAX(DTHR) MAX FROM TBTESTE_WRK WITH(NOLOCK); -- 12:55:03 - 12:56:48 = 00:01:45 = 10.000 registros
		SELECT COUNT(1), MIN(DTHR) MIN, MAX(DTHR) MAX FROM TBTESTE_API WITH(NOLOCK); -- 12:48:54 - 12:50:03 = 00:01:09 = 10.000 registros
	*/

	log.Print("FIM")
}

type Situacao struct {
	Codigo    string `json:"codigo"`
	Descricao string `json:"descricao"`
}
