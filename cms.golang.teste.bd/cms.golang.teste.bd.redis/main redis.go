package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/chrismarsilva/rinha-backend-2025/dtos"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// go run main.go

func main_redis() {
	log.Println("Starting Rinha Backend 2025...")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //  "localhost:6379",
		Password: "123",            // no password set
		//DB:           0,                // use default DB
		PoolSize:     100,
		MinIdleConns: 25,
		ReadTimeout:  100 * time.Millisecond,
		WriteTimeout: 100 * time.Millisecond,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("Redis connect:", err)
	}

	ctx := context.Background()

	payment := dtos.PaymentProcessorPaymentRequest{CorrelationID: uuid.New(), Amount: 80.0, RequestedAt: time.Now().UTC()}
	payload, _ := json.Marshal(payment)

	//---------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------

	// //Define os campos especificados para seus respectivos valores no hash armazenado em key.
	// client.HSet(ctx, "payments", payment.CorrelationID.String(), payload).Err()

	// // Retorna todos os campos e valores do hash armazenado em key.
	// payments, err := client.HGetAll(ctx, "payments").Result()

	// for _, paymentJson := range payments {
	// 	var paymentData dtos.PaymentProcessorPaymentRequest
	// 	json.Unmarshal([]byte(paymentJson), &paymentData)

	// 	log.Println("Payment data JSON:", paymentJson)
	// 	log.Println("Payment data Struct:", paymentData)
	// 	log.Println("")
	// }

	//---------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------

	// default - fallback
	data := map[string]interface{}{"service": "fallback", "timestamp": time.Now().Unix()}
	client.HMSet(ctx, "healthy_processor_status", data)

	m, _ := client.HGetAll(ctx, "healthy_processor_status").Result()
	log.Println("Healthy Processor Status:", m)

	//---------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------

	// Insere todos os valores especificados no início da lista
	client.LPush(ctx, "payments:queue", payload).Err()
	// client.LPush(ctx, "payments_pending", payload)

	// processingQueue := fmt.Sprintf("payments:processing:%d", 1)

	// //Retorna e remove atomicamente o último elemento da lista
	// result, err := client.RPopLPush(ctx, "payments:queue", processingQueue).Result()
	// log.Println("RPopLPush result:", result, "error:", err)

	// // //Remove as primeiras ocorrências de elementos iguais a elementda lista armazenada
	// // client.LRem(ctx, processingQueue, 1, result)

	// // //Insere todos os valores no início da lista armazenada
	// client.LPush(ctx, "payments:queue", result)
	// client.LRem(ctx, processingQueue, 1, result)

	// // client.LRem(ctx, processingQueue, 1, result)

	//---------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------

	// result, err := client.SetNX(ctx, "health_check_lock", "locked", 10*time.Second).Result()
	// log.Println("1-Health check lock result:", result, "error:", err)
	// result, err = client.SetNX(ctx, "health_check_lock", "locked", 10*time.Second).Result()
	// log.Println("2-Health check lock result:", result, "error:", err)

	// // 2025/07/24 10:05:10 1-Health check lock result: true error: <nil>
	// // 2025/07/24 10:05:10 2-Health check lock result: false error: <nil>

	// resultDel, err := client.Del(ctx, "health_check_lock").Result()
	// log.Println("3-Health check lock delete result:", resultDel, "error:", err)
	// resultDel, err = client.Del(ctx, "health_check_lock").Result()
	// log.Println("4-Health check lock delete result:", resultDel, "error:", err)
	// // 2025/07/24 10:05:10 3-Health check lock delete result: 1 error: <nil>
	// // 2025/07/24 10:05:10 4-Health check lock delete result: 0 error: <nil>

	// result, err = client.SetNX(ctx, "health_check_lock", "locked", 10*time.Second).Result()
	// log.Println("5-Health check lock result:", result, "error:", err)
	// // 2025/07/24 10:05:10 5-Health check lock result: true error: <nil>

	//---------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------

	// client.XAdd(ctx, &redis.XAddArgs{
	// 	Stream: "tickets",
	// 	MaxLen: 0,
	// 	//MaxLenApprox: 0,
	// 	ID: "",
	// 	Values: map[string]interface{}{
	// 		"whatHappened": string("ticket received"),
	// 		"ticketID":     int(rand.Intn(100000000)),
	// 		"ticketData":   string("some ticket data"),
	// 	},
	// })

	// // set, err := a.db.SetNX(ctx, "correlation:" + payment.CorrelationId, "1", 1*time.Minute).Result()
	// // if err != nil {
	// // 	slog.Warn("Redis error", "err", err)
	// // } else if !set {
	// // 	slog.Debug("Duplicate correlationId, skipping", "correlationId", payment.CorrelationId)x
	// // }

	// subject := "tickets"
	// consumersGroup := "tickets-consumer-group"
	// client.XGroupCreate(ctx, subject, consumersGroup, "0").Err()

	// for {
	// 	entries, err := client.XReadGroup(ctx,
	// 		&redis.XReadGroupArgs{
	// 			Group:    consumersGroup,
	// 			Consumer: uuid.New().String(),
	// 			Streams:  []string{subject, ">"},
	// 			Count:    2,
	// 			Block:    0,
	// 			NoAck:    false,
	// 		}).Result()

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	for i := 0; i < len(entries[0].Messages); i++ {
	// 		messageID := entries[0].Messages[i].ID
	// 		values := entries[0].Messages[i].Values
	// 		eventDescription := fmt.Sprintf("%v", values["whatHappened"])
	// 		ticketID := fmt.Sprintf("%v", values["ticketID"])
	// 		ticketData := fmt.Sprintf("%v", values["ticketData"])

	// 		if eventDescription == "ticket received" {
	// 			err := handleNewTicket(ticketID, ticketData)
	// 			if err != nil {
	// 				log.Fatal(err)
	// 			}
	// 			client.XAck(ctx, subject, consumersGroup, messageID)
	// 		}
	// 	}
	// }

	//---------------------------------------------------------------------------------
	//---------------------------------------------------------------------------------

	log.Println("Finishing Rinha Backend 2025...")
}

func handleNewTicket(ticketID string, ticketData string) error {
	log.Printf("Handling new ticket id : %s data %s\n", ticketID, ticketData)
	// time.Sleep(100 * time.Millisecond)
	return nil
}
