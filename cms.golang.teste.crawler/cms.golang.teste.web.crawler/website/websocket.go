package website

import (
	"time"
	"context"
	"net/http"
	"html/template"

	"nhooyr.io/websocket"
	"github.com/ChrisMarSilva/cms.golang.teste.web.crawler/crawler"
)


func websocketHandle() func (http.ResponseWriter, *http.Request) {

	tmpl, err := template.ParseFiles("website/templates/websocket.html")
	if err != nil {
		panic(err)
	}

	return func (w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		website := r.FormValue("website")
		if website == "" {
			http.Error(w, "website não pode ser vazio", http.StatusBadRequest)
			return
		}
		
		// done := make(chan bool)
		wb := crawler.NewWebCrawler()
		go wb.VisitLink(website)
		// for log := range wb.Log() {
		// 	fmt.Println("chan", log)
		// }
		// <-done

		c, err := websocket.Accept(w, r, nil)
		if website == "" {
			http.Error(w, "erro no websocket.Accept()", http.StatusBadRequest)
			return
		}

		err = subscriber(r.Context(), c, wb.Log())
		if err != nil {
			http.Error(w, "erro no subscriber()", http.StatusBadRequest)
			return
		}

		tmpl.Execute(w, nil) 
	}

}

func subscriber(ctx context.Context, c *websocket.Conn, logs <-chan string) error {
	ctx = c.CloseRead(ctx)
	for {
		select {
		case msg := <- logs:
			err := writeTimeout(ctx, c, msg)
			if err != nil {
				return err
			}
		case <- ctx.Done():
			return ctx.Err()
		}
	}
}

func writeTimeout(ctx context.Context, c *websocket.Conn, msg string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second * 3)
	defer cancel()
	return c.Write(ctx, websocket.MessageText, []byte(msg))
}

