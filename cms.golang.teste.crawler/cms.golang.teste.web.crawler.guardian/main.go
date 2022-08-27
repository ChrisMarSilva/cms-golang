package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
	
	"github.com/PuerkitoBio/goquery"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.web.crawler.guardian
// go get github.com/PuerkitoBio/goquery
// go mod tidy

// go run main.go

func main() {

	baseDomain := "https://www.theguardian.com"
	worklist := make(chan []string)
	var n int
	n++

	go func() { 
		worklist <- []string{"https://www.theguardian.com"} 
	}()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string, baseURL string) {
					foundLinks := Crawl(link, baseDomain)
					if foundLinks != nil {
						worklist <- foundLinks
					}
				}(link, baseDomain)
			}
		}
	}

}

var tokens = make(chan struct{}, 5) // Channel working as a semaphore - using 5 or more tokens likely to overload target site

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func Crawl(targetURl string, baseURL string) []string {
	fmt.Println(targetURl)
	tokens <- struct{}{}
	resp, _ := getRequest(targetURl)
	<-tokens
	links := discoverLinks(resp, baseURL)
	foundUrls := []string{}
	for _, link := range links {
		ok, correctLink := resolveRelativeLinks(link, baseURL)
		if ok {
			if correctLink != "" {
				foundUrls = append(foundUrls, correctLink)
			}
		}
	}
	ParseHTML(resp)
	return foundUrls
}

func getRequest(targetUrl string) (*http.Response, error) {

	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("User-Agent", randomUserAgent())

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func discoverLinks(response *http.Response, baseURL string) []string {

	if response != nil {
		doc, _ := goquery.NewDocumentFromResponse(response)
		foundUrls := []string{}
		if doc != nil {
			doc.Find("a").Each(func(i int, s *goquery.Selection) {
				res, _ := s.Attr("href")
				foundUrls = append(foundUrls, res)
			})
		}
		return foundUrls
	}

	return []string{}
}

func resolveRelativeLinks(href string, baseUrl string) (bool, string) {
	resultHref := checkRelative(href, baseUrl)
	baseParse, _ := url.Parse(baseUrl)
	resultParse, _ := url.Parse(resultHref)
	if baseParse != nil && resultParse != nil {
		if baseParse.Host == resultParse.Host {
			return true, resultHref
		}
		return false, ""
	}
	return false, ""
}

func checkRelative(href string, baseUrl string) string {
	if strings.HasPrefix(href, "/") {
		return fmt.Sprintf("%s%s", baseUrl, href)
	}
	return href
}

func ParseHTML(response *http.Response) {
	// Overwrite method and get the stuff you want from the page
}
