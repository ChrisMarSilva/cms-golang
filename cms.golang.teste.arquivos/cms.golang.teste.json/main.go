package main

import (
	"bytes"
	"encoding/json"
	"log"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.json
// go mod tidy

// go run main.go

type Message map[string]interface{}

func serialize(msg Message) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)
	return b.Bytes(), err
}

func deserialize(b []byte) (Message, error) {
	var msg Message
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return msg, err
}

func main() {

	var msg Message
	b, err := serialize(msg)
	if err != nil {
		log.Println(err)
		return
	}

	msg2, err := deserialize(b)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(msg2)

}
