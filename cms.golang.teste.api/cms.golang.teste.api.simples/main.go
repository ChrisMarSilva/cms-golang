package main

// go mod init github.com/chrismarsilva/cms.golang.teste.api.simples
// go get -u github.com/gorilla/mux
// go get -u github.com/prometheus/client_golang
// go get -u github.com/bytedance/sonic
// go get -u go.opentelemetry.io/otel
// go get -u go.opentelemetry.io/otel/exporters/stdout/stdoutmetric
// go get -u go.opentelemetry.io/otel/exporters/stdout/stdouttrace
// go get -u go.opentelemetry.io/otel/propagation
// go get -u go.opentelemetry.io/otel/sdk/metric
// go get -u go.opentelemetry.io/otel/sdk/resource
// go get -u go.opentelemetry.io/otel/sdk/trace
// go get -u go.opentelemetry.io/otel/semconv/v1.24.0
// go get -u go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
// go get -u go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp
// go get -u github.com/mattn/go-sqlite3
// go mod tidy
// go run main.go
// go run .

// go install github.com/cosmtrek/air@latest
// air init
// air

// docker-compose down
// docker-compose up -d --build

// docker rm -f $(docker ps -a -q)
// docker run -it rinha-backend-2024-api01:latest

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	service      = "myapp"
	environment  = "production"
	tracer       trace.Tracer
	otlpEndpoint = "jaeger:4318" // "127.0.0.1:4318"
)

func init() {
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tp := newTraceProvider(ctx)
	tracer = tp.Tracer(service)

	defer func(ctx context.Context) {
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	reg := prometheus.NewRegistry()
	reg.MustRegister()
	m := NewMetrics(reg)
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}) // promhttp.HandlerOpts{}

	router := mux.NewRouter().StrictSlash(true)

	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			route := mux.CurrentRoute(r)
			path, _ := route.GetPathTemplate()

			w.Header().Set("Content-Type", "application/json")
			//w.WriteHeader(http.StatusBadRequest)

			start := time.Now()
			timer := prometheus.NewTimer(m.httpDuration.WithLabelValues(path))
			//rw := NewResponseWriter(w)
			//next.ServeHTTP(rw, r)
			next.ServeHTTP(w, r)
			elapsed := time.Since(start).Seconds()

			statusCode := w.Header().Get("Status-Code")
			if statusCode == "" {
				statusCode = strconv.Itoa(http.StatusBadRequest)
			}

			m.responseStatus.WithLabelValues(statusCode).Inc()
			m.totalRequests.WithLabelValues(path).Inc()
			//m.totalRequests.With(prometheus.Labels{"path": r.URL.Path}).Inc()
			//m.httpDuration.WithLabelValues(path)
			//m.httpDuration.Observe(elapsed)

			log.Printf("%s %s %v %v", r.Method, r.URL.Path, statusCode, elapsed)

			timer.ObserveDuration()
		})
	})

	//router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/user", CreateUserHandler).Methods("POST")
	router.HandleFunc("/user", GetUsersHandler).Methods("GET")
	router.HandleFunc("/user/{id}", GetUserHandler).Methods("GET")
	router.HandleFunc("/user/{id}", UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/user/{id}", DeleteUserHandler).Methods("DELETE")

	routerMetrics := mux.NewRouter().StrictSlash(true)
	routerMetrics.Handle("/metricsOld", promhttp.Handler())
	routerMetrics.Handle("/metrics", promHandler)

	go func() {
		log.Println("Localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	}()

	go func() {
		log.Println("Localhost:8081")
		log.Fatal(http.ListenAndServe(":8081", routerMetrics))
	}()

	select {}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// _, span := tracer.Start(r.Context(), "HTTP GET /")
	// defer span.End()

	json.NewEncoder(w).Encode("Hello, World!")
	// w.Write([]byte("Hello, World!"))
	//fmt.Fprintf(w, "Hello, World!")
	w.WriteHeader(http.StatusOK)
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func db(ctx context.Context) (*sql.DB, error) {
	_, span_Banco := tracer.Start(ctx, "sql.Open)")
	db, err := sql.Open("sqlite3", "./banco.db")
	if err != nil {
		span_Banco.SetStatus(codes.Error, err.Error())
		span_Banco.End()
		return nil, err
	}
	span_Banco.End()

	// _, span_CREATE := tracer.Start(ctx, "db.Exe(CREATE TABLE)")
	// const create string = `CREATE TABLE IF NOT EXISTS users ( ID INTEGER NOT NULL PRIMARY KEY, Name TEXT );`
	// _, err = db.Exec(create)
	// if err != nil {
	// 	span_CREATE.SetStatus(codes.Error, err.Error())
	// 	span_CREATE.End()
	// 	return nil, err
	// }
	// span_CREATE.End()

	return db, err
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span_Main := tracer.Start(r.Context(), "HTTP GET /user")
	defer span_Main.End()

	db, err := db(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, span_SELECT := tracer.Start(ctx, "db.Query(SELECT)")
	rows, err := db.Query("SELECT ID, Name FROM users ORDER BY ID")
	if err != nil {
		span_SELECT.SetStatus(codes.Error, err.Error())
		span_SELECT.End()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	span_SELECT.End()
	defer rows.Close()

	_, span_Next := tracer.Start(ctx, "rows.Next()")
	users := []*User{}
	for rows.Next() {
		var user User // user := User{}
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			span_Next.SetStatus(codes.Error, err.Error())
			span_Next.End()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, &user)
	}
	span_Next.End()

	_, span_Json := tracer.Start(ctx, "json.NewEncoder().Encode()")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		span_Json.SetStatus(codes.Error, err.Error())
		span_Json.End()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	span_Json.End()

	w.WriteHeader(http.StatusOK)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span_Main := tracer.Start(r.Context(), "HTTP GET /user/id")
	defer span_Main.End()

	// params := mux.Vars(r)
	// id := params["id"]
	id := r.PathValue("id")
	span_Main.SetAttributes(attribute.Key("id").String(id))

	db, err := db(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, span_SELECT := tracer.Start(ctx, "db.Query(SELECT)")
	rows, err := db.Query("SELECT ID, Name FROM users WHERE ID = ?", id)
	if err != nil {
		span_SELECT.SetStatus(codes.Error, err.Error())
		span_SELECT.End()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	span_SELECT.End()
	defer rows.Close()

	user := User{}

	_, span_Next := tracer.Start(ctx, "rows.Next()")
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			span_Next.SetStatus(codes.Error, err.Error())
			span_Next.End()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	span_Next.End()

	_, span_Json := tracer.Start(ctx, "json.NewEncoder().Encode()")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		span_Json.SetStatus(codes.Error, err.Error())
		span_Json.End()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	span_Json.End()

	w.WriteHeader(http.StatusOK)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span_Main := tracer.Start(r.Context(), "HTTP POST /user")
	defer span_Main.End()

	db, err := db(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, span_INSERT := tracer.Start(r.Context(), "db.Exec(INSERT)")
	_, err = db.Exec("INSERT INTO users (ID, Name) VALUES (?, ?)", user.ID, user.Name)
	if err != nil {
		span_INSERT.SetStatus(codes.Error, err.Error())
		span_INSERT.End()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	span_INSERT.End()

	_, span_Json := tracer.Start(ctx, "json.NewEncoder().Encode()")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		span_Json.SetStatus(codes.Error, err.Error())
		span_Json.End()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	span_Json.End()

	w.WriteHeader(http.StatusCreated)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span_Main := tracer.Start(r.Context(), "HTTP PUT /user/id")
	defer span_Main.End()

	// params := mux.Vars(r)
	// id := params["id"]
	id := r.PathValue("id")
	span_Main.SetAttributes(attribute.Key("id").String(id))

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := db(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, span_UPDATE := tracer.Start(ctx, "db.Query(UPDATE)")
	_, err = db.Exec("UPDATE users SET name = ? WHERE ID = ?", user.Name, id)
	if err != nil {
		span_UPDATE.SetStatus(codes.Error, err.Error())
		span_UPDATE.End()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	span_UPDATE.End()

	_, span_Json := tracer.Start(ctx, "json.NewEncoder().Encode()")
	err = json.NewEncoder(w).Encode("")
	if err != nil {
		span_Json.SetStatus(codes.Error, err.Error())
		span_Json.End()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	span_Json.End()

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span_Main := tracer.Start(r.Context(), "HTTP GET /user/id")
	defer span_Main.End()

	// params := mux.Vars(r)
	// id := params["id"]
	id := r.PathValue("id")
	span_Main.SetAttributes(attribute.Key("id").String(id))

	db, err := db(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	_, span_SELECT := tracer.Start(ctx, "db.Query(DELETE)")
	_, err = db.Exec("DELETE FROM users WHERE ID = ?", id)
	if err != nil {
		span_SELECT.SetStatus(codes.Error, err.Error())
		span_SELECT.End()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	span_SELECT.End()

	_, span_Json := tracer.Start(ctx, "json.NewEncoder().Encode()")
	err = json.NewEncoder(w).Encode("")
	if err != nil {
		span_Json.SetStatus(codes.Error, err.Error())
		span_Json.End()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	span_Json.End()

	w.WriteHeader(http.StatusNoContent)
}

func newTraceProvider(ctx context.Context) *sdktrace.TracerProvider {
	insecureOpt := otlptracehttp.WithInsecure()
	endpointOpt := otlptracehttp.WithEndpoint(otlpEndpoint)
	exp, err := otlptracehttp.New(ctx, insecureOpt, endpointOpt)
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
		return nil
	}

	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
			semconv.ServiceVersionKey.String("1.0.0"),
			semconv.DeploymentEnvironmentKey.String(environment),
		),
	)
	if err != nil {
		log.Fatalf("failed to initialize exporter 2: %v", err)
		return nil
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}))

	return tp
}

type metrics struct {
	totalRequests  *prometheus.CounterVec
	responseStatus *prometheus.CounterVec
	httpDuration   *prometheus.HistogramVec
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		totalRequests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: service,
			Name:      "http_requests_total",
			Help:      "Number of get requests.",
		}, []string{"path"}),
		responseStatus: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: service,
			Name:      "response_status",
			Help:      "Status of HTTP response",
		}, []string{"status"}),
		httpDuration: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: service,
			Name:      "http_response_time_seconds",
			Help:      "Duration of HTTP requests.",
		}, []string{"path"}),
	}

	reg.MustRegister(m.totalRequests)
	reg.MustRegister(m.responseStatus)
	reg.MustRegister(m.httpDuration)

	return m
}

/*

var (
	// users        []User
	//dvs          []Device
	//version      string
	service      = "myapp" // "cms.teste.trace"
	environment  = "production"
	tracer       trace.Tracer
	otlpEndpoint string
)

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

type metrics struct {
	// users         prometheus.Gauge
	// devices       prometheus.Gauge
	// info          *prometheus.GaugeVec
	// upgrades      *prometheus.CounterVec
	// duration      *prometheus.HistogramVec
	// loginDuration prometheus.Summary
	totalRequests  *prometheus.CounterVec
	responseStatus *prometheus.CounterVec
	httpDuration   *prometheus.HistogramVec
}

type registerDevicesHandler struct {
	metrics *metrics
}

type manageDevicesHandler struct {
	metrics *metrics
}

func init() {
	//version = "2.10.5"

	// dvs = []Device{
	// 	{1, "5F-33-CC-1F-43-82", "2.1.6"},
	// 	{2, "EF-2B-C4-F5-D6-34", "2.1.6"},
	// }

	// users = append(users, User{ID: 1, Name: "John Doe"})
	// users = append(users, User{ID: 2, Name: "Koko Snow"})
	// users = append(users, User{ID: 3, Name: "Francis Sunday"})

	otlpEndpoint = "127.0.0.1:4318" // "127.0.0.1:4318" // "http://localhost:4318/" // "http://jaeger:4318" // os.Getenv("OTLP_ENDPOINT") // OTEL_EXPORTER_OTLP_ENDPOINT
	if otlpEndpoint == "" {
		log.Fatalln("You MUST set OTLP_ENDPOINT env variable!")
	}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) //ctx := context.Background()
	defer cancel()

	exp, err := newOTLPExporter(ctx) // newConsoleExporter() // For testing to print out traces to the console
	if err != nil {
		log.Fatalf("failed to initialize exporter: %v", err)
	}

	tp := newTraceProvider(exp) // Create a new tracer provider with a batch span processor and the given exporter.
	otel.SetTracerProvider(tp)
	tracer = tp.Tracer(service)       // Finally, set the tracer that can be used for this package.
	defer func(ctx context.Context) { // Cleanly shutdown and flush telemetry when the application exits. // Handle shutdown properly so nothing leaks. defer func() { _ = tp.Shutdown(ctx) }()
		ctx, cancel = context.WithTimeout(ctx, time.Second*5) // Do not make the application hang when it is shutdown.
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	reg := prometheus.NewRegistry()
	reg.MustRegister() // reg.MustRegister(collectors.NewGoCollector())

	m := NewMetrics(reg)
	//m.devices.Set(float64(0))                                 // float64(len(dvs))
	//m.users.Set(float64(0))                                   // float64(len(users))
	//m.info.With(prometheus.Labels{"version": "1.0.0"}).Set(1) // m.info.With(prometheus.Labels{"version": version}).Set(1)

	// rdh := registerDevicesHandler{metrics: m}
	// mdh := manageDevicesHandler{metrics: m}
	// lh := loginHandler{}
	// mlh := middleware(lh, m)
	// promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})

	// router.HandleFunc("/devices", getDevices).Methods("GET")
	// router.HandleFunc("/devices", createDevice).Methods("POST")
	// router.Handle("/devices", rdh)
	// router.Handle("/devices", mdh)
	// router.Handle("/login", mlh)

		span_Banco.SetStatus(codes.Error, err.Error())
		span_Banco.SetAttributes(attribute.Key("erro-Banco").String(err.Error()))
		span_Banco.RecordError(err)
}

// Console Exporter, only for testing
func newConsoleExporter() (oteltrace.SpanExporter, error) {
	return stdouttrace.New()
}

// OTLP Exporter
func newOTLPExporter(ctx context.Context) (oteltrace.SpanExporter, error) {
	insecureOpt := otlptracehttp.WithInsecure()             // Change default HTTPS -> HTTP
	endpointOpt := otlptracehttp.WithEndpoint(otlpEndpoint) // Update default OTLP reciver endpoint
	return otlptracehttp.New(ctx, insecureOpt, endpointOpt)
}


func sleep(ms int) {
	rand.Seed(time.Now().UnixNano())
	now := time.Now()
	n := rand.Intn(ms + now.Second())
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func (rdh registerDevicesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getDevices(w, r, rdh.metrics)
	case "POST":
		createDevice(w, r, rdh.metrics)
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (mdh manageDevicesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		upgradeDevice(w, r, mdh.metrics)
	default:
		w.Header().Set("Allow", "PUT")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func getDevices(w http.ResponseWriter, r *http.Request, m *metrics) {
	//now := time.Now()

	ctx, span := tracer.Start(r.Context(), "HTTP GET /devices")
	defer span.End()

	// Simulate Database call to fetch connected devices.
	db(ctx)

	// b, err := sonic.Marshal(dvs) //b, err := json.Marshal(dvs)
	//  if err != nil {
	//  	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// sleep(200)

	// labels := prometheus.Labels{
	// 	//"client": "Server A", // defines the client server
	// 	//"server": "Server B", // defines the outbound request server
	// 	"method": "GET", // HTTP method
	// 	//"route":  "/",        // Request route
	// 	"status": "200", // Response status
	// }

	//m.duration.With(prometheus.Labels{"method": "GET", "status": "200"}).Observe(time.Since(now).Seconds())
	//m.upgrades.With(prometheus.Labels{"type": "router"}).Inc()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//w.Write(b)
	w.Write([]byte("Devices"))
}

func createDevice(w http.ResponseWriter, r *http.Request, m *metrics) {
	_, span := tracer.Start(r.Context(), "HTTP GET /devices")
	defer span.End()

	var dv Device
	err := json.NewDecoder(r.Body).Decode(&dv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//dvs = append(dvs, dv)
	sleep(1000)

	//m.devices.Set(float64(len(dvs)))
	//m.upgrades.With(prometheus.Labels{"type": "router"}).Inc()

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Device created!"))
}

func upgradeDevice(w http.ResponseWriter, r *http.Request, m *metrics) {
	_, span := tracer.Start(r.Context(), "HTTP GET /devices")
	defer span.End()

	path := strings.TrimPrefix(r.URL.Path, "/devices/")

	id, err := strconv.Atoi(path)
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}

	var dv Device
	err = json.NewDecoder(r.Body).Decode(&dv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// for i := range dvs {
	// 	if dvs[i].ID == id {
	// 		dvs[i].Firmware = dv.Firmware
	// 	}
	// }

	sleep(1000)
	//m.upgrades.With(prometheus.Labels{"type": "router"}).Inc()

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Upgrading..."))
}

type loginHandler struct {
}

func middleware(next http.Handler, m *metrics) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//now := time.Now()
		next.ServeHTTP(w, r)
		//m.loginDuration.Observe(time.Since(now).Seconds())
	})
}
func (l loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), "HTTP GET /login")
	defer span.End()

	sleep(200)
	w.Write([]byte("Welcome to the app!"))
}


func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		totalRequests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: service,
			Name:      "http_requests_total",
			Help:      "Number of get requests.",
		}, []string{"path"}),
		responseStatus: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: service,
			Name:      "response_status",
			Help:      "Status of HTTP response",
		}, []string{"status"}),
		httpDuration: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: service,
			Name:      "http_response_time_seconds",
			Help:      "Duration of HTTP requests.",
		}, []string{"path"}),

		// devices: prometheus.NewGauge(prometheus.GaugeOpts{
		// 	Namespace: service,
		// 	Name:      "connected_devices",
		// 	Help:      "Number of currently connected devices.",
		// }),
		// users: prometheus.NewGauge(prometheus.GaugeOpts{
		// 	Namespace: service,
		// 	Name:      "connected_users",
		// 	Help:      "Number of currently connected users.",
		// }),
		// info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
		// 	Namespace: service,
		// 	Name:      "info",
		// 	Help:      "Information about the My App environment.",
		// }, []string{"version"}),
		// upgrades: prometheus.NewCounterVec(prometheus.CounterOpts{
		// 	Namespace: service,
		// 	Name:      "device_upgrade_total",
		// 	Help:      "Number of upgraded devices.",
		// }, []string{"type"}), // []string{"client", "server", "method", "route", "status"}), //
		// duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
		// 	Namespace: service,
		// 	Name:      "request_duration_seconds",
		// 	Help:      "Duration of the request.",
		// 	Buckets:   []float64{0.1, 0.15, 0.2, 0.25, 0.3}, // []float64{0.1, 0.5, 1, 2, 5}, //
		// }, []string{"status", "method"}), // []string{"client", "server", "method", "route", "status"}), //
		// loginDuration: prometheus.NewSummary(prometheus.SummaryOpts{
		// 	Namespace:  service,
		// 	Name:       "login_request_duration_seconds",
		// 	Help:       "Duration of the login request.",
		// 	Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		// }),
	}

	// var totalRequests = prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Name: "http_requests_total",
	// 		Help: "Number of get requests.",
	// 	},
	// 	[]string{"path"},
	// )

	// var responseStatus = prometheus.NewCounterVec(
	// 	prometheus.CounterOpts{
	// 		Name: "response_status",
	// 		Help: "Status of HTTP response",
	// 	},
	// 	[]string{"status"},
	// )

	// var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	// 	Name: "http_response_time_seconds",
	// 	Help: "Duration of HTTP requests.",
	// }, []string{"path"})

	 reg.MustRegister(m.devices, m.users, m.info, m.upgrades, m.duration, m.loginDuration)
	 reg.MustRegister(m.devices)
	reg.MustRegister(m.users)
	 reg.MustRegister(m.info)
	reg.MustRegister(m.upgrades)
	reg.MustRegister(m.duration)
	reg.MustRegister(m.loginDuration)
	reg.MustRegister(m.totalRequests)  // prometheus.Register(totalRequests)
	reg.MustRegister(m.responseStatus) // prometheus.Register(responseStatus)
	reg.MustRegister(m.httpDuration)   // prometheus.Register(httpDuration)

	return m
}
*/
