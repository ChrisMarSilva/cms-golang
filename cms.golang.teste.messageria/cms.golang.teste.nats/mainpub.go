package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Event struct {
	ID  int    `json:"id"`
	Msg string `json:"msg"`
}

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to NATS at " + nats.DefaultURL)

	i := 0

	var start time.Time

	start = time.Now()
	for {
		ev := Event{ID: i, Msg: "Hello, Subscriber!"}
		b, _ := json.Marshal(ev)
		nc.Publish("event", b)
		i++
		if i >= 1000 {
			break
		}
	}
	log.Println("Publish=", time.Since(start), i)

	// for range time.Tick(1 * time.Milliseconds) {
	// 	ev := Event{ID: i, Msg: "Hello, Subscriber!"}
	// 	b, _ := json.Marshal(ev)
	// 	nc.Publish("event", b)
	// 	log.Println("Event published", string(b))
	// 	i++
	// }

	fmt.Println("FIM")
}
