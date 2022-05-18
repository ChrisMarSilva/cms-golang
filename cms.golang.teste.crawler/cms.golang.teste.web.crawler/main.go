package main

import (
	"fmt"
	"flag"

	"github.com/ChrisMarSilva/cms.golang.teste.web.crawler/website"
	"github.com/ChrisMarSilva/cms.golang.teste.web.crawler/crawler"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.web.crawler
// go get golang.org/x/net/html
// go get github.com/mattn/go-sqlite3
// go get github.com/PuerkitoBio/goquery
// go get github.com/gorilla/websocket
// go get nhooyr.io/websocket
// go mod tidy

// go run main.go -action=batata
// go run main.go -url=https://linuxtips.com.br/
// go run main.go

var (
	link    string
	action  string
)

func init() {
	flag.StringVar(&link, "url", "https://aprendagolang.com.br/", "url para iniciar")
	flag.StringVar(&action, "action", "website", "qual serviço iniciar") // website // webcrawler
}

func main() {

	flag.Parse()

	// var err error
	// db.Delete()
	// db.SelectTotal()

	switch action {
	case "website": 
		website.Run()
	case "webcrawler": 
		done := make(chan bool)
		wb := crawler.NewWebCrawler()
		go wb.VisitLink(link)

		for log := range wb.Log() {
			fmt.Println("chan", log)
		}
		<-done

	default:
		fmt.Printf("action '%s' não reconhecida!", action)
	}

}