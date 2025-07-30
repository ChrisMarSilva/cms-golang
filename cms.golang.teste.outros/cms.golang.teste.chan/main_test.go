package main

import (
	"testing"
	"testing/synctest"
)

// GOEXPERIMENT=synctest  go test ./...

func TestMainEvents(t *testing.T) {
	synctest.Run(func() {
		processor := NewEventProcessor(10)
		processor.Start()

		processor.Process("event1")
		processor.Process("event2")
		processor.Process("event3")

		synctest.Wait()
		processor.Stop()

		processed := processor.GetAllEvents()

		if !processed["event1"] {
			t.Error("Event 'event1' was not processed")
		}

		if !processed["event2"] {
			t.Error("Event 'event2' was not processed")
		}

		if !processed["event3"] {
			t.Error("Event 'event3' was not processed")
		}

		if len(processed) != 3 {
			t.Errorf("Expected 3 processed events, got %d", len(processed))
		}
	})
}
