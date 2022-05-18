package main

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	_ "time"

	"github.com/nats-io/nats.go"
)

type Event struct {
	ID  int    `json:"id"`
	Msg string `json:"msg"`
}

func main() {

	conn, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to NATS at " + nats.DefaultURL)

	i := 0
	// var start time.Time

	conn.Subscribe("event", func(msg *nats.Msg) {
		// if i == 0 {
		// 	start = time.Now()
		// }
		var ev Event
		json.Unmarshal(msg.Data, &ev)
		i++
		log.Println("Event received", i, ev)
		// if i >= 1000 {
		// 	log.Println("Subscribe=", time.Since(start), i)
		// 	i = 0
		// }
	})

	log.Println("depois Subscribe")
	runtime.Goexit()

	fmt.Println("FIM")
}
