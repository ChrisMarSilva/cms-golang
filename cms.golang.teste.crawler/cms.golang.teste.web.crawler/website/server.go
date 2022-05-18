package website

import (
	"log"
	"net/http"
)

func Run() {
	log.Println("Server run port 8000")
	http.HandleFunc("/", indexHandle())
	http.HandleFunc("/busca", websocketHandle())
	log.Fatal(http.ListenAndServe(":8000", nil))
}
