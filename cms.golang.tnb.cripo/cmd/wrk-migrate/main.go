package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	nurl "net/url"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/atomic"
)

var DefaultMigrationsTable = "schema_migrations"

var (
	ErrDatabaseDirty  = fmt.Errorf("database is dirty")
	ErrNilConfig      = fmt.Errorf("no config")
	ErrNoDatabaseName = fmt.Errorf("no database name")
)

func init() {
	database.Register("sqlite3", &Sqlite{})
}

func main() {
	db, err := sql.Open("sqlite3", "./banco.db")
	if err != nil {
		log.Fatal("1:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("2:", err)
	}

	config := &Config{
		MigrationsTable: "my_migration_table",
	}

	driver, err := WithInstance(db, config)
	if err != nil {
		log.Fatal("3:", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "ql", driver)
	if err != nil {
		log.Fatal("4:", err)
	}

	err = m.Up()
	if err != nil {
		log.Fatal("5:", err)
	}

	_, err = db.Query(fmt.Sprintf("SELECT * FROM %s", config.MigrationsTable))
	if err != nil {
		log.Fatal("6:", err)
	}

	log.Println("ok")
}

type Config struct {
	MigrationsTable string
	DatabaseName    string
	NoTxWrap        bool
}

type Sqlite struct {
	db       *sql.DB
	isLocked atomic.Bool
	config   *Config
}

func WithInstance(instance *sql.DB, config *Config) (database.Driver, error) {
	if config == nil {
		return nil, ErrNilConfig
	}

	if err := instance.Ping(); err != nil {
		return nil, err
	}

	if len(config.MigrationsTable) == 0 {
		config.MigrationsTable = DefaultMigrationsTable
	}

	mx := &Sqlite{
		db:     instance,
		config: config,
	}
	if err := mx.ensureVersionTable(); err != nil {
		return nil, err
	}

	return mx, nil
}

func (m *Sqlite) ensureVersionTable() (err error) {
	if err = m.Lock(); err != nil {
		return err
	}

	defer func() {
		if e := m.Unlock(); e != nil {
			if err == nil {
				err = e
			}
		}
	}()

	query := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (version uint64,dirty bool);
  CREATE UNIQUE INDEX IF NOT EXISTS version_unique ON %s (version);
  `, m.config.MigrationsTable, m.config.MigrationsTable)

	if _, err := m.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (m *Sqlite) Open(url string) (database.Driver, error) {
	purl, err := nurl.Parse(url)
	if err != nil {
		return nil, err
	}
	dbfile := strings.Replace(migrate.FilterCustomQuery(purl).String(), "sqlite3://", "", 1)
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return nil, err
	}

	qv := purl.Query()

	migrationsTable := qv.Get("x-migrations-table")
	if len(migrationsTable) == 0 {
		migrationsTable = DefaultMigrationsTable
	}

	noTxWrap := false
	if v := qv.Get("x-no-tx-wrap"); v != "" {
		noTxWrap, err = strconv.ParseBool(v)
		if err != nil {
			return nil, fmt.Errorf("x-no-tx-wrap: %s", err)
		}
	}

	mx, err := WithInstance(db, &Config{
		DatabaseName:    purl.Path,
		MigrationsTable: migrationsTable,
		NoTxWrap:        noTxWrap,
	})
	if err != nil {
		return nil, err
	}
	return mx, nil
}

func (m *Sqlite) Drop() (err error) {
	query := `SELECT name FROM sqlite_master WHERE type = 'table';`
	tables, err := m.db.Query(query)
	if err != nil {
		return &database.Error{OrigErr: err, Query: []byte(query)}
	}
	defer tables.Close()

	tableNames := make([]string, 0)
	for tables.Next() {
		var tableName string
		if err := tables.Scan(&tableName); err != nil {
			return err
		}
		if len(tableName) > 0 {
			tableNames = append(tableNames, tableName)
		}
	}
	if err := tables.Err(); err != nil {
		return &database.Error{OrigErr: err, Query: []byte(query)}
	}

	if len(tableNames) > 0 {
		for _, t := range tableNames {
			query := "DROP TABLE " + t
			err = m.executeQuery(query)
			if err != nil {
				return &database.Error{OrigErr: err, Query: []byte(query)}
			}
		}
		query := "VACUUM"
		_, err = m.db.Query(query)
		if err != nil {
			return &database.Error{OrigErr: err, Query: []byte(query)}
		}
	}

	return nil
}

func (m *Sqlite) Run(migration io.Reader) error {
	migr, err := io.ReadAll(migration)
	if err != nil {
		return err
	}
	query := string(migr[:])

	if m.config.NoTxWrap {
		return m.executeQueryNoTx(query)
	}
	return m.executeQuery(query)
}

func (m *Sqlite) executeQuery(query string) error {
	tx, err := m.db.Begin()
	if err != nil {
		return &database.Error{OrigErr: err, Err: "transaction start failed"}
	}
	if _, err := tx.Exec(query); err != nil {
		err := tx.Rollback()
		return &database.Error{OrigErr: err, Query: []byte(query)}
	}
	if err := tx.Commit(); err != nil {
		return &database.Error{OrigErr: err, Err: "transaction commit failed"}
	}
	return nil
}

func (m *Sqlite) executeQueryNoTx(query string) error {
	if _, err := m.db.Exec(query); err != nil {
		return &database.Error{OrigErr: err, Query: []byte(query)}
	}
	return nil
}

func (m *Sqlite) SetVersion(version int, dirty bool) error {
	tx, err := m.db.Begin()
	if err != nil {
		return &database.Error{OrigErr: err, Err: "transaction start failed"}
	}

	query := "DELETE FROM " + m.config.MigrationsTable
	if _, err := tx.Exec(query); err != nil {
		return &database.Error{OrigErr: err, Query: []byte(query)}
	}

	if version >= 0 || (version == database.NilVersion && dirty) {
		query := fmt.Sprintf(`INSERT INTO %s (version, dirty) VALUES (?, ?)`, m.config.MigrationsTable)
		if _, err := tx.Exec(query, version, dirty); err != nil {
			err := tx.Rollback()
			return &database.Error{OrigErr: err, Query: []byte(query)}
		}
	}

	if err := tx.Commit(); err != nil {
		return &database.Error{OrigErr: err, Err: "transaction commit failed"}
	}

	return nil
}

func (m *Sqlite) Version() (version int, dirty bool, err error) {
	query := "SELECT version, dirty FROM " + m.config.MigrationsTable + " LIMIT 1"
	err = m.db.QueryRow(query).Scan(&version, &dirty)
	if err != nil {
		return database.NilVersion, false, nil
	}
	return version, dirty, nil
}

func (m *Sqlite) Close() error {
	return m.db.Close()
}

func (m *Sqlite) Lock() error {
	if !m.isLocked.CAS(false, true) {
		return database.ErrLocked
	}
	return nil
}

func (m *Sqlite) Unlock() error {
	if !m.isLocked.CAS(true, false) {
		return database.ErrNotLocked
	}
	return nil
}



package migrate

import (
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	postgresmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jailtonjunior94/financial/pkg/logger"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	ErrMigrateVersion        = errors.New("error checking migration version")
	ErrPostgresMigrateDriver = errors.New("unable to instantiate postgres migration driver")
)

type Migrate struct {
	logger  logger.Logger
	migrate *migrate.Migrate
}

func NewMigrate(logger logger.Logger, db *sql.DB, migratePath, dbName string) (*Migrate, error) {
	if db == nil {
		return nil, ErrPostgresMigrateDriver
	}

	driver, err := postgresmigrate.WithInstance(db, &postgresmigrate.Config{})
	if err != nil {
		return nil, ErrPostgresMigrateDriver
	}

	m, err := migrate.NewWithDatabaseInstance(migratePath, dbName, driver)
	if err != nil {
		return nil, err
	}
	return &Migrate{logger: logger, migrate: m}, nil
}

func (m *Migrate) ExecuteMigration() error {
	_, _, err := m.migrate.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		m.logger.Error(err.Error())
		return ErrMigrateVersion
	}

	err = m.migrate.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		m.logger.Error(err.Error())
		return nil
	}

	if err != nil {
		m.logger.Error(err.Error())
		return err
	}
	return nil
}


package migrate

import (
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/jailtonjunior94/financial/pkg/logger"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewMigrateMySql(logger logger.Logger, db *sql.DB, migratePath, dbName string) (*Migrate, error) {
	if db == nil {
		return nil, ErrPostgresMigrateDriver
	}

	driver, err := mysqlMigrate.WithInstance(db, &mysqlMigrate.Config{})
	if err != nil {
		return nil, ErrPostgresMigrateDriver
	}

	m, err := migrate.NewWithDatabaseInstance(migratePath, dbName, driver)
	if err != nil {
		return nil, err
	}
	return &Migrate{logger: logger, migrate: m}, nil
}

func (m *Migrate) ExecuteMigrationMySql() error {
	_, _, err := m.migrate.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		m.logger.Info(err.Error())
		return ErrMigrateVersion
	}

	err = m.migrate.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		m.logger.Error(err.Error())
		return nil
	}

	if err != nil {
		m.logger.Error(err.Error())
		return err
	}
	return nil
}