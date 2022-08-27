package main

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

// go mod init github.com/chrismarsilva/cms.golang.tnb.foxbit
// go get github.com/gorilla/websocket
// go mod tidy

// go run main.go

func main() {

	// log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "wss", Host: "api.foxbit.com.br", Path: "/"}
	log.Printf("connecting to %s", u.String())

	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer ws.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			log.Println("ws.ReadMessage()")
			_, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("222 - recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	authenticated := false

	for {
		select {

		case <-done:
			return

		case t := <-ticker.C:

			//-----------

			// log.Println("TextMessage")
			err := ws.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("TextMessage WriteMessage Err:", err)
				return
			}

			//-----------

			// var payloadAuth PayloadAuth
			// payloadAuth.UserName = "email@gmail.com"
			// payloadAuth.Password = "senha"
			// outPayloadAuth, _ := json.Marshal(payloadAuth)

			if !authenticated {
				var payloadAuth PayloadAuth2
				payloadAuth.UserId = 10000
				payloadAuth.SessionToken = "SessionToken"
				outPayloadAuth, _ := json.Marshal(payloadAuth)

				var frameAuth MessageFrame
				frameAuth.m = 0 //MessageType ( 0_Request / 1_Reply / 2_Subscribe / 3_Event / 4_Unsubscribe / Error )
				frameAuth.i = 0 //Sequence Number
				frameAuth.n = "WebAuthenticateUser"
				frameAuth.o = string(outPayloadAuth)
				outFrameAuth, _ := json.Marshal(frameAuth)

				log.Println("WebAuthenticateUser")
				err = ws.WriteMessage(websocket.TextMessage, []byte(outFrameAuth))
				if err != nil {
					log.Println("WebAuthenticateUser WriteMessage Err:", err)
					return
				}
				authenticated = true
			}

			//-----------

			var payloadOrderHist PayloadOrderHist
			payloadOrderHist.OMSId = 1
			payloadOrderHist.AccountId = 258442 // 81
			outPayloadOrderHist, _ := json.Marshal(payloadOrderHist)

			var frameOrderHist MessageFrame
			frameOrderHist.m = 0 //MessageType ( 0_Request / 1_Reply / 2_Subscribe / 3_Event / 4_Unsubscribe / Error )
			frameOrderHist.i = 0 //Sequence Number
			frameOrderHist.n = "GetOrderHistory"
			frameOrderHist.o = string(outPayloadOrderHist)
			outFrameOrderHist, _ := json.Marshal(frameOrderHist)

			log.Println("GetOrderHistory")
			err = ws.WriteMessage(websocket.TextMessage, []byte(outFrameOrderHist))
			if err != nil {
				log.Println("GetOrderHistory WriteMessage Err:", err)
				return
			}

			//-----------

		case <-interrupt:
			log.Println("interrupt")
			err := ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}

	}

}

type MessageFrame struct {
	m int    //MessageType ( 0_Request/1_Reply/2_Subscribe /3_Event /4_Unsubscribe /Error )
	i int    //Sequence Number
	n string //Endpoint
	o string //Payload
}

type PayloadAuth struct {
	UserName string
	Password string
}

type PayloadAuth2 struct {
	UserId       int
	SessionToken string
}

type PayloadOrderHist struct {
	OMSId     int
	AccountId int
}
