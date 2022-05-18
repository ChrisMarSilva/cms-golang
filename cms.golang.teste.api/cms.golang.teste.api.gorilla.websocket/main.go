package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.api.websocket
// go get github.com/gorilla/mux
// go get github.com/gorilla/websocket 
// go mod tidy

// go run main.go

func main() {

	fmt.Println("websocket - Port 9100")

	router := mux.NewRouter()
	router.HandleFunc("/socket", WsEndpoint)
	log.Fatal(http.ListenAndServe(":9100", router))

}

var (
	wsUpgrader = websocket.Upgrader {
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
	wsConn *websocket.Conn
)

type Message struct {
	Greeting string `json:"greeting"`
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {

	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	var err error
	wsConn, err = wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("could not upgrade: %s\n", err.Error())
		return
	}
	defer wsConn.Close()

	iNum := 0

	for {
		var msg Message
		err := wsConn.ReadJSON(&msg)
		if err != nil {
			fmt.Printf("error reading JSON: %s\n", err.Error())
			break
		}
		iNum++
		fmt.Printf("Message Received: %s\n", msg.Greeting)
		SendMessage("Hello, " + msg.Greeting + "! #" +strconv.Itoa(iNum))
	}

}

func SendMessage(msg string) {
	err := wsConn.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		fmt.Printf("error sending message: %s\n", err.Error())
	}
}
