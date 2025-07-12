package main

import (
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// go run ./cmd/channel/main.go

var (
	NumWorkers   = 1   // número de workers // 10
	maxQueueSize = 100 // tamanho do buffer // 5000
	batchSize    = 4   // tamanho do lote //  1000
	batchSleep   = 500 // espera antes de drenar lote (milissegundos) // 2000
	retryDelay   = 1
)

func main() {
	//log.Println("[main] started...")

	batchChannel := make(chan Job, maxQueueSize)
	for i := 1; i <= NumWorkers; i++ {
		go func(idx int) {
			processBatch(idx, batchChannel) // RunWorker
		}(i)
	}

	// time.Sleep(100 * time.Millisecond) // wait for db is up

	for i := 0; i < 10; i++ {
		go func(idx int) {
			//batchChannel <- Job{name: "#" + strconv.Itoa(idx), Pessoa: &Pessoa{Id: uuid.New(), Nome: "Pessoa " + strconv.Itoa(idx)}}
			batchChannel <- Job{ID: uuid.New(), Name: "Pessoa " + strconv.Itoa(idx)}
		}(i)
		time.Sleep(1 * time.Second) // intervalos entre envios
	}

	// <-time.After(time.Hour)
	for {
		time.Sleep(time.Second) // Aguarda 1 segundo
	}

	//log.Println("[main] finished...")
}

type Job struct {
	// name   string
	// Pessoa *Pessoa
	ID   uuid.UUID
	Name string
}

func processBatch(idx int, batchCh chan Job) {
	// slog.Info("Starting worker(", idx, ")...")
	// defer slog.Info("Finishing worker(", idx, ")...")

	var useDefault = true
	go resetUseDefault(&useDefault)

	for {
		firstJob := <-batchCh // Pega o primeiro job (bloca até existir)
		batch := []Job{firstJob}
		time.Sleep(time.Duration(batchSleep) * time.Millisecond) // Aguarda um tempo para coletar mais jobs

	collectLoop:
		for i := 1; i < batchSize; i++ {
			select {
			case job := <-batchCh:
				batch = append(batch, job)
			default:
				break collectLoop // No more jobs available, stop collecting for this batch
			}
		}

		//log.Printf("Processing batch of %d jobs, initial useDefault=%v", len(batch), useDefault)
		useDefault = processJob(batch, useDefault, batchCh)
	}
}

// resetLoop reseta useDefault para true a cada retryDelay segundos
func resetUseDefault(useDefault *bool) {
	ticker := time.NewTicker(time.Duration(retryDelay) * time.Second) // retryDelay // time.NewTicker(time.Duration(100) * time.Millisecond) //
	defer ticker.Stop()

	for range ticker.C {
		if !*useDefault {
			log.Println("[resetLoop] Timeout: reset useDefault = true")
		}
		*useDefault = true
	}
}

func processJob(batch []Job, useDefault bool, batchCh chan<- Job) bool {
	// slog.Info("Starting processJob(", len(batch), ")...")
	// defer slog.Info("Finishing processJob(", len(batch), ")...")

	for _, job := range batch {
		if useDefault {
			if err := processJobDefault(job); err != nil {
				//log.Printf("[processJob] Default falhou no job %s, usando fallback", job.ID)
				//_ = processJobFallback(job)

				if err2 := processJobFallback(job); err2 != nil {
					log.Printf("[processJob] Fallback também falhou no job %s, reenfileirar", job.Name)
					job.Name = job.Name + " (retry)"
					batchCh <- job
				}
				useDefault = false
			}
		} else {
			if err := processJobFallback(job); err != nil {
				log.Printf("[processJob] Fallback falhou no job %s, reenfileirar", job.Name)
				job.Name = job.Name + " (retry)"
				batchCh <- job
			}
		}
	}

	return useDefault
}

func processJobDefault(job Job) error {
	//log.Printf("[DEFAULT] Processing job %s (%s)", job.ID, job.Name)

	time.Sleep(200 * time.Millisecond)
	if job.Name == "Pessoa 3" || job.Name == "Pessoa 4" { // if time.Now().UnixNano()%2 == 0 {
		// log.Printf("[DEFAULT] Job %s FAILED", job.ID)
		return errors.New("default error")
	}

	log.Printf("[DEFAULT] Job %s %s SUCCEEDED", job.Name, job.ID)
	return nil
}

func processJobFallback(job Job) error {
	// log.Printf("[FALLBACK] Processing job %s (%s)", job.ID, job.Name)

	time.Sleep(200 * time.Millisecond)
	if job.Name == "Pessoa 4" {
		return errors.New("fallback error")
	}

	log.Printf("[FALLBACK] Job %s %s SUCCEEDED", job.Name, job.ID)
	return nil
}
