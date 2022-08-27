package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"math/rand"
	"time"
	"github.com/PuerkitoBio/goquery"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.web.crawler.sitemap
// go get github.com/PuerkitoBio/goquery
// go mod tidy

// go run main.go

func main() {
	parser := DefaultParser{}
	url := "https://www.quicksprout.com/sitemap.xml"
	results := ScrapeSitemap(url, parser, 10)
	for _, res := range results {
		fmt.Println(" ====> res:", res)
	}
}

type SeoData struct {
	URL             string
	Title           string
	H1              string
	MetaDescription string
	StatusCode      int
}

type Parser interface {
	GetSeoData(resp *http.Response) (SeoData, error)
}

type DefaultParser struct {
	
}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func ScrapeSitemap(url string, parser Parser, concurrency int) []SeoData {
	results := extractSiteMapURLs(url)
	res := scrapeURLs(results, parser, concurrency)
	return res
}

func extractSiteMapURLs(startURL string) []string {

	worklist := make(chan []string)
	toCraw := []string{}
	var n int 
	n++

	go func() { 
		worklist <- []string{startURL} 
	}()

	for ; n>0; n-- {
		link := <- worklist
		for _, link := range link {
			n++
			go func(link string){
				response, err := makeRequest(link)
				if err != nil {
					log.Printf("Error retrieving URL: %s", link)
				}
				urls, err := extractUrls(response)
				if err != nil {
					log.Printf("Error extracting document from response, URL: %s", link)
				}
				sitemapFiles, pages := isSitemap(urls)
				if sitemapFiles != nil {
					worklist <- sitemapFiles
				}
				for _, page := range pages {
					toCraw =  append(toCraw, page)
				}
			}(link)
		}
	}

	return toCraw
}

func extractUrls(response *http.Response) ([]string, error){

	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}

	results := []string{}
	sel := doc.Find("loc")

	for i := range sel.Nodes {
		loc := sel.Eq(i)
		result := loc.Text()
		results = append(results, result) // results = append(results, sel.Eq(i).Text())
	}

	return results, nil
}

func isSitemap(urls []string) ([]string, []string) {

	sitemapFiles := []string{}
	pages := []string{}

	for _, page := range urls {
		foundSitemap := strings.Contains(page, "xml")
		if foundSitemap == true {
			// fmt.Println("Found Sitemap", page)
			sitemapFiles = append(sitemapFiles, page)
		} else {
			pages = append(pages, page)
		}
	}

	return sitemapFiles, pages
}

func scrapeURLs(urls []string, parser Parser, concurrency int) []SeoData {

	tokens := make(chan struct{}, concurrency)
	var n int
	n++
	worklist := make(chan []string)
	results := []SeoData{}

	go func (){
		worklist <- urls
	}()

	for ; n > 0; n-- {
		list := <-worklist
		for _, url := range list {
			if url != "" {
				n++
				go func(url string, token chan struct{}) {
					log.Printf("Requesting URL: %s", url)
					res, err := scrapePage(url, tokens, parser)
					if err != nil {
						log.Printf("Encountered error, URL: %s", url)
					} else {
						results = append(results, res)
					}
					worklist <- []string{}
				}(url, tokens)
				// break
			}
		}
	}

	return results
}

func scrapePage(url string, token chan struct{}, parser Parser) (SeoData, error) {

	res, err := crawlPage(url, token)
	if err != nil {
		return SeoData{}, err
	}

	data, err := parser.GetSeoData(res)
	if err != nil {
		return SeoData{}, err
	}

	return data, nil
}

func crawlPage(url string, tokens chan struct{}) (*http.Response, error) {

	tokens <- struct{}{}
	resp, err := makeRequest(url)
	<-tokens

	if err != nil {
		return nil, err
	}

	return resp, err
}

func makeRequest(url string) (*http.Response, error){

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	userAgent := randomUserAgent()
	req.Header.Set("user-Agent", userAgent)

	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func (d DefaultParser) GetSeoData(resp *http.Response) (SeoData, error) {

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return SeoData{}, err
	}

	result := SeoData{}
	result.URL = resp.Request.URL.String()
	result.Title = doc.Find("title").First().Text()
	result.H1 = doc.Find("h1").First().Text()
	result.MetaDescription, _ = doc.Find("meta[name^=description]").Attr("content")
	result.StatusCode = resp.StatusCode

	return result, nil
}














