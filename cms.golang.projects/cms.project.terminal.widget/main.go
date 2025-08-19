package main

// go mod init github.com/chrismarsilva/cms.project.terminal.widget
// go get -u github.com/rivo/tview@master
// go mod tidy

// go run main.go

// go install github.com/air-verse/air@latest
// air init
// air

import (
	"log/slog"

	"github.com/rivo/tview"
)

func main() {
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")

	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		slog.Error("Error running program", slog.Any("err", err))
	}
}
