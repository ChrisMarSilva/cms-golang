 DROP TABLE IF EXISTS go_device;
DROP TABLE IF EXISTS python_device;

   CREATE TABLE "go_device" ("id" serial primary key, "uuid" character varying(255), "mac" character varying(255), "firmware" character varying(255), "created_at" timestamp NOT NULL, "updated_at" timestamp NOT NULL);
    CREATE TABLE "python_device" ("id" serial primary key, "uuid" character varying(255), "mac" character varying(255), "firmware" character varying(255), "created_at" timestamp NOT NULL, "updated_at" timestamp NOT NULL);
    COMMIT;



server.go


package main

import (
	"context"
	"net/http"

	mon "github.com/antonputra/go-utils/monitoring"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/jackc/pgx/v5/pgxpool"
	jsoniter "github.com/json-iterator/go"
	"github.com/prometheus/client_golang/prometheus"
)

// json - jsoniter object
var json = jsoniter.ConfigFastest

type server struct {
	db    *pgxpool.Pool
	cache *memcache.Client
	cfg   *Config
	m     *mon.Metrics
}

type resp struct {
	Msg string `json:"msg"`
}

func newServer(ctx context.Context, cfg *Config, reg *prometheus.Registry) *server {
	m := mon.NewMetrics(reg)
	s := server{cfg: cfg, m: m}
	s.db = dbConnect(ctx, cfg)
	s.cache = cacheConnect(cfg)

	return &s
}

func renderJSON(w http.ResponseWriter, value any, status int) {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := enc.Encode(value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


routes.go


package main

import (
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/antonputra/go-utils/util"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
)

func (s *server) getHealth(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "OK")
}

func (s *server) getDevices(w http.ResponseWriter, req *http.Request) {
	device := []Device{
		{
			Id:        1,
			Uuid:      "9add349c-c35c-4d32-ab0f-53da1ba40a2a",
			Mac:       "EF-2B-C4-F5-D6-34",
			Firmware:  "2.1.5",
			CreatedAt: "2024-05-28T15:21:51.137Z",
			UpdatedAt: "2024-05-28T15:21:51.137Z",
		},
		{
			Id:        2,
			Uuid:      "d2293412-36eb-46e7-9231-af7e9249fffe",
			Mac:       "E7-34-96-33-0C-4C",
			Firmware:  "1.0.3",
			CreatedAt: "2024-01-28T15:20:51.137Z",
			UpdatedAt: "2024-01-28T15:20:51.137Z",
		},
		{
			Id:        3,
			Uuid:      "eee58ca8-ca51-47a5-ab48-163fd0e44b77",
			Mac:       "68-93-9B-B5-33-B9",
			Firmware:  "4.3.1",
			CreatedAt: "2024-08-28T15:18:21.137Z",
			UpdatedAt: "2024-08-28T15:18:21.137Z",
		},
	}

	renderJSON(w, &device, 200)
}

func (s *server) saveDevice(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	decoder := json.NewDecoder(req.Body)
	d := new(Device)
	err := decoder.Decode(&d)
	if err != nil {
		s.m.Errors.With(prometheus.Labels{"op": "decode", "db": "postgres"}).Add(1)
		util.Warn(err, "failed to decode device")
		renderJSON(w, resp{Msg: "failed to decode device"}, 400)
		return
	}

	// to match elixir
	now := time.Now().Format(time.RFC3339Nano)
	d.Uuid = uuid.New().String()
	d.CreatedAt = now
	d.UpdatedAt = now

	err = d.insert(ctx, s.db, s.m)
	if err != nil {
		s.m.Errors.With(prometheus.Labels{"op": "insert", "db": "postgres"}).Add(1)
		util.Warn(err, "failed to save device in postgres")
		renderJSON(w, resp{Msg: "failed to save device in postgres"}, 400)
		return
	}
	slog.Debug("device saved in postgres", "id", d.Id, "mac", d.Mac, "firmware", d.Firmware)

	err = d.set(s.cache, s.m, s.cfg.CacheConfig.Expiration)
	if err != nil {
		s.m.Errors.With(prometheus.Labels{"op": "set", "db": "memcache"}).Add(1)
		util.Warn(err, "failed to save device in memcache")
		renderJSON(w, resp{Msg: "failed to save device in memcache"}, 400)
		return
	}

	renderJSON(w, &d, 201)
}


/main.go



package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	mon "github.com/antonputra/go-utils/monitoring"
	"github.com/prometheus/client_golang/prometheus"
)

func main() {
	cp := flag.String("config", "", "path to the config")
	flag.Parse()

	ctx, done := context.WithCancel(context.Background())
	defer done()

	cfg := new(Config)
	cfg.LoadConfig(*cp)

	if cfg.Debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	reg := prometheus.NewRegistry()
	mon.StartPrometheusServer(cfg.MetricsPort, reg)

	s := newServer(ctx, cfg, reg)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/devices", s.getDevices)
	mux.HandleFunc("POST /api/devices", s.saveDevice)
	mux.HandleFunc("GET /healthz", s.getHealth)

	appPort := fmt.Sprintf(":%d", cfg.AppPort)
	log.Printf("Starting the web server on port %d", cfg.AppPort)
	log.Fatal(http.ListenAndServe(appPort, mux))
}


/device.go



package main

import (
	"context"
	"log/slog"
	"time"

	mon "github.com/antonputra/go-utils/monitoring"
	"github.com/antonputra/go-utils/util"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
)

type Device struct {
	Id        int    `json:"id"`
	Uuid      string `json:"uuid"`
	Mac       string `json:"mac"`
	Firmware  string `json:"firmware"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (d *Device) insert(ctx context.Context, db *pgxpool.Pool, m *mon.Metrics) (err error) {
	now := time.Now()
	defer func() {
		if err == nil {
			m.Duration.With(prometheus.Labels{"op": "insert", "db": "postgres"}).Observe(time.Since(now).Seconds())
		}
	}()

	// Execute the query to create a new device record (pgx automatically prepares and caches statements by default).
	sql := `INSERT INTO "go_device" (uuid, mac, firmware, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err = db.QueryRow(ctx, sql, d.Uuid, d.Mac, d.Firmware, d.CreatedAt, d.UpdatedAt).Scan(&d.Id)

	return util.Annotate(err, "db.Exec failed")
}

func (d *Device) set(mc *memcache.Client, m *mon.Metrics, exp int32) (err error) {
	b, err := json.Marshal(d)
	if err != nil {
		return util.Annotate(err, "json.Marshal failed")
	}

	now := time.Now()
	err = mc.Set(&memcache.Item{Key: d.Uuid, Value: b, Expiration: exp})
	if err != nil {
		return util.Annotate(err, "mc.Set failed")
	}
	m.Duration.With(prometheus.Labels{"op": "set", "db": "memcache"}).Observe(time.Since(now).Seconds())

	slog.Debug("device saved in memcache", "uuid", d.Uuid, "value", string(b))

	return nil
}


db.go


package main

import (
	"context"
	"fmt"

	"github.com/antonputra/go-utils/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func dbConnect(ctx context.Context, cfg *Config) *pgxpool.Pool {
	url := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?pool_max_conns=%d",
		cfg.DatabaseConfig.User, cfg.DatabaseConfig.Password, cfg.DatabaseConfig.Host, cfg.DatabaseConfig.Database, cfg.DatabaseConfig.MaxConnections)

	dbpool, err := pgxpool.New(ctx, url)
	util.Fail(err, "failed to create connection pool")

	return dbpool
}


config.yaml


---
debug: true
appPort: 8080
metricsPort: 8081
database:
  user: go_app
  password: devops123
  host: 192.168.50.74
  database: mydb
  maxConnections: 500
cache:
  host: 192.168.50.19
  expirationS: 20



/config.go


package main

import (
	"os"

	"github.com/antonputra/go-utils/util"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Debug          bool           `yaml:"debug"`
	AppPort        int            `yaml:"appPort"`
	MetricsPort    int            `yaml:"metricsPort"`
	DatabaseConfig DatabaseConfig `yaml:"database"`
	CacheConfig    CacheConfig    `yaml:"cache"`
}

type DatabaseConfig struct {
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	Host           string `yaml:"host"`
	Database       string `yaml:"database"`
	MaxConnections int    `yaml:"maxConnections"`
}

type CacheConfig struct {
	Host       string `yaml:"host"`
	Expiration int32  `yaml:"expirationS"`
}

func (c *Config) LoadConfig(path string) {
	f, err := os.ReadFile(path)
	util.Fail(err, "failed to read config")

	err = yaml.Unmarshal(f, c)
	util.Fail(err, "failed to parse config")
}


/cache.go


package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func cacheConnect(cfg *Config) *memcache.Client {
	mc := memcache.New(fmt.Sprintf("%s:11211", cfg.CacheConfig.Host))

	return mc
}
