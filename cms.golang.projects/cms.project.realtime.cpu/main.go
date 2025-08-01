package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"cms.project.realtime.cpu/hardware"
	"github.com/coder/websocket"
)

// go mod init github.com/chrismarsilva/cms.project.realtime.cpu
// go get -u "github.com/shirou/gopsutil/v4"
// go get -u "github.com/shirou/gopsutil/v4/disk"
// go get -u "github.com/shirou/gopsutil/v4/cpu"
// go get -u "github.com/coder/websocket"
// go get -u "nhooyr.io/websocket"
// go get -u "github.com/cosmtrek/air@latest"
// go mod tidy
// air init
// air
// go run main.go

type server struct {
	subScriberMessageBuffer int
	mux                     http.ServeMux
	subscribersMutex        sync.Mutex
	subscribers             map[*subscriber]struct{}
}

type subscriber struct {
	msgs chan []byte
}

func NewServer() *server {
	s := &server{
		subScriberMessageBuffer: 10,
		subscribers:             make(map[*subscriber]struct{}),
	}

	s.mux.Handle("/", http.FileServer(http.Dir("./pages")))
	s.mux.HandleFunc("/ws", s.subscriberHandler)
	return s
}

func (s *server) subscriberHandler(w http.ResponseWriter, r *http.Request) {
	err := s.subscribe(r.Context(), w, r)
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *server) subscribe(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	subscriber := &subscriber{msgs: make(chan []byte, s.subScriberMessageBuffer)}
	s.addSubscriber(subscriber)

	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		return err
	}
	defer c.CloseNow()

	ctx = c.CloseRead(ctx)
	for {
		select {
		case msg := <-subscriber.msgs:
			ctx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()

			//log.Println("Sending message to subscriber:", string(msg))
			err := c.Write(ctx, websocket.MessageText, msg)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *server) addSubscriber(subscriber *subscriber) {
	s.subscribersMutex.Lock()
	defer s.subscribersMutex.Unlock()

	s.subscribers[subscriber] = struct{}{}
}

func (s *server) publishMsg(msg []byte) {
	s.subscribersMutex.Lock()
	defer s.subscribersMutex.Unlock()

	for subscriber := range s.subscribers {
		subscriber.msgs <- msg
	}
}

func main() {
	log.Println("Starting monitor server on port 8080")
	srv := NewServer()

	go func(s *server) {
		for {
			systemSection, err := hardware.GetSystemSection()
			if err != nil {
				log.Println(err)
			}

			diskSection, err := hardware.GetDiskSection()
			if err != nil {
				log.Println(err)
			}

			cpuSection, err := hardware.GetCpuSection()
			if err != nil {
				log.Println(err)
			}

			msg := []byte(`
			<div hx-swap-oob="innerHTML:#update-timestamp"><p><i style="color: green" class="fa fa-circle"></i> ` + time.Now().Format(time.DateTime) + `</p></div>
			<div hx-swap-oob="innerHTML:#system-data">` + systemSection + `</div>
			<div hx-swap-oob="innerHTML:#cpu-data">` + cpuSection + `</div>
			<div hx-swap-oob="innerHTML:#disk-data">` + diskSection + `</div>`)
			srv.publishMsg(msg)

			time.Sleep(1 * time.Second)
		}
	}(srv)

	log.Fatal(http.ListenAndServe(":8080", &srv.mux))
}
