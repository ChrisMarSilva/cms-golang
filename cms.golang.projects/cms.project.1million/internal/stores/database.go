package stores

import (
	"context"
	"log/slog"
	"os"

	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	//pgxslog "github.com/mcosta74/pgx-slog"
	pgxslog "github.com/timtoronto634/pgx-slog"
	//"github.com/pgx-contrib/pgxtrace"
	// otelpgx "github.com/quantumsheep/otelpgxpool"
	"github.com/exaring/otelpgx"
)

type Database struct {
	logger *slog.Logger
	config *utils.Config
	Conn   *pgxpool.Pool
}

func NewDatabase(logger *slog.Logger, config *utils.Config) *Database {

	// databaseUrl = fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=%s",
	// 	s.cfg.Database.Driver,
	// 	s.cfg.Database.User,
	// 	s.cfg.Database.Pass,
	// 	s.cfg.Database.Host,
	// 	s.cfg.Database.Port,
	// 	s.cfg.Database.Name,
	// 	s.cfg.Database.SslMode,
	// )

	pgxConfig, err := pgxpool.ParseConfig(config.DatabaseUrl)
	if err != nil {
		logger.Error("Unable to parse database config", slog.Any("error", err))
		os.Exit(1)
	}

	pgxConfig.MaxConns = 3 // Defina o número máximo de conexões
	pgxConfig.MinConns = 3 // Defina o número mínimo de conexões
	// pgxConfig.MaxConnLifetime = defaultMaxConnLifetime
	pgxConfig.MaxConnIdleTime = 0 // Desative o tempo máximo de inatividade da conexão
	// pgxConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	// pgxConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	// pgxConfig.ConnConfig.Tracer = pgxslog.NewTracer(logger)
	// pgxConfig.ConnConfig.Tracer = pgxtrace.CompositeQueryTracer{}
	// pgxConfig.ConnConfig.Tracer = otelpgx.NewTracer()
	// pgxConfig.ConnConfig.Tracer = otelpgx.NewTracer()

	pgxConfig.ConnConfig.Tracer = &MultiQueryTracer{
		Tracers: []pgx.QueryTracer{
			otelpgx.NewTracer(),
			pgxslog.NewTracer(logger),
			//&tracelog.TraceLog{Logger: logger, LogLevel: tracelog.LogLevelTrace},
		},
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		logger.Error("Unable to connect to database", slog.Any("error", err))
		os.Exit(1)
	}

	if err = conn.Ping(context.Background()); err != nil {
		logger.Error("Unable to ping database", slog.Any("error", err))
		os.Exit(1)
	}

	// 	conn.Config().MaxConns = 1000     // Define o número máximo de conexões abertas// Defina o número máximo de conexões
	// 	conn.Config().MinConns = 100      // Defina o número mínimo de conexões
	// 	conn.Config().MaxConnIdleTime = 0 // Define o tempo máximo de ociosidade para 5 minutos// Desative o tempo máximo de inatividade da conexão
	// 	conn.Config().MaxConnLifetime = 0 // Define o tempo máximo de vida (lifetime) para 30 minutos// Desative o tempo máximo de espera para adquirir uma conexão
	// 	conn.Config().MaxConnIdleTime = time.Minute * 3

	return &Database{
		logger: logger,
		config: config,
		Conn:   conn,
	}
}

func (d *Database) Close() {
	if d.Conn != nil {
		d.Conn.Close()
	}
}

type MultiQueryTracer struct {
	Tracers []pgx.QueryTracer
}

func (m *MultiQueryTracer) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	for _, t := range m.Tracers {
		ctx = t.TraceQueryStart(ctx, conn, data)
	}

	return ctx
}

func (m *MultiQueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	for _, t := range m.Tracers {
		t.TraceQueryEnd(ctx, conn, data)
	}
}
