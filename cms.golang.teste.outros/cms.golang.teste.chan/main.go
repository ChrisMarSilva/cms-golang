package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.chan
// go mod tidy

// go run main.go

func main() {
	fmt.Println(time.Now().Format(time.RFC3339), "- Starting event processing...")

	processor := NewEventProcessor(10, 1*time.Second)
	processor.Start()

	processor.Process("fast-1")
	processor.Process("fast-2")
	processor.Process("fast3")
	processor.Process("slow-1")
	processor.Process("slow-2")

	time.Sleep(2 * time.Second)
	processor.Stop()

	fmt.Println(time.Now().Format(time.RFC3339), "- Processed events...")
	processedEvents := processor.GetProcessedEvents()
	for event, processed := range processedEvents {
		fmt.Printf("%s -     Event: %s, Processed: %v\n", time.Now().Format(time.RFC3339), event, processed)
	}

	fmt.Println(time.Now().Format(time.RFC3339), "- TimeOut events...")
	timeOutEvents := processor.GetTimeOutEvents()
	for event, timeOut := range timeOutEvents {
		fmt.Printf("%s -     Event: %s, TimeOut: %v\n", time.Now().Format(time.RFC3339), event, timeOut)
	}

	fmt.Println(time.Now().Format(time.RFC3339), "- Event processing completed...")
}

type EventProcessor struct {
	events    chan string
	mu        sync.Mutex
	processed map[string]bool
	timeOut   map[string]bool
	timeOutMu time.Duration
}

func NewEventProcessor(bufferSize int, timeout time.Duration) *EventProcessor {
	return &EventProcessor{
		events:    make(chan string, bufferSize),
		processed: make(map[string]bool),
		timeOut:   make(map[string]bool),
		timeOutMu: timeout,
	}
}

func (ep *EventProcessor) Start() {
	go func() {
		num := 0
		for event := range ep.events {
			num++
			go ep.handleEvent(event)
		}
	}()
}

func (ep *EventProcessor) handleEvent(event string) {
	ctx, cancel := context.WithTimeout(context.Background(), ep.timeOutMu)
	defer cancel()

	doneProcessing := make(chan struct{})

	go func() {
		if !strings.HasPrefix(event, "slow") {
			ep.mu.Lock()
			defer ep.mu.Unlock()

			ep.processed[event] = true
			doneProcessing <- struct{}{}
		}
	}()

	select {
	case <-doneProcessing:
		return
	case <-ctx.Done():
		ep.mu.Lock()
		defer ep.mu.Unlock()

		ep.timeOut[event] = true
		return
	}
}

func (ep *EventProcessor) Stop() {
	close(ep.events)
}

func (ep *EventProcessor) Process(event string) {
	ep.events <- event
}

func (ep *EventProcessor) GetProcessedEvents() map[string]bool {
	ep.mu.Lock()
	defer ep.mu.Unlock()

	result := make(map[string]bool)
	for k, v := range ep.processed {
		result[k] = v
	}

	return result
}

func (ep *EventProcessor) GetTimeOutEvents() map[string]bool {
	ep.mu.Lock()
	defer ep.mu.Unlock()

	result := make(map[string]bool)
	for k, v := range ep.timeOut {
		result[k] = v
	}

	return result
}
