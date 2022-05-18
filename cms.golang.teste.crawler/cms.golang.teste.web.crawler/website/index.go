package website

import (
	"log"
	"net/http"
	"html/template"

	"github.com/ChrisMarSilva/cms.golang.teste.web.crawler/db"
	"github.com/ChrisMarSilva/cms.golang.teste.web.crawler/model"
)


type DataLinks struct {
	Links []model.VisitedLink
}

func indexHandle() func (http.ResponseWriter, *http.Request) {

	tmpl, err := template.ParseFiles("website/templates/index.html")
	if err != nil {
		panic(err)
	}

	return func (w http.ResponseWriter, r *http.Request) {
		
		links, err := db.SelectAll()
		if err != nil {
			log.Println("Erro db.SelectAll():", err)
			return
		}

		data := DataLinks{Links: links}
	
		// for iIndx, visitedLink := range links {
		// 	log.Println("TABELA", iIndx, "====", visitedLink.WebSite, "====", visitedLink.Link, "====", visitedLink.VisitedDate)
		// }

		tmpl.Execute(w, data) 
	}

}


