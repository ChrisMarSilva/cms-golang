package main

import (
	"io"
	"log"
	"os"
	"time"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.ler.dir.and.upload
// go get -u xxxxxxxxx
// go mod tidy

// go run main.go

func main() {
	log.Println("INI")
	var start time.Time = time.Now()

	dir, err := os.Open("C:\\Users\\chris\\AppData\\Local\\Temp")
	if err != nil {
		log.Println("Error os.Open(): " + err.Error())
		return
	}

	for {
		files, err := dir.Readdir(1) // ler 1 arquivo por vez
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Error os.Readdir(): " + err.Error())
			continue
		}
		log.Println("Name: " + files[0].Name())
	}

	log.Println("FIM:", time.Since(start))
}
