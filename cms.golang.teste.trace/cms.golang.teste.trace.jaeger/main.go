package main

import (
	"context"
	"log"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	spanlog "github.com/opentracing/opentracing-go/log"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	metrics "github.com/uber/jaeger-lib/metrics"
)

// https://opentracing.io/guides/golang/quick-start/
// cd "C:\Users\chris\Desktop\CMS GoLang\cms.golang.teste.outros\cms.golang.teste.trace.jaeger"

// go mod init github.com/chrismarsilva/cms.golang.teste.jaeger
// go get github.com/uber/jaeger-client-go
// go get github.com/opentracing/opentracing-go
// go get package github.com/uber/jaeger-client-go
// go get github.com/jaegertracing/jaeger-client-go
// go get github.com/pkg/errors
// go mod tidy

// go run main.go

// func init() {

// }

func main() {

	log.Println("Ini")

	ctx := context.Background()

	cfg := jaegercfg.Configuration{
		ServiceName: "cms.golang.teste.trace.jaeger",
		Sampler:     &jaegercfg.SamplerConfig{Type: jaeger.SamplerTypeConst, Param: 1},
		Reporter:    &jaegercfg.ReporterConfig{LogSpans: true, BufferFlushInterval: 1 * time.Second},
	}

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaegerlog.NullLogger), jaegercfg.Metrics(metrics.NullFactory))
	if err != nil {
		log.Println("ERROR: cannot init Jaeger: ", err)
		return
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	//span1 := opentracing.SpanFromContext(ctx)

	// // span := tracer.StartSpan("say-hello")
	// span, ctx := opentracing.StartSpanFromContext(ctx, "say-hello")
	// time.Sleep(1 * time.Second)
	// span.Finish()

	//parentSpan := tracer.StartSpan("parent")
	parentSpan, ctx := opentracing.StartSpanFromContext(ctx, "parent")
	parentSpan.SetTag("error", true)
	parentSpan.SetTag("error1", 1)
	parentSpan.SetTag("error2", "ssss")
	parentSpan.SetBaggageItem("error3", "123")
	parentSpan.BaggageItem("error4")
	parentSpan.LogFields(spanlog.String("event", "getResponse"), spanlog.String("value", "string(body)"))
	parentSpan.LogFields(spanlog.Error(err))
	parentSpan.LogEvent("AAA")
	defer parentSpan.Finish()
	time.Sleep(1 * time.Second)

	//logger.For(ctx).Info("Searching for nearby drivers", zap.String("location", "aki"))

	childSpan := tracer.StartSpan("child1", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(1 * time.Second)
	childSpan1_1 := tracer.StartSpan("child1.1", opentracing.ChildOf(childSpan.Context()))
	time.Sleep(1 * time.Second)
	childSpan1_1.Finish()
	childSpan.Finish()

	childSpan2 := tracer.StartSpan("child2", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(1 * time.Second)
	childSpan2.Finish()

	childSpan3 := tracer.StartSpan("child3", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(1 * time.Second)
	childSpan3_1 := tracer.StartSpan("child3.1", opentracing.ChildOf(childSpan3.Context()))
	time.Sleep(1 * time.Second)
	childSpan3_2 := tracer.StartSpan("child3.2", opentracing.ChildOf(childSpan3.Context()))
	time.Sleep(1 * time.Second)
	childSpan3_2.Finish()
	childSpan3_1.Finish()
	childSpan3.Finish()

	childSpan4 := tracer.StartSpan("child4", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(1 * time.Second)
	childSpan4.Finish()

	Teste1(ctx)
	Teste2(ctx, tracer, parentSpan)
	Teste3()

	// // clientSpan := tracer.StartSpan("clientspan")
	// clientSpan, ctx := opentracing.StartSpanFromContext(ctx, "clientspan")
	// defer clientSpan.Finish()
	// time.Sleep(time.Second)

	// url := "https://jsonplaceholder.typicode.com/todos"
	// req, _ := http.NewRequest("GET", url, nil)

	// ext.SpanKindRPCClient.Set(clientSpan)
	// ext.HTTPUrl.Set(clientSpan, url)
	// ext.HTTPMethod.Set(clientSpan, "GET")

	// tracer.Inject(clientSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	// resp, _ := http.DefaultClient.Do(req)
	// fmt.Println(resp)

	log.Println("Fim")
}

func Teste1(ctx context.Context) {

	tracer := opentracing.GlobalTracer()

	// parentSpan := tracer.StartSpan("parent5")
	parentSpan, ctx := opentracing.StartSpanFromContext(ctx, "parent5")
	defer parentSpan.Finish()
	parentSpan.LogFields(spanlog.String("event", "soft error"), spanlog.String("type", "cache timeout"), spanlog.Int("waited.millis", 1500))
	time.Sleep(1 * time.Second)

	childSpan5 := tracer.StartSpan("child5", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(1 * time.Second)
	childSpan5.Finish()

}

func Teste2(ctx context.Context, tracer opentracing.Tracer, parentSpan opentracing.Span) {

	sp11 := opentracing.StartSpan("parent6", opentracing.ChildOf(parentSpan.Context()))
	defer sp11.Finish()

	childSpan6 := tracer.StartSpan("child6", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(1 * time.Second)
	childSpan6.Finish()

	childSpan6 = tracer.StartSpan("child6", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(1 * time.Second)
	childSpan6.Finish()

	//sp := opentracing.StartSpan("parent7")
	sp, ctx := opentracing.StartSpanFromContext(ctx, "parent7")
	time.Sleep(1 * time.Second)
	sp1 := opentracing.StartSpan("child7.1", opentracing.ChildOf(sp.Context()))
	time.Sleep(1 * time.Second)
	sp1.Finish()
	sp.Finish()

}

func Teste3() {

	ctx := context.Background()
	tracer := opentracing.GlobalTracer()

	// parentSpan := tracer.StartSpan("parent8")
	parentSpan, ctx := opentracing.StartSpanFromContext(ctx, "parent8")
	defer parentSpan.Finish()
	time.Sleep(1 * time.Second)

	childSpan5 := tracer.StartSpan("child8", opentracing.ChildOf(parentSpan.Context()))
	time.Sleep(1 * time.Second)
	childSpan5.Finish()

}
