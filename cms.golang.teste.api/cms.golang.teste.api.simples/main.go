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
// go get -u get github.com/mattn/go-sqlite3
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
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	oteltrace "go.opentelemetry.io/otel/sdk/trace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

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

type loginHandler struct {
}

var (
	// users        []User
	//dvs          []Device
	//version      string
	service      = "myapp" // "cms.teste.trace"
	environment  = "production"
	tracer       trace.Tracer
	otlpEndpoint string
)

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

func init() {
	//version = "2.10.5"

	// dvs = []Device{
	// 	{1, "5F-33-CC-1F-43-82", "2.1.6"},
	// 	{2, "EF-2B-C4-F5-D6-34", "2.1.6"},
	// }

	// users = append(users, User{ID: 1, Name: "John Doe"})
	// users = append(users, User{ID: 2, Name: "Koko Snow"})
	// users = append(users, User{ID: 3, Name: "Francis Sunday"})

	otlpEndpoint = "127.0.0.1:4318" // "127.0.0.1:4318" // "http://localhost:16686/api/traces" // "http://localhost:4318/" // "http://jaeger:4318" // os.Getenv("OTLP_ENDPOINT") // OTEL_EXPORTER_OTLP_ENDPOINT
	if otlpEndpoint == "" {
		log.Fatalln("You MUST set OTLP_ENDPOINT env variable!")
	}
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

	router := mux.NewRouter()

	//router.Use(prometheusMiddleware)
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			route := mux.CurrentRoute(r)
			path, _ := route.GetPathTemplate()

			// 	start := time.Now()
			timer := prometheus.NewTimer(m.httpDuration.WithLabelValues(path))
			rw := NewResponseWriter(w)
			next.ServeHTTP(rw, r)
			// 	elapsed := time.Since(start).Seconds()

			statusCode := rw.statusCode

			m.responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
			m.totalRequests.WithLabelValues(path).Inc()

			timer.ObserveDuration()
		})
	})

	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/user", GetUsersHandler).Methods("GET")
	router.HandleFunc("/user/{id}", GetUserHandler).Methods("GET")
	router.HandleFunc("/user/{id}", CreateUserHandler).Methods("POST")
	router.HandleFunc("/user/{id}", DeleteUserHandler).Methods("DELETE")
	// router.HandleFunc("/devices", getDevices).Methods("GET")
	// router.HandleFunc("/devices", createDevice).Methods("POST")
	// router.Handle("/devices", rdh)
	// router.Handle("/devices", mdh)
	// router.Handle("/login", mlh)

	routerMetrics := mux.NewRouter()
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

func db(ctx context.Context) (*sql.DB, error) {
	_, span_Banco := tracer.Start(ctx, "sql.Open)")
	db, err := sql.Open("sqlite3", "./banco.db")
	if err != nil {
		span_Banco.SetStatus(codes.Error, err.Error())
		span_Banco.SetAttributes(attribute.Key("erro-Banco").String(err.Error()))
		span_Banco.RecordError(err)
		span_Banco.End()
		return nil, err
	}
	span_Banco.End()

	_, span_CREATE := tracer.Start(ctx, "db.Exe(CREATE TABLE)")
	const create string = `CREATE TABLE IF NOT EXISTS users ( ID INTEGER NOT NULL PRIMARY KEY, Name TEXT );`
	_, err = db.Exec(create)
	if err != nil {
		span_CREATE.SetStatus(codes.Error, err.Error())
		span_CREATE.SetAttributes(attribute.Key("erro-CREATE").String(err.Error()))
		span_CREATE.RecordError(err)
		span_CREATE.End()
		return nil, err
	}
	span_CREATE.End()

	return db, err
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), "HTTP GET /")
	defer span.End()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World!")
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span_Main := tracer.Start(r.Context(), "HTTP GET /user")
	defer span_Main.End()

	db, err := db(ctx)
	if err != nil {
		return
	}
	defer db.Close()

	_, span_SELECT := tracer.Start(ctx, "db.Query(SELECT)")
	rows, err := db.Query("SELECT ID, Name FROM users ORDER BY ID")
	if err != nil {
		span_SELECT.SetStatus(codes.Error, err.Error())
		span_SELECT.SetAttributes(attribute.Key("erro-SELECT").String(err.Error()))
		span_SELECT.RecordError(err)
		span_SELECT.End()
		return
	}
	span_SELECT.End()
	defer rows.Close()

	_, span_Next := tracer.Start(ctx, "rows.Next()")
	data := []User{}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			span_Next.SetStatus(codes.Error, err.Error())
			span_Next.SetAttributes(attribute.Key("erro-Next").String(err.Error()))
			span_Next.RecordError(err)
			span_Next.End()
			return
		}
		data = append(data, user)
	}
	span_Next.End()

	json.NewEncoder(w).Encode(data)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), "HTTP GET /user")
	defer span.End()

	//params := mux.Vars(r)
	//id := params["id"]

	// for _, item := range users {
	// 	if strconv.Itoa(item.ID) == id {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }

	json.NewEncoder(w).Encode(&User{})
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), "HTTP GET /user")
	defer span.End()

	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.Atoi(idStr)

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	user.ID = id
	// users = append(users, user)

	//json.NewEncoder(w).Encode(users)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), "HTTP GET /user")
	defer span.End()

	//params := mux.Vars(r)
	//id := params["id"]

	// for index, item := range users {
	// 	if strconv.Itoa(item.ID) == id {
	// 		users = append(users[:index], users[index+1:]...)
	// 		break
	// 	}
	// }

	// json.NewEncoder(w).Encode(users)
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

func (l loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), "HTTP GET /login")
	defer span.End()

	sleep(200)
	w.Write([]byte("Welcome to the app!"))
}

func middleware(next http.Handler, m *metrics) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//now := time.Now()
		next.ServeHTTP(w, r)
		//m.loginDuration.Observe(time.Since(now).Seconds())
	})
}

func sleep(ms int) {
	rand.Seed(time.Now().UnixNano())
	now := time.Now()
	n := rand.Intn(ms + now.Second())
	time.Sleep(time.Duration(n) * time.Millisecond)
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

	// reg.MustRegister(m.devices, m.users, m.info, m.upgrades, m.duration, m.loginDuration)
	// reg.MustRegister(m.devices)
	// reg.MustRegister(m.users)
	// reg.MustRegister(m.info)
	// reg.MustRegister(m.upgrades)
	// reg.MustRegister(m.duration)
	// reg.MustRegister(m.loginDuration)
	reg.MustRegister(m.totalRequests)  // prometheus.Register(totalRequests)
	reg.MustRegister(m.responseStatus) // prometheus.Register(responseStatus)
	reg.MustRegister(m.httpDuration)   // prometheus.Register(httpDuration)

	return m
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

// TracerProvider is an OpenTelemetry TracerProvider.
// It provides Tracers to instrumentation so it can trace operational flow through a system.
func newTraceProvider(exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	// exp, err := jaeger.New(
	//     jaeger.WithAgentEndpoint(
	//         jaeger.WithAgentHost("localhost"),
	//         jaeger.WithAgentPort("6831"),
	//     ),
	// )
	// if err != nil {
	//     return err
	// }

	// Ensure default SDK resources and the required service name are set.
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service), // emconv.ServiceNameKey.String(service),
			semconv.ServiceVersionKey.String("1.0.0"),
			semconv.DeploymentEnvironmentKey.String(environment), // attribute.String("environment", environment),
			// attribute.Int64("ID", id),
		),
	)
	if err != nil {
		panic(err)
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	)

	// tp := tracesdk.NewTracerProvider(
	//     tracesdk.WithBatcher(exporter),
	//     tracesdk.WithResource(resource.NewWithAttributes(
	//         semconv.SchemaURL,
	//         semconv.ServiceNameKey.String("server"),
	//         semconv.ServiceVersionKey.String("1.0.0"),
	//         semconv.DeploymentEnvironmentKey.String("local"),
	//     )),
	// )

	// defer func() {
	//     if err := tp.Shutdown(context.Background()); err != nil {
	//         log.Fatal(err)
	//     }
	// }()

	// otel.SetTracerProvider(tp)

	// otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
	//     propagation.TraceContext{},
	//     propagation.Baggage{},
	//     propagators.Jaeger{},
	// ))
}
