package main

import (
	"context"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// go run ./cmd/channel/main.go

const (
	maxQueueSize = 5000 // tamanho do buffer // 5000 // 100
	batchSize    = 1000 // tamanho do lote //  1000 // 4
	batchSleep   = 100  // espera antes de drenar lote (milissegundos) // 2000 // 500
	retryDelay   = 5    // tempo de espera para resetar useDefault (segundos) // 5 // 1
)

var (
	numWorkers = 1 // número de workers // 10
	useDefault atomic.Bool
	jobPool    = sync.Pool{
		New: func() interface{} {
			return make([]Job, 0, batchSize)
		},
	}
	logger *zap.Logger
)

func init() {
	useDefault.Store(true)
}

func main() {
	//log.Println("[main] started...")

	logger := zap.Must(zap.NewDevelopment())
	defer logger.Sync() // flushes buffer, if any
	zap.ReplaceGlobals(logger)

	runtime.GOMAXPROCS(runtime.NumCPU())
	numWorkers = runtime.NumCPU() // runtime.NumCPU() // 10

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// initMetrics()

	jobs := make(chan Job, maxQueueSize)
	for i := 1; i <= numWorkers; i++ {
		go func(idx int) {
			processWorkers(ctx, idx, jobs) // RunWorker
		}(i)
	}

	go resetUseDefault()
	//go metricsServer()

	go func() {
		for i := 0; i < 10; i++ { // 1_000_000
			//go func(idx int) {
			//batchChannel <- Job{name: "#" + strconv.Itoa(idx), Pessoa: &Pessoa{Id: uuid.New(), Nome: "Pessoa " + strconv.Itoa(idx)}}
			jobs <- Job{ID: uuid.New()} // , Name: "Pessoa " + strconv.Itoa(idx)
			//}(i)
			//time.Sleep(1 * time.Second) // intervalos entre envios
		}
		close(jobs)
	}()

	zap.L().Info("processing_batch",
		zap.Int("size", len(jobs)),
		zap.Bool("useDefault", useDefault.Load()),
	)

	<-ctx.Done()
	// <-time.After(time.Hour)
	// for {
	// 	time.Sleep(time.Second) // Aguarda 1 segundo
	// }

	//log.Println("[main] finished...")
}

type Job struct {
	// Pessoa *Pessoa
	ID uuid.UUID
	//Name string
}

func resetUseDefault() {
	ticker := time.NewTicker(time.Duration(retryDelay) * time.Second)
	defer ticker.Stop() // useDefault *bool

	for range ticker.C {
		defaultMode := useDefault.Load()
		if !defaultMode { // !*useDefault
			log.Println("[resetUseDefault] useDefault reset to tru")
			//*useDefault = true
			useDefault.Store(true)
		}
	}
}

func processWorkers(ctx context.Context, idx int, jobs chan Job) {
	// slog.Info("Starting worker(", idx, ")...")
	// defer slog.Info("Finishing worker(", idx, ")...")

	for {
		//firstJob := <-jobs // Pega o primeiro job (bloca até existir)
		//batch := []Job{firstJob}
		//time.Sleep(time.Duration(batchSleep) * time.Millisecond) // Aguarda um tempo para coletar mais jobs

		select {
		case <-ctx.Done():
			return
		case first, ok := <-jobs:
			if !ok {
				return
			}

			batch := jobPool.Get().([]Job)[:0]
			batch = append(batch, first)

			timer := time.NewTimer(time.Duration(batchSleep) * time.Millisecond)

		batchLoop:
			for len(batch) < batchSize {
				select {
				case job, ok2 := <-jobs:
					if !ok2 {
						break batchLoop
					}
					batch = append(batch, job)
				case <-timer.C:
					break batchLoop
				default:
					break batchLoop
				}
			}

			timer.Stop()

			//log.Printf("Processing batch of %d jobs, initial useDefault=%v", len(batch), useDefault)
			processJob(batch, jobs)
			jobPool.Put(batch)
		}
	}
}

func processJob(batch []Job, jobs chan<- Job) {
	// slog.Info("Starting processJob(", len(batch), ")...")
	// defer slog.Info("Finishing processJob(", len(batch), ")...")

	defaultMode := useDefault.Load()

	for _, job := range batch {
		if defaultMode {
			if err := processJobDefault(job); err != nil {
				//log.Printf("[processJob] Default falhou no job %s, usando fallback", job.ID)

				if err2 := processJobFallback(job); err2 != nil {
					log.Printf("[processJob] Fallback também falhou no job %s, reenfileirar", job.ID.String())
					// go func(j JobLocal) {
					//job.Name = job.Name + " (retry)"
					jobs <- job
					//}(job)
				}
				defaultMode = false
				useDefault.Store(false)
			}
		} else {
			if err := processJobFallback(job); err != nil {
				log.Printf("[processJob] Fallback falhou no job %s, reenfileirar", job.ID.String())
				// go func(j JobLocal) {
				//job.Name = job.Name + " (retry)"
				jobs <- job
				//}(job)
			}
		}
	}
}

func processJobDefault(job Job) error {
	//log.Printf("[DEFAULT] Processing job %s (%s)", job.ID, job.Name)

	time.Sleep(200 * time.Millisecond)
	// if job.Name == "Pessoa 3" || job.Name == "Pessoa 4" { // if time.Now().UnixNano()%2 == 0 { // if j.ID.ID()%5 == 0 {
	// 	// log.Printf("[DEFAULT] Job %s FAILED", job.ID)
	// 	return errors.New("default error")
	// }

	//log.Printf("[DEFAULT] Job %s SUCCEEDED", job.ID)
	return nil
}

func processJobFallback(job Job) error {
	// log.Printf("[FALLBACK] Processing job %s (%s)", job.ID, job.Name)

	time.Sleep(200 * time.Millisecond)
	// if job.Name == "Pessoa 4" { //  if j.ID.ID()%13 == 0 {
	// 	return errors.New("fallback error")
	// }

	//log.Printf("[FALLBACK] Job %s SUCCEEDED", job.ID)
	return nil
}
