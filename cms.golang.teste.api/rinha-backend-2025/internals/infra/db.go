package infra

// import (
//     "log"
//     "time"

//     "github.com/jmoiron/sqlx"
//     "github.com/lib/pq"
//     "github.com/prometheus/client_golang/prometheus"

//     "github.com/dlmiddlecote/sqlstats"
// )

// var DBStatsCollector prometheus.Collector

// func NewDBOrFatal(conn string) *sqlx.DB {
//     db, err := sqlx.Open("postgres", conn)
//     if err != nil {
//         log.Fatal(err)
//     }
//     db.SetMaxOpenConns(20)
//     db.SetMaxIdleConns(10)
//     db.SetConnMaxLifetime(time.Hour)

//     collector := sqlstats.NewStatsCollector("rinha_db", db.DB)
//     prometheus.MustRegister(collector)
//     DBStatsCollector = collector
//     return db
// }
