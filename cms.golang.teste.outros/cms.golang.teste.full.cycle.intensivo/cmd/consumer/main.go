package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chrismarsilva/cms.golang.teste.intensivo/internal/order/infra/database"
	"github.com/chrismarsilva/cms.golang.teste.intensivo/internal/order/usecase"
	"github.com/chrismarsilva/cms.golang.teste.intensivo/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := database.NewOrderRepository(db)
	uc := usecase.CalculateFinalPriceUseCase{OrderRepository: repository}

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	out := make(chan amqp.Delivery, 100) // channel
	go rabbitmq.Consume(ch, out)         // T2

	// // T1
	// for msg := range out {
	// 	var inputDTO usecase.OrderInputDTO
	// 	err := json.Unmarshal(msg.Body, &inputDTO)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	outputDTO, err := uc.Execute(inputDTO)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	msg.Ack(false)
	// 	fmt.Println(outputDTO)
	// 	time.Sleep(1 * time.Millisecond)
	// }

	//forever := make(chan bool) // channel
	var qtdWorkers int = 150 // 5 // 500 // 15_000 panic: database is locked
	for i := 1; i <= qtdWorkers; i++ {
		go worker(out, &uc, i)
	}
	//<-forever

	http.HandleFunc("/", getTotal)
	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		getTotalUC := usecase.GetTotalUseCase{OrderRepository: repository}

		total, err := getTotalUC.Execute()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		json.NewEncoder(w).Encode(total)
	})

	log.Fatal(http.ListenAndServe(":8080", nil)) // ao chamar o Server HTTP, ele cria uma Thread
}

func getTotal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	var total int = 0
	payload := map[string]string{"total": strconv.Itoa(total)}
	json.NewEncoder(w).Encode(payload)
}

func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerID int) {
	for msg := range deliveryMessage {
		var inputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			panic(err)
		}

		outputDTO, err := uc.Execute(inputDTO)
		if err != nil {
			panic(err)
		}

		msg.Ack(false)
		fmt.Printf("Worker %d has processed order %s\n", workerID, outputDTO.ID)
		// time.Sleep(500 * time.Millisecond)
		time.Sleep(1 * time.Second)
	}
}
