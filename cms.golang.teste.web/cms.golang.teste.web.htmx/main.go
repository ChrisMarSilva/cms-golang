package main

// go work init ./
// go mod init github.com/chrismarsilva/cms.golang.teste.web.htmx
// go get -u github.com/gofiber/fiber/v3
// go get -u github.com/chasefleming/elem-go
// go get -u github.com/chasefleming/elem-go/htmx
// go get -u github.com/chasefleming/elem-go/style
// go get -u github.com/a-h/templ
// go mod tidy
// go run main.go
// go run .

// go install github.com/cosmtrek/air@latest
// air init
// air

import (
	"fmt"
	"log"
	"strconv"

	"github.com/chasefleming/elem-go"
	"github.com/chasefleming/elem-go/attrs"
	"github.com/chasefleming/elem-go/htmx"
	"github.com/chasefleming/elem-go/styles"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/utils/v2"
)

// Todo model
type Todo struct {
	ID    int
	Title string
	Done  bool
}

type FormData struct {
	Name  string
	Email string
}

var (
	count    int
	formData = FormData{}
	todos    = []Todo{
		{ID: 1, Title: "First task", Done: false},
		{ID: 2, Title: "Second task", Done: true},
	}
)

func init() {
	count = 0
}

func main() {
	app := fiber.New()
	//app.Static("/", "./assets")
	//app.Static("*", "./assets/index.html")

	app.Get("/index.html", func(c fiber.Ctx) error {
		//return web.HTML(http.StatusOK, html, "index.html", data, nil)
		return c.SendFile("./assets/index.html")
	})

	app.Post("/increment", func(c fiber.Ctx) error {
		count++
		return c.SendString(fmt.Sprintf("%d", count))
	})

	app.Post("/decrement", func(c fiber.Ctx) error {
		count--
		return c.SendString(fmt.Sprintf("%d", count))
	})

	app.Post("/submit-form", func(c fiber.Ctx) error {
		formData.Name = c.FormValue("name")
		formData.Email = c.FormValue("email")
		return c.SendString(fmt.Sprintf("Name: %s, Email: %s", formData.Name, formData.Email))
	})

	app.Post("/toggle/:id", toggleTodoRoute)
	app.Post("/add", addTodoRoute)
	app.Get("/", renderTodosRoute)

	log.Fatal(app.Listen(":3000"))
}

func renderTodosRoute(c fiber.Ctx) error {
	c.Type("html")
	return c.SendString(renderTodos(todos))
}

func toggleTodoRoute(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var updatedTodo Todo
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Done = !todo.Done
			updatedTodo = todos[i]
			break
		}
	}
	c.Type("html")
	return c.SendString(createTodoNode(updatedTodo).Render())
}

func addTodoRoute(c fiber.Ctx) error {
	newTitle := utils.CopyString(c.FormValue("newTodo"))
	if newTitle != "" {
		todos = append(todos, Todo{ID: len(todos) + 1, Title: newTitle, Done: false})
	}
	//return c.Redirect("/")
	return c.SendString("")
}

func createTodoNode(todo Todo) elem.Node {
	checkbox := elem.Input(attrs.Props{
		attrs.Type:    "checkbox",
		attrs.Checked: strconv.FormatBool(todo.Done),
		htmx.HXPost:   "/toggle/" + strconv.Itoa(todo.ID),
		htmx.HXTarget: "#todo-" + strconv.Itoa(todo.ID),
	})

	return elem.Li(attrs.Props{
		attrs.ID: "todo-" + strconv.Itoa(todo.ID),
	}, checkbox, elem.Span(attrs.Props{
		attrs.Style: styles.Props{
			styles.TextDecoration: elem.If(todo.Done, "line-through", "none"),
		}.ToInline(),
	}, elem.Text(todo.Title)))
}

func renderTodos(todos []Todo) string {
	inputButtonStyle := styles.Props{
		styles.Width:           "100%",
		styles.Padding:         "10px",
		styles.MarginBottom:    "10px",
		styles.Border:          "1px solid #ccc",
		styles.BorderRadius:    "4px",
		styles.BackgroundColor: "#f9f9f9",
	}

	buttonStyle := styles.Props{
		styles.BackgroundColor: "#007BFF",
		styles.Color:           "white",
		styles.BorderStyle:     "none",
		styles.BorderRadius:    "4px",
		styles.Cursor:          "pointer",
		styles.Width:           "100%",
		styles.Padding:         "8px 12px",
		styles.FontSize:        "14px",
		styles.Height:          "36px",
		styles.MarginRight:     "10px",
	}

	listContainerStyle := styles.Props{
		styles.ListStyleType: "none",
		styles.Padding:       "0",
		styles.Width:         "100%",
	}

	centerContainerStyle := styles.Props{
		styles.MaxWidth:        "300px",
		styles.Margin:          "40px auto",
		styles.Padding:         "20px",
		styles.Border:          "1px solid #ccc",
		styles.BoxShadow:       "0px 0px 10px rgba(0,0,0,0.1)",
		styles.BackgroundColor: "#f9f9f9",
	}

	headContent := elem.Head(nil,
		elem.Script(attrs.Props{attrs.Src: "https://unpkg.com/htmx.org@1.9.11"}),
	)

	bodyStyle2 := styles.Props{
		styles.BackgroundColor: "#f4f4f4",
		styles.FontFamily:      "Arial, sans-serif",
		styles.Height:          "100vh",
		styles.Display:         "flex",
		styles.FlexDirection:   "column",
		styles.AlignItems:      "center",
		styles.JustifyContent:  "center",
	}

	buttonStyle2 := styles.Props{
		styles.Padding:         "10px 20px",
		styles.BackgroundColor: "#007BFF",
		styles.Color:           "#fff",
		styles.BorderColor:     "#007BFF",
		styles.BorderRadius:    "5px",
		styles.Margin:          "10px",
		styles.Cursor:          "pointer",
	}
	bodyContent := elem.Div(
		attrs.Props{attrs.Style: centerContainerStyle.ToInline()},
		elem.Body(attrs.Props{ attrs.Style: bodyStyle2.ToInline() },

		elem.H1(nil, elem.Text("Counter App")),
		elem.Div(attrs.Props{attrs.ID: "count"}, elem.Text("0")),
		elem.Button(attrs.Props{
			htmx.HXPost:   "/increment",
			htmx.HXTarget: "#count",
			attrs.Style:   buttonStyle2.ToInline(),
		}, elem.Text("+")),

		elem.Button(attrs.Props{
			htmx.HXPost:   "/decrement",
			htmx.HXTarget: "#count",
			attrs.Style:   buttonStyle2.ToInline(),
		}, elem.Text("-")),

		elem.H1(nil, elem.Text("Simple Form App")),
		elem.Form(attrs.Props{
			attrs.Action: "/submit-form",
			attrs.Method: "POST",
			htmx.HXPost:  "/submit-form",
			htmx.HXSwap:  "outerHTML",
		},
			elem.Label(attrs.Props{attrs.For: "name"}, elem.Text("Name: ")),
			elem.Input(attrs.Props{
				attrs.Type: "text",
				attrs.Name: "name",
				attrs.ID:   "name",
			}),
			elem.Br(nil),
			elem.Label(attrs.Props{attrs.For: "email"}, elem.Text("Email: ")),
			elem.Input(attrs.Props{
				attrs.Type: "email",
				attrs.Name: "email",
				attrs.ID:   "email",
			}),
			elem.Br(nil),
			elem.Input(attrs.Props{
				attrs.Type:  "submit",
				attrs.Value: "Submit",
			}),
		),
		elem.Div(attrs.Props{attrs.ID: "response"}, elem.Text("")),

		elem.H1(nil, elem.Text("Todo List")),
		elem.Form(
			attrs.Props{attrs.Method: "post", attrs.Action: "/add"},
			elem.Input(
				attrs.Props{
					attrs.Type:        "text",
					attrs.Name:        "newTodo",
					attrs.Placeholder: "Add new task...",
					attrs.Style:       inputButtonStyle.ToInline(),
				},
			),
			elem.Button(
				attrs.Props{
					attrs.Type:  "submit",
					attrs.Style: buttonStyle.ToInline(),
				},
				elem.Text("Add"),
			),
		),
		elem.Ul(
			attrs.Props{attrs.Style: listContainerStyle.ToInline()},
			elem.TransformEach(todos, createTodoNode)...,
		),
	)

	htmlContent := elem.Html(nil, headContent, bodyContent)
	return htmlContent.Render()
}
