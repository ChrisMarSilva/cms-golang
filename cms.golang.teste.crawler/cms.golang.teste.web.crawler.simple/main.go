package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	
	"golang.org/x/net/html"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.web.crawler.simple
// go get golang.org/x/net/html
// go mod tidy

// go run main.go https://www.wikipedia.org
// go run main.go https://www.wikipedia.org https://www.medium.com
// go run main.go

func main() {

	foundUrls := make(map[string]bool)
	seedUrls := os.Args[1:]

	chUrls := make(chan string)
	chFinished := make(chan bool)

	for _, url := range seedUrls {
		go crawl(url, chUrls, chFinished)
	}

	for c := 0; c<len(seedUrls); {
		select {
		case url := <- chUrls:
			foundUrls[url] = true
		case <-chFinished:
			c++
		}
	}

	fmt.Println("\nFound", len(foundUrls), "unique urls:\n")

	for url, _ := range foundUrls {
		fmt.Println("-" + url)
	}

	close(chUrls)

}

func crawl(url string, ch chan string, chFinished chan bool) {

	res, err := http.Get(url)

	defer func(){
		chFinished <- true 
	}()

	if err != nil {
		fmt.Println("ERRO: failed to crawl:", url)
		return
	}

	b := res.Body
	defer b.Close()

	z := html.NewTokenizer(b)

	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := z.Token()
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}
			ok, url := getHref(t)
			if !ok {
				continue
			}
			hasProto := strings.Index(url, "http") == 0
			if hasProto {
				ch <- url
			}

		}
	}

}

func getHref(t html.Token) (ok bool, href string) {
	ok = false
	href = ""
	for _, a := range t.Attr {
		if a.Key == "href" {
			ok = true
			href = a.Val
			break
		}
	}
	return // ok, href
}
