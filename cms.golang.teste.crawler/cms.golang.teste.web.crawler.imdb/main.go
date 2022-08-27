package main

import (
	"fmt"
	"log"
	"strings"
	"encoding/json"
	"flag"
	
	"github.com/gocolly/colly/v2"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.web.crawler.imdb
// go get github.com/gocolly/colly/v2
// go mod tidy

// go run main.go

func main() {

	month := flag.Int("month", 1, "Month to fetch bitchdays for")
	day := flag.Int("day", 1, "Day to fetch bitchdays for")
	flag.Parse()

	crawl(*month, *day)

}

type Star struct {
	Name      string
	Photo     string
	JobTitle  string
	BirthDate string
	Bio       string
	TopMovies []Movie
}

type Movie struct {
	Title string
	Year  string
}

func crawl(month int, day int) {

	c := colly.NewCollector(colly.AllowedDomains("imdb.com", "www.imdb.com"))

	infoCollector := c.Clone()

	c.OnHTML(".mode-detail", func(e *colly.HTMLElement) {
		link := e.ChildAttr("div.lister-item-image > a", "href")
		profileUrl := e.Request.AbsoluteURL(link)
		infoCollector.Visit(profileUrl)
		// fmt.Printf("[profileUrl] Link found: %q -> %s\n", e.Text, link)
	})

	c.OnHTML("a.lister-page-next", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		nextPage := e.Request.AbsoluteURL(link)
		c.Visit(nextPage)
		// fmt.Printf("[nextPage] Link found: %q -> %s\n", e.Text, link)
	})

	infoCollector.OnHTML("#content-2-wide", func(e *colly.HTMLElement) {
		tmpProfile := Star {}
		tmpProfile.Name = e.ChildText("h1.header > span.itemprop")
		tmpProfile.Photo = e.ChildAttr("#name-poster", "src")
		tmpProfile.JobTitle = e.ChildText("#name-job-categories > a > span.itemprop")
		tmpProfile.BirthDate = e.ChildAttr("#name-born-info time", "datetime")
		tmpProfile.Bio = strings.TrimSpace(e.ChildText("#name-bio-text > div.name.trivia-bio-text > div.inline"))
		// tmpProfile.TopMovies = []Movie{}

		e.ForEach("div.knownfor-title", func(_ int, kf *colly.HTMLElement) {
			tmpMovie := Movie{}
			tmpMovie.Title = kf.ChildText("div.knownfor-title-role > a.knownfor-ellipsis")
			tmpMovie.Year = kf.ChildText("div.knownfor-year > span.knownfor-ellipsis")
			tmpProfile.TopMovies = append(tmpProfile.TopMovies, tmpMovie)
		})

		js, err := json.MarshalIndent(tmpProfile, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(js))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	infoCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting profile URL", r.URL.String())
	})

	startUrl := fmt.Sprintf("https://www.imdb.com/search/name/?birth_monthday=%d-%d", month, day)
	c.Visit(startUrl)

}
