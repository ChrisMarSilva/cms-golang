package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// go mod init github.com/chrismarsilva/cms.golang.teste.gui2
// go get -u fyne.io/fyne/v2
// go get -u  fyne.io/fyne/v2/cmd/fyne_demo/
// go mod tidy

// go run main.go
// go build main.go

var client *http.Client

func main() {
	a := app.New()
	w := a.NewWindow("Table widget")
	w.Resize(fyne.NewSize(400, 400))
	table := widget.NewTable(
		func() (int, int) { return 3, 3 },
		func() fyne.CanvasObject { return widget.NewLabel("....") },
		func(i widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(fmt.Sprintf("%d %d", i.Col, i.Row))
		},
	)
	w.SetContent(table)
	w.ShowAndRun()
}

func main2() {
	client = &http.Client{Timeout: 10 * time.Second}

	a := app.New()
	win := a.NewWindow("Get Useless Fact")
	win.Resize(fyne.NewSize(800, 300))

	title := canvas.NewText("Get Your Useless Facts", color.White)
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	factText := widget.NewLabel("")
	factText.Wrapping = fyne.TextWrapWord

	button := widget.NewButton("Get Fact", func() {
		fact, err := getRandomFact()
		if err != nil {
			dialog.ShowError(err, win)
		} else {
			factText.SetText(fact.Text)
		}
	})

	hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	vBox := container.New(layout.NewVBoxLayout(), title, hBox, widget.NewSeparator(), factText)

	win.SetContent(vBox)
	win.ShowAndRun()

}

type randomFact struct {
	Text string `json:"text"`
}

func getRandomFact() (randomFact, error) {
	var fact randomFact
	resp, err := client.Get("https://uselessfacts.jsph.pl/random.json?language=en")
	if err != nil {
		return randomFact{}, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&fact)
	if err != nil {
		return randomFact{}, err
	}

	return fact, nil
}
