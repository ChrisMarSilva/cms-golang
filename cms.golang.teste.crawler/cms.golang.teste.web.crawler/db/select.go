package db

import (
	"log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/ChrisMarSilva/cms.golang.teste.web.crawler/model"
)

func SelectAll() ([]model.VisitedLink, error) {

	db, err := getConnection()
	if err != nil {
		return []model.VisitedLink{}, err
	}
	defer db.Close()

	rows, err := db.Query("select website, link, data from links order by data desc")
	if err != nil {
		return []model.VisitedLink{}, err
	}
	defer rows.Close()

	var links []model.VisitedLink

	for rows.Next() {
		var visitedLink model.VisitedLink

		err = rows.Scan(&visitedLink.WebSite, &visitedLink.Link, &visitedLink.VisitedDate)
		if err != nil {
			return []model.VisitedLink{}, err
		}

		if len(visitedLink.WebSite) > 50 {
			visitedLink.WebSite = visitedLink.WebSite[:50]
		}
		
		if len(visitedLink.Link) > 50 {
			visitedLink.Link = visitedLink.Link[:50]
		}

		links = append(links, visitedLink)
		// log.Println("TABELA", visitedLink.WebSite, "====", visitedLink.Link, "====", visitedLink.VisitedDate)
	}

	err = rows.Err()
	if err != nil {
		return []model.VisitedLink{}, err
	}

	// log.Println(links)
	// log.Println(len(links))

	return links, nil

}

func SelectExist(linkSearch string) error {

	db, err := getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("select id, website, link, data from links where link = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	var id int
	var website string
	var link string
	var data string

	err = stmt.QueryRow(linkSearch).Scan(&id, &website, &link, &data)
	if err != nil {
		return err
	}

	// log.Println("ACHOU", id, website, link, data)

	return nil

}

func SelectTotal() error {

	db, err := getConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	var tot int

	err = db.QueryRow("select count(1) from links").Scan(&tot)
	if err != nil {
		return err
	}

	log.Println("TOTAL", tot)

	return nil

}
