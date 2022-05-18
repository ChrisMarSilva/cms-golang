package crawler

import (
	"fmt"
	"time"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
	"github.com/ChrisMarSilva/cms.golang.teste.web.crawler/db"
	"github.com/ChrisMarSilva/cms.golang.teste.web.crawler/model"
)

type webCrawler struct {
	log chan string
}

func NewWebCrawler() *webCrawler {
	wb := &webCrawler{
		log: make(chan string, 10),
	}
	return wb
}

func (wb *webCrawler) Log() chan string {
	return wb.log
}

func (wb *webCrawler) VisitLink(link string) {

	wb.log <- fmt.Sprintf("visitando site: %s", link)

	resp, err := http.Get(link)
	if err != nil {
		wb.log <- fmt.Sprintf("[error] http.Get: %s", link)
		return
	}

	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		wb.log <- fmt.Sprintf("[error] status diferente de 200: %d", resp.StatusCode)
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		wb.log <- fmt.Sprintf("[error] html.Parse: %s", link)
		return
	}

	wb.extractLinks(doc)

}

func (wb *webCrawler) extractLinks(node *html.Node) {

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {

			if attr.Key != "href" {
				continue
			}

			linkLocal, err := url.Parse(attr.Val)
			if err != nil || linkLocal.Host == "" || linkLocal.Scheme == "" || linkLocal.Scheme == "tel" || linkLocal.Scheme == "mailto" || linkLocal.Scheme == "javascript" || linkLocal.Scheme == "ws" || linkLocal.Scheme == "wss" {
				continue
			}

			if linkLocal.Scheme != "http" && linkLocal.Scheme != "https" {
				continue
			}
			err = db.SelectExist(linkLocal.String())
			if err == nil {
				wb.log <- fmt.Sprintf("link já visitado: %s", linkLocal.String())
				continue
			}

			visitedLink := model.VisitedLink{
				WebSite:     linkLocal.Host,
				Link:        linkLocal.String(),
				VisitedDate: time.Now().Format("2006-01-02 15:04:05.000"),
			}
			
			err = db.Insert(visitedLink)
			if err != nil {
				wb.log <- fmt.Sprintf("[error] db.Insert: %s", err)
			}

			go wb.VisitLink(linkLocal.String())

		} 
	} 

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		wb.extractLinks(c)
	}

}
