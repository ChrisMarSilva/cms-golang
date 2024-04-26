package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func GetDatabase() (*sql.DB, error) {
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	db, err := sql.Open("sqlite3", "./banco.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

type Database struct {
	*sqlx.DB
}


func Connect() (*Database, error) {
	config := configs.New()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln("error connecting to database:", err)
	}

	return &Database{db}, nil
}



func ConnectDB(config *Config) {
var err error
dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
log.Fatal("Failed to connect to the Database! \n", err.Error())
os.Exit(1)
}

DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
DB.Logger = logger.Default.LogMode(logger.Info)

log.Println("Running Migrations")
err = DB.AutoMigrate(&models.User{})
if err != nil {
log.Fatal("Migration Failed:  \n", err.Error())
os.Exit(1)
}

log.Println("🚀 Connected Successfully to the Database")
}


var Db *gorm.DB

func InitDb() *gorm.DB {
Db = connectDB()
return Db
}

func connectDB() *gorm.DB {
var err error
host := os.Getenv("DB_HOST")
username := os.Getenv("DB_USER")
password := os.Getenv("DB_PASSWORD")
dbname := os.Getenv("DB_NAME")
port := os.Getenv("DB_PORT")

dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
//log.Println("dsn : ", dsn)
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

if err != nil {
log.Fatal("Error connecting to database :", err)
return nil
}
log.Println("`Successfully connected to the database")

return db
}


package driver

import (
"database/sql"

_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
SQL *sql.DB
}

var dbConn = &DB{}

func ConnectSQL(dsn string)(*DB, error){
db, err := sql.Open("pgx", dsn)

if err := db.Ping(); err != nil {
panic(err)
}

dbConn.SQL = db

return dbConn, err
}




type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectSQL(dsn string)(*DB, error){
	db, err := sql.Open("pgx", dsn)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	dbConn.SQL = db

	return dbConn, err
}




import (
	"github.com/jackc/pgx/v5"x
	"github.com/jackc/pgx/v5/pgxpool"
)

var defaultMaxConns = int32(4)
var defaultMinConns = int32(0)
var defaultMaxConnLifetime = time.Hour
var defaultMaxConnIdleTime = time.Minute * 30
var defaultHealthCheckPeriod = time.Minute

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())



// NewWithConfig creates a new Pool. config must have been created by [ParseConfig].
func NewWithConfig(ctx context.Context, config *Config) (*Pool, error) {
	// Default values are set in ParseConfig. Enforce initial creation by ParseConfig rather than setting defaults from
	// zero values.
	if !config.createdByParseConfig {
		panic("config must be created by ParseConfig")
	}

	p := &Pool{
		config:                config,
		beforeConnect:         config.BeforeConnect,
		afterConnect:          config.AfterConnect,
		beforeAcquire:         config.BeforeAcquire,
		afterRelease:          config.AfterRelease,
		beforeClose:           config.BeforeClose,
		minConns:              config.MinConns,
		maxConns:              config.MaxConns,
		maxConnLifetime:       config.MaxConnLifetime,
		maxConnLifetimeJitter: config.MaxConnLifetimeJitter,
		maxConnIdleTime:       config.MaxConnIdleTime,
		healthCheckPeriod:     config.HealthCheckPeriod,
		healthCheckChan:       make(chan struct{}, 1),
		closeChan:             make(chan struct{}),
	}

	var err error
	p.p, err = puddle.NewPool(
		&puddle.Config[*connResource]{
			Constructor: func(ctx context.Context) (*connResource, error) {
				atomic.AddInt64(&p.newConnsCount, 1)
				connConfig := p.config.ConnConfig.Copy()

				// Connection will continue in background even if Acquire is canceled. Ensure that a connect won't hang forever.
				if connConfig.ConnectTimeout <= 0 {
					connConfig.ConnectTimeout = 2 * time.Minute
				}

				if p.beforeConnect != nil {
					if err := p.beforeConnect(ctx, connConfig); err != nil {
						return nil, err
					}
				}

				conn, err := pgx.ConnectConfig(ctx, connConfig)
				if err != nil {
					return nil, err
				}

				if p.afterConnect != nil {
					err = p.afterConnect(ctx, conn)
					if err != nil {
						conn.Close(ctx)
						return nil, err
					}
				}

				jitterSecs := rand.Float64() * config.MaxConnLifetimeJitter.Seconds()
				maxAgeTime := time.Now().Add(config.MaxConnLifetime).Add(time.Duration(jitterSecs) * time.Second)

				cr := &connResource{
					conn:       conn,
					conns:      make([]Conn, 64),
					poolRows:   make([]poolRow, 64),
					poolRowss:  make([]poolRows, 64),
					maxAgeTime: maxAgeTime,
				}

				return cr, nil
			},
			Destructor: func(value *connResource) {
				ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
				conn := value.conn
				if p.beforeClose != nil {
					p.beforeClose(conn)
				}
				conn.Close(ctx)
				select {
				case <-conn.PgConn().CleanupDone():
				case <-ctx.Done():
				}
				cancel()
			},
			MaxSize: config.MaxConns,
		},
	)
	if err != nil {
		return nil, err
	}

	go func() {
		p.createIdleResources(ctx, int(p.minConns))
		p.backgroundHealthCheck()
	}()

	return p, nil
}

var Connection *pgxpool.Pool

func Connect() error {
	var err error
	Connection, err = pgxpool.Connect(context.Background(), config.DatabaseURL)

	return err
}

func Close() {
	if Connection == nil {
		return
	}
	Connection.Close()
}



ctxTimeout, ctxCancel := context.WithTimeout(context.Background(), time.Second*3)
defer ctxCancel()

db, err := sqlx.ConnectContext(ctxTimeout, "postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",  postgresConfig.Host, postgresConfig.Port, postgresConfig.Username, postgresConfig.Password, postgresConfig.DBName, postgresConfig.SSLMode))
 if err != nil {
	 panic(err)
 }
 