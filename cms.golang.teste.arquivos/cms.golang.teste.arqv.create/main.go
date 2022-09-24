package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.arqv.create
// go get -u xxxxxxxxx
// go mod tidy

// go run main.go

func main() {
	log.Println("INI")
	var start time.Time = time.Now()

	CreateFileWriteString("teste01.txt")
	CreateFileWriteBytes("teste02.txt", []byte("Hello I am a byte slice!"))
	CreateFileWriteLines("teste03.txt", []string{"Linha01", "Linha02", "Linha03"})
	CreateFileWriteAppend("teste03.txt", "Hello1")
	CreateFileWriteAppend("teste03.txt", "Hello2")
	CreateFileWriteAppend("teste04.txt", "Hello1")
	CreateFileWriteAppend("teste04.txt", "Hello2")

	log.Println("FIM:", time.Since(start))
}

func CreateFileWriteString(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Println("Error os.Create(): " + err.Error())
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hey new file!")
	if err != nil {
		log.Println("Error file.WriteString(): " + err.Error())
		return
	}
}

func CreateFileWriteBytes(filename string, data []byte) {
	file, err := os.Create(filename)
	if err != nil {
		log.Println("Error os.Create(): " + err.Error())
		return
	}
	defer file.Close()

	size, err := file.Write(data)
	if err != nil {
		log.Println("Error file.Write(): " + err.Error())
		return
	}
	log.Printf("Wrote %d bytes to file.\n", size)
}

func CreateFileWriteLines(filename string, lines []string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Println("Error os.Create(): " + err.Error())
		return
	}
	defer file.Close()

	for _, val := range lines {
		_, err = fmt.Fprintln(file, val)
		if err != nil {
			log.Println("Error fmt.Fprintln(): " + err.Error())
			return
		}
	}
}

func CreateFileWriteAppend(filename string, data string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644|os.ModeAppend|os.ModePerm|fs.ModeAppend) // os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, fs.ModeAppend) // os.O_CREATE|
	if err != nil {
		log.Println("Error os.OpenFile(): " + err.Error())
		return
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, data)
	if err != nil {
		log.Println("Error fmt.Fprintln(): " + err.Error())
		return
	}
}
