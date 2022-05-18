package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly/v2"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.web.scraping
// go get -u github.com/gocolly/colly/v2
// go get -u github.com/PuerkitoBio/goquery
// go get -u github.com/go-sql-driver/mysql
// go mod tidy

// go run main.go

func main() {

	log.Println("INI")

	// TesteSiteTamoNaBolsa()
	TesteSiteFilmes()

	ExampleScrape()

	log.Println("FIM")
}

func ExampleScrape() {

	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})

}

func TesteSiteFilmes() {

	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)
	c.UserAgent = "xy"
	c.AllowURLRevisit = true

	// c.WithTransport(&http.Transport{
	// 	Proxy: http.ProxyFromEnvironment,
	// 	DialContext: (&net.Dialer{
	// 		Timeout:   30 * time.Second,
	// 		KeepAlive: 30 * time.Second,
	// 		DualStack: true,
	// 	}).DialContext,
	// 	MaxIdleConns:          100,
	// 	IdleConnTimeout:       90 * time.Second,
	// 	TLSHandshakeTimeout:   10 * time.Second,
	// 	ExpectContinueTimeout: 1 * time.Second,
	// }

	// c := colly.NewCollector(
	// 	colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	// )

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML(`div[id=content]`, func(row *colly.HTMLElement) {
		row.ForEach("article", func(_ int, el *colly.HTMLElement) { // `article[id=post-139597]`

			title := el.ChildText("a[href]")
			iIdx := strings.Index(title, " Torrent")
			if iIdx > 0 {
				title = title[:iIdx]
			}

			// imdb := el.DOM.Find(".entry-content").Text()
			// iIdx = strings.Index(imdb, "IMDb:")
			// if iIdx == 0 {
			// 	iIdx = strings.Index(imdb, "IMDB:")
			// }
			// if iIdx > 0 {
			// 	imdb = imdb[iIdx+6 : iIdx+6+10]
			// 	// iIdx = strings.Index(imdb, "Ano de")
			// 	// if iIdx > 0 {
			// 	// 	imdb = imdb[:iIdx]
			// 	// }
			// }

			log.Println("Titulo:", title)
			//log.Println("IMDb:", imdb)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println("Finished")
	})

	for i := 1; i <= 2; i++ {
		err := c.Visit("https://comandotorrent.net/page/" + strconv.Itoa(i) + "/")
		if err != nil {
			log.Fatalln(err)
		}
	}

}

func TesteSiteTamoNaBolsa() {

	var start time.Time

	c := colly.NewCollector()

	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })

	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")
	// 	log.Printf("Link found: %q -> %s\n", e.Text, link)
	// 	c.Visit(e.Request.AbsoluteURL(link))
	// })

	var empresas []Empresa
	c.OnHTML("table tbody", func(tbl *colly.HTMLElement) {
		tbl.ForEach("tr", func(_ int, row *colly.HTMLElement) {

			sData := ""
			sEmpresa := ""
			sCodigos := ""

			// sData := row.Attr("td:nth-child(0)")
			// sEmpresa := row.Attr("td:nth-child(1)")
			// sCodigos := row.Attr("td:nth-child(2)")

			row.ForEach("td", func(_ int, el *colly.HTMLElement) {
				switch el.Index {
				case 0:
					sData = el.Text
				case 1:
					sEmpresa = el.Text
				case 2:
					sCodigos = el.Text
				}
			})

			if sData != "Data" {

				sEmpresa = strings.TrimSpace(sEmpresa)
				sEmpresa = strings.Replace(sEmpresa, "Resultado do ", "", -1)
				sEmpresa = strings.Replace(sEmpresa, "Resultado da ", "", -1)
				sEmpresa = strings.ToUpper(sEmpresa)

				sData = strings.TrimSpace(sData)
				sData = sData[6:10] + sData[3:5] + sData[0:2] // 29/03/2022 => // 20220329

				sCodigos = strings.TrimSpace(sCodigos)
				sCodigos = strings.ToUpper(sCodigos)

				lstCodigo := strings.Split(sCodigos, ",")

				if len(lstCodigo) > 0 {
					for _, codigo := range lstCodigo {
						empresa := &Empresa{Nome: sEmpresa, Codigo: strings.TrimSpace(codigo), Data: sData, Hora: ""}
						empresas = append(empresas, *empresa)
					}
				} else {
					empresa := &Empresa{Nome: sEmpresa, Codigo: sCodigos, Data: sData, Hora: ""}
					empresas = append(empresas, *empresa)
				}

				//log.Println("Dt:", sData, "; Empr:", sEmpresa, "; Cod:", sCodigo)

			}

		})
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	// err := c.Visit("http://go-colly.org/")
	err := c.Visit("https://comoinvestir.thecap.com.br/agenda-divulgacao-resultados-bovespa-4t21-2021")
	if err != nil {
		log.Fatalln(err)
	}

	dsn := "tamonabo_rootcms:senha@tcp(nspro44.hostgator.com.br:3306)/database?parseTime=true"
	dsn = "root:senha@tcp(localhost:3306)/database"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close() // adiar o fechamento até depois que a função principal terminar

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("empresas", len(empresas))

	start = time.Now()
	//tx, _ := db.Begin()
	// db.Exec("DELETE FROM TBEMPRESA_FINAN_AGENDA")
	// db.Exec("TRUNCATE TABLE TBEMPRESA_FINAN_AGENDA")
	// for _, empresa := range empresas {
	// 	db.Exec("INSERT INTO TBEMPRESA_FINAN_AGENDA (NOME, CODIGO, DIVULGACAO, HORARIO) VALUES (?, ?, ?, ?)", empresa.Nome, empresa.Codigo, empresa.Data, empresa.Hora)
	// }
	db.Exec("UPDATE TBEMPRESA_FINAN_AGENDA SET IDEMPRESA = ( SELECT MAX(IDEMPRESA) FROM TBEMPRESA_ATIVO A WHERE A.CODIGO = TBEMPRESA_FINAN_AGENDA.CODIGO )")
	//tx.Commit()
	log.Println("db.Exec=", time.Since(start))

}

type Empresa struct {
	Nome   string `json:"nome" db:"NOME"`
	Codigo string `json:"codigo" db:"CODIGO"`
	Data   string `json:"data" db:"DIVULGACAO"`
	Hora   string `json:"hora" db:"HORARIO"`
}
