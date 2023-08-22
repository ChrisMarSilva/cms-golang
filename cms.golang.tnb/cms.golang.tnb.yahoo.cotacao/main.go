package main

import (
	//"io/ioutil"
	//"regexp"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/piquette/finance-go/quote"
)

// go get -u github.com/gocolly/colly
// go get -u github.com/gocolly/colly/v2
// go mod tidy
// go run main.go

func main() {
	symbol := "ITSA4.SA"
	// Teste_FinanceGo(symbol)
	// Teste_YahooFinance(symbol)
	//Teste_YahooFinance_Colly(symbol)
	Teste_YahooFinance_Colly2(symbol)
}

func Teste_FinanceGo(symbol string) {
	qq, err := quote.Get(symbol)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(qq)
	fmt.Printf("------- %v -------\n", qq.ShortName)
	fmt.Printf("Current Price: $%v\n", qq.Ask)
	fmt.Printf("52wk High: $%v\n", qq.FiftyTwoWeekHigh)
	fmt.Printf("52wk Low: $%v\n", qq.FiftyTwoWeekLow)
}

func Teste_YahooFinance(symbol string) {
	url := "https://finance.yahoo.com/quote/" + symbol

	res, err := http.Get(url)
	if err != nil {
		//log.Fatalln(err)
		log.Println("Erro ao fazer a requisição:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		//log.Fatalf("status code %d: %s", res.StatusCode, res.Status)
		log.Printf("status code: %s", res.Status)
		return
	}

	/*
		// body, err := ioutil.ReadAll(response.Body)
		// if err != nil {
		// 	fmt.Println("Erro ao ler a resposta:", err)
		// 	return
		// }

		// // Procurando o padrão da cotação usando expressões regulares
		// pattern := regexp.MustCompile(`"regularMarketPrice":{"raw":([\d.]+)`)
		// matches := pattern.FindStringSubmatch(string(body))

		// // Verificando se a cotação foi encontrada
		// if len(matches) != 2 {
		// 	fmt.Println("Não foi possível obter a cotação de ITSA4")
		// 	return
		// }

		// price := matches[1] // Extraindo o valor da cotação
		// fmt.Printf("Cotação de ITSA4: %s\n", price)
	*/

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("Erro ao ler a resposta:", err)
		return //log.Fatalln(err)
	}

	marketPrice := doc.Find("[data-field=regularMarketPrice][data-symbol=" + symbol + "]").Text()
	if marketPrice == "" {
		log.Fatalf("cannot access market price")
	}

	fmt.Printf("%s price: %s\n", symbol, marketPrice)
}

func Teste_YahooFinance_Colly(symbol string) {
	c := colly.NewCollector()

	headers := map[string]string{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36",
		"Accept-Language": "en-US,en;q=0.9",
	}

	c.OnRequest(func(r *colly.Request) {
		for key, value := range headers {
			r.Headers.Set(key, value)
		}
	})

	c.OnHTML("span[data-reactid='50']", func(e *colly.HTMLElement) {
		price := e.Text
		fmt.Printf("Cotação de ITSA4: %s\n", price)
	})

	url := "https://finance.yahoo.com/quote/" + symbol
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Erro ao visitar a página:", err)
	}
}

func Teste_YahooFinance_Colly2(symbol string) {
	// Cria uma nova instância do colly
	c := colly.NewCollector()

	// Define um cliente HTTP personalizado para controlar o redirecionamento
	// client := &http.Client{
	// 	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	// 		return http.ErrUseLastResponse
	// 	},
	// 	Timeout: time.Second * 10,
	// }

	// Define os cabeçalhos personalizados
	headers := map[string]string{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36",
		"Accept-Language": "en-US,en;q=0.9",
	}

	// Define os cabeçalhos personalizados para a requisição
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", headers["User-Agent"])
		r.Headers.Set("Accept-Language", headers["Accept-Language"])
		//r.Client = client
	})

	// Procura o valor da cotação na página do Yahoo Finance
	// c.OnHTML("#quote-header-info span[data-reactid='50']", func(e *colly.HTMLElement) {
	c.OnHTML("fin-streamer[data-symbol='ITSA4.SA'][data-field='regularMarketPrice']", func(e *colly.HTMLElement) {
		price := strings.TrimSpace(e.Text)
		fmt.Printf("Cotação de ITSA4: %s\n", price)
	})

	c.OnHTML("fin-streamer[data-symbol='ITSA4.SA'][data-field='regularMarketChangePercent']", func(e *colly.HTMLElement) {
		percente := strings.TrimSpace(e.Text)
		fmt.Printf("Percentual de ITSA4: %s\n", percente)
	})

	// Manipula possíveis erros durante o scraping
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Erro durante o scraping:", err)
	})

	// Manipula redirecionamentos
	c.OnResponse(func(r *colly.Response) {
		if r.Request.URL.Host == "consent.yahoo.com" {
			// Aceita o consentimento de cookies do Yahoo Finance automaticamente
			err := c.Visit("https://consent.yahoo.com/collectConsent?sessionId=3_cc-session_178170f0-87b5-4b9e-86df-3adec25b594d")
			if err != nil {
				log.Println("Erro ao aceitar o consentimento de cookies:", err)
			}
		}
	})

	// Visita a página de ITSA4 no Yahoo Finance
	url := "https://br.financas.yahoo.com/quote/" + symbol
	err := c.Visit(url)
	if err != nil {
		log.Println("Erro ao visitar a página:", err)
	}
}
