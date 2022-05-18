package main

import (
	"github.com/ChrisMarSilva/cms.golang.tnb.api/server"
)

// go mod init github.com/ChrisMarSilva/cms.golang.tnb.api
// go mod tidy

// go run main.go

func main() {
	s := server.NewServer()
	s.Run()
}
