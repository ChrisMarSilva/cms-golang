package utils

// import (
// 	"context"
// 	"fmt"
// 	"math/rand"
// 	"net/http"
// 	"os"
// 	"time"

// 	"google.golang.org/grpc"

// 	"github.com/go-kit/kit/log"
// 	"github.com/gorilla/mux"
// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/promauto"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// 	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
// 	"go.opentelemetry.io/otel"
// 	"go.opentelemetry.io/otel/exporters/otlp"
// 	"go.opentelemetry.io/otel/exporters/otlp/otlpgrpc"
// 	"go.opentelemetry.io/otel/propagation"
// 	"go.opentelemetry.io/otel/sdk/resource"
// 	sdktrace "go.opentelemetry.io/otel/sdk/trace"
// 	"go.opentelemetry.io/otel/semconv"
// 	"go.opentelemetry.io/otel/trace"
// )

// var metricRequestLatency = promauto.NewHistogram(prometheus.HistogramOpts{
// 	Namespace: "demo",
// 	Name:      "request_latency_seconds",
// 	Help:      "Request Latency",
// 	Buckets:   prometheus.ExponentialBuckets(.0001, 2, 50),
// })

// // global vars...gasp!
// var addr = "127.0.0.1:8000"
// var tracer trace.Tracer
// var httpClient http.Client
// var logger log.Logger

// func main() {
// 	flush := initTracer()
// 	defer flush()

// 	// initiate globals
// 	tracer = otel.Tracer("demo-app")
// 	httpClient = http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}
// 	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
// 	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

// 	// create and start server
// 	server := instrumentedServer(handler)

// 	fmt.Println("listening...")
// 	server.ListenAndServe()
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	longRunningProcess(ctx)

// 	// check cache
// 	if shouldExecute(40) {
// 		url := "http://" + addr + "/"

// 		resp, err := instrumentedGet(ctx, url)
// 		defer resp.Body.Close()
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 	}

// 	// query database
// 	if shouldExecute(40) {
// 		url := "http://" + addr + "/"

// 		resp, err := instrumentedGet(ctx, url)
// 		defer resp.Body.Close()
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 	}
// }

// func shouldExecute(percent int) bool {
// 	return rand.Int()%100 < percent
// }

// func longRunningProcess(ctx context.Context) {
// 	ctx, sp := tracer.Start(ctx, "Long Running Process")
// 	defer sp.End()

// 	time.Sleep(time.Millisecond * 50)
// 	sp.AddEvent("halfway done!")
// 	time.Sleep(time.Millisecond * 50)
// }

// /***
// Tracing
// ***/
// // Initializes an OTLP exporter, and configures the trace provider
// func initTracer() func() {
// 	ctx := context.Background()

// 	driver := otlpgrpc.NewDriver(
// 		otlpgrpc.WithInsecure(),
// 		otlpgrpc.WithEndpoint("tempo:55680"),
// 		otlpgrpc.WithDialOption(grpc.WithBlock()), // useful for testing
// 	)
// 	exp, err := otlp.NewExporter(ctx, driver)
// 	handleErr(err, "failed to create exporter")

// 	res, err := resource.New(ctx,
// 		resource.WithAttributes(
// 			// the service name used to display traces in backends
// 			semconv.ServiceNameKey.String("demo-service"),
// 		),
// 	)
// 	handleErr(err, "failed to create resource")

// 	bsp := sdktrace.NewBatchSpanProcessor(exp)
// 	tracerProvider := sdktrace.NewTracerProvider(
// 		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
// 		sdktrace.WithResource(res),
// 		sdktrace.WithSpanProcessor(bsp),
// 	)

// 	// set global propagator to tracecontext (the default is no-op).
// 	otel.SetTextMapPropagator(propagation.TraceContext{})
// 	otel.SetTracerProvider(tracerProvider)

// 	return func() {
// 		// Shutdown will flush any remaining spans.
// 		handleErr(tracerProvider.Shutdown(ctx), "failed to shutdown TracerProvider")
// 	}
// }

// /***
// Server
// ***/
// func instrumentedServer(handler http.HandlerFunc) *http.Server {
// 	// OpenMetrics handler : metrics and exemplars
// 	omHandleFunc := func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()

// 		handler.ServeHTTP(w, r)

// 		ctx := r.Context()
// 		traceID := trace.SpanContextFromContext(ctx).TraceID.String()

// 		metricRequestLatency.(prometheus.ExemplarObserver).ObserveWithExemplar(
// 			time.Since(start).Seconds(), prometheus.Labels{"traceID": traceID},
// 		)

// 		// log the trace id with other fields so we can discover traces through logs
// 		logger.Log("msg", "http request", "traceID", traceID, "path", r.URL.Path, "latency", time.Since(start))
// 	}

// 	// OTel handler : traces
// 	otelHandler := otelhttp.NewHandler(http.HandlerFunc(omHandleFunc), "http")

// 	r := mux.NewRouter()
// 	r.Handle("/", otelHandler)
// 	r.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{
// 		EnableOpenMetrics: true,
// 	}))

// 	return &http.Server{
// 		Handler: r,
// 		Addr:    "0.0.0.0:8000",
// 	}
// }

// /***
// Client
// ***/
// func instrumentedGet(ctx context.Context, url string) (*http.Response, error) {
// 	// create http request
// 	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return httpClient.Do(req)
// }

// func handleErr(err error, message string) {
// 	if err != nil {
// 		panic(fmt.Sprintf("%s: %s", err, message))
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"

// 	"github.com/opentracing/opentracing-go"
// 	"github.com/spf13/cobra"
// 	"github.com/spf13/viper"
// 	jaegerClientConfig "github.com/uber/jaeger-client-go/config"
// 	jaegerClientZapLog "github.com/uber/jaeger-client-go/log/zap"
// 	_ "go.uber.org/automaxprocs"
// 	"go.uber.org/zap"

// 	agentApp "github.com/jaegertracing/jaeger/cmd/agent/app"
// 	agentRep "github.com/jaegertracing/jaeger/cmd/agent/app/reporter"
// 	agentGrpcRep "github.com/jaegertracing/jaeger/cmd/agent/app/reporter/grpc"
// 	"github.com/jaegertracing/jaeger/cmd/all-in-one/setupcontext"
// 	collectorApp "github.com/jaegertracing/jaeger/cmd/collector/app"
// 	collectorFlags "github.com/jaegertracing/jaeger/cmd/collector/app/flags"
// 	"github.com/jaegertracing/jaeger/cmd/docs"
// 	"github.com/jaegertracing/jaeger/cmd/env"
// 	"github.com/jaegertracing/jaeger/cmd/flags"
// 	queryApp "github.com/jaegertracing/jaeger/cmd/query/app"
// 	"github.com/jaegertracing/jaeger/cmd/query/app/querysvc"
// 	"github.com/jaegertracing/jaeger/cmd/status"
// 	"github.com/jaegertracing/jaeger/internal/metrics/expvar"
// 	"github.com/jaegertracing/jaeger/internal/metrics/fork"
// 	"github.com/jaegertracing/jaeger/internal/metrics/jlibadapter"
// 	"github.com/jaegertracing/jaeger/pkg/config"
// 	"github.com/jaegertracing/jaeger/pkg/metrics"
// 	"github.com/jaegertracing/jaeger/pkg/tenancy"
// 	"github.com/jaegertracing/jaeger/pkg/version"
// 	metricsPlugin "github.com/jaegertracing/jaeger/plugin/metrics"
// 	ss "github.com/jaegertracing/jaeger/plugin/sampling/strategystore"
// 	"github.com/jaegertracing/jaeger/plugin/storage"
// 	"github.com/jaegertracing/jaeger/ports"
// 	"github.com/jaegertracing/jaeger/storage/dependencystore"
// 	metricsstoreMetrics "github.com/jaegertracing/jaeger/storage/metricsstore/metrics"
// 	"github.com/jaegertracing/jaeger/storage/spanstore"
// 	storageMetrics "github.com/jaegertracing/jaeger/storage/spanstore/metrics"
// )

// // all-in-one/main is a standalone full-stack jaeger backend, backed by a memory store
// func main() {
// 	setupcontext.SetAllInOne()

// 	svc := flags.NewService(ports.CollectorAdminHTTP)

// 	if os.Getenv(storage.SpanStorageTypeEnvVar) == "" {
// 		os.Setenv(storage.SpanStorageTypeEnvVar, "memory") // other storage types default to SpanStorage
// 	}
// 	storageFactory, err := storage.NewFactory(storage.FactoryConfigFromEnvAndCLI(os.Args, os.Stderr))
// 	if err != nil {
// 		log.Fatalf("Cannot initialize storage factory: %v", err)
// 	}
// 	strategyStoreFactoryConfig, err := ss.FactoryConfigFromEnv(os.Stderr)
// 	if err != nil {
// 		log.Fatalf("Cannot initialize sampling strategy store factory config: %v", err)
// 	}
// 	strategyStoreFactory, err := ss.NewFactory(*strategyStoreFactoryConfig)
// 	if err != nil {
// 		log.Fatalf("Cannot initialize sampling strategy store factory: %v", err)
// 	}

// 	fc := metricsPlugin.FactoryConfigFromEnv()
// 	metricsReaderFactory, err := metricsPlugin.NewFactory(fc)
// 	if err != nil {
// 		log.Fatalf("Cannot initialize metrics store factory: %v", err)
// 	}

// 	v := viper.New()
// 	command := &cobra.Command{
// 		Use:   "jaeger-all-in-one",
// 		Short: "Jaeger all-in-one distribution with agent, collector and query in one process.",
// 		Long: `Jaeger all-in-one distribution with agent, collector and query. Use with caution this version
// by default uses only in-memory database.`,
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			if err := svc.Start(v); err != nil {
// 				return err
// 			}
// 			logger := svc.Logger // shortcut
// 			metricsFactory := fork.New("internal",
// 				expvar.NewFactory(10), // backend for internal opts
// 				svc.MetricsFactory.Namespace(metrics.NSOptions{Name: "jaeger"}))
// 			version.NewInfoMetrics(metricsFactory)

// 			tracerCloser := initTracer(svc)

// 			storageFactory.InitFromViper(v, logger)
// 			if err := storageFactory.Initialize(metricsFactory, logger); err != nil {
// 				logger.Fatal("Failed to init storage factory", zap.Error(err))
// 			}

// 			spanReader, err := storageFactory.CreateSpanReader()
// 			if err != nil {
// 				logger.Fatal("Failed to create span reader", zap.Error(err))
// 			}
// 			spanWriter, err := storageFactory.CreateSpanWriter()
// 			if err != nil {
// 				logger.Fatal("Failed to create span writer", zap.Error(err))
// 			}
// 			dependencyReader, err := storageFactory.CreateDependencyReader()
// 			if err != nil {
// 				logger.Fatal("Failed to create dependency reader", zap.Error(err))
// 			}

// 			metricsQueryService, err := createMetricsQueryService(metricsReaderFactory, v, logger, metricsFactory)
// 			if err != nil {
// 				logger.Fatal("Failed to create metrics reader", zap.Error(err))
// 			}

// 			ssFactory, err := storageFactory.CreateSamplingStoreFactory()
// 			if err != nil {
// 				logger.Fatal("Failed to create sampling store factory", zap.Error(err))
// 			}

// 			strategyStoreFactory.InitFromViper(v, logger)
// 			if err := strategyStoreFactory.Initialize(metricsFactory, ssFactory, logger); err != nil {
// 				logger.Fatal("Failed to init sampling strategy store factory", zap.Error(err))
// 			}
// 			strategyStore, aggregator, err := strategyStoreFactory.CreateStrategyStore()
// 			if err != nil {
// 				logger.Fatal("Failed to create sampling strategy store", zap.Error(err))
// 			}

// 			aOpts := new(agentApp.Builder).InitFromViper(v)
// 			repOpts := new(agentRep.Options).InitFromViper(v, logger)
// 			grpcBuilder, err := agentGrpcRep.NewConnBuilder().InitFromViper(v)
// 			if err != nil {
// 				logger.Fatal("Failed to configure connection for grpc", zap.Error(err))
// 			}
// 			cOpts, err := new(collectorFlags.CollectorOptions).InitFromViper(v, logger)
// 			if err != nil {
// 				logger.Fatal("Failed to initialize collector", zap.Error(err))
// 			}
// 			qOpts, err := new(queryApp.QueryOptions).InitFromViper(v, logger)
// 			if err != nil {
// 				logger.Fatal("Failed to configure query service", zap.Error(err))
// 			}

// 			tm := tenancy.NewManager(&cOpts.GRPC.Tenancy)

// 			// collector
// 			c := collectorApp.New(&collectorApp.CollectorParams{
// 				ServiceName:    "jaeger-collector",
// 				Logger:         logger,
// 				MetricsFactory: metricsFactory,
// 				SpanWriter:     spanWriter,
// 				StrategyStore:  strategyStore,
// 				Aggregator:     aggregator,
// 				HealthCheck:    svc.HC(),
// 				TenancyMgr:     tm,
// 			})
// 			if err := c.Start(cOpts); err != nil {
// 				log.Fatal(err)
// 			}

// 			// agent
// 			// if the agent reporter grpc host:port was not explicitly set then use whatever the collector is listening on
// 			if len(grpcBuilder.CollectorHostPorts) == 0 {
// 				grpcBuilder.CollectorHostPorts = append(grpcBuilder.CollectorHostPorts, cOpts.GRPC.HostPort)
// 			}
// 			agentMetricsFactory := metricsFactory.Namespace(metrics.NSOptions{Name: "agent", Tags: nil})
// 			builders := map[agentRep.Type]agentApp.CollectorProxyBuilder{
// 				agentRep.GRPC: agentApp.GRPCCollectorProxyBuilder(grpcBuilder),
// 			}
// 			cp, err := agentApp.CreateCollectorProxy(agentApp.ProxyBuilderOptions{
// 				Options: *repOpts,
// 				Logger:  logger,
// 				Metrics: agentMetricsFactory,
// 			}, builders)
// 			if err != nil {
// 				logger.Fatal("Could not create collector proxy", zap.Error(err))
// 			}
// 			agent := startAgent(cp, aOpts, logger, metricsFactory)

// 			// query
// 			querySrv := startQuery(
// 				svc, qOpts, qOpts.BuildQueryServiceOptions(storageFactory, logger),
// 				spanReader, dependencyReader, metricsQueryService,
// 				metricsFactory, tm,
// 			)

// 			svc.RunAndThen(func() {
// 				agent.Stop()
// 				_ = cp.Close()
// 				_ = c.Close()
// 				_ = querySrv.Close()
// 				if closer, ok := spanWriter.(io.Closer); ok {
// 					if err := closer.Close(); err != nil {
// 						logger.Error("Failed to close span writer", zap.Error(err))
// 					}
// 				}
// 				if err := storageFactory.Close(); err != nil {
// 					logger.Error("Failed to close storage factory", zap.Error(err))
// 				}
// 				_ = tracerCloser.Close()
// 			})
// 			return nil
// 		},
// 	}

// 	command.AddCommand(version.Command())
// 	command.AddCommand(env.Command())
// 	command.AddCommand(docs.Command(v))
// 	command.AddCommand(status.Command(v, ports.CollectorAdminHTTP))

// 	config.AddFlags(
// 		v,
// 		command,
// 		svc.AddFlags,
// 		storageFactory.AddPipelineFlags,
// 		agentApp.AddFlags,
// 		agentRep.AddFlags,
// 		agentGrpcRep.AddFlags,
// 		collectorFlags.AddFlags,
// 		queryApp.AddFlags,
// 		strategyStoreFactory.AddFlags,
// 		metricsReaderFactory.AddFlags,
// 	)

// 	if err := command.Execute(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// func startAgent(
// 	cp agentApp.CollectorProxy,
// 	b *agentApp.Builder,
// 	logger *zap.Logger,
// 	baseFactory metrics.Factory,
// ) *agentApp.Agent {
// 	agent, err := b.CreateAgent(cp, logger, baseFactory)
// 	if err != nil {
// 		logger.Fatal("Unable to initialize Jaeger Agent", zap.Error(err))
// 	}

// 	logger.Info("Starting agent")
// 	if err := agent.Run(); err != nil {
// 		logger.Fatal("Failed to run the agent", zap.Error(err))
// 	}

// 	return agent
// }

// func startQuery(
// 	svc *flags.Service,
// 	qOpts *queryApp.QueryOptions,
// 	queryOpts *querysvc.QueryServiceOptions,
// 	spanReader spanstore.Reader,
// 	depReader dependencystore.Reader,
// 	metricsQueryService querysvc.MetricsQueryService,
// 	baseFactory metrics.Factory,
// 	tm *tenancy.Manager,
// ) *queryApp.Server {
// 	spanReader = storageMetrics.NewReadMetricsDecorator(spanReader, baseFactory.Namespace(metrics.NSOptions{Name: "query"}))
// 	qs := querysvc.NewQueryService(spanReader, depReader, *queryOpts)
// 	server, err := queryApp.NewServer(svc.Logger, qs, metricsQueryService, qOpts, tm, opentracing.GlobalTracer())
// 	if err != nil {
// 		svc.Logger.Fatal("Could not start jaeger-query service", zap.Error(err))
// 	}
// 	go func() {
// 		for s := range server.HealthCheckStatus() {
// 			svc.SetHealthCheckStatus(s)
// 		}
// 	}()
// 	if err := server.Start(); err != nil {
// 		svc.Logger.Fatal("Could not start jaeger-query service", zap.Error(err))
// 	}
// 	return server
// }

// func initTracer(svc *flags.Service) io.Closer {
// 	logger := svc.Logger
// 	traceCfg := &jaegerClientConfig.Configuration{
// 		ServiceName: "jaeger-query",
// 		Sampler: &jaegerClientConfig.SamplerConfig{
// 			Type:  "const",
// 			Param: 1.0,
// 		},
// 		RPCMetrics: true,
// 	}
// 	traceCfg, err := traceCfg.FromEnv()
// 	if err != nil {
// 		logger.Fatal("Failed to read tracer configuration", zap.Error(err))
// 	}
// 	tracer, closer, err := traceCfg.NewTracer(
// 		jaegerClientConfig.Metrics(jlibadapter.NewAdapter(svc.MetricsFactory)),
// 		jaegerClientConfig.Logger(jaegerClientZapLog.NewLogger(logger)),
// 	)
// 	if err != nil {
// 		logger.Fatal("Failed to initialize tracer", zap.Error(err))
// 	}
// 	opentracing.SetGlobalTracer(tracer)
// 	return closer
// }

// func createMetricsQueryService(
// 	metricsReaderFactory *metricsPlugin.Factory,
// 	v *viper.Viper,
// 	logger *zap.Logger,
// 	metricsReaderMetricsFactory metrics.Factory,
// ) (querysvc.MetricsQueryService, error) {
// 	if err := metricsReaderFactory.Initialize(logger); err != nil {
// 		return nil, fmt.Errorf("failed to init metrics reader factory: %w", err)
// 	}

// 	// Ensure default parameter values are loaded correctly.
// 	metricsReaderFactory.InitFromViper(v, logger)
// 	reader, err := metricsReaderFactory.CreateMetricsReader()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create metrics reader: %w", err)
// 	}

// 	// Decorate the metrics reader with metrics instrumentation.
// 	return metricsstoreMetrics.NewReadMetricsDecorator(reader, metricsReaderMetricsFactory), nil
// }

// package endpointmetrics

// import (
// 	"fmt"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// 	"medium/m/v2/internal/observability/metrics/countermetrics"
// 	"medium/m/v2/internal/observability/metrics/histogrammetrics"
// 	"net/http"
// )

// const (
// 	// Labels
// 	endpoint            string = "endpoint"
// 	verb                string = "verb"
// 	pattern             string = "pattern"
// 	failed              string = "failed"
// 	error               string = "error"
// 	responseCode        string = "response_code"
// 	isAvailabilityError string = "is_availability_error"
// 	isReliabilityError  string = "is_reliability_error"

// 	// Names
// 	endpointRequestCounter string = "endpoint_request_counter"
// 	endpointRequestLatency string = "endpoint_request_latency"
// )

// type Metrics struct {
// 	// Metric
// 	Latency float64

// 	// Labels
// 	Endpoint             string
// 	Verb                 string
// 	Pattern              string
// 	ResponseCode         int
// 	Failed               bool
// 	Error                string
// 	HasAvailabilityError bool
// 	HasReliabilityError  bool
// }

// func Send(metrics Metrics) {
// 	labels := map[string]string{
// 		endpoint:            metrics.Endpoint,
// 		verb:                metrics.Verb,
// 		pattern:             metrics.Pattern,
// 		responseCode:        fmt.Sprintf("%d", metrics.ResponseCode),
// 		failed:              fmt.Sprintf("%v", metrics.Failed),
// 		error:               metrics.Error,
// 		isAvailabilityError: fmt.Sprintf("%v", metrics.HasAvailabilityError),
// 		isReliabilityError:  fmt.Sprintf("%v", metrics.HasReliabilityError),
// 	}

// 	countermetrics.Increment(countermetrics.Metric{
// 		Name:   endpointRequestCounter,
// 		Labels: labels,
// 	})

// 	histogrammetrics.Observe(histogrammetrics.Metric{
// 		Name:  endpointRequestLatency,
// 		Value: float64(metrics.Latency),
// 		Labels: map[string]string{
// 			endpoint: metrics.Endpoint,
// 		},
// 	})
// }

// func Start() {
// 	fmt.Println("starting prometheus")
// 	http.Handle("/metrics", promhttp.Handler())

// 	go func() {
// 		http.ListenAndServe(":2112", nil)
// 	}()
// 	fmt.Println("started prometheus")
// }

// ---

// package countermetrics

// import (
// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/promauto"
// 	"medium/m/v2/internal/observability/metrics"
// )

// type Metric struct {
// 	Name   string
// 	Labels map[string]string
// }

// var createdMetrics = make(map[string]*prometheus.CounterVec)

// func Increment(metric Metric) {
// 	go func() {
// 		labelsKey := metrics.GetLabelsKey(metric.Labels)

// 		opts := prometheus.CounterOpts{
// 			Name: metric.Name,
// 		}

// 		if createdMetrics[metric.Name] == nil {
// 			counter := promauto.NewCounterVec(opts, labelsKey)
// 			createdMetrics[metric.Name] = counter
// 		}

// 		counter := createdMetrics[metric.Name]
// 		counter.With(metric.Labels).Inc()
// 	}()
// }

// func (m *metricsHandlerAdapter) execute(w http.ResponseWriter, r *http.Request) {
// 	start := time.Now()

// 	response, err := m.handler.Handle(r)

// 	latency := time.Since(start).Seconds()
// 	m.metrify(response, err, latency)

// 	if err != nil {
// 		http.Error(w, err.Name(), err.Code())
// 		return
// 	}

// 	encode.WriteJsonResponse(w, response.Object(), response.Code())
// }

// func (m *metricsHandlerAdapter) metrify(response apiresponse.ApiResponse, err apierror.ApiError, latencyInMs float64) {
// 	metrics := endpointmetrics.Metrics{
// 		Latency:  latencyInMs,
// 		Endpoint: m.handler.Name(),
// 		Verb:     m.handler.Verb(),
// 		Pattern:  m.handler.Pattern(),
// 	}

// 	if err != nil {
// 		metrics.Failed = true
// 		metrics.Error = err.Name()
// 		metrics.ResponseCode = err.Code()
// 		if err.Code() >= 500 {
// 			metrics.HasReliabilityError = false
// 			metrics.HasAvailabilityError = true
// 		} else {
// 			metrics.HasReliabilityError = true
// 			metrics.HasAvailabilityError = false
// 		}
// 	} else {
// 		metrics.Failed = false
// 		metrics.ResponseCode = response.Code()
// 	}

// 	endpointmetrics.Send(metrics)
// }

// package main

// import (
// 	"flag"
// 	"log"
// 	"net/http"

// 	"github.com/prometheus/client_golang/prometheus/collectors"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// )

// var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

// func main() {
// 	flag.Parse()

// 	// Create non-global registry.
// 	reg := prometheus.NewRegistry()

// 	// Add go runtime metrics and process collectors.
// 	reg.MustRegister(
// 		collectors.NewGoCollector(),
// 		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
// 	)

// 	// Expose /metrics HTTP endpoint using the created custom registry.
// 	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
// 	log.Fatal(http.ListenAndServe(*addr, nil))
// }

// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/prometheus/client_golang/examples/middleware/httpmiddleware"
// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/collectors"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// )

// func main() {
// 	// Create non-global registry.
// 	registry := prometheus.NewRegistry()

// 	// Add go runtime metrics and process collectors.
// 	registry.MustRegister(
// 		collectors.NewGoCollector(),
// 		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
// 	)

// 	// Expose /metrics HTTP endpoint using the created custom registry.
// 	http.Handle(
// 		"/metrics",
// 		httpmiddleware.New(
// 			registry, nil).
// 			WrapHandler("/metrics", promhttp.HandlerFor(
// 				registry,
// 				promhttp.HandlerOpts{}),
// 			))

// 	log.Fatalln(http.ListenAndServe(":8080", nil))
// }

// // Copyright 2022 The Prometheus Authors
// // Licensed under the Apache License, Version 2.0 (the "License");
// // you may not use this file except in compliance with the License.
// // You may obtain a copy of the License at
// //
// // http://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS,
// // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// // See the License for the specific language governing permissions and
// // limitations under the License.

// // Package httpmiddleware is adapted from
// // https://github.com/bwplotka/correlator/tree/main/examples/observability/ping/pkg/httpinstrumentation
// package httpmiddleware

// import (
// 	"net/http"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/promauto"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// )

// type Middleware interface {
// 	// WrapHandler wraps the given HTTP handler for instrumentation.
// 	WrapHandler(handlerName string, handler http.Handler) http.HandlerFunc
// }

// type middleware struct {
// 	buckets  []float64
// 	registry prometheus.Registerer
// }

// // WrapHandler wraps the given HTTP handler for instrumentation:
// // It registers four metric collectors (if not already done) and reports HTTP
// // metrics to the (newly or already) registered collectors.
// // Each has a constant label named "handler" with the provided handlerName as
// // value.
// func (m *middleware) WrapHandler(handlerName string, handler http.Handler) http.HandlerFunc {
// 	reg := prometheus.WrapRegistererWith(prometheus.Labels{"handler": handlerName}, m.registry)

// 	requestsTotal := promauto.With(reg).NewCounterVec(
// 		prometheus.CounterOpts{
// 			Name: "http_requests_total",
// 			Help: "Tracks the number of HTTP requests.",
// 		}, []string{"method", "code"},
// 	)
// 	requestDuration := promauto.With(reg).NewHistogramVec(
// 		prometheus.HistogramOpts{
// 			Name:    "http_request_duration_seconds",
// 			Help:    "Tracks the latencies for HTTP requests.",
// 			Buckets: m.buckets,
// 		},
// 		[]string{"method", "code"},
// 	)
// 	requestSize := promauto.With(reg).NewSummaryVec(
// 		prometheus.SummaryOpts{
// 			Name: "http_request_size_bytes",
// 			Help: "Tracks the size of HTTP requests.",
// 		},
// 		[]string{"method", "code"},
// 	)
// 	responseSize := promauto.With(reg).NewSummaryVec(
// 		prometheus.SummaryOpts{
// 			Name: "http_response_size_bytes",
// 			Help: "Tracks the size of HTTP responses.",
// 		},
// 		[]string{"method", "code"},
// 	)

// 	// Wraps the provided http.Handler to observe the request result with the provided metrics.
// 	base := promhttp.InstrumentHandlerCounter(
// 		requestsTotal,
// 		promhttp.InstrumentHandlerDuration(
// 			requestDuration,
// 			promhttp.InstrumentHandlerRequestSize(
// 				requestSize,
// 				promhttp.InstrumentHandlerResponseSize(
// 					responseSize,
// 					handler,
// 				),
// 			),
// 		),
// 	)

// 	return base.ServeHTTP
// }

// // New returns a Middleware interface.
// func New(registry prometheus.Registerer, buckets []float64) Middleware {
// 	if buckets == nil {
// 		buckets = prometheus.ExponentialBuckets(0.1, 1.5, 5)
// 	}

// 	return &middleware{
// 		buckets:  buckets,
// 		registry: registry,
// 	}
// }

// // Copyright 2022 The Prometheus Authors
// // Licensed under the Apache License, Version 2.0 (the "License");
// // you may not use this file except in compliance with the License.
// // You may obtain a copy of the License at
// //
// // http://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS,
// // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// // See the License for the specific language governing permissions and
// // limitations under the License.

// //go:build go1.17
// // +build go1.17

// // A minimal example of how to include Prometheus instrumentation.
// package main

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"regexp"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/prometheus/client_golang/prometheus/collectors"
// 	"github.com/prometheus/client_golang/prometheus/promhttp"
// )

// var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

// func main() {
// 	flag.Parse()

// 	// Create a new registry.
// 	reg := prometheus.NewRegistry()

// 	// Add Go module build info.
// 	reg.MustRegister(collectors.NewBuildInfoCollector())
// 	reg.MustRegister(collectors.NewGoCollector(
// 		collectors.WithGoCollectorRuntimeMetrics(collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")}),
// 	))

// 	// Expose the registered metrics via HTTP.
// 	http.Handle("/metrics", promhttp.HandlerFor(
// 		reg,
// 		promhttp.HandlerOpts{
// 			// Opt into OpenMetrics to support exemplars.
// 			EnableOpenMetrics: true,
// 		},
// 	))
// 	fmt.Println("Hello world from new Go Collector!")
// 	log.Fatal(http.ListenAndServe(*addr, nil))
// }
