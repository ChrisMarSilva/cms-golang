package main

import (
	"log"
	"strings"
	"time"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.data.hora
// go get -u XXXXXXXXXXXX
// go mod tidy

// go run main.go

func main() {
	log.Println("INI")
	var start time.Time = time.Now()

	now := time.Now()

	log.Println("")
	log.Println("yyyy-MM-dd:", now.Format("2006-01-02"))
	log.Println("yy-MM-dd:", now.Format("06-01-02"))

	log.Println("")
	log.Println("yyyyMMdd:", now.Format("20060102"))
	log.Println("yyMMdd:", now.Format("060102"))

	log.Println("")
	log.Println("HHmmss:", now.Format("150405"))
	log.Println("HHmmssfff:", strings.Replace(now.Format("150405.000"), ".", "", -1))

	log.Println("")
	log.Println("yyyyMMddHHmmss:", now.Format("20060102150405"))
	log.Println("yyyyMMddHHmmssfff:", strings.Replace(now.Format("20060102150405.000"), ".", "", -1))

	log.Println("")
	log.Println("yyyy-MM-dd HH:mm:ss:", now.Format("2006-01-02 15:04:05"))
	log.Println("yyyy-MM-dd HH:mm:ss.fff:", now.Format("2006-01-02 15:04:05.000"))

	log.Println("")
	log.Println("FIM:", time.Since(start))
}
