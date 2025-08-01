package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

// go mod init cms.project.todo
// go get -u xxxxxxx
// go mod tidy

// go run main.go

type ToDo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
}

var todos = []ToDo{
	{ID: 1, Title: "Learn Go", IsCompleted: true, CreatedAt: time.Now()},
	{ID: 2, Title: "Learn HTMX  ", IsCompleted: false, CreatedAt: time.Now()},
	{ID: 3, Title: "Learn Alpine  ", IsCompleted: false, CreatedAt: time.Now()},
	{ID: 4, Title: "Learn Tailwind", IsCompleted: false, CreatedAt: time.Now()},
}

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["index.html"] = template.Must(template.ParseFiles("index.html"))
}

func main() {
	//http.HandleFunc("/", indexHandler)
	http.HandleFunc("/", h1Handler)
	http.HandleFunc("/add-film/", h2Handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := templates["index.html"]
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

type Film struct {
	Title    string
	Director string
}

func h1Handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}
	tmpl.Execute(w, films)
}

func h2Handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	// htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
	// tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}
