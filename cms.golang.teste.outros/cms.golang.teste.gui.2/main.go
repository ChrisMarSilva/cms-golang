package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.gui2
// go get -u fyne.io/fyne/v2
// go get -u fyne.io/fyne/v2/cmd/fyne_demo/
// go get -u fyne.io/fyne/v2/cmd/fyne
// go mod tidy

// go run main.go
// go build main.go

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
