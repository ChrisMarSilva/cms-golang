package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
	"go.elastic.co/apm/v2"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.trace.elastic.logrus
// go get -u github.com/gin-gonic/gin
// go get -u github.com/rs/zerolog
// go get -u github.com/elastic/go-elasticsearch/v8
// go get -u go.elastic.co/apm/v2
// go mod tidy

// go run main.go

func main() {

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	for {
		logger.Info().Msg("ping ping ping")
		tracer, err := apm.NewTracer("service-name", "1.0.0")
		if err != nil {
			log.Fatal(err)
		}
		//defer tracer.Close()
		api := &api{tracer: tracer}
		api.handleOrder(context.Background(), "fish fingers")
		api.handleOrder02(context.Background(), "detergent")
		tracer.Flush(nil)
		tracer.Close()
		time.Sleep(1 * time.Millisecond)
	}

}

type api struct {
	tracer *apm.Tracer
}

func (api *api) handleOrder(ctx context.Context, product string) {

	tx := api.tracer.StartTransaction("order", "request")
	defer tx.End()
	ctx = apm.ContextWithTransaction(ctx, tx)

	tx.Context.SetLabel("product", product)

	time.Sleep(10 * time.Millisecond)
	storeOrder(ctx, product)
	time.Sleep(20 * time.Millisecond)
}

func (api *api) handleOrder02(ctx context.Context, product string) {

	tx := api.tracer.StartTransaction("testeeee", "request")
	defer tx.End()
	ctx = apm.ContextWithTransaction(ctx, tx)

	tx.Context.SetLabel("product", product)

	time.Sleep(10 * time.Millisecond)
	storeOrder(ctx, product)
	time.Sleep(20 * time.Millisecond)
}

func storeOrder(ctx context.Context, product string) {
	span, _ := apm.StartSpan(ctx, "store_order", "rpc")
	defer span.End()

	time.Sleep(10 * time.Millisecond)
	storeOrder_02(ctx, product)
	time.Sleep(20 * time.Millisecond)

	time.Sleep(50 * time.Millisecond)
}

func storeOrder_02(ctx context.Context, product string) {
	span, _ := apm.StartSpan(ctx, "store_order_02", "rpc")
	defer span.End()

	time.Sleep(10 * time.Millisecond)
	storeOrder_03(ctx, product)
	time.Sleep(20 * time.Millisecond)

	time.Sleep(50 * time.Millisecond)
}

func storeOrder_03(ctx context.Context, product string) {
	span, _ := apm.StartSpan(ctx, "store_order_03", "rpc")
	defer span.End()

	time.Sleep(50 * time.Millisecond)
}
